package valid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRegisterName(t *testing.T) {
	testCases := []struct {
		name string
		want bool
	}{
		{
			name: "",
			want: false,
		},
		{
			name: "a",
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("name=%s", tc.name), func(t *testing.T) {
			ok := CheckRegisterName(tc.name)
			assert.Equal(t, tc.want, ok)
		})
	}
}

func TestCheckPassword(t *testing.T) {
	testCases := []struct {
		password string
		want     bool
	}{
		{
			password: "",
			want:     false,
		},
		{
			password: "a",
			want:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("password=%s", tc.password), func(t *testing.T) {
			ok := CheckPassword(tc.password)
			assert.Equal(t, tc.want, ok)
		})
	}
}

func TestCheckPassword2(t *testing.T) {
	testCases := []struct {
		password  string
		password2 string
		want      bool
	}{
		{
			password:  "",
			password2: "",
			want:      false,
		},
		{
			password:  "",
			password2: "a",
			want:      false,
		},
		{
			password:  "a",
			password2: "",
			want:      false,
		},
		{
			password:  "a",
			password2: "b",
			want:      false,
		},
		{
			password:  "a",
			password2: "a",
			want:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("password=%s password2=%s", tc.password, tc.password2), func(t *testing.T) {
			ok := CheckPassword2(tc.password, tc.password2)
			assert.Equal(t, tc.want, ok)
		})
	}
}
