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
