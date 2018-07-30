// ipify-api/models
//
// This package contains all models used in the ipify service.

package models

// IPAddress is a struct we use to represent JSON API responses.
type IPAddress struct {
	IP string `json:"ip"`
}

// RequestInfo is a struct we use to represent JSON API responses.
type RequestInfo struct {
	IP string `json:"ip"`
	UserAgent string `json:"user-agent"`
	Referrer string `json:"referer"`
}

// TimeUTC is a struct we use to represent JSON API responses.
type TimeUTC struct {
	UTC string `json:"utc"`
}
