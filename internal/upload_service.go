package internal

import (
	"github.com/google/uuid"
	"path/filepath"
)

type UploadService struct {
	cfg           NamingConfig
	maxFileSizeMB int64
}

func NewUploadService(cfg NamingConfig, maxFileSizeMB int64) *UploadService {
	return &UploadService{cfg, maxFileSizeMB}
}

func (u *UploadService) GetNewFilename(uploadedFilename string) string {
	randomName := uuid.New().String()
	fullName := randomName + filepath.Ext(uploadedFilename)
	return fullName
}

func (u *UploadService) GetFullFSPath(filename string) string {
	return filepath.Join(u.cfg.FSRoot, filename)
}

func (u *UploadService) GetURL(filename string) string {
	return u.cfg.WebURLRoot + filename
}

func (u *UploadService) FSRoot() string {
	return u.cfg.FSRoot
}

func (u *UploadService) MaxFileSizeBytes() int64 {
	return u.maxFileSizeMB * 1024 * 1024
}

func (u *UploadService) MakeErrTooBigFile() error {
	return ErrTooBigFile(u.maxFileSizeMB)
}
