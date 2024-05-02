package infoModels

import (
	"time"
)

type Record struct {
	ID          uint      `json:"id"`
	URL         string    `json:"url"`
	Type        string    `json:"type"`
	Time        time.Time `json:"time"`
	PatientName string    `json:"patientName"`
}
