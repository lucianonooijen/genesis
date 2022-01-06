package mailer

// SendAccountCreated notifies a user of successful account creation.
func (m *Mailer) SendAccountCreated(toEmail, firstName string) error {
	html, err := m.hermes.GenerateHTML(accountCreatedTemplate(firstName))
	if err != nil {
		return err
	}

	// TODO: Add button for account/email confirmation
	return m.sendEmail(toEmail, firstName, "Genesis account aangemaakt", html)
}

// SendLoginSuccess notifies a user that someone (hopefully the user him-/herself!) has logged in to their account.
func (m *Mailer) SendLoginSuccess(toEmail, firstName string) error {
	html, err := m.hermes.GenerateHTML(accountLoginTemplate(firstName))
	if err != nil {
		return err
	}

	return m.sendEmail(toEmail, firstName, "Succesvolle inlog met jouw Genesis account", html)
}

// SendPasswordResetToken sends a password reset token to a user.
func (m *Mailer) SendPasswordResetToken(toEmail, firstName, resetToken string) error {
	html, err := m.hermes.GenerateHTML(passwordResetStartTemplate(firstName, resetToken))
	if err != nil {
		return err
	}

	return m.sendEmail(toEmail, firstName, "Je wachtwoord voor Genesis herstellen", html)
}

// SendPasswordResetConfirmation sends a confirmation to the user that his/her password has been reset.
func (m *Mailer) SendPasswordResetConfirmation(toEmail, firstName string) error {
	html, err := m.hermes.GenerateHTML(passwordResetCompleteTemplate(firstName))
	if err != nil {
		return err
	}

	return m.sendEmail(toEmail, firstName, "Het wachtwoord van jouw Genesis account is opnieuw ingesteld", html)
}
