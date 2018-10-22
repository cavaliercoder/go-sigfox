package sigfox

import (
	"net/http"
)

// A Server holds server data.
type Server struct{}

// ListenAndServe listens on the TCP network address and then
// handle requests on incoming connections.
func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}

// ListenAndServeTLS listens on the TCP network address and then
// handle requests on incoming TLS connections.
func (s *Server) ListenAndServeTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, nil)
}

// HandleUplink registers the handler for the given pattern.
func (s *Server) HandleUplink(pattern string, handler UplinkHandler) {
	http.Handle(pattern, handler)
}
