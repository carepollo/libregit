package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// hash and salt plain string password, returns the hashed password
func HashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// check if plain string password is the equivalent of a given hash
func CheckPassword(plain string, hashed string) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plain))
	return err == nil
}
