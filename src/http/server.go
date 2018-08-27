package http

import (
	"log"
	nethttp "net/http"
)

// Listen to the http
func (s *Server) Listen() error {
	log.Printf("HTTP listening to %s", s.Listener)
	return s.Server.ListenAndServe()
}

/*
func (s *Server) Handler() nethttp.Handler {

	return func(resp nethttp.ResponseWriter, req *nethttp.Request) {

	}

}
*/

// NewServer creates a new instance
func NewServer(listener string) *Server {
	server := &Server{
		Listener: listener,
	}

	server.Server = &nethttp.Server{
		Addr: listener}

	return server
}
