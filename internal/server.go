package internal

import (
	"net/http"
)
import "github.com/go-chi/chi/v5"

type Server struct {
	r chi.Router
}

func NewServer() http.Handler {
	srv := &Server{chi.NewRouter()}
	srv.defineEndpoints()
	return srv
}

func (s *Server) defineEndpoints() {
	s.r.Route("/api/v1/upload", func(r chi.Router) {
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
