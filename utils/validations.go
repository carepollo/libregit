package utils

import "regexp"

// validate through REGEX if a string is a valid email address
func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	return matched && err == nil
}

// checks if a string has at least 8 characters at least 1 number and at least one uppercase
func ValidatePassword(password string) bool {
	pattern := `^[A-Za-z0-9]*[A-Z]+[A-Za-z0-9]*[0-9]+[A-Za-z0-9]*$`
	matched, err := regexp.MatchString(pattern, password)
	return matched && err == nil
}

// check if the given string is valid to be used as username or reponame for the application
// (git actions, url, etc). A valid username can only:
// have letters, numbers and the characters - or _, have a minimum of 3
// characters and maximum 25
func ValidateName(username string) bool {
	pattern := "^[a-zA-Z0-9_-]{3,25}$"
	matched, err := regexp.MatchString(pattern, username)
	return matched && err == nil
}
