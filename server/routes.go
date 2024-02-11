package server

import "net/http"

func (s *server) routedMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.handleHealth())

	mux.HandleFunc("GET /greet/{name}", s.handleGreetGet())
	mux.HandleFunc("POST /greet", s.handleGreetPost())

	return mux
}
