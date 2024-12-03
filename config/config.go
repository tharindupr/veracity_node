package config

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
)

// Config structure for application settings
type Config struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`

	CACertPath    string `mapstructure:"ca_cert_path"`    // Path to the CA certificate file
	CACertContent string `mapstructure:"ca_cert_content"` // Content of the CA certificate file
}

// LoadConfig loads the configuration from a YAML file and environment variables
func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigName("config")              // Name of the config file (without extension)
	viper.SetConfigType("yaml")                // Specify file type
	viper.AddConfigPath(configPath)            // Path to look for the config file
	viper.AddConfigPath(".")                   // Fallback to current directory
	viper.AutomaticEnv()                       // Enable environment variable overrides
	viper.SetEnvPrefix("APP")                  // Prefix for env vars
	viper.BindEnv("ca_cert_path", "APP_CA_CERT_PATH") // Bind specific environment variable

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Warning: Config file not found: %v\n", err)
	}

	// Unmarshal into the Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	// Read the content of the CA certificate file
	if config.CACertPath != "" {
		certContent, err := os.ReadFile(config.CACertPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read CA certificate file: %w", err)
		}
		config.CACertContent = string(certContent)
	}

	return &config, nil
}
