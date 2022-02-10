package user

import (
	"context"

	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/server/internal/data/pushnotifications"
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

// SendUserPushNotification sends a push notification to a user based on the userID.
func SendUserPushNotification(s *interactors.Services, userID int32, data pushnotifications.PushNotificationData) error {
	log := s.BaseLogger.
		Named("domains/user/SendUserPushNotification").
		With(zap.Int32("userID", userID), zap.Any("pushnotificationdata", data))
	ctx := context.TODO()

	log.Debug("fetching user push notification tokens")

	toks, err := s.Database.GetPushTokensForUser(ctx, userID)
	if err != nil {
		return err
	}

	tokLen := len(toks)
	if tokLen == 0 {
		log.Warn("no push notification tokens for user found")

		return nil
	}

	log.Info("sending user push notifications asynchronously", zap.Int("target_count", tokLen))

	s.PushNotifications.SendNotificationAsync(toks, data)

	return nil
}
