package pushnotifications

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/go-playground/validator/v10"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	"go.uber.org/zap"
	"google.golang.org/api/option"

	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/config"
)

// PushService is the struct containing the methods for sending push notifications.
type PushService struct {
	// iOS
	apnsClient *apns2.Client `validate:"required"`
	apnsTopic  string        `validate:"required"`

	// Android
	fcmClient    *messaging.Client `validate:"required"`
	fcmChannelID string            `validate:"required"`

	// Shared
	logger *zap.Logger `validate:"required"`
}

// New returns an instance of PushService.
func New(c *config.Config, logger *zap.Logger) (*PushService, error) {
	ctx := context.TODO()
	log := logger.Named("pushnotifications")
	l := log.Named("New")

	// Init notification service for Android
	l.Debug("init notification service for Android")

	fcmCredentials, err := c.FcmCredentialsDecoded()
	if err != nil {
		return nil, err
	}

	firebaseClient, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON(fcmCredentials))
	if err != nil {
		return nil, err
	}

	fcmClient, err := firebaseClient.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	// Init notification service for iOS
	l.Debug("init notification service for iOS")

	apnsKeyRaw, err := c.ApnsKeyDecoded()
	if err != nil {
		return nil, err
	}

	apnsKey, err := token.AuthKeyFromBytes(apnsKeyRaw)
	if err != nil {
		return nil, err
	}

	apnsToken := &token.Token{
		TeamID:  c.ApnsTeamID,
		KeyID:   c.ApnsKeyID,
		AuthKey: apnsKey,
	}

	// To test push notifications for builds running from XCode, .Development() must be used
	// This is done when c.IsDevMode is set to true.
	// When c.IsDevMode is false, .Production() will be used.
	var apnsClient *apns2.Client

	if c.IsDevMode {
		logger.Info("registering apnsClient using the Development servers")

		apnsClient = apns2.NewTokenClient(apnsToken).Development()
	} else {
		logger.Info("registering apnsClient using the Production servers")

		apnsClient = apns2.NewTokenClient(apnsToken).Production()
	}

	l.Debug("both notification services have been created")

	// Create and return instance with validation
	cn := PushService{
		apnsClient:   apnsClient,
		apnsTopic:    c.ApnsTopic,
		fcmClient:    fcmClient,
		fcmChannelID: c.FcmChannelID,
		logger:       log,
	}
	validate := validator.New()

	err = validate.Struct(cn)
	if err != nil {
		return nil, err
	}

	l.Info("notification service creation and struct validation successful")

	return &cn, nil
}
