package certificate

import (
	"context"

	"github.com/hammi85/swerve/src/db"
	"golang.org/x/crypto/acme/autocert"
)

// newPersistentCertCache creates a new persistent cache based on dynamo db
func newPersistentCertCache(d *db.DynamoDB) *persistentCertCache {
	c := &persistentCertCache{
		db: d,
	}

	return c
}

// Get cert by domain name
func (c *persistentCertCache) Get(ctx context.Context, key string) ([]byte, error) {
	var (
		data   []byte
		done   = make(chan struct{})
		err    error
		domain *db.Domain
	)

	// fetch domain by name. Look for cert data. Copy and exit
	go func() {
		defer close(done)
		if domain, err = c.db.FetchByDomain(key); domain != nil && len(domain.Certificate) > 0 {
			data = make([]byte, len(domain.Certificate))
			copy(data, domain.Certificate)
		}
	}()

	// handle context timeouts and errors
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
	}

	// if we got no data then we missed the cache
	if data == nil || err != nil {
		return nil, autocert.ErrCacheMiss
	}

	// return cert data
	return data, nil
}

// Put a cert to the cache
func (c *persistentCertCache) Put(ctx context.Context, key string, data []byte) error {
	var (
		done = make(chan struct{})
		err  error
	)

	go func() {
		defer close(done)
		err = c.db.UpdateCertificateData(key, data)
	}()

	// handle context timeouts and errors
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	}

	return err
}

// Delete a domain from
func (c *persistentCertCache) Delete(ctx context.Context, key string) error {
	var (
		done = make(chan struct{})
		err  error
	)

	go func() {
		defer close(done)
		_, err = c.db.DeleteByDomain(key)
	}()

	// handle context timeouts and errors
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	}

	return err
}
