package helpers

import "unicode"

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLetter := false
	hasNumber := false

	for _, char := range password {
		if unicode.IsLetter(char) {
			hasLetter = true
		}
		if unicode.IsNumber(char) {
			hasNumber = true
		}

		if hasLetter && hasNumber {
			return true
		}
	}

	return false
}
