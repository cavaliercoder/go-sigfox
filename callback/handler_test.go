package callback

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type testHandler struct {
	t *testing.T
}

func (c *testHandler) HandleCallback(cb Callback) error {
	return nil
}

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(HTTPHandler(&testHandler{t}))

	defer ts.Close()

	http.Get(ts.URL)
}

func TestHandlerFunc(t *testing.T) {
	ts := httptest.NewServer(HTTPHandlerFunc(func(cb Callback) error {
		return nil
	}))

	defer ts.Close()

	http.Get(ts.URL)
}
