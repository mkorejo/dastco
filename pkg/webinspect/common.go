package webinspect

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// InvokeWebInspectAPI abstracts the HTTP request/response flow and returns the HTTP response body.
func InvokeWebInspectAPI(client *http.Client, request *http.Request) []byte {
	resp, err := client.Do(request)
	if err != nil {
		log.Error("Error sending the HTTP request: ", err)
		return nil
	}
	log.Info("Response code: ", resp.StatusCode)

	// Close resp.Body after reading
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the HTTP response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading the HTTP response: ", err)
		return nil
	}
	return body
}
