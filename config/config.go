package config

import (
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	// TemplateExt stores the extension used for the template files
	TemplateExt = ".gohtml"

	// StaticDir stores the name of the directory that will serve static files
	StaticDir = "static"

	// StaticPrefix stores the URL prefix used when serving static files
	StaticPrefix = "files"
)

type environment string

const (
	// EnvLocal represents the local environment
	EnvLocal environment = "local"

	// EnvTest represents the test environment
	EnvTest environment = "test"

	// EnvDevelop represents the development environment
	EnvDevelop environment = "dev"

	// EnvStaging represents the staging environment
	EnvStaging environment = "staging"

	// EnvQA represents the qa environment
	EnvQA environment = "qa"

	// EnvProduction represents the production environment
	EnvProduction environment = "prod"
)

// SwitchEnvironment sets the environment variable used to dictate which environment the application is
// currently running in.
// This must be called prior to loading the configuration in order for it to take effect.
func SwitchEnvironment(env environment) {
	if err := os.Setenv("PAGODA_APP_ENVIRONMENT", string(env)); err != nil {
		panic(err)
	}
}

type (
	// Config stores complete configuration
	Config struct {
		HTTP     HTTPConfig
		App      AppConfig
		Cache    CacheConfig
		Database DatabaseConfig
		Mail     MailConfig
		Storage  StorageConfig
	}

	// HTTPConfig stores HTTP configuration
	HTTPConfig struct {
		Hostname     string
		Port         uint16
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		IdleTimeout  time.Duration
		TLS          struct {
			Enabled     bool
			Certificate string
			Key         string
		}
	}

	// AppConfig stores application configuration
	AppConfig struct {
		Name          string
		Environment   environment
		EncryptionKey string
		Timeout       time.Duration
		PasswordToken struct {
			Expiration time.Duration
			Length     int
		}
		EmailVerificationTokenExpiration time.Duration
	}

	// CacheConfig stores the cache configuration
	CacheConfig struct {
		Hostname     string
		Port         uint16
		Password     string
		Database     int
		TestDatabase int
		Expiration   struct {
			StaticFile time.Duration
			Page       time.Duration
		}
	}

	// DatabaseConfig stores the database configuration
	DatabaseConfig struct {
		Hostname     string
		Port         uint16
		User         string
		Password     string
		Database     string
		TestDatabase string
	}

	// MailConfig stores the mail configuration
	MailConfig struct {
		Hostname     string
		HttpPort     uint16
		SmtpPort     uint16
		User         string
		Password     string
		FromAddress  string
		ResendAPIKey string
	}

	StorageConfig struct {
		AppBucketName             string
		StaticFilesBucketName     string
		S3Endpoint                string
		S3AccessKey               string
		S3SecretKey               string
		S3UseSSL                  bool
		ProfilePhotoMaxFileSizeMB int64
		PhotosMaxFileSizeMB       int64
	}
)

// GetConfig loads and returns configuration
func GetConfig() (Config, error) {
	var c Config

	// Common config loading
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")

	// Check the environment variable PAGODA_APP_ENVIRONMENT
	env := os.Getenv("PAGODA_APP_ENVIRONMENT")
	if env == string(EnvProduction) {
		// Load env variables for production
		viper.SetEnvPrefix("pagoda")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	}

	// Load the config file
	if err := viper.ReadInConfig(); err != nil {
		return c, err
	}

	// Unmarshal the config
	if err := viper.Unmarshal(&c); err != nil {
		return c, err
	}

	return c, nil
}
