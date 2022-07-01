package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

const envsPath = "./configs"

type CustomerDataPlatform struct {
	Host                      string `mapstructure:"CUSTOMER_DATA_PLATFORM_HOST"`
	ApiKey                    string `mapstructure:"CUSTOMER_DATA_PLATFORM_API_KEY"`
	SegmentWritekey           string `mapstructure:"SPREADSHEET_SEGMENT_WRITE_KEY"`
	GoogleServiceAccountEmail string `mapstructure:"GOOGLE_SERVICE_ACCOUNT_EMAIL"`
	GoogleServiceAccountId    string `mapstructure:"GOOGLE_SERVICE_ACCOUNT_PRIVATE_KEY_ID"`
}

type Config struct {
	CustomerDataPlatform `mapstructure:",squash"`
}

func NewConfig() *Config {
	config, err := LoadEnvConfig()
	if err != nil {
		log.Fatalln(err)
	}
	return &config
}

func LoadEnvConfig() (config Config, err error) {
	env := os.Getenv("ENV")
	viper.AddConfigPath(envsPath)

	if env != "" {
		viper.SetConfigName(env)
		viper.SetConfigType("env")
	} else {
		viper.SetConfigName("envs")
		viper.SetConfigType("yml")
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
