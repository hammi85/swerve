package https

import (
	"crypto/tls"
	"net/http"

	"github.com/hammi85/swerve/src/certificate"
	"github.com/hammi85/swerve/src/log"
)

// Listen to the https
func (s *Server) Listen() error {
	log.Infof("HTTPS listening to %s", s.Listener)
	return s.Server.ListenAndServeTLS("", "")
}

// RedirectHandler redirects the request to the domain redirect location
func (s *Server) RedirectHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostHeader := r.Header.Get("Host")
		domain, err := s.certManager.GetDomain(hostHeader)

		if domain != nil && err == nil {
			http.Redirect(w, r, domain.Redirect, domain.RedirectCode)
			return
		}

		http.NotFound(w, r)
	})
}

// NewServer creates a new instance
func NewServer(listener string, certManager *certificate.Manager) *Server {
	server := &Server{
		certManager: certManager,
		Listener:    listener,
	}

	server.Server = &http.Server{
		Addr: listener,
		TLSConfig: &tls.Config{
			GetCertificate: server.certManager.GetCertificate,
		},
		Handler: server.RedirectHandler(),
	}

	return server
}
