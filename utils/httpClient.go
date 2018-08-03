package utils

import (
	"net/http"
	"time"
)

// NewHTTPClient is a helper function to retrieve a previously configurated HTTP Client.
func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 30,
	}
}
