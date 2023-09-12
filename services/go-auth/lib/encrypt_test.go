package lib

import (
	"testing"
)

func TestHashAndValidate(t *testing.T) {
	password := "securePassword123"

	hashedPassword, err := Hash(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if !ValidatePSWHash(password, hashedPassword) {
		t.Error("Password validation failed for the correct password")
	}

	if ValidatePSWHash("wrong_password", hashedPassword) {
		t.Error("Password validation passed for the wrong password")
	}
}
