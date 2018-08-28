package certificate

import (
	"golang.org/x/crypto/acme/autocert"
)

// NewManager creates a new instance
func NewManager() *Manager {
	return &Manager{
		acmeManager: &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("yourdomainname.com"),
			Cache:      autocert.DirCache("cache-path"),
		},
	}
}
