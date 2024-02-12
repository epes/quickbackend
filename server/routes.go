package server

import "net/http"

func (s *server) routedMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.handleHealth())

	mux.HandleFunc("GET /api/greet/{name}", s.handleGreetGet())
	mux.HandleFunc("POST /api/greet", s.handleGreetPost())

	return mux
}
