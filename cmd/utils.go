package dastco

import (
	"crypto/tls"
	"net/http"
)

// createHTTPClient creates a generic HTTP client.
func createHTTPClient() *http.Client {
	client := &http.Client{}
	if insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return client
}
