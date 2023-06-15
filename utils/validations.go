package utils

import (
	"regexp"
	"strings"
)

// validate through REGEX if a string is a valid email address
func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	return matched && err == nil
}

// checks if a string has at least 8 characters, has 1 either number or special char
func ValidatePassword(password string) bool {
	if len(strings.TrimSpace(password)) == 0 {
		return false
	}

	if len(password) < 8 {
		return false
	}

	// At least one special character or number
	pattern := "[^a-zA-Z]"
	matched, err := regexp.MatchString(pattern, password)
	return matched && err == nil
}

// check if the given string is valid to be used as User for the application (git actions, url, etc).
// A valid username can only have letters, numbers and the characters - or _, have a minimum of 2
// characters and maximum 25
func ValidateUsername(username string) bool {
	pattern := "^[a-zA-Z0-9_-]{2,25}$"
	matched, err := regexp.MatchString(pattern, username)
	return matched && err == nil
}
