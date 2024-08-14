package helper

import (
	"strings"
)

var symbols = []string{
	"@",
	"#",
	"$",
	"%",
	"^",
	"&",
	"*",
	"(",
	")",
	"-",
	"_",
	"=",
	"+",
	"[",
	"]",
	"{",
	"}",
	"|",
	"\\",
	";",
	":",
	"'",
	"\"",
	",",
	".",
	"/",
	"<",
	">",
	"?",
	"~",
	"!",
	"`",
}

var badDisplayNameOptions = []string{
	"admin",
	"administrator",
	"root",
}

func DidUserDisplayNamePassValidation(displayName string) bool {
	var valid = true
	if len(displayName) < 5 {
		valid = false
	}

	for _, value := range badDisplayNameOptions {
		if strings.ToLower(displayName) == value {
			valid = false
		}
	}

	for _, value := range symbols {
		if strings.Contains(displayName, value) {
			valid = false
		}
	}

	return valid
}
