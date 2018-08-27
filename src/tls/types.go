package tls

import "net/http"

// Server model
type Server struct {
	Server   *http.Server
	Listener string
}
