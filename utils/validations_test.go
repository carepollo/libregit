package utils

import "testing"

func TestEmailValidation(t *testing.T) {
	input := []string{
		"this is a test",
		"maik@",
		"maik@maik",
		"maik@maik.",
		"maik@maik.a",
		"maik@maik.a.b",
		"maik@maik.a.b.c",
		"maik@maik.a.b.",
		".maik@maik.a",
		".maik@maik.a.",
		"mail@example.com",
	}

	tests := []bool{
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		true,
	}

	for index, expected := range tests {
		actual := ValidateEmail(input[index])
		if expected != actual {
			t.Fatalf(
				"Test ValidateEmail have unmatched results, expected %v, got %v on value %s",
				expected,
				actual,
				input[index],
			)
		}
	}
}

func TestPasswordValidation(t *testing.T) {
	input := []string{
		GeneratePassword(),
		GeneratePassword(),
		GeneratePassword(),
		GeneratePassword(),
		"abc123",
		"abC123",
		"abCdg123",
		"aaaaaaaa",
		"aaaaaaa1",
		"aaaaaaaA",
		"aaaaaaA1",
		"   aB1    ",
	}

	tests := []bool{
		true,
		true,
		true,
		true,
		false,
		false,
		true,
		false,
		false,
		false,
		true,
		false,
	}

	for index, expected := range tests {
		actual := ValidatePassword(input[index])
		if actual != expected {
			t.Fatalf(
				"Test ValidatePassword have unexpected results, expected %v, got %v on %s",
				expected,
				actual,
				input[index],
			)
		}
	}
}

func TestNameValidation(t *testing.T) {
	input := []string{
		"user",
		"user-name",
		"UserName",
		"user_name",
		"user name",
		"user$name",
		" username",
		"username1",
		"a",
		" a",
		" 1",
		" a ",
		" a1",
	}

	tests := []bool{
		true,
		true,
		true,
		true,
		false,
		false,
		false,
		true,
		false,
		false,
		false,
		false,
	}

	for index, expected := range tests {
		actual := ValidateName(input[index])
		if actual != expected {
			t.Fatalf(
				"Test ValidateName expected %v and got %v on value %s",
				expected,
				actual,
				input[index],
			)
		}
	}
}
