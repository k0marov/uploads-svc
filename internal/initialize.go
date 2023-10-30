package internal

import (
	"log"
	"net/http"
)

func InitializeAndStart(cfg AppConfig) {
	namer := NewNamingService(cfg.Naming)
	srv := NewServer(namer)
	log.Print(http.ListenAndServe(cfg.HTTPServer.Host, srv))
}
