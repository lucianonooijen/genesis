package entities

// PushNotificationRegister contains the data to register a push notification token for a user.
type PushNotificationRegister struct {
	Platform string `json:"platform" validate:"required,eq=Android|eq=iOS"`
	Token    string `json:"token" validate:"required"`
}
