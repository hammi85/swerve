package tls

import "net/http"

type TLSServer struct {
	Server   *http.Server
	Listener string
}
