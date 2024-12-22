package server

import (
	"net/http"

	"calc_service/internal/handler"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Run() error {
	http.HandleFunc("/api/v1/calculate", handler.CalculateHandler)
	return http.ListenAndServe(s.addr, nil)
}
