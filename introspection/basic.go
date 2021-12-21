package introspection

import (
	"encoding/base64"
	"errors"
	"strings"
)

func BasicAuthorization(auth string) ([]string, error) {
	if auth == "" {
		return nil, errors.New("missing authorization header")
	}

	parts := strings.Fields(auth)
	auth_type := parts[0]
	creds := parts[1]
	if creds == "" {
		return nil, errors.New("corrupted authorization header")
	}
	if auth_type != "Basic" {
		return nil, errors.New("unsupported authorization type")
	}

	decoded, _ := base64.StdEncoding.DecodeString(creds)
	decoded_parts := strings.Split(string(decoded), ":")
	client_id := decoded_parts[0]
	client_secret := decoded_parts[1]
	if client_id == "" || client_secret == "" {
		return nil, errors.New("corrupted authorization credentials")
	}

	return []string{client_id, client_secret}, nil
}
