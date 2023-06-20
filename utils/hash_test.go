package utils

import "testing"

func TestHashedAndSalted(t *testing.T) {
	plainString := "passwordxd"
	one, _ := HashAndSalt(plainString)
	two, _ := HashAndSalt(plainString)

	if one == plainString || two == plainString {
		t.Fatalf(
			"Test HashAndSalt the password was not hashed, outputs one=%s two=%s given input %s",
			one,
			two,
			plainString,
		)
	}

	if one == two {
		t.Fatalf(
			"Test HashAndSalt the password are not being salted appropiately, two same inputs return same output, input %s output %s",
			plainString,
			one, // since they are the same doesn't matter which to print out
		)
	}
}

func TestCheckPassword(t *testing.T) {
	plain := "password"
	hashed, _ := HashAndSalt(plain)

	if !CheckPassword(plain, hashed) {
		t.Fatalf(
			"Test CheckPassword plain string is not accepted as its hashed equivalent, plain %s and hash %s",
			plain,
			hashed,
		)
	}
}
