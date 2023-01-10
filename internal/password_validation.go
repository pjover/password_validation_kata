package internal

import (
	"unicode"
)

type Validation uint

const (
	One Validation = iota
	Two
)

func ValidatePassword(validation Validation, password string) bool {
	switch validation {
	case One:
		return validation1(password)
	case Two:
		return validation2(password)
	}

	return false
}

func validation1(password string) bool {
	if len(password) < 8 {
		return false
	}
	if !check_upper(password) {
		return false
	}

	if !check_lower(password) {
		return false
	}

	if !check_underscore(password) {
		return false
	}
	return check_contains_digit(password)
}

func validation2(password string) bool {
	if len(password) < 7 {
		return false
	}
	if !check_upper(password) {
		return false
	}

	if !check_lower(password) {
		return false
	}
	return check_contains_digit(password)
}

func check_contains_digit(password string) bool {
	for i := 0; i < len(password); i++ {
		if unicode.IsDigit(rune(password[i])) {
			return true
		}
	}
	return false
}

func check_underscore(password string) bool {
	for i := 0; i < len(password); i++ {
		if password[i] == '_' {
			return true
		}
	}
	return false
}
func check_upper(password string) bool {
	for i := 0; i < len(password); i++ {
		if unicode.IsUpper(rune(password[i])) {
			return true
		}
	}
	return false
}
func check_lower(password string) bool {
	for i := 0; i < len(password); i++ {
		if unicode.IsLower(rune(password[i])) {
			return true
		}
	}
	return false
}
