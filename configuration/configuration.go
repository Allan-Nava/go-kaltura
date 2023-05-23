package configuration

import (
	"github.com/caarlos0/env/v6"
)


type Configuration struct {
	IsDebug    bool `env:"IS_DEBUG"`
	AdminSecret string `env:"ADMIN_SECRET"`
	UniqueUserId string `env:"UNIQUE_USER_ID"`
	BaseUrl    string `env:"BASE_URL"`
	ApiVersion string `env:"ApiVersion"`
	ClientTag 	string 
}


func GetConfiguration() *Configuration {
	configuration := Configuration{}
	err := env.Parse(&configuration)
	if err != nil {
		panic("failed to read configuration")
	}
	configuration.ClientTag = "v0.1.0"
	//
	return &configuration
}