package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config is the application configuration
type Config struct {
	// Application configuration
	IsDevMode  bool   `mapstructure:"development"` // Not marked required as it would fail when `false` is given as value TODO: Add to docs
	ServerPort int    `mapstructure:"api_port" validate:"required"`
	JWTSecret  string `mapstructure:"jwt_secret" validate:"required"`

	// Database configuration
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
}

// LoadConfig reads in config file and ENV variables if set.
func LoadConfig() (Config, error) {
	v := viper.New()

	// For now we only support a .env file in the current directory or env variables.
	// Later on, using AddConfigPath and SetConfigName we can allow
	// other methods of application configuration and utilize the
	// full power of Viper for creating a true 12-factor application
	v.SetConfigFile(".env")
	v.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", v.ConfigFileUsed())
	}

	// Marshal Viper config to struct and return
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	// Validate config
	validate := validator.New()
	err := validate.Struct(config)
	return config, err
}

// DatabaseConnectionString returns database connection string for connection to PostgreSQL
func (c Config) DatabaseConnectionString() string {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		c.DatabaseHost, c.DatabasePort, c.DatabaseUser, c.DatabasePass, c.DatabaseName)
	if !c.DatabaseSSL {
		connStr += " sslmode=disable"
	}
	return connStr
}
