package sigfox

import "net/http"

// callbackHandler is a request Handler that invokes a user defined handler for
// a specific SIGFOX callback type.
type callbackHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// callbackHandlerFunc is a function signature that invokes a user defined
// handler for a SIGFOX callback base type.
type callbackHandlerFunc func(*callback)

func (f callbackHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := http.StatusInternalServerError

	// parse request body
	cb := &callback{}
	if s, err := parseCallback(r, cb); err != nil {
		// TODO: make error available to user
		dprintf("Error parsing callbackL %v", err)
		status = s
	} else {
		status = http.StatusNoContent
	}

	// respond to sigfox server
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)

	// invoke user handler async
	go f(cb)
}
