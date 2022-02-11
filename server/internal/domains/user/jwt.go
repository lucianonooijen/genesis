package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/jwt"
	"git.bytecode.nl/bytecode/genesis/server/internal/interactors"
)

// HTTPJwtFunc is the function used in the Gin HTTP middleware.
type HTTPJwtFunc func(jwt string) (*database.User, error)

func jwtForUser(s *interactors.Services, user *database.User) (string, error) {
	uniqueUserIdentifier := user.UserUuid.String()

	return s.JWT.CreateJWT(uniqueUserIdentifier, user.PasswordUuid)
}

// GetUserByJwt validates the JWT and returns the user if it's valid.
func GetUserByJwt(s *interactors.Services, jwtString string) (*database.User, error) {
	log := s.BaseLogger.Named("domains/user/GetUserByJwt")

	log.Debug("extracting audience (user uuid) from jwtString")

	aud, err := jwt.ExtractAudience(jwtString)
	if err != nil {
		return nil, err
	}

	log.Info("extracted user uuid from jwt", zap.String("user_uuid", aud))

	userUUID, err := uuid.Parse(aud)
	if err != nil {
		return nil, err
	}

	log.Debug("fetching user from database")

	user, err := s.Database.GetUserByUuid(context.TODO(), userUUID)
	if err != nil {
		return nil, err
	}

	log.Debug("user fetched from database", zap.Int32("user_id", user.ID))

	log.Debug("validating jwtString, including the jwt key id with the user password uuid")

	userUUIDString, err := s.JWT.ValidateJWT(jwtString, user.PasswordUuid)
	if err != nil {
		return nil, err
	}

	if userUUIDString != aud {
		return nil, fmt.Errorf("userUUIDString %s is not the same as the aud %s", userUUIDString, aud)
	}

	return &user, err
}

// GenerateUserJwtMiddleware generates a closure wrapped around GetUserByJwt to generate the HTTPJwtFunc used by Gin.
func GenerateUserJwtMiddleware(s *interactors.Services) HTTPJwtFunc {
	return func(jwt string) (*database.User, error) {
		return GetUserByJwt(s, jwt)
	}
}
