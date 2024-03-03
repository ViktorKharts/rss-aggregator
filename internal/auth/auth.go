package auth

import (
	"net/http"
	"strings"
	"errors"
)

func GetApiKey(h http.Header) (string, error) {
	header := h.Get("Authorization")
	if header == "" {
		return "", errors.New("No API Key was provided")
	}

	splitAuth := strings.Split(header, " ")
	if len(splitAuth) > 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("Malformed API Key")
	}

	return splitAuth[1], nil
}

