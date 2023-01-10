package internal

import (
	"unicode"
)

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	if check_all_chars(password) {
		return true
	}
	return false
}

func check_all_chars(password string) bool {
	for i := 0; i < len(password); i++ {
		if check_capital(rune(password[i])) {
			return true
		}
	}
	return false
}

func check_capital(letter rune) bool {
	return unicode.IsUpper(letter)
}
