package tls

import (
	"crypto/tls"
	"log"
	"net/http"
)

// Listen to the https
func (s *Server) Listen() error {
	log.Printf("HTTPS listening to %s", s.Listener)
	return s.Server.ListenAndServeTLS("", "")
}

// GetCertificate bypass
func (s *Server) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	log.Printf("GetCertificate %#v", hello)
	// extreme magic happens here
	return nil, nil
}

// NewServer creates a new instance
func NewServer(listener string) *Server {
	tlsServer := &Server{
		Listener: listener,
	}

	tlsServer.Server = &http.Server{
		Addr: listener,
		TLSConfig: &tls.Config{
			GetCertificate: tlsServer.GetCertificate,
		},
	}

	return tlsServer
}
