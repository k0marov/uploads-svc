package internal

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type HTTPServerConfig struct {
	Host string `default:":8080"`
}

type NamingConfig struct {
	FSRoot     string `required:"true"`
	WebURLRoot string `required:"true"`
}

type AppConfig struct {
	HTTPServer    HTTPServerConfig
	Naming        NamingConfig
	MaxFileSizeMB int64 `default:"10"`
}

func ReadConfigFromEnv() AppConfig {
	var cfg AppConfig
	err := envconfig.Process("uploads", &cfg)
	if err != nil {
		log.Panicf("while parsing app config from env: %v", err)
	}
	return cfg
}
