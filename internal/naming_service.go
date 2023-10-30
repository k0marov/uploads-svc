package internal

import (
	"github.com/google/uuid"
	"path/filepath"
)

type NamingService struct {
}

func NewNamingService() *NamingService {
	return &NamingService{}
}

func (n NamingService) GetNewFilename(uploadedFilename string) string {
	randomName := uuid.New().String()
	fullName := randomName + filepath.Ext(uploadedFilename)
	return fullName
}

func (n NamingService) GetFullFSPath(filename string) string {
	//TODO implement me
	panic("implement me")
}

func (n NamingService) GetURL(filename string) string {
	//TODO implement me
	panic("implement me")
}
