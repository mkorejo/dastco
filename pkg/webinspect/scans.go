package webinspect

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Scan contains the details of a scan in WebInspect.
type Scan struct {
	ID        string `json:"ID"`
	Name      string `json:"Name"`
	StartTime string `json:"StartTime"`
	Status    string `json:"Status"`
}

// ScanStatus contains the status of a scan in WebInspect.
type ScanStatus struct {
	Status string `json:"scanStatus"`
}

// GetScanStatus returns the status of a specific scan.
func GetScanStatus(client *http.Client, url string, username string, password string, scanID string) ScanStatus {
	req, err := http.NewRequest("GET", url+"/scanner/scans/"+scanID, nil)
	if err != nil {
		log.Error("Building the HTTP request: ", err)
	}
	req.SetBasicAuth(username, password)

	body := InvokeWebInspectAPI(client, req)

	// Unmarshal the JSON
	s := ScanStatus{}
	if err := json.Unmarshal(body, &s); err != nil {
		log.Error("Reading JSON: ", err)
	}
	return s
}

// ListScans returns all WebInspect scans in the database.
func ListScans(client *http.Client, url string, username string, password string) []Scan {
	req, err := http.NewRequest("GET", url+"/scanner/scans/", nil)
	if err != nil {
		log.Error("Building the HTTP request:", err)
	}
	req.SetBasicAuth(username, password)

	body := InvokeWebInspectAPI(client, req)

	// Unmarshal the JSON
	s := []Scan{}
	if err := json.Unmarshal(body, &s); err != nil {
		log.Error("Reading JSON:", err)
	}
	return s
}

// StartScan ...
func StartScan(client *http.Client, url string, username string, password string) {

}
