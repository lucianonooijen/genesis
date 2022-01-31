package user

import (
	"context"

	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/server/internal/entities"
	"git.bytecode.nl/bytecode/genesis/server/internal/interactors"
)

// GetUserProfile gets the user profile.
func GetUserProfile(s *interactors.Services, userID int32) (*entities.UserProfile, error) {
	logger := s.BaseLogger.Named("domains/user/GetUserProfile").With(zap.Int32("user_id", userID))
	ctx := context.TODO()

	logger.Debug("fetching user profile")

	firstName, err := s.Database.GetUserProfileByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	profile := entities.UserProfile{
		FirstName: firstName,
	}

	return &profile, nil
}

// UpdateUserProfile updates the user profile and returns the updated profile.
func UpdateUserProfile(s *interactors.Services, userID int32, newProfile entities.UserProfile) (*entities.UserProfile, error) {
	logger := s.BaseLogger.Named("domains/user/UpdateUserProfile").With(zap.Int32("user_id", userID))
	ctx := context.TODO()

	err := s.Database.UpdateUserProfile(ctx, database.UpdateUserProfileParams{
		FirstName: newProfile.FirstName,
		ID:        userID,
	})
	if err != nil {
		return nil, err
	}

	logger.Debug("returning results from UpdateUserProfile")

	return GetUserProfile(s, userID)
}

// DeleteAccount deletes the account based on userID.
func DeleteAccount(s *interactors.Services, userID int32, password string) error {
	logger := s.BaseLogger.Named("domains/user/DeleteAccount").With(zap.Int32("user_id", userID))
	ctx := context.TODO()

	logger.Debug("fetching user from database")

	user, err := s.Database.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	logger.Debug("comparing passwords")

	err = s.PassHash.ComparePassToHash(password, user.PasswordHash)
	if err != nil {
		return err
	}

	logger.Info("starting deletion of user account after password verification")

	logger.Debug("deleting password_forgot")

	if err = s.Database.DeletePasswordForgotByUserId(ctx, userID); err != nil {
		return err
	}

	logger.Debug("deleting user")

	if err = s.Database.DeleteUser(ctx, userID); err != nil {
		return err
	}

	logger.Info("user deleted from database")

	return nil
}
