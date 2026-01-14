package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// Config holds all configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	App      AppConfig
}

type ServerConfig struct {
	Port  int    `validate:"required,min=1,max=65535"`
	Stage string `validate:"required,oneof=dev prod staging"`
}

type DatabaseConfig struct {
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
	URI      string `validate:"required,uri"`
	MaxPool  int    `validate:"min=1,max=100"`
}

type AppConfig struct {
	IsProd   bool
	LogLevel string `validate:"oneof=debug info warn error"`
	GinMode  string `validate:"oneof=debug release test"`
}

var Envs *Config
var validate *validator.Validate

// LoadConfig initializes and validates configuration
func LoadConfig() error {
	// Load .env file (optional in production)
	_ = godotenv.Load()

	config := &Config{
		Server: ServerConfig{
			Port:  getEnvAsInt("PORT", 3000),
			Stage: getEnv("STAGE", "dev"),
		},
		Database: DatabaseConfig{
			User:     getEnv("MONGO_USER", ""),
			Password: getEnv("MONGO_PASSWORD", ""),
			Database: getEnv("MONGO_DATABASE", ""),
			URI:      getEnv("MONGO_URI", ""),
			MaxPool:  getEnvAsInt("MONGO_MAX_POOL", 10),
		},
		App: AppConfig{
			IsProd:   getEnv("STAGE", "dev") == "prod",
			LogLevel: getEnv("LOG_LEVEL", "info"),
			GinMode:  getEnv("GIN_MODE", "debug"),
		},
	}

	// Initialize validator
	validate = validator.New()

	// Register custom validators
	registerCustomValidators()

	// Validate configuration
	if err := validate.Struct(config); err != nil {
		return formatValidationError(err)
	}

	Envs = config
	return nil
}

// registerCustomValidators adds custom validation rules
func registerCustomValidators() {
	// Example: validate stage-specific requirements
	validate.RegisterValidation("secure_if_prod", func(fl validator.FieldLevel) bool {
		stage := fl.Parent().FieldByName("Stage").String()
		secure := fl.Field().Bool()
		if stage == "prod" {
			return secure
		}
		return true
	})
}

// formatValidationError provides clear error messages
func formatValidationError(err error) error {
	var errors []string

	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()

		switch tag {
		case "required":
			errors = append(errors, fmt.Sprintf("%s is required", field))
		case "min":
			errors = append(errors, fmt.Sprintf("%s must be at least %s", field, err.Param()))
		case "max":
			errors = append(errors, fmt.Sprintf("%s must be at most %s", field, err.Param()))
		case "oneof":
			errors = append(errors, fmt.Sprintf("%s must be one of: %s", field, err.Param()))
		case "uri":
			errors = append(errors, fmt.Sprintf("%s must be a valid URI", field))
		default:
			errors = append(errors, fmt.Sprintf("%s failed %s validation", field, tag))
		}
	}

	return fmt.Errorf("config validation errors:\n  - %s", strings.Join(errors, "\n  - "))
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// func getEnvAsBool(key string, defaultValue bool) bool {
// 	if value := os.Getenv(key); value != "" {
// 		if boolValue, err := strconv.ParseBool(value); err == nil {
// 			return boolValue
// 		}
// 	}
// 	return defaultValue
// }

// MustLoad panics if config fails to load (use in main.go)
func MustLoad() *Config {
	if err := LoadConfig(); err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}
	return Envs
}
