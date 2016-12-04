package sigfox

import (
	"log"
	"net/http"
)

type Server struct {
	logger *log.Logger // TODO: implement access logs
}

func (c *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}

func (c *Server) ListenAndServerTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, nil)
}

func (c *Server) HandleUplink(pattern string, handler UplinkHandler) {
	http.Handle(pattern, handler)
}
