package utils

import (
	"testing"

	"backend/packages/db"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser_ValidUser(t *testing.T) {
	user := db.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	errs := ValidateUser(user)
	assert.Empty(t, errs)
}

func TestValidateUser_InvalidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
	}{
		{"empty email", ""},
		{"no domain", "john@"},
		{"no at sign", "johnexample.com"},
		{"invalid tld", "john@example.c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := db.User{
				Name:     "John",
				Email:    tt.email,
				Password: "password123",
			}
			errs := ValidateUser(user)
			assert.Contains(t, errs, "Invalid email")
		})
	}
}

func TestValidateUser_InvalidPassword(t *testing.T) {
	user := db.User{
		Name:     "John",
		Email:    "john@example.com",
		Password: "abc",
	}

	errs := ValidateUser(user)
	assert.Contains(t, errs, "Invalid password, Password should be more than 4 characters")
}

func TestValidateUser_InvalidName(t *testing.T) {
	user := db.User{
		Name:     "",
		Email:    "john@example.com",
		Password: "password123",
	}

	errs := ValidateUser(user)
	assert.Contains(t, errs, "Invalid name, please enter a name")
}

func TestValidateUser_MultipleErrors(t *testing.T) {
	user := db.User{
		Name:     "",
		Email:    "invalid",
		Password: "abc",
	}

	errs := ValidateUser(user)
	assert.Len(t, errs, 3)
}

func TestValidatePasswordReset_Valid(t *testing.T) {
	reset := db.ResetPassword{
		Password:        "newpassword",
		ConfirmPassword: "newpassword",
	}

	valid, msg := ValidatePasswordReset(reset)
	assert.True(t, valid)
	assert.Empty(t, msg)
}

func TestValidatePasswordReset_TooShort(t *testing.T) {
	reset := db.ResetPassword{
		Password:        "abc",
		ConfirmPassword: "abc",
	}

	valid, msg := ValidatePasswordReset(reset)
	assert.False(t, valid)
	assert.Contains(t, msg, "more than 4 characters")
}

func TestValidatePasswordReset_Mismatch(t *testing.T) {
	reset := db.ResetPassword{
		Password:        "password1",
		ConfirmPassword: "password2",
	}

	valid, msg := ValidatePasswordReset(reset)
	assert.False(t, valid)
	assert.Contains(t, msg, "passwords must match")
}

func TestGetHash(t *testing.T) {
	hash, err := GetHash("password123")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, "password123", hash)
}

func TestComparePassword(t *testing.T) {
	hash, _ := GetHash("password123")

	assert.True(t, ComparePassword(hash, "password123"))
	assert.False(t, ComparePassword(hash, "wrongpassword"))
}
