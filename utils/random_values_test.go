package utils

import "testing"

func TestRandomInts(t *testing.T) {
	inRange := 5
	input := []int{
		RandomInt(inRange),
		RandomInt(inRange),
		RandomInt(inRange),
		RandomInt(inRange),
		RandomInt(inRange),
	}

	for _, value := range input {
		if value > inRange || value < 0 {
			t.Fatalf(
				"Test RandomInt failed, the value %v is out of range, max %v",
				value,
				inRange,
			)
		}

	}

	if input[0] == input[1] && input[0] == input[2] {
		t.Fatalf(
			"Test RandomInt failed due to values not being random enough: %v %v %v",
			input[0],
			input[1],
			input[2],
		)
	}
}

func TestGeneratePassword(t *testing.T) {
	for i := 0; i < 3; i++ {
		one := GeneratePassword()
		two := GeneratePassword()
		three := GeneratePassword()

		if one == two || one == three {
			t.Fatalf(
				"Test GeneratePassword passwords are not randomly generated, values %s %s %s are equal",
				one,
				two,
				three,
			)
		}
	}
}
