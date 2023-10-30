package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUploadService interface {
	// GetNewFilename uses uploadedFilename to preserve extension if possible
	GetNewFilename(uploadedFilename string) string
	GetFullFSPath(filename string) string
	GetURL(filename string) string
	FSRoot() string
}

type Server struct {
	r                *gin.Engine
	svc              IUploadService
	maxFileSizeBytes int64
}

func NewServer(storage IUploadService, maxFileSizeMB int64) http.Handler {
	maxFileSizeBytes := maxFileSizeMB << 20 // convert to bytes
	srv := &Server{gin.New(), storage, maxFileSizeBytes}
	srv.defineEndpoints()
	return srv
}

func (s *Server) defineEndpoints() {
	group := s.r.Group("/api/v1/uploads")
	group.POST("/", s.HandleUpload)
	group.Static("/", s.svc.FSRoot())
}

func (s *Server) HandleUpload(c *gin.Context) {
	// limits request size
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, s.maxFileSizeBytes)

	if c.ContentType() != "multipart/form-data" {
		WriteErrorResponse(c.Writer, ErrInvalidContentType)
		return
	}
	gotFile, err := c.FormFile("file")
	if err != nil {
		if maxBytesErr, ok := err.(*http.MaxBytesError); ok {
			err = ErrTooBigFile(maxBytesErr.Limit / 1024 / 1024)
		} else if errors.Is(err, http.ErrMissingFile) {
			err = ErrNoFileProvided
		}
		WriteErrorResponse(c.Writer, err)
		return
	}
	newFilename := s.svc.GetNewFilename(gotFile.Filename)
	err = c.SaveUploadedFile(gotFile, s.svc.GetFullFSPath(newFilename))
	if err != nil {
		WriteErrorResponse(c.Writer, err)
		return
	}
	c.String(http.StatusCreated, s.svc.GetURL(newFilename))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
