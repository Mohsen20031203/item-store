package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	TokenSymmetricKey        string    `mapstructure:"JWT_SECRET_KEY"`
	RefreshTokenSymmetricKey string    `mapstructure:"JWT_REFRESH_SECRET_KEY"`
	SwaggerUser              string    `mapstructure:"SWAGGER_USER"`
	SwaggerPass              string    `mapstructure:"SWAGGER_PASS"`
	LvlConfig                LvlConfig `mapstructure:"LVL_CONFIG"`
	ApiPort                  string    `mapstructure:"API_PORT"`
}

type LvlConfig struct {
	Path string `mapstructure:"LVL_PATH"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	swaggerUser := config.SwaggerUser
	swaggerPass := config.SwaggerPass
	if swaggerUser == "" || swaggerPass == "" {
		log.Fatal("Environment variables SWAGGER_USER and SWAGGER_PASS must be set")
	}
	return
}
