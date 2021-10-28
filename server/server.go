package server

import (
	"log"
	"net/http"

	"github.com/bazeeko/mechta-test-task/route"
)

type Server struct {
	*http.Server
}

func NewServer(addr string, h *route.Handler) *Server {
	return &Server{
		&http.Server{
			Addr:    addr,
			Handler: h.Router,
		},
	}
}

func (s *Server) Run() error {
	log.Println("API is running at", s.Addr)
	return s.ListenAndServe()
}
