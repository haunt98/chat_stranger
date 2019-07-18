package valid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRegisterName(t *testing.T) {
	testCases := []struct {
		name   string
		wantOK bool
	}{
		{
			name:   "",
			wantOK: false,
		},
		{
			name:   "a",
			wantOK: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("name=%s", tc.name), func(t *testing.T) {
			ok, _ := CheckRegisterName(tc.name)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}

func TestCheckPassword(t *testing.T) {
	testCases := []struct {
		password string
		wantOK   bool
	}{
		{
			password: "",
			wantOK:   false,
		},
		{
			password: "a",
			wantOK:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("password=%s", tc.password), func(t *testing.T) {
			ok, _ := CheckPassword(tc.password)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}

func TestCheckFullName(t *testing.T) {
	testCases := []struct {
		name   string
		wantOK bool
	}{
		{
			name:   "",
			wantOK: false,
		},
		{
			name:   "a",
			wantOK: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("name=%s", tc.name), func(t *testing.T) {
			ok, _ := CheckFullName(tc.name)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}
