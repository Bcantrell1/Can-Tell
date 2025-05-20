package verifier

import (
	"testing"
	"time"
)

func TestNewVerifier(t *testing.T) {
	verifier := NewVerifier()

	if verifier.fromEmail != defaultEmail {
		t.Errorf("Expected fromEmail to be %s, got %s", defaultEmail, verifier.fromEmail)
	}

	if verifier.fromName != defaultName {
		t.Errorf("Expected fromName to be %s, got %s", defaultName, verifier.fromName)
	}

	if verifier.connectTimeout != 10*time.Second {
		t.Errorf("Expected connectTimeout to be 10 seconds, got %v", verifier.connectTimeout)
	}

	if verifier.operationTimeout != 10*time.Second {
		t.Errorf("Expected operationTimeout to be 10 seconds, got %v", verifier.operationTimeout)
	}
}

func TestValidate(t *testing.T) {
	verifier := NewVerifier()

	validateTests := []struct {
		email      string
		obtainable string
		valid      bool
	}{
		{"bribri546@gmail.com", obtainableYes, true},
		{"bribri546@invalid_domain", obtainableNo, false},
		{"invalid_email", obtainableNo, false},
		{"", obtainableNo, false},
	}
	for _, test := range validateTests {
		result, err := verifier.Validate(test.email)
		if err != nil {
			t.Errorf("Unexpected error for email %s: %v", test.email, err)
			continue
		}

		if result.Obtainable != test.obtainable {
			t.Errorf("Expected obtainable to be %s for email %s, got %s", test.obtainable, test.email, result.Obtainable)
		}

		if result.Structure.Valid != test.valid {
			t.Errorf("Expected valid to be %v for email %s, got %v", test.valid, test.email, result.Structure.Valid)
		}
	}
}
