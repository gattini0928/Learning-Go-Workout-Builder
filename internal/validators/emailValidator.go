package validators

import (
	"errors"
	"net/mail"
	"strings"
)

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
        return errors.New("Email inválido")
    }
	emailParts := strings.Split(email, "@")
	domain := emailParts[len(emailParts)-1]
	charToFind := "."

	if !strings.Contains(domain, charToFind) {
        return errors.New("Email inválido")
	}
	return nil
}