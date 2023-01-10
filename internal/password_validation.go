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
		errors = validation1(password)
	case Two:
		errors = validation2(password)
	case Three:
		errors = validation3(password)
	}
	if len(errors) == 1 {
		return []error{}
	}
	return errors
}

func validation1(password string) (errors []error) {
	errors = []error{}
	err := checkLength(password, 8)
	if err != nil {
		errors = append(errors, err)
	}

	err = checkUpper(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = checkLower(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = checkContainsDigit(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = checkUnderscore(password)
	if err != nil {
		errors = append(errors, err)
	}
	return errors
}

func validation2(password string) (errors []error) {
	errors = []error{}
	err := checkLength(password, 6)
	if err != nil {
		errors = append(errors, err)
	}

	err = checkUpper(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = checkLower(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = checkContainsDigit(password)
	if err != nil {
		errors = append(errors, err)
	}
	return errors
}
func validation3(password string) (errors []error) {
	errors = []error{}
	err := checkLength(password, 16)
	if err != nil {
		errors = append(errors, err)
	}

	err = checkUpper(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = checkLower(password)
	if err != nil {
		errors = append(errors, err)
	}

	err = checkUnderscore(password)
	if err != nil {
		errors = append(errors, err)
	}
	return errors
}

func checkLength(password string, max int) error {
	if len(password) > max {
		return nil
	}
	return errors.New("invalid length")
}

func checkContainsDigit(password string) error {
	for i := 0; i < len(password); i++ {
		if unicode.IsDigit(rune(password[i])) {
			return nil
		}
	}
	return errors.New("a digit is required")
}

func checkUnderscore(password string) error {
	for i := 0; i < len(password); i++ {
		if password[i] == '_' {
			return nil
		}
	}
	return errors.New("an underscore char is required")
}
func checkUpper(password string) error {
	for i := 0; i < len(password); i++ {
		if unicode.IsUpper(rune(password[i])) {
			return nil
		}
	}
	return errors.New("an upper case char is required")
}
func checkLower(password string) error {
	for i := 0; i < len(password); i++ {
		if unicode.IsLower(rune(password[i])) {
			return nil
		}
	}
	return errors.New("a lower case char is required")
}
