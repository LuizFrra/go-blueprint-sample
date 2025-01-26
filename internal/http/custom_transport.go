package http

import (
	"log"
	"net/http"
)

type LoggingTransport struct {
	Transport http.RoundTripper
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("Client Request: %s %s", req.Method, req.URL)
	return t.Transport.RoundTrip(req)
}
