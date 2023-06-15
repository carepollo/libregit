package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// hash and salt plain string password, returns the hashed password
func HashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// check if plain string password is the equivalent of a given hash
func CheckPassword(hashed string, plain []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	return err == nil
}
