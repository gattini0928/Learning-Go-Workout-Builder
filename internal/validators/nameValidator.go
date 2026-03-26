package validators

import (
	"errors"
	"strings"
	"unicode"
)

func ValidateName(name string) error {
	parts := strings.Fields(name)
	if len(parts) < 2 {
		return errors.New("nome incompleto")
	}

	if len(parts[0]) < 2 || len(parts[1]) < 2 {
		return errors.New("nome ou sobrenome incompletos")
	}

	if parts[0] == parts[1] {
		return errors.New("nome e sobrenome iguais")
	}

	for _, char := range name {
		if !unicode.IsLetter(char) && char != ' ' {
			return errors.New("nome só pode conter letras")
		}
	}

	return nil
}