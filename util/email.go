package util

import (
	"net"
	"regexp"
	"strings"
)

// Email represents util for email
type Email string

// Validate validates email, returns TRUE if ok, otherwise FALSE
func (e Email) Validate() bool {
	email := string(e)
	if len(email) < 3 && len(email) > 254 {
		return false
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(email) {
		return false
	}

	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {

		return false
	}

	return true
}
