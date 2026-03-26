package validators

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"unicode"
	"fmt"
)

func ValidatePassword(password string) error {

	if password == "" {
		return errors.New("senha em branco")
	}

	if len(password) < 8 {
		return errors.New("senha deve ter pelo menos 8 caracteres")
	}

	specials := func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	}

	rules := map[string]func(rune) bool {
		"letra maiúscula": unicode.IsUpper,
		"letra minúscula": unicode.IsLower,
		"dígito": unicode.IsDigit,
		"caractere especial": specials,
	}

	for ruleName, ruleFunc := range rules {
		found := false
		for _, char := range password {
			if ruleFunc(char) {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("sua senha deve conter pelo menos um(a) %s", ruleName)
		}
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}