package util

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// Config stores all configuration of the application.
// The values are read by Viper from a config file or environment variables.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Load configuration from a file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Automatically read environment variables
	viper.AutomaticEnv()

	// Unmarshal configuration into the Config struct
	err = viper.Unmarshal(&config)
	return
}
