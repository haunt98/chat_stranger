package valid

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheckSignUpSubmit(t *testing.T) {
	testCases := []struct {
		inRegName  string
		inPassword string
		inShowName string
		ok         bool
	}{
		{
			inRegName:  "reg",
			inPassword: "pass",
			inShowName: "name",
			ok:         true,
		},
		{
			ok: false,
		},
		{
			inRegName: "reg2",
			ok:        false,
		},
		{
			inRegName:  "reg3",
			inPassword: "pass3",
			ok:         false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			err := CheckSignUpSubmit(tc.inShowName, tc.inRegName, tc.inPassword)
			if tc.ok {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestCheckLogInSubmit(t *testing.T) {
	testCases := []struct {
		inRegName  string
		inPassword string
		ok         bool
	}{
		{
			inRegName:  "reg",
			inPassword: "pass",
			ok:         true,
		},
		{
			ok: false,
		},
		{
			inRegName: "reg2",
			ok:        false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			err := CheckLogInSubmit(tc.inRegName, tc.inPassword)
			if tc.ok {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestCheckUpdateInfoSubmit(t *testing.T) {
	testCases := []struct {
		inShowName  string
		inGender    string
		inBrithYear int
		ok          bool
	}{
		{
			inShowName:  "name",
			inGender:    "gender",
			inBrithYear: 2000,
			ok:          true,
		},
		{
			ok: false,
		},
		{
			inShowName: "name2",
			ok:         false,
		},
		{
			inShowName:  "name3",
			inGender:    "gender3",
			inBrithYear: 0,
			ok:          false,
		},
		{
			inShowName:  "name3",
			inGender:    "gender3",
			inBrithYear: time.Now().Year() + 1,
			ok:          false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			err := CheckUpdateInfoSubmit(tc.inShowName, tc.inGender, tc.inBrithYear)
			if tc.ok {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
