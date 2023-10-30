package internal

import (
	"log"
	"net/http"
)

func InitializeAndStart(cfg AppConfig) {
	svc := NewUploadService(cfg.Naming)
	srv := NewServer(svc, cfg.MaxFileSizeMB)
	log.Print(http.ListenAndServe(cfg.HTTPServer.Host, srv))
}
