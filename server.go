package sigfox

import (
	"net/http"
)

// A Server defines parameters for running an HTTP server.
// The zero value for Server is a valid configuration.
type Server http.Server

func (c *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}

func (c *Server) ListenAndServerTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, nil)
}

func (c *Server) HandleUplink(pattern string, handler UplinkHandler) {
	http.Handle(pattern, handler)
}
