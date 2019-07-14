package token

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	_, ok := Create(SignClaims{}, "")

	assert.True(t, ok)
}

func TestVerify(t *testing.T) {
	s, ok := Create(SignClaims{
		ID:             1,
		Role:           "a",
		StandardClaims: jwt.StandardClaims{},
	}, "b")

	assert.True(t, ok)

	signClaims, ok := Verify(s, "b")

	assert.True(t, ok)
	assert.Equal(t, 1, signClaims.ID)
	assert.Equal(t, "a", signClaims.Role)
}
