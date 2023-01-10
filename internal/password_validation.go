package internal

import (
	"errors"
	"unicode"
)

type Validation uint

const (
	One Validation = iota
	Two
	Three
)

func ValidatePassword(validation Validation, password string) (errors []error) {
	switch validation {
	case One:
		return validation1(password)
	case Two:
		return validation2(password)
	case Three:
		return validation3(password)
	}
	return []error{}
}

func validation1(password string) (errors []error) {
	errors = []error{}
	err := check_length(password, 8)
	if err != nil {
		errors = append(errors, err)
	}

	err = check_upper(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = check_lower(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = check_contains_digit(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = check_underscore(password)
	if err != nil {
		errors = append(errors, err)
	}
	return errors
}

func validation2(password string) (errors []error) {
	errors = []error{}
	err := check_length(password, 6)
	if err != nil {
		errors = append(errors, err)
	}

	err = check_upper(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = check_lower(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = check_contains_digit(password)
	if err != nil {
		errors = append(errors, err)
	}
	return errors
}
func validation3(password string) (errors []error) {
	errors = []error{}
	err := check_length(password, 16)
	if err != nil {
		errors = append(errors, err)
	}

	err = check_upper(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = check_lower(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = check_underscore(password)
	if err != nil {
		errors = append(errors, err)
	}
	return errors
}

func check_length(password string, max int) error {
	if len(password) > max {
		return nil
	}
	return errors.New("invalid length")
}

func check_contains_digit(password string) error {
	for i := 0; i < len(password); i++ {
		if unicode.IsDigit(rune(password[i])) {
			return nil
		}
	}
	return errors.New("a digit is required")
}

func check_underscore(password string) error {
	for i := 0; i < len(password); i++ {
		if password[i] == '_' {
			return nil
		}
	}
	return errors.New("an underscore char is required")
}
func check_upper(password string) error {
	for i := 0; i < len(password); i++ {
		if unicode.IsUpper(rune(password[i])) {
			return nil
		}
	}
	return errors.New("an upper case char is required")
}
func check_lower(password string) error {
	for i := 0; i < len(password); i++ {
		if unicode.IsLower(rune(password[i])) {
			return nil
		}
	}
	return errors.New("a lower case char is required")
}
