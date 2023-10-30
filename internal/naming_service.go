package internal

import (
	"github.com/google/uuid"
	"path/filepath"
)

type NamingService struct {
	cfg NamingConfig
}

func NewNamingService(cfg NamingConfig) *NamingService {
	return &NamingService{cfg}
}

func (n *NamingService) GetNewFilename(uploadedFilename string) string {
	randomName := uuid.New().String()
	fullName := randomName + filepath.Ext(uploadedFilename)
	return fullName
}

func (n *NamingService) GetFullFSPath(filename string) string {
	return filepath.Join(n.cfg.FSRoot, filename)
}

func (n *NamingService) GetURL(filename string) string {
	return n.cfg.WebURLRoot + filename
}
