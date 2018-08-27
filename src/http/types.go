package http

import (
	nethttp "net/http"
)

// Server model
type Server struct {
	Server   *nethttp.Server
	Listener string
}
