package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from the Authorization header of an HTTP request.
// Expected format: "Authorization: ApiKey {apikey}"
func GetAPIKey(headers http.Header) (string, error) {
	const authScheme = "ApiKey"

	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization header missing")
	}

	parts := strings.Fields(val)
	if len(parts) != 2 || parts[0] != authScheme {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}
