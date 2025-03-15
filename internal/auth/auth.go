package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIkey extracts an API Key from
// the headers of the HTTP Request
// Example:
// Authorization: ApiKey {insert API Key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No authentication info found")
	}

	vals := strings.Split(val, " ")
	// The val is separated by space in between (ApiKey {insert API Key here}) where we need the second value
	if len(vals) != 2 {
		return "", errors.New("Malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part of auth header")
	}
	return vals[1], nil
}
