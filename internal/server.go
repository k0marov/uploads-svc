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
	MaxFileSizeBytes() int64
	MakeErrTooBigFile() error
}

type Server struct {
	r   *gin.Engine
	svc IUploadService
}

func NewServer(svc IUploadService) http.Handler {
	srv := &Server{gin.New(), svc}
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
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, s.svc.MaxFileSizeBytes())

	if c.ContentType() != "multipart/form-data" {
		WriteErrorResponse(c.Writer, ErrInvalidContentType)
		return
	}
	gotFile, err := c.FormFile("file")
	if err != nil {
		if _, ok := err.(*http.MaxBytesError); ok {
			err = s.svc.MakeErrTooBigFile()
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
