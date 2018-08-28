package https

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/hammi85/swerve/src/certificate"
)

// Listen to the https
func (s *Server) Listen() error {
	log.Printf("HTTPS listening to %s", s.Listener)
	return s.Server.ListenAndServeTLS("", "")
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
	}

	return server
}
