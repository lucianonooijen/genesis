package user

import (
	"context"

	"git.bytecode.nl/bytecode/genesis/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/internal/entities"
	"git.bytecode.nl/bytecode/genesis/internal/interactors"
)

// CreateUser creates a user
func CreateUser(s *interactors.Services, newUser entities.NewUserRequest) (*entities.NewUserResponse, error) {
	log := s.BaseLogger.Named("domains/user/CreateUser")
	ctx := context.TODO()

	// Hash password
	hashedPassword, err := s.PassHash.PlainTextToHash(newUser.Password)
	if err != nil {
		return nil, err
	}
	log.Debug("hashed password successfully")

	// Create user
	userToDB := database.CreateUserParams{
		Email:        newUser.Email,
		PasswordHash: hashedPassword,
		FirstName:    newUser.FirstName,
	}
	err = s.Database.CreateUser(ctx, userToDB)
	if err != nil {
		return nil, err
	}
	log.Debug("saved user to database")

	// Fetch new user
	user, err := s.Database.GetUserByEmail(ctx, newUser.Email)
	if err != nil {
		return nil, err
	}
	log.Debug("fetched new user from database")

	// Send confirmation email
	if err := s.Mailer.SendAccountCreated(user.Email, user.FirstName); err != nil {
		return nil, err
	}
	log.Debug("confirmation email has been sent")

	// Generate JWT
	jwt, err := s.JWT.CreateJWT(string(user.ID))
	if err != nil {
		return nil, err
	}
	log.Debug("jwt has been generated")

	// Return
	log.Debug("create user done")
	return &entities.NewUserResponse{JWT: jwt}, nil
}
