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
