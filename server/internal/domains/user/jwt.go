package user

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/server/internal/interactors"
)

// HTTPJwtFunc is the function used in the Gin HTTP middleware.
type HTTPJwtFunc func(jwt string) (*database.User, error)

func jwtForUser(s *interactors.Services, user *database.User) (string, error) {
	uniqueUserIdentifier := user.UserUuid.String()

	return s.JWT.CreateJWT(uniqueUserIdentifier)
}

// GetUserByJwt validates the JWT and returns the user if it's valid.
func GetUserByJwt(s *interactors.Services, jwt string) (*database.User, error) {
	log := s.BaseLogger.Named("domains/user/GetUserByJwt")
	log.Debug("validating jwt")

	userUUIDString, err := s.JWT.ValidateJWT(jwt)
	if err != nil {
		return nil, err
	}

	userUUID, err := uuid.Parse(userUUIDString)
	if err != nil {
		return nil, err
	}

	log.Info("validated JWT, detected user uuid", zap.String("userUUIDString", userUUIDString))
	log = log.With(zap.String("user_email", userUUIDString))

	log.Debug("fetching user from database")

	user, err := s.Database.GetUserByUuid(context.TODO(), userUUID)
	if err != nil {
		return nil, err
	}

	log.Debug("user fetched from database", zap.Int32("user_id", user.ID))

	return &user, err
}

// GenerateUserJwtMiddleware generates a closure wrapped around GetUserByJwt to generate the HTTPJwtFunc used by Gin.
func GenerateUserJwtMiddleware(s *interactors.Services) HTTPJwtFunc {
	return func(jwt string) (*database.User, error) {
		return GetUserByJwt(s, jwt)
	}
}
