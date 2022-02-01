package pushnotifications

// PushNotificationData is the struct used for defining platform-agnostic push notification data.
type PushNotificationData struct {
	Title   string
	Message string
	Custom  map[string]string
}
