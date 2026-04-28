package utils

import "testing"

func TestHashPasswordAndCheck(t *testing.T) {
	pwd := "S3cretPass!"
	hash, err := HashPassword(pwd)
	if err != nil {
		t.Fatalf("HashPassword error: %v", err)
	}
	if ok := CheckPassword(pwd, hash); !ok {
		t.Fatalf("CheckPassword expected true")
	}
	if ok := CheckPassword("wrong", hash); ok {
		t.Fatalf("CheckPassword expected false for wrong password")
	}
}

func TestGenerateAndParseJWT(t *testing.T) {
	username := "alice"
	token, err := GenerateJWT(username)
	if err != nil {
		t.Fatalf("GenerateJWT error: %v", err)
	}

	parsed, err := ParseJWT(token)
	if err != nil {
		t.Fatalf("ParseJWT error: %v", err)
	}
	if parsed != username {
		t.Fatalf("ParseJWT mismatch: got %q want %q", parsed, username)
	}
}

func TestParseJWTInvalid(t *testing.T) {
	_, err := ParseJWT("Bearer invalid.token.value")
	if err == nil {
		t.Fatalf("ParseJWT expected error for invalid token")
	}
}
