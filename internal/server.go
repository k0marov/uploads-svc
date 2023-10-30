package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type INamingService interface {
	// GetNewFilename uses uploadedFilename to preserve extension if possible
	GetNewFilename(uploadedFilename string) string
	GetFullFSPath(filename string) string
	GetURL(filename string) string
	FSRoot() string
}

type Server struct {
	r     *gin.Engine
	namer INamingService
}

func NewServer(storage INamingService) http.Handler {
	srv := &Server{gin.New(), storage}
	//srv.r.MaxMultipartMemory = 8 << 20 // 8 MiB
	srv.defineEndpoints()
	return srv
}

func (s *Server) defineEndpoints() {
	group := s.r.Group("/api/v1/uploads")
	group.POST("/", s.HandleUpload)
	group.Static("/", s.namer.FSRoot())
}

func (s *Server) HandleUpload(c *gin.Context) {
	if c.ContentType() != "multipart/form-data" {
		WriteErrorResponse(c.Writer, ErrInvalidContentType)
		return
	}
	gotFile, err := c.FormFile("file")
	if err != nil {
		if errors.Is(err, multipart.ErrMessageTooLarge) {
			err = ErrTooBigFile
		} else if errors.Is(err, http.ErrMissingFile) {
			err = ErrNoFileProvided
		}
		WriteErrorResponse(c.Writer, err)
		return
	}
	newFilename := s.namer.GetNewFilename(gotFile.Filename)
	err = c.SaveUploadedFile(gotFile, s.namer.GetFullFSPath(newFilename))
	if err != nil {
		WriteErrorResponse(c.Writer, err)
		return
	}
	c.String(http.StatusCreated, s.namer.GetURL(newFilename))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
