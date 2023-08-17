package validator

import (
	"regexp"
)

func IsEntryIdValid(input string) bool {
	pattern := "^[A-Za-z0-9]{22}$"
	regex := regexp.MustCompile(pattern)
	
	return regex.MatchString(input)
}