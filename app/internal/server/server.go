package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, handler http.Handler) *Server {
	server := &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: handler,
		},
	}
	return server
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
