package helper

import "strings"

func DidUserEmailPassValidation(email string) bool {
	valid := true
	if len(email) == 0 {
		valid = false
	}

	if !strings.Contains(email, "@") {
		valid = false
	}

	if !strings.Contains(email, ".") {
		valid = false
	}

	return valid
}
