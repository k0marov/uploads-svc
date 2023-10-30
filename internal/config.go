package internal

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type HTTPServerConfig struct {
	Host string `default:"127.0.0.1:8002"`
}

type AppConfig struct {
	HTTPServer HTTPServerConfig
}

func ReadConfigFromEnv() AppConfig {
	var cfg AppConfig
	err := envconfig.Process("upload", &cfg)
	if err != nil {
		log.Panicf("while parsing app config from env: %w", err)
	}
	return cfg
}
