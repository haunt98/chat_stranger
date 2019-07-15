package token

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		claims interface{}
		secret string
		wantOK bool
	}{
		{
			claims: SignClaims{},
			secret: "",
			wantOK: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("SignClaims=%v secret=%s", tc.claims, tc.secret), func(t *testing.T) {
			_, ok := Create(tc.claims.(SignClaims), tc.secret)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}

func TestVerify(t *testing.T) {
	testCases := []struct {
		s          string
		secret     string
		wantClaims *SignClaims
		wantOK     bool
	}{
		{
			wantClaims: nil,
			wantOK:     false,
		},
		{
			wantClaims: &SignClaims{},
			wantOK:     true,
		},
	}

	for _, tc := range testCases {
		if tc.wantOK {
			s, ok := Create(tc.wantClaims, tc.secret)
			assert.True(t, ok)
			tc.s = s
		}

		t.Run(fmt.Sprintf("s=%s secret=%s", tc.s, tc.secret), func(t *testing.T) {
			claims, ok := Verify(tc.s, tc.secret)
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantClaims, claims)
		})
	}
}
