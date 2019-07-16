package repository

import (
	"fmt"
	"testing"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/configwrap"
	"github.com/1612180/chat_stranger/internal/pkg/variable"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestMemberGorm_NewMemberRepo(t *testing.T) {
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

	testCases := []struct {
		db             *gorm.DB
		wantMemberRepo MemberRepo
	}{
		{
			db:             db,
			wantMemberRepo: &memberGorm{db: db},
		},
	}

	for _, tc := range testCases {
		t.Run("NewRepo", func(t *testing.T) {
			memberRepo := NewMemberRepo(tc.db)
			assert.Equal(t, tc.wantMemberRepo, memberRepo)
		})
	}
}

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
		userID     int
		roomID     int
		haveMember bool
		wantOK     bool
	}{
		{
			userID:     1,
			roomID:     1,
			haveMember: false,
			wantOK:     false,
		},
		{
			userID:     1,
			roomID:     1,
			haveMember: true,
			wantOK:     true,
		},
	}

	for _, tc := range testCases {
		if !tc.haveMember {
			if err := db.DropTableIfExists(&model.Member{}).Error; err != nil {
				t.Error(err)
			}
		} else {
			if err := db.AutoMigrate(&model.Member{}).Error; err != nil {
				t.Error(err)
			}
		}

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
		userID     int
		haveMember bool
		wantOK     bool
	}{
		{
			userID:     1,
			haveMember: false,
			wantOK:     false,
		},
		{
			userID:     1,
			haveMember: true,
			wantOK:     true,
		},
	}

	for _, tc := range testCases {
		if !tc.haveMember {
			if err := db.DropTableIfExists(&model.Member{}).Error; err != nil {
				t.Error(err)
			}
		} else {
			if err := db.AutoMigrate(&model.Member{}).Error; err != nil {
				t.Error(err)
			}
		}

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
		roomID     int
		haveMember bool
		wantCount  int
		wantOK     bool
	}{
		{
			roomID:     1,
			haveMember: false,
			wantCount:  0,
			wantOK:     false,
		},
		{
			roomID:     1,
			haveMember: true,
			wantCount:  0,
			wantOK:     true,
		},
	}

	for _, tc := range testCases {
		if !tc.haveMember {
			if err := db.DropTableIfExists(&model.Member{}).Error; err != nil {
				t.Error(err)
			}
		} else {
			if err := db.AutoMigrate(&model.Member{}).Error; err != nil {
				t.Error(err)
			}
		}

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
		userID     int
		haveMember bool
		wantCount  int
		wantOK     bool
	}{
		{
			userID:     1,
			haveMember: false,
			wantCount:  0,
			wantOK:     false,
		},
		{
			userID:     1,
			haveMember: true,
			wantCount:  0,
			wantOK:     true,
		},
	}

	for _, tc := range testCases {
		if !tc.haveMember {
			if err := db.DropTableIfExists(&model.Member{}).Error; err != nil {
				t.Error(err)
			}
		} else {
			if err := db.AutoMigrate(&model.Member{}).Error; err != nil {
				t.Error(err)
			}
		}

		t.Run(fmt.Sprintf("userID=%d", tc.userID), func(t *testing.T) {
			count, ok := memberGorm.CountByUser(tc.userID)
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantCount, count)
		})
	}
}
