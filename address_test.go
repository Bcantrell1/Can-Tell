package verifier

import (
	"testing"
)

var (
	examples = []struct {
		email   string
		correct bool
	}{
		{email: "contact@business.org", correct: true},
		{email: "help@microsoft.com", correct: true},
		{email: " customer.service@amazon.com", correct: false},
		{email: "notifications@twitter.com", correct: true},
		{email: "🕵🏼@outlook.com", correct: false},
		{email: "newsletter@nytimes.com", correct: true},
		{email: "john_doe@university.edu", correct: true},
		{email: "info@company.io", correct: true},
	}
)

func TestCheckAddressSyntax(t *testing.T) {
	var verifier = NewVerifier()
	for _, is := range examples {
		emailAddress := verifier.ParseEmail(is.email)
		if !emailAddress.Valid && is.correct == true {
			t.Errorf(`"%s" check failed with an unexpected error`, is.email)
		}
		if emailAddress.Valid && is.correct == false {
			t.Errorf(`"%s" => incorrect email address`, is.email)
		}
	}
}
