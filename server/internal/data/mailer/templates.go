package mailer

import "github.com/matcornic/hermes/v2"

func addDefaultValuesToTemplate(template *hermes.Email) {
	template.Body.Greeting = "Beste"
	template.Body.Signature = "Met hartelijke groet"
}

func accountCreatedTemplate(firstName string) hermes.Email {
	email := hermes.Email{
		Body: hermes.Body{
			Name: firstName,
			Intros: []string{
				"Welkom bij Genesis! Je account is succesvol aangemaakt",
				"Je kan nu inloggen met de door jou ingestelde inloggegevens",
			},
			Outros: []string{
				"Heb je hulp nodig, of een vraag? Antwoord gerust op deze email. We helpen je graag.",
			},
		},
	}
	addDefaultValuesToTemplate(&email)

	return email
}

func accountLoginTemplate(firstName string) hermes.Email {
	email := hermes.Email{
		Body: hermes.Body{
			Name: firstName,
			Intros: []string{
				"Er is zojuist succesvol ingelogd met jouw Genesis account.",
				"Mocht je dit niet zelf zijn geweest, laat het ons zo snel mogelijk weten.",
			},
			Outros: []string{
				"Heb je hulp nodig, of een vraag? Antwoord gerust op deze email. We helpen je graag.",
			},
		},
	}
	addDefaultValuesToTemplate(&email)

	return email
}

func passwordResetStartTemplate(firstName, resetToken string) hermes.Email {
	email := hermes.Email{
		Body: hermes.Body{
			Name: firstName,
			Intros: []string{
				"Deze email bevat jouw unieke code om je wachtwoord te herstellen.",
				"Gebruik deze code alleen zelf, medewerkers van Genesis vragen nooit om deze code.",
			},
			Dictionary: []hermes.Entry{{
				Key:   "Reset token",
				Value: resetToken,
			}},
			Outros: []string{
				"Als je geen aanvraag voor een wachtwoordherstel hebt gedaan, kan je deze mail negeren.",
				"Heb je hulp nodig, of een vraag? Antwoord gerust op deze email. We helpen je graag.",
			},
		},
	}
	addDefaultValuesToTemplate(&email)

	return email
}

func passwordResetCompleteTemplate(firstName string) hermes.Email {
	email := hermes.Email{
		Body: hermes.Body{
			Name: firstName,
			Intros: []string{
				"Het wachtwoord van jouw Genesis account is opnieuw ingesteld",
				"Ben je dit niet zelf geweest? Neem dan zo snel mogelijk contact op met ons.",
			},
			Outros: []string{
				"Heb je hulp nodig, of een vraag? Antwoord gerust op deze email. We helpen je graag.",
			},
		},
	}
	addDefaultValuesToTemplate(&email)

	return email
}
