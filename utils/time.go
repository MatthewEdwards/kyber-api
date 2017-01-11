package utils

import (
	"time"
)

// GetTimestamp will return the current time as a ISO 8601 string
func GetTimestamp() string {
	t := time.Now()
	timestamp := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		time.Local,
	).Format("2006-01-02 15:04:05")

	return timestamp
}
