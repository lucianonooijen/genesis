package config

import (
	b64 "encoding/base64"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config is the application configuration.
type Config struct {
	// Application configuration
	IsDevMode            bool   `mapstructure:"development"` // Not marked required as it would fail when `false` is given as value
	ServerHostname       string `mapstructure:"api_hostname" validate:"required"`
	ServerHostnameCanary string `mapstructure:"api_hostname_canary"`
	ServerPort           int    `mapstructure:"api_port" validate:"required"`
	JWTSecret            string `mapstructure:"jwt_secret" validate:"required"`

	// DBConn configuration
	DatabaseHost string `mapstructure:"db_host" validate:"required"`
	DatabasePort int    `mapstructure:"db_port" validate:"required"`
	DatabaseName string `mapstructure:"db_name" validate:"required"`
	DatabaseUser string `mapstructure:"db_user" validate:"required"`
	DatabasePass string `mapstructure:"db_pass" validate:"required"`
	DatabaseSSL  bool   `mapstructure:"db_ssl"` // Not marked required as it would fail when `false` is given as value

	// Sentry configuration
	SentryDSN         string `mapstructure:"sentry_dsn" validate:"required"`
	SentryEnvironment string `mapstructure:"sentry_environment" validate:"required"`

	// Email configuration
	EmailSenderName  string `mapstructure:"email_sender_name" validate:"required"`
	EmailSenderEmail string `mapstructure:"email_sender_email" validate:"required"`
	SendinblueAPIKey string `mapstructure:"sendinblue_api_key" validate:"required"`

	// Push notifications and data for connecting to Apple and Google services,
	// APNS = Apple, FCM = Google
	ApnsTopic            string `mapstructure:"apns_topic" validate:"required"` // Bundle identifyer
	ApnsTeamID           string `mapstructure:"apns_team_id" validate:"required"`
	ApnsKeyID            string `mapstructure:"apns_key_id" validate:"required"`
	ApnsKeyBase64        string `mapstructure:"apns_key_base64" validate:"required"`
	FcmChannelID         string `mapstructure:"fcm_channel_id" validate:"required"`
	FcmCredentialsBase64 string `mapstructure:"fcm_credentials_base64" validate:"required"`
}

// LoadConfig reads in config file and ENV variables if set.
func LoadConfig() (Config, error) {
	v := viper.New()

	var config Config

	// Loop over the Config type and add all mapstructure field values to v.BindEnv
	// so that Viper will read them, this means you can define application config
	// both in the config.yml file or in the environment variables
	cf := reflect.TypeOf(config)
	for i := 0; i < cf.NumField(); i++ {
		field := cf.Field(i)
		tagValue := field.Tag.Get("mapstructure")

		if err := v.BindEnv(tagValue); err != nil {
			return config, err
		}
	}

	// For now we only support a config.yml file in the current directory or env variables.
	// Later on, using AddConfigPath and SetConfigName we can allow
	// other methods of application configuration and utilize the
	// full power of Viper for creating a true 12-factor application
	v.SetConfigFile("config.yml")
	v.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", v.ConfigFileUsed())
	}

	// Marshal Viper config to struct and return
	if err := v.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	// Validate config
	validate := validator.New()
	err := validate.Struct(config)

	return config, err
}

// DatabaseConnectionString returns database connection string for connection to PostgreSQL.
func (c *Config) DatabaseConnectionString() string {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		c.DatabaseHost, c.DatabasePort, c.DatabaseUser, c.DatabasePass, c.DatabaseName)

	if !c.DatabaseSSL {
		connStr += " sslmode=disable"
	}

	return connStr
}

// ApnsKeyDecoded decodes ApnsKeyBase64.
func (c *Config) ApnsKeyDecoded() ([]byte, error) {
	return b64.StdEncoding.DecodeString(c.ApnsKeyBase64)
}

// FcmCredentialsDecoded decodes FcmCredentialsBase64.
func (c *Config) FcmCredentialsDecoded() ([]byte, error) {
	return b64.StdEncoding.DecodeString(c.FcmCredentialsBase64)
}
