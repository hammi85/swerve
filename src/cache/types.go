package cache

// Entry model
type Entry struct {
	Domain string
	Cert   []byte
}

// Cache interface
type Cache interface {
	GetByDomain(string) (*Entry, error)
	Set(*Entry)
}
