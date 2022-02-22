package mailer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/matcornic/hermes/v2"
	"go.uber.org/zap"
)

// Mailer contains methods for interactors.Instance to send emails to users of Genesis.
type Mailer struct {
	logger            *zap.Logger
	fromEmail         string
	fromName          string
	apiKey            string
	staticFileURLBase string
	hermes            hermes.Hermes
}

// New returns a new Mailer instance as a interactors.Instance.
func New(loggerBase *zap.Logger, fromEmail, fromName, apiKey, staticFileURLBase string) (*Mailer, error) {
	if fromEmail == "" || fromName == "" || apiKey == "" || staticFileURLBase == "" {
		return nil, errors.New("arguments cannot have default values")
	}

	mailer := Mailer{
		logger:            loggerBase.Named("mailer"),
		fromEmail:         fromEmail,
		fromName:          fromName,
		apiKey:            apiKey,
		staticFileURLBase: staticFileURLBase,
		hermes: hermes.Hermes{
			Product: hermes.Product{
				Name:        "Genesis",
				Link:        "https://bytecode.nl",
				Logo:        fmt.Sprintf("%s/logo.png", staticFileURLBase), // NOTE: This breaks on localhost due to SendInBlue serving the images. Can this be fixed?
				Copyright:   fmt.Sprintf("Copyright © %d Genesis. All rights reserved.", time.Now().Year()),
				TroubleText: "Als je problemen hebt met de knop '{ACTION}', knip en plak de URL hieronder in je webbrowser.",
			},
		},
	}

	return &mailer, nil
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

func (m *Mailer) sendEmail(toMail, toName, subject, htmlContents string) error {
	url := "https://api.sendinblue.com/v3/smtp/email"

	reqBody := emailRequest{
		Sender: contact{
			Name:  m.fromName,
			Email: m.fromEmail,
		},
		To:          []contact{{Name: toName, Email: toMail}},
		Subject:     subject,
		HTMLContent: htmlContents,
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

	m.logger.Debug("making email sending request")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			m.logger.Error("error when closing res.Body", zap.Error(err))
		}
	}()

	body, _ := io.ReadAll(res.Body)
	m.logger.Debug("received email request response", zap.ByteString("response", body))

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("did not receive 2xx status code while sending email, got %d", res.StatusCode)
	}

	return nil
}
