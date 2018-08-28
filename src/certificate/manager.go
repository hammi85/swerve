package certificate

import (
	"crypto/tls"
	"net/http"

	"github.com/hammi85/swerve/src/db"
	"golang.org/x/crypto/acme/autocert"
)

// NewManager creates a new instance
func NewManager(d *db.DynamoDB) *Manager {
	return &Manager{
		acmeManager: &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("yourdomainname.com"),
			Cache:      newPersistentCertCache(d),
		},
	}
}

// GetCertificate wrapper for the cert getter
func (m *Manager) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return m.acmeManager.GetCertificate(hello)
}

// Serve http.Handler bridge
func (m *Manager) Serve(fallback http.Handler, w http.ResponseWriter, r *http.Request) {
	m.acmeManager.HTTPHandler(fallback).ServeHTTP(w, r)
}
