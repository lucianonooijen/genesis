package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/internal/entities"
	"git.bytecode.nl/bytecode/genesis/internal/interactors"
)

var (
	// ErrPasswordResetIsUsed is the error indicating the password reset token has already been used.
	ErrPasswordResetIsUsed = fmt.Errorf("password reset token is already used, please request a new token")

	// ErrPasswordResetExpired is the error indicating that the reset token has been expired.
	ErrPasswordResetExpired = fmt.Errorf("password reset token is expired, please request a new token")
)

// PasswordResetStart starts the password reset procedure, by creating a password reset database entry
// and sending an email to the known email of the user.
func PasswordResetStart(s *interactors.Services, email string) error {
	log := s.BaseLogger.Named("domains/user/PasswordResetStart").With(zap.String("email", email))
	ctx := context.TODO()

	log.Info("entering password reset start procedure")

	// Fetch user from DB
	log.Debug("fetching user by email")

	user, err := s.Database.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	// Generate password reset database entry
	log.Debug("creating password reset entry for user", zap.Int32("user_id", user.ID))

	if err = s.Database.AddPasswordForgotForUser(ctx, user.ID); err != nil {
		return err
	}

	// Fetch the new password reset database entry
	log.Debug("fetching password reset entry for user", zap.Int32("user_id", user.ID))

	passReset, err := s.Database.GetPasswordResetByUserId(ctx, user.ID)
	if err != nil {
		return err
	}

	// Send email with password reset token to user
	log.Info("sending password reset token to user",
		zap.String("first_name", user.FirstName),
		zap.String("reset_token", passReset.ResetToken.String()))

	if err = s.Mailer.SendPasswordResetToken(email, user.FirstName, passReset.ResetToken.String()); err != nil {
		return err
	}

	// Done
	log.Info("completed initialization of password reset for user")

	return nil
}

// PasswordResetComplete completes the password reset procedure, sets the new password if the reset token is valid,
// sends an email notification to the user and returns the JWT
// nolint:funlen,gocyclo // hard to break up, better readable like this
func PasswordResetComplete(s *interactors.Services, resetToken uuid.UUID, password string) (*entities.JwtResponse, error) {
	log := s.BaseLogger.Named("domains/user/PasswordResetComplete").With(zap.String("reset_token", resetToken.String()))
	ctx := context.TODO()

	log.Info("starting password reset completion")

	// Fetch the new password reset database entry
	log.Debug("fetching password reset entry by token", zap.String("reset_token", resetToken.String()))

	passReset, err := s.Database.GetPasswordResetByResetToken(ctx, resetToken)
	if err != nil {
		return nil, err
	}

	// Check if password reset is valid
	log.Debug("checking if password reset is valid")

	if passReset.IsUsed {
		log.Info("password reset is already used")

		return nil, ErrPasswordResetIsUsed
	}

	if time.Now().After(passReset.ValidUntil) {
		log.Info("password reset token is expired")

		return nil, ErrPasswordResetExpired
	}

	// Fetch user from DB
	log.Debug("fetching user by id", zap.Int32("user_id", passReset.UserID))

	user, err := s.Database.GetUserByID(ctx, passReset.UserID)
	if err != nil {
		return nil, err
	}

	// Start transaction
	log.Debug("starting database transaction")

	dbTx, err := s.DBConn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func(dbTx *sql.Tx) {
		log.Debug("rolling back database transaction in defer block")

		err = dbTx.Rollback()
		if err != nil {
			log.Error("error rolling back dbTx", zap.Error(err))
		}
	}(dbTx)

	// tx is a Database instance, using the dbTx database transaction
	tx := s.Database.WithTx(dbTx)

	// Mark password reset as used
	log.Debug("marking reset token as used")

	err = tx.MarkResetTokenUsed(ctx, resetToken)
	if err != nil {
		return nil, err
	}

	// Generate new hash and save to database
	log.Debug("generating new password hash")

	passHash, err := s.PassHash.PlainTextToHash(password)
	if err != nil {
		return nil, err
	}

	log.Debug("saving hash to database")

	err = tx.UpdateUserPassword(ctx, database.UpdateUserPasswordParams{
		ID:           user.ID,
		PasswordHash: passHash,
	})
	if err != nil {
		return nil, err
	}

	// Commit transaction
	log.Debug("committing database transaction")

	err = dbTx.Commit()
	if err != nil {
		return nil, err
	}

	// Send confirmation email
	log.Debug("sending password reset confirmation to user",
		zap.String("email", user.Email),
		zap.String("first_name", user.FirstName))

	if err = s.Mailer.SendPasswordResetConfirmation(user.Email, user.FirstName); err != nil {
		return nil, err
	}

	// Generate JWT
	jwt, err := jwtForUser(s, &user)
	if err != nil {
		return nil, err
	}

	// Done
	log.Info("completed password reset for user")

	return &entities.JwtResponse{JWT: jwt}, nil
}
