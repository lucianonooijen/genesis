package pushnotifications

import (
	"fmt"

	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
)

// SendNotification sends a notification to a list of targets.
func (pn PushService) SendNotification(targets []database.UserPushToken, notificationData PushNotificationData) []error {
	log := pn.logger.
		Named("SendNotification").
		With(zap.Int("target_count", len(targets)), zap.Any("notification", notificationData), zap.Any("targets", targets))

	log.Info("sending notification")

	var errors []error

	for _, target := range targets {
		switch target.Platform {
		case "iOS":
			log.Debug("sending notification for iOS", zap.String("token", target.Token))

			if err := pn.SendNotificationIos(target.Token, notificationData); err != nil {
				errors = append(errors, err)
			}
		case "Android":
			log.Debug("sending notification for Android", zap.String("token", target.Token))

			if err := pn.SendNotificationAndroid(target.Token, notificationData); err != nil {
				errors = append(errors, err)
			}
		default:
			errors = append(errors, fmt.Errorf("unsupported platform %s", target.Platform))
		}
	}

	return errors
}

// SendNotificationAsync is the exact same as SendNotification but runs in a goroutine.
// It does not return an error and should thus not be trusted to be error-safe.
// Errors are logged (and sent to Sentry), but not returned
// TODO: Send notifications concurrently to speed up delivery.
func (pn PushService) SendNotificationAsync(targets []database.UserPushToken, notificationData PushNotificationData) {
	log := pn.logger.Named("SendNotificationAsync")

	go func() {
		errs := pn.SendNotification(targets, notificationData)

		for _, err := range errs {
			log.Error("notification sending error", zap.Error(err))
		}
	}()
}
