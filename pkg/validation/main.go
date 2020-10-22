package validation

import (
	"strings"
	"regexp"
	"unicode"
	"net"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	passRegex = regexp.MustCompile("")
)

// IsEmail ...
func IsEmail(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}

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

// IsPassword ...
func IsPassword(pass string) bool {
	var (
		upp, low, num, sym bool
		sum                uint8
	)
 
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			sum++
		case unicode.IsLower(char):
			low = true
			sum++
		case unicode.IsNumber(char):
			num = true
			sum++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			sum++
		default:
			return false
		}
	}
 
	if !upp || !low || !num || !sym || sum < 8 || sum > 16{
		return false
	}
 
	return true
}