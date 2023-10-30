package internal

import (
	"log"
	"net/http"
)

func InitializeAndStart(cfg AppConfig) {
	svc := NewUploadService(cfg.Naming, cfg.MaxFileSizeMB)
	srv := NewServer(svc)
	log.Print(http.ListenAndServe(cfg.HTTPServer.Host, srv))
}
