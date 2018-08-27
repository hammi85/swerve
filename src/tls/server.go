package tls

import (
	"crypto/tls"
	"log"
	"net/http"
)

// Listen to the https
func (s *TLSServer) Listen() error {
	return s.Server.ListenAndServeTLS("", "")
}

// GetCertificate bypass
func (s *TLSServer) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {

	log.Printf("GetCertificate %#v", hello)

	return nil, nil
}

// NewTLSServer creates a new instance
func NewTLSServer(listener string) *TLSServer {
	tlsServer := &TLSServer{
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
