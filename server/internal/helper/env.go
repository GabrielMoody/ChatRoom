package helper

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func LoadEnv() *viper.Viper {
	v := viper.New()

	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AutomaticEnv()

	err := v.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return v
}
