package user

import (
	"context"

	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/server/internal/interactors"
)

// SaveUserPushNotificationToken saves a notification token for the given platform for the user to the database.
func SaveUserPushNotificationToken(s *interactors.Services, userID int32, platform, token string) error {
	log := s.BaseLogger.
		Named("domains/user/SaveUserPushNotificationToken").
		With(zap.Int32("userID", userID), zap.String("platform", platform), zap.String("token", token))
	ctx := context.TODO()

	log.Debug("scanning platform into database.MobilePlatform type")

	var platformEnum database.MobilePlatform

	err := platformEnum.Scan(platform) // error handling not required as the incoming request is already validated to be "Android"|"iOS"
	if err != nil {
		return err
	}

	log.Info("adding user push token to database")

	err = s.Database.AddUserPushToken(ctx, database.AddUserPushTokenParams{
		Userid:   userID,
		Platform: platformEnum,
		Token:    token,
	})

	return err
}
