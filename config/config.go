package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func BuildConfiguration() Configurations {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var configuration = Configurations{}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration

}

type Configurations struct {
	Server Server
}

type Server struct {
	Port         string
	WriteTimeOut time.Duration
	ReadTimeOut  time.Duration
}
