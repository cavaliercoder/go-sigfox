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
	status := http.StatusNoContent
	defer func() {
		if err := recover(); err != nil {
			// TODO: handle me
			status = http.StatusInternalServerError
		}

		if r.Body != nil {
			r.Body.Close()
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(status)
	}()

	// parse request
	cb := &callback{}
	if err := parseCallback(r, cb); err != nil {
		switch err.(type) {
		case ContentTypeError, RequestBodyError:
			status = http.StatusBadRequest

		case MethodError:
			status = http.StatusMethodNotAllowed

		default:
			// TODO: handler unexpected parser errors
			panic(err)
		}

		return
	}

	// invoke user handler
	f(cb)
}
