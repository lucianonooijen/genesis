package pushnotifications

import (
	"context"

	"firebase.google.com/go/messaging"
	"go.uber.org/zap"
)

// SendNotificationAndroid sends a notification to an Android device.
func (pn PushService) SendNotificationAndroid(token string, notification PushNotificationData) error {
	ctx := context.TODO()
	log := pn.logger.
		Named("SendNotificationAndroid").
		With(zap.String("token", token), zap.Any("notification", notification))

	log.Info("sending Android push notification")

	// Build the notification message struct
	message := &messaging.Message{
		Token: token,
		Data:  notification.Custom,
		Notification: &messaging.Notification{
			Title: notification.Title,
			Body:  notification.Message,
		},
		Android: &messaging.AndroidConfig{
			Data:     notification.Custom,
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				Title:                 notification.Title,
				Body:                  notification.Message,
				Priority:              messaging.PriorityMax,
				DefaultSound:          true,
				DefaultLightSettings:  true,
				DefaultVibrateTimings: true,
				Visibility:            messaging.VisibilityPrivate,
				ChannelID:             pn.fcmChannelID,
			},
		},
	}

	// Send the notification message
	res, err := pn.fcmClient.Send(ctx, message)

	log.Info("response received", zap.Any("response", res))

	// If the token is not registered, the function should not error, only log a warning
	if messaging.IsRegistrationTokenNotRegistered(err) {
		log.Warn("registration-token-not-registered: Android token is not not registered")

		return nil
	}

	return err
}
