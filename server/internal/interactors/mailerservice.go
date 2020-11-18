package interactors

// MailerService contains all methods for sending emails to users of Genesis
type MailerService interface {
	SendAccountCreated(toEmail, firstName string) error
	SendLoginSuccess(toEmail, firstName string) error
	//SendPasswordForgotStart(toEmail, firstName string) error // TODO: Implement
	//SendPasswordForgotCompleted(toEmail, firstName string) error // TODO: Implement
}
