package validator

import (
	"regexp"
)

func IsEntryIdValid(input string) bool {
	pattern := "^[A-Za-z0-9]{22}$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)
}

func IsSpaceIdValid(input string) bool {
	pattern := "^[A-Za-z0-9]{12}$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)
}

func IsAccessTokenValid(input string) bool {
	pattern := "^[A-Za-z0-9]{43}$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)
}
