package certificate

import (
	"github.com/hammi85/swerve/src/db"
	"golang.org/x/crypto/acme/autocert"
)

// Manager wraps around autocert and injects a cache
type Manager struct {
	certCache   *persistentCertCache
	acmeManager *autocert.Manager
}

// persistentCertCache certificate cache
type persistentCertCache struct {
	autocert.Cache
	db *db.DynamoDB
}
