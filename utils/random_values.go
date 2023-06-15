package utils

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// generates a random integer number within given range
func RandomInt(max int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max)
}

// generates a random password of 8 characters long with random uppercase and lowercase letters
// and random numbers
func GeneratePassword() string {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	result := ""
	var index int
	turned := false
	rand.NewSource(time.Now().UnixNano())

	// add random characters that are either numbers or letters
	for i := 1; i <= 8; i++ {
		letterOrNumber := RandomInt(2)
		possibleLetter := RandomInt(len(letters) - 1)
		possibleNumber := RandomInt(len(numbers) - 1)
		choose := letterOrNumber%2 == 0
		value := ""
		if choose {
			value = letters[possibleLetter]
		} else {
			value = numbers[possibleNumber]
		}
		result += value
	}

	// iterate until at least one letter is turned uppercase
	index = RandomInt(len(result) - 1)
	for !turned {
		isNumeric := unicode.IsDigit(rune(result[index]))
		if !isNumeric {
			upperized := strings.ToUpper(string(result[index]))
			result = strings.Replace(result, string(result[index]), upperized, 1)
			turned = true
		}
		index = RandomInt(len(result) - 1)
	}

	return result
}
