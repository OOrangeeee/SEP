package infoModels

import (
	"time"
)

type Record struct {
	URL         string    `json:"url"`
	Type        string    `json:"type"`
	Time        time.Time `json:"time"`
	PatientName string    `json:"patientName"`
}
