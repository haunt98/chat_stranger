package jwt

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	s, ok := Create(SignClaims{
		ID:             1,
		Role:           "user",
		StandardClaims: jwt.StandardClaims{},
	}, "secret")

	assert.Equal(t, true, ok)
	assert.NotEqual(t, "", s)
}

func TestVerify(t *testing.T) {
	s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiUm9sZSI6InVzZXIifQ.viQGhwxNYw6HjXmyVygxlgGj5Ue1_SYvcO8ApOU7hII"

	signClaims, ok := Verify(s, "secret")
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, signClaims.ID)
	assert.Equal(t, "user", signClaims.Role)
}
