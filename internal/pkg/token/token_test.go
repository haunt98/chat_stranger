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
		ok     bool
	}{
		{
			claims: AccountClaims{},
			secret: "",
			ok:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("AccountClaims=%v secret=%s", tc.claims, tc.secret), func(t *testing.T) {
			_, err := Create(tc.claims.(AccountClaims), tc.secret)
			if tc.ok {
				assert.Nil(t, err)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	s1, _ := Create(AccountClaims{ID: 1, Role: "a"}, "b")

	testCases := []struct {
		s             string
		secret        string
		accountClaims AccountClaims
		ok            bool
	}{
		{
			s:             s1,
			secret:        "b",
			accountClaims: AccountClaims{ID: 1, Role: "a"},
			ok:            true,
		},
		{
			s:  "",
			ok: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%s secret=%s", tc.s, tc.secret), func(t *testing.T) {
			accountClaims, err := Verify(tc.s, tc.secret)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.accountClaims, accountClaims)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
