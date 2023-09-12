package lib

import (
	"os"
	"testing"
)

func TestEncryptAndValidateJWT(t *testing.T) {
	os.Setenv("SECRET_KEY", "SECRET_KEY")

	username := "testuser"
	userID := "1234567890"
	email := "arol@arol.dev"

	token, err := EncryptJWT(username, userID, email)
	if err != nil {
		t.Fatalf("failed to encrypt JWT: %v", err)
	}

	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("failed to validate JWT: %v", err)
	}
	if claims.Username != username {
		t.Errorf("expected username %v, got %v", username, claims.Username)
	}
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	_, err := ValidateJWT("invalidToken")
	if err == nil {
		t.Error("expected error due to invalid token, got nil")
	}
}

func TestValidateJWT_NoSecretKey(t *testing.T) {
	os.Unsetenv("SECRET_KEY")
	username := "testuser"
	userID := "1234567890"
	email := "arol@arol.dev"

	token, err := EncryptJWT(username, userID, email)
	if err != nil {
		t.Fatalf("failed to encrypt JWT without secret key: %v", err)
	}

	_, err = ValidateJWT(token)
	if err == nil {
		t.Error("expected error due to missing secret key, got nil")
	}
}
