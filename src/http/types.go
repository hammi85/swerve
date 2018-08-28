package http

import (
	nethttp "net/http"
)

type fallbackHandler struct{}

// Server model
type Server struct {
	Server   *nethttp.Server
	Listener string
}
