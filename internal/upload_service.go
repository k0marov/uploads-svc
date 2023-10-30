package internal

import (
	"github.com/google/uuid"
	"path/filepath"
)

type UploadService struct {
	cfg NamingConfig
}

func NewUploadService(cfg NamingConfig) *UploadService {
	return &UploadService{cfg}
}

func (n *UploadService) GetNewFilename(uploadedFilename string) string {
	randomName := uuid.New().String()
	fullName := randomName + filepath.Ext(uploadedFilename)
	return fullName
}

func (n *UploadService) GetFullFSPath(filename string) string {
	return filepath.Join(n.cfg.FSRoot, filename)
}

func (n *UploadService) GetURL(filename string) string {
	return n.cfg.WebURLRoot + filename
}

func (n *UploadService) FSRoot() string {
	return n.cfg.FSRoot
}
