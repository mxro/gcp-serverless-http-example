package hello

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHTTP(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Content-Type", "text/html")

	rr := httptest.NewRecorder()
	HelloHTTP(rr, req)

	if got := rr.Body.String(); strings.Contains(got, "Hello There!") {
		t.Errorf("HelloHTTP() = %q, want %q", got, "Hello There")
	}
}
