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

func TestRoomGorm_Exist(t *testing.T) {
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

	roomGorm := roomGorm{db: db}

	// create data
	migrate(db, t)

	if err := db.Create(&model.Room{}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		id     int
		wantOK bool
	}{
		{
			id:     0,
			wantOK: false,
		},
		{
			id:     1,
			wantOK: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("id=%d", tc.id), func(t *testing.T) {
			ok := roomGorm.Exist(tc.id)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}

func TestRoomGorm_FindEmpty(t *testing.T) {
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

	roomGorm := roomGorm{db: db}

	// create data
	migrate(db, t)

	testCases := []struct {
		wantRoom *model.Room
		wantOK   bool
	}{
		{
			wantRoom: nil,
			wantOK:   false,
		},
		{
			wantRoom: &model.Room{ID: 1},
			wantOK:   true,
		},
	}

	for _, tc := range testCases {
		if tc.wantOK {
			if err := db.Create(&model.Room{}).Error; err != nil {
				t.Error(err)
			}
		}

		t.Run(fmt.Sprintf("%t", tc.wantOK), func(t *testing.T) {
			room, ok := roomGorm.FindEmpty()
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantRoom, room)
		})
	}
}

func TestRoomGorm_FindNext(t *testing.T) {
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

	roomGorm := roomGorm{db: db}

	// create data
	migrate(db, t)

	if err := db.Create(&model.Room{}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		old      int
		wantRoom *model.Room
		wantOK   bool
	}{
		{
			old:      1,
			wantRoom: nil,
			wantOK:   false,
		},
		{
			old:      2,
			wantRoom: &model.Room{ID: 1},
			wantOK:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("old=%d", tc.old), func(t *testing.T) {
			room, ok := roomGorm.FindNext(tc.old)
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantRoom, room)
		})
	}
}

func TestRoomGorm_FindByUser(t *testing.T) {
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

	roomGorm := roomGorm{db: db}

	// create data
	migrate(db, t)

	if err := db.Create(&model.Room{}).Error; err != nil {
		t.Error(err)
	}

	if err := db.Create(&model.Member{
		UserID: 1,
		RoomID: 1,
	}).Error; err != nil {
		t.Error(err)
	}

	if err := db.Create(&model.Member{
		UserID: 2,
		RoomID: 1,
	}).Error; err != nil {
		t.Error(err)
	}

	if err := db.Create(&model.Member{
		UserID: 2,
		RoomID: 2,
	}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		userID   int
		wantRoom *model.Room
		wantOK   bool
	}{
		{
			userID:   0,
			wantRoom: nil,
			wantOK:   false,
		},
		{
			userID:   1,
			wantRoom: &model.Room{ID: 1},
			wantOK:   true,
		},
		{
			userID:   2,
			wantRoom: nil,
			wantOK:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("userID=%d", tc.userID), func(t *testing.T) {
			room, ok := roomGorm.FindByUser(tc.userID)
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantRoom, room)
		})
	}
}

func TestRoomGorm_Create(t *testing.T) {
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

	roomGorm := roomGorm{db: db}

	// create data
	migrate(db, t)

	testCases := []struct {
		wantRoom *model.Room
		wantOK   bool
	}{
		{
			wantRoom: &model.Room{ID: 1},
			wantOK:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%t", tc.wantOK), func(t *testing.T) {
			room, ok := roomGorm.Create()
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantRoom, room)
		})
	}
}
