package validator

import (
	"errors"
	"unicode"
)

func ValidatePassword(pass string) error {
	var (
		hasUpper, hasLower, hasNumber, hasSpecial bool
		length                                    int
	)

	for _, char := range pass {
		length++
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if length < 8 {
		return errors.New("La contraseña tiene que tener minimo 8 caracteres.")
	}
	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return errors.New("la contraseña debe contener mayúsculas, minúsculas, números y caracteres especiales")
	}
	return nil
}
