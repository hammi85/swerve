package http

import (
	nethttp "net/http"

	"github.com/hammi85/swerve/src/certificate"
	"github.com/hammi85/swerve/src/log"
)

// Listen to the http
func (s *Server) Listen() error {
	log.Infof("HTTP listening to %s", s.Listener)
	return s.Server.ListenAndServe()
}

// handle normal redirect request on http
func (s *Server) handleRedirect(w nethttp.ResponseWriter, r *nethttp.Request) {
	hostHeader := r.Header.Get("Host")
	domain, err := s.certManager.GetDomain(hostHeader)

	if domain != nil && err == nil {
		nethttp.Redirect(w, r, domain.Redirect, domain.RedirectCode)
		return
	}

	nethttp.NotFound(w, r)
}

// Handler for requests
func (s *Server) Handler() nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		s.certManager.Serve(nethttp.HandlerFunc(s.handleRedirect), w, r)
	})
}

// NewServer creates a new instance
func NewServer(listener string, certManager *certificate.Manager) *Server {
	server := &Server{
		Listener:    listener,
		certManager: certManager,
	}

	server.Server = &nethttp.Server{
		Addr:    listener,
		Handler: server.Handler(),
	}

	return server
}
