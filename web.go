package jwt

import (
	"crypto/rsa"
	"errors"
	"net/http"
	"strings"
)

// MIMEType is the IANA registered media type.
const MIMEType = "application/jwt"

// OAuthURN is the IANA registered OAuth URI.
const OAuthURN = "urn:ietf:params:oauth:token-type:jwt"

var (
	errAuthHeader = errors.New("want Authorization header")
	errAuthSchema = errors.New("want Bearer schema")
)

// HMACCheckHeader applies HMACCheck on a HTTP requests.
// Specifically it looks for the Bearer schema in the Authorization header.
func HMACCheckHeader(r *http.Request, secret []byte) (*Claims, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return nil, errAuthHeader
	}
	if !strings.HasPrefix(auth, "Bearer ") {
		return nil, errAuthSchema
	}
	return HMACCheck(auth[7:], secret)
}

// RSACheckHeader applies RSACheck on a HTTP requests.
// Specifically it looks for the Bearer schema in the Authorization header.
func RSACheckHeader(r *http.Request, key *rsa.PublicKey) (*Claims, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return nil, errAuthHeader
	}
	if !strings.HasPrefix(auth, "Bearer ") {
		return nil, errAuthSchema
	}
	return RSACheck(auth[7:], key)
}