package internal

import (
	"log"
	"net/http"
)

func InitializeAndStart(cfg AppConfig) {
	srv := NewServer()
	log.Print(http.ListenAndServe(cfg.HTTPServer.Host, srv))
}
