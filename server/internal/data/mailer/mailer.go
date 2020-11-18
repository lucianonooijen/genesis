package mailer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"git.bytecode.nl/bytecode/genesis/internal/interactors"
	"git.bytecode.nl/bytecode/genesis/internal/utils/logger"
	"github.com/matcornic/hermes/v2"
)

var log = logger.New("mailer")

// Mailer contains methods for interactors.MailerService to send emails to users of Genesis
type Mailer struct {
	fromEmail         string
	fromName          string
	apiKey            string
	staticFileURLBase string
	hermes            hermes.Hermes
}

// New returns a new Mailer instance as a interactors.MailerService
func New(fromEmail, fromName, apiKey, staticFileURLBase string) (interactors.MailerService, error) {
	if fromEmail == "" || fromName == "" || apiKey == "" || staticFileURLBase == "" {
		return nil, errors.New("arguments cannot have default values")
	}
	mailer := Mailer{
		fromEmail:         fromEmail,
		fromName:          fromName,
		apiKey:            apiKey,
		staticFileURLBase: staticFileURLBase,
		hermes: hermes.Hermes{
			Product: hermes.Product{
				Name:        "Genesis",
				Link:        "https://bytecode.nl",                         // TODO: Load this from the configuration (?)
				Logo:        fmt.Sprintf("%s/logo.png", staticFileURLBase), // TODO: This breaks on localhost due to SendInBlue serving the images. Can this be fixed?
				Copyright:   fmt.Sprintf("Copyright © %d Genesis. All rights reserved.", time.Now().Year()),
				TroubleText: "Als je problemen hebt met de knop '{ACTION}', knip en plak de URL hieronder in je webbrowser.",
			},
		},
	}
	return mailer, nil
}

/**
 * Email sending handler
 */

type contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type emailRequest struct {
	Sender      contact   `json:"sender"`
	To          []contact `json:"to"`
	Subject     string    `json:"subject"`
	HTMLContent string    `json:"htmlContent"`
}

func (m Mailer) sendEmail(toMail string, toName string, subject string, HTMLContents string) error {
	url := "https://api.sendinblue.com/v3/smtp/email"
	// TODO: Support reply-to (?)
	reqBody := emailRequest{
		Sender: contact{
			Name:  m.fromName,
			Email: m.fromEmail,
		},
		To:          []contact{{Name: toName, Email: toMail}},
		Subject:     subject,
		HTMLContent: HTMLContents,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(jsonBody))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", m.apiKey)

	log.Debug("making email sending request")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Error(err)
		}
	}()

	body, _ := ioutil.ReadAll(res.Body)
	log.Trace("received email response: " + string(body))
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("did not receive 2xx status code while sending email, got %d", res.StatusCode)
	}
	return nil
}
