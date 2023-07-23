package utils

import (
	"fmt"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr bool
	}{
		{"Success", "secret", false},
		{"Error", "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hash, err := GenerateHash(test.input)
			if (err != nil) != test.expectedErr {
				t.Errorf("GenerateHash(%q) error = %v, expectedErr %v", test.input, err, test.expectedErr)
			}
			if !test.expectedErr {
				if len(hash) == 0 {
					t.Errorf("GenerateHash(%q) = %q, expected non-empty string", test.input, hash)
				}
			}
		})
	}
}

func TestValidateHash(t *testing.T) {
	secret := "secret"
	poorHash := "fail"
	hash, err := GenerateHash(secret)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		secret  string
		hash    string
		isValid bool
	}{
		{"Success", secret, hash, true},
		{"Invalid", secret, hash[:len(hash)-len(poorHash)] + poorHash, false},
		{"Invalid", "invalid", "invalid", false},
	}
	for index, test := range tests {
		t.Run(fmt.Sprintf("Test %d | %s", index, test.name), func(t *testing.T) {
			if got, _ := ValidateHash(test.secret, test.hash); got != test.isValid {
				t.Errorf("ValidateHash(%q, %q) = %v, expected %v", test.secret, test.hash, got, test.isValid)
			}
		})
	}
}
