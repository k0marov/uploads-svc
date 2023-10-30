package internal

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type HTTPServerConfig struct {
	Host string `default:"127.0.0.1:8002"`
}

type NamingConfig struct {
	FSRoot     string `required:"true"`
	WebURLRoot string `required:"true"`
}

type AppConfig struct {
	HTTPServer HTTPServerConfig
	Naming     NamingConfig
}

func ReadConfigFromEnv() AppConfig {
	var cfg AppConfig
	err := envconfig.Process("uploads", &cfg)
	if err != nil {
		log.Panicf("while parsing app config from env: %v", err)
	}
	return cfg
}
