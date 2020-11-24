package validation

import (
	"strings"
	"regexp"
	"net"
	"unicode"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	//passRegex = regexp.MustCompile("^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{7,31}$")
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
		uppercasePresent bool
		lowercasePresent bool
		numberPresent bool
		specialCharPresent bool
		passLen int
	)

	const (
		minPassLength = 8
		maxPassLength = 30
	)

	for _, ch := range pass {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}

	return lowercasePresent && uppercasePresent && numberPresent && specialCharPresent && minPassLength <= passLen && passLen <= maxPassLength
}