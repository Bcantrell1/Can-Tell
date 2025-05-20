package verifier

import (
	"regexp"
	"strings"
)

type Structure struct {
	Username string `json:"username"`
	Domain   string `json:"domain"`
	Valid    bool   `json:"valid"`
}

var emailCheck = regexp.MustCompile(emailRegex)

func (v *Verifier) ParseEmail(email string) Structure {

	isEmailValid := IsEmailValid(email)
	if !isEmailValid {
		return Structure{Valid: false}
	}

	index := strings.LastIndex(email, "@")
	username := email[:index]
	domain := strings.ToLower(email[index+1:])

	return Structure{
		Username: username,
		Domain:   domain,
		Valid:    isEmailValid,
	}
}

func IsEmailValid(email string) bool {
	return emailCheck.MatchString(email)
}
