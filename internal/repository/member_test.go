package repository

import (
	"fmt"
	"testing"

	"github.com/1612180/chat_stranger/internal/pkg/configwrap"
	"github.com/1612180/chat_stranger/internal/pkg/variable"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestMemberGorm_Create(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	memberGorm := memberGorm{db: db}

	// create data
	migrate(db, t)

	testCases := []struct {
		userID int
		roomID int
		wantOK bool
	}{
		{
			userID: 1,
			roomID: 1,
			wantOK: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("userID=%d roomID=%d", tc.userID, tc.roomID), func(t *testing.T) {
			ok := memberGorm.Create(tc.userID, tc.roomID)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}

func TestMemberGorm_Delete(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	memberGorm := memberGorm{db: db}

	// create data
	migrate(db, t)

	testCases := []struct {
		userID int
		wantOK bool
	}{
		{
			userID: 1,
			wantOK: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("userID=%d", tc.userID), func(t *testing.T) {
			ok := memberGorm.Delete(tc.userID)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}

func TestMemberGorm_CountByRoom(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	memberGorm := memberGorm{db: db}

	// create data
	migrate(db, t)

	testCases := []struct {
		roomID    int
		wantCount int
		wantOK    bool
	}{
		{
			roomID:    1,
			wantCount: 0,
			wantOK:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("roomID=%d", tc.roomID), func(t *testing.T) {
			count, ok := memberGorm.CountByRoom(tc.roomID)
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantCount, count)
		})
	}
}

func TestMemberGorm_CountByUser(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	memberGorm := memberGorm{db: db}

	// create data
	migrate(db, t)

	testCases := []struct {
		userID    int
		wantCount int
		wantOK    bool
	}{
		{
			userID:    1,
			wantCount: 0,
			wantOK:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("userID=%d", tc.userID), func(t *testing.T) {
			count, ok := memberGorm.CountByUser(tc.userID)
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantCount, count)
		})
	}
}
