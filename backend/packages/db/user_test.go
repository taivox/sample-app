package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUser_HashPassword(t *testing.T) {
	user := &User{
		Password: "plaintext123",
	}

	err := user.HashPassword()
	assert.NoError(t, err)
	assert.NotEqual(t, "plaintext123", user.Password)

	// Verify the hash is valid bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("plaintext123"))
	assert.NoError(t, err)
}

func TestUser_HashPassword_EmptyPassword(t *testing.T) {
	user := &User{
		Password: "",
	}

	err := user.HashPassword()
	assert.NoError(t, err)
	assert.NotEmpty(t, user.Password)
}

func TestUser_HashPassword_DifferentHashesForSamePassword(t *testing.T) {
	user1 := &User{Password: "samepassword"}
	user2 := &User{Password: "samepassword"}

	_ = user1.HashPassword()
	_ = user2.HashPassword()

	// bcrypt generates different hashes due to random salt
	assert.NotEqual(t, user1.Password, user2.Password)
}
