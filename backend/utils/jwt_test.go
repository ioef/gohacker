package utils

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	userID := "test-user-123"
	username := "testuser"
	secret := "test-secret"
	expiration := 24 * time.Hour

	token, err := GenerateToken(userID, username, secret, expiration)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	if token == "" {
		t.Error("Token should not be empty")
	}
}

func TestValidateToken(t *testing.T) {
	userID := "test-user-123"
	username := "testuser"
	secret := "test-secret"
	expiration := 24 * time.Hour

	// Generate a valid token
	token, err := GenerateToken(userID, username, secret, expiration)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Validate the token
	claims, err := ValidateToken(token, secret)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected UserID=%s, got %s", userID, claims.UserID)
	}

	if claims.Username != username {
		t.Errorf("Expected Username=%s, got %s", username, claims.Username)
	}
}

func TestValidateTokenWithWrongSecret(t *testing.T) {
	userID := "test-user-123"
	username := "testuser"
	secret := "test-secret"
	wrongSecret := "wrong-secret"
	expiration := 24 * time.Hour

	// Generate a token with one secret
	token, err := GenerateToken(userID, username, secret, expiration)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Try to validate with wrong secret
	_, err = ValidateToken(token, wrongSecret)
	if err == nil {
		t.Error("ValidateToken should fail with wrong secret")
	}
}

func TestValidateTokenExpired(t *testing.T) {
	userID := "test-user-123"
	username := "testuser"
	secret := "test-secret"
	expiration := -1 * time.Hour // Already expired

	// Generate an expired token
	token, err := GenerateToken(userID, username, secret, expiration)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Try to validate expired token
	_, err = ValidateToken(token, secret)
	if err == nil {
		t.Error("ValidateToken should fail with expired token")
	}
}
