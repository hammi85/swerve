package certificate

import (
	"golang.org/x/crypto/acme/autocert"
)

// Manager wraps around autocert and injects a cache
type Manager struct {
	acmeManager *autocert.Manager
}
