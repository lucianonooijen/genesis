package pushnotifications

import (
	"fmt"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"go.uber.org/zap"
)

// SendNotificationIos sends a notification to an iOS device.
func (pn PushService) SendNotificationIos(token string, notification PushNotificationData) error {
	log := pn.logger.
		Named("SendNotificationIos").
		With(zap.String("token", token), zap.Any("notification", notification))

	log.Info("sending iOS push notification")

	// Generate payload
	pushPayload := payload.
		NewPayload().
		AlertTitle(notification.Title).
		AlertBody(notification.Message)

	for key, val := range notification.Custom {
		pushPayload = pushPayload.Custom(key, val)
	}

	notificationPayload := &apns2.Notification{
		DeviceToken: token,
		Topic:       pn.apnsTopic,
		Payload:     pushPayload,
	}

	// Send the push notification to Apple
	res, err := pn.apnsClient.Push(notificationPayload)
	if err != nil {
		return err
	}

	// Handle errors if they occur
	log.Info("iOS notification pushed",
		zap.String("res.ApnsID", res.ApnsID), zap.Int("res.StatusCode", res.StatusCode), zap.String("res.Reason", res.Reason))

	if !res.Sent() {
		if res.Reason == apns2.ReasonBadDeviceToken || res.Reason == apns2.ReasonUnregistered {
			// We just log as a warning to count them in Sentry, but they are expected, so no error is returned
			log.Warn("iOS notification was not sent",
				zap.String("res.ApnsID", res.ApnsID), zap.Int("res.StatusCode", res.StatusCode), zap.String("res.Reason", res.Reason))

			return nil
		}

		return fmt.Errorf("iOS notification not sent, code %d and reason '%s'", res.StatusCode, res.Reason)
	}

	return nil
}
