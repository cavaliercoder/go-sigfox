package callback

import (
	"net/http"
)

type Handler interface {
	HandleCallback(Callback) error
}

type HandlerFunc func(Callback) error

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// close request stream before exiting
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()

	// parse request
	if cb, err := ParseUplinkCallback(r); err != nil {
		// return 400 for bad request content types
		if _, ok := err.(ContentTypeError); ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// return 405 for invalid request methods
		if _, ok := err.(MethodError); ok {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// TODO: handler unexpected parser errors
		panic(err)
	} else {
		if err := f(cb); err != nil {
			// TODO: handle handler errors
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusNoContent)
}

func HTTPHandlerFunc(f HandlerFunc) http.HandlerFunc {
	return HandlerFunc(f).ServeHTTP
}

func HTTPHandler(f Handler) http.Handler {
	return HandlerFunc(f.HandleCallback)
}

// use can register handler with http.Handle("/", callback.HTTPHandler(myHandler))
// or http.HandleFunc("/", HTTPHandleFunc(myHandlerFunc))
