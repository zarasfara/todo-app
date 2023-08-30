package server

import (
	"net/http"

	"github.com/zarasfara/pet-adoption-platoform/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer:  &http.Server{
			Addr:    ":" + cfg.HTTP.Port,
			Handler: handler,
			MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		},
	}

}

func (s Server) Run() error {
	return s.httpServer.ListenAndServe()
}
