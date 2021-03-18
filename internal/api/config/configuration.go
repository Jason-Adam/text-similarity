package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds pointer to Configuration struct
var Config *Configuration

// Configuration holds all configurations (currently just server)
type Configuration struct {
	Server ServerConfiguration
}

// ServerConfiguration holds all variables for configuring the server
type ServerConfiguration struct {
	Port string
	Mode string
}

// Setup reads the config yaml and hydrates the Configuration struct
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file, %s", err)
	}

	if uErr := viper.Unmarshal(&configuration); uErr != nil {
		log.Fatalf("unable to decode into struct, %v", uErr)
	}

	Config = configuration
}

// GetConfig returns the pointer to the hydrated Configuration struct
func GetConfig() *Configuration {
	return Config
}
