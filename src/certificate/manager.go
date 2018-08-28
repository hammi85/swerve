package certificate

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"

	"github.com/hammi85/swerve/src/db"
	"golang.org/x/crypto/acme/autocert"
)

var (
	errHostNotConfigured = errors.New("acme/autocert: host not configured")
)

// NewManager creates a new instance
func NewManager(d *db.DynamoDB) *Manager {
	manager := &Manager{}

	manager.acmeManager = &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: manager.allowHostPolicy,
		Cache:      newPersistentCertCache(d),
	}

	return manager
}

// allowHostPolicy decides which host shall pass
func (m *Manager) allowHostPolicy(_ context.Context, host string) error {
	if !m.certCache.IsDomainAcceptable(host) {
		return errHostNotConfigured
	}

	return nil
}

// GetCertificate wrapper for the cert getter
func (m *Manager) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return m.acmeManager.GetCertificate(hello)
}

// Serve http.Handler bridge
func (m *Manager) Serve(fallback http.Handler, w http.ResponseWriter, r *http.Request) {
	m.acmeManager.HTTPHandler(fallback).ServeHTTP(w, r)
}
