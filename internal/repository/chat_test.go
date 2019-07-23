package repository

import (
	"fmt"
	"testing"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/config"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestNewChatRepo(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	assert.Equal(t, &defaultChatRepo{db: db}, NewChatRepo(db))
}

func TestDefaultChatRepo_CreateRoom(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	chatRepo := defaultChatRepo{db: db}

	migrate(db, t)

	testCases := []struct {
		outRoom model.Room
		ok      bool
	}{
		{
			outRoom: model.Room{ID: 1},
			ok:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			room, err := chatRepo.CreateRoom()
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outRoom, room)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultChatRepo_FindRooms(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	chatRepo := defaultChatRepo{db: db}

	migrate(db, t)
	tempRoom := model.Room{}
	if err := db.Create(&tempRoom).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		outRooms []model.Room
		ok       bool
	}{
		{
			outRooms: []model.Room{tempRoom},
			ok:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			rooms, err := chatRepo.FindRooms()
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outRooms, rooms)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultChatRepo_FindRoomOfUser(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	chatRepo := defaultChatRepo{db: db}

	migrate(db, t)
	var tempUser model.User
	if err := db.Create(&tempUser).Error; err != nil {
		t.Error(err)
	}
	var tempRoom model.Room
	if err := db.Create(&tempRoom).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.Member{UserID: tempUser.ID, RoomID: tempRoom.ID}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inUserID int
		outRoom  model.Room
		ok       bool
	}{
		{
			inUserID: tempUser.ID,
			outRoom:  tempRoom,
			ok:       true,
		},
		{
			inUserID: 0,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			room, err := chatRepo.FindRoomOfUser(tc.inUserID)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outRoom, room)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultChatRepo_CountMembersInRoom(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	chatRepo := defaultChatRepo{db: db}

	migrate(db, t)
	tempUser := model.User{}
	if err := db.Create(&tempUser).Error; err != nil {
		t.Error(err)
	}
	tempRoom := model.Room{}
	if err := db.Create(&tempRoom).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.Member{RoomID: tempRoom.ID, UserID: tempUser.ID}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inRoomID int
		outCount int
		ok       bool
	}{
		{
			inRoomID: tempRoom.ID,
			outCount: 1,
			ok:       true,
		},
		{
			inRoomID: 0,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			count, err := chatRepo.CountMembersInRoom(tc.inRoomID)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outCount, count)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultChatRepo_FindUsersInRoom(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	chatRepo := defaultChatRepo{db: db}

	migrate(db, t)
	tempUser := model.User{ShowName: "a"}
	if err := db.Create(&tempUser).Error; err != nil {
		t.Error(err)
	}
	tempRoom := model.Room{}
	if err := db.Create(&tempRoom).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.Member{RoomID: tempRoom.ID, UserID: tempUser.ID}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inRoomID int
		outUsers []model.User
		ok       bool
	}{
		{
			inRoomID: tempRoom.ID,
			outUsers: []model.User{tempUser},
			ok:       true,
		},
		{
			inRoomID: 0,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			users, err := chatRepo.FindUsersInRoom(tc.inRoomID)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outUsers, users)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultChatRepo_CreateMember(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	chatRepo := defaultChatRepo{db: db}

	migrate(db, t)
	tempUser := model.User{}
	if err := db.Create(&tempUser).Error; err != nil {
		t.Error(err)
	}
	tempRoom := model.Room{}
	if err := db.Create(&tempRoom).Error; err != nil {
		t.Error(err)
	}
	tempRoom2 := model.Room{}
	if err := db.Create(&tempRoom2).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.Member{UserID: tempUser.ID, RoomID: tempRoom2.ID}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inUserID  int
		inRoomID  int
		outMember model.Member
		ok        bool
	}{
		{
			inUserID:  tempUser.ID,
			inRoomID:  tempRoom.ID,
			outMember: model.Member{UserID: tempUser.ID, RoomID: tempRoom.ID},
			ok:        true,
		},
		{
			inUserID:  tempUser.ID,
			inRoomID:  tempRoom2.ID,
			outMember: model.Member{UserID: tempUser.ID, RoomID: tempRoom2.ID},
			ok:        false,
		},
		{
			inUserID: 0,
			inRoomID: tempRoom.ID,
			ok:       false,
		},
		{
			inUserID: tempUser.ID,
			inRoomID: 0,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			member, err := chatRepo.CreateMember(tc.inUserID, tc.inRoomID)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outMember.UserID, member.UserID)
				assert.Equal(t, tc.outMember.RoomID, member.RoomID)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultChatRepo_DeleteMember(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	chatRepo := defaultChatRepo{db: db}

	migrate(db, t)
	tempUser := model.User{}
	if err := db.Create(&tempUser).Error; err != nil {
		t.Error(err)
	}
	tempRoom := model.Room{}
	if err := db.Create(&tempRoom).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.Member{UserID: tempUser.ID, RoomID: tempRoom.ID}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inUserID int
		inRoomID int
		ok       bool
	}{
		{
			inUserID: tempUser.ID,
			inRoomID: tempRoom.ID,
			ok:       true,
		},
		{
			inUserID: 0,
			inRoomID: tempRoom.ID,
			ok:       false,
		},
		{
			inUserID: tempUser.ID,
			inRoomID: 0,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			err := chatRepo.DeleteMember(tc.inUserID, tc.inRoomID)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, gorm.ErrRecordNotFound, db.Where("user_id = ? AND room_id = ?", tc.inUserID, tc.inRoomID).First(&model.Member{}).Error)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultChatRepo_CreateMessage(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	chatRepo := defaultChatRepo{db: db}

	migrate(db, t)
	tempUser := model.User{}
	if err := db.Create(&tempUser).Error; err != nil {
		t.Error(err)
	}
	tempRoom := model.Room{}
	if err := db.Create(&tempRoom).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inRoomID int
		inUserID int
		inBody   string
		outMsg   model.Message
		ok       bool
	}{
		{
			inRoomID: tempRoom.ID,
			inUserID: tempUser.ID,
			inBody:   "hi",
			outMsg:   model.Message{RoomID: 1, UserID: tempUser.ID, UserShowName: tempUser.ShowName, Body: "hi"},
			ok:       true,
		},
		{
			inRoomID: tempRoom.ID,
			inUserID: 0,
			ok:       false,
		},
		{
			inRoomID: 0,
			inUserID: tempUser.ID,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			msg, err := chatRepo.CreateMessage(tc.inRoomID, tc.inUserID, tc.inBody)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outMsg.RoomID, msg.RoomID)
				assert.Equal(t, tc.outMsg.UserID, msg.UserID)
				assert.Equal(t, tc.outMsg.UserShowName, msg.UserShowName)
				assert.Equal(t, tc.outMsg.Body, msg.Body)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultChatRepo_DeleteMessagesOfRoom(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	chatRepo := defaultChatRepo{db: db}

	migrate(db, t)
	tempRoom := model.Room{}
	if err := db.Create(&tempRoom).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.Message{RoomID: tempRoom.ID}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inRoomID int
		ok       bool
	}{
		{
			inRoomID: tempRoom.ID,
			ok:       true,
		},
		{
			inRoomID: 0,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			err := chatRepo.DeleteMessagesOfRoom(tc.inRoomID)
			if tc.ok {
				assert.Nil(t, err)
				var msgs []model.Message
				if err := db.Where("room_id = ?", tc.inRoomID).Error; err != nil {
					t.Error(err)
				}
				assert.Equal(t, 0, len(msgs))
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
