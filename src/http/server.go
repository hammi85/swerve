package http

import (
	"log"
	nethttp "net/http"

	"github.com/hammi85/swerve/src/certificate"
)

// Listen to the http
func (s *Server) Listen() error {
	log.Printf("HTTP listening to %s", s.Listener)
	return s.Server.ListenAndServe()
}

// handle normal redirect request on http
func (s *Server) handleRedirect(w nethttp.ResponseWriter, r *nethttp.Request) {
	//hostHeader := r.Header.Get("Host")
}

// Handler for requests
func (s *Server) Handler() nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		s.CertManager.Serve(nethttp.HandlerFunc(s.handleRedirect), w, r)
	})
}

// NewServer creates a new instance
func NewServer(listener string, certManager *certificate.Manager) *Server {
	server := &Server{
		Listener:    listener,
		CertManager: certManager,
	}

	server.Server = &nethttp.Server{
		Addr:    listener,
		Handler: server.Handler(),
	}

	return server
}
