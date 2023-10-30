package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	r *gin.Engine
}

func NewServer() http.Handler {
	srv := &Server{gin.New()}
	srv.defineEndpoints()
	return srv
}

func (s *Server) defineEndpoints() {
	group := s.r.Group("/api/v1/upload")
	group.POST("/", s.HandleUpload)
	group.GET("/:id", s.GetFileById)
}

func (s *Server) HandleUpload(c *gin.Context) {
	// TODO: implement HandleUpload
}

func (s *Server) GetFileById(c *gin.Context) {
	// TODO: implement GetFileById
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
