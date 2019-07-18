package repository

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/variable"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//go:generate $GOPATH/bin/mockgen -destination=../mock/mock_repository/mock_room.go -source=room.go

type RoomRepository interface {
	Exist(id int) bool
	FindEmpty() (*model.Room, bool)
	FindNext(old int) (*model.Room, bool)
	FindSameGender(old int, gender string) (*model.Room, bool)
	FindSameBirthYear(old int, year int) (*model.Room, bool)
	FindByUser(userID int) (*model.Room, bool)
	Create() (*model.Room, bool)
}

// implement

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomGorm{db: db}
}

type roomGorm struct {
	db *gorm.DB
}

func (g *roomGorm) Exist(id int) bool {
	if err := g.db.Where("id = ?", id).First(&model.Room{}).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "exist",
		}).Error(err)
		return false
	}
	return true
}

func (g *roomGorm) fetchAll() ([]*model.Room, bool) {
	var rooms []*model.Room
	if err := g.db.Find(&rooms).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "fetch all",
		}).Error(err)
		return nil, false
	}
	return rooms, true
}

func (g *roomGorm) FindEmpty() (*model.Room, bool) {
	rooms, ok := g.fetchAll()
	if !ok {
		return nil, false
	}

	// find empty room
	for _, room := range rooms {
		count, ok := countByRoom(g.db, room.ID)
		if !ok {
			return nil, false
		}
		if count < variable.LimitRoom {
			return room, true
		}
	}
	return nil, false
}

func (g *roomGorm) FindNext(old int) (*model.Room, bool) {
	rooms, ok := g.fetchAll()
	if !ok {
		return nil, false
	}

	// find empty room and not old room
	for _, room := range rooms {
		count, ok := countByRoom(g.db, room.ID)
		if !ok {
			return nil, false
		}
		if count < variable.LimitRoom && room.ID != old {
			return room, true
		}
	}
	return nil, false
}

func (g *roomGorm) FindSameGender(old int, gender string) (*model.Room, bool) {
	rooms, ok := g.fetchAll()
	if !ok {
		return nil, false
	}

	// find empty room and same gender
	for _, room := range rooms {
		count, ok := countByRoom(g.db, room.ID)
		if !ok {
			return nil, false
		}

		if count >= variable.LimitRoom {
			return nil, false
		}

		if room.ID == old {
			return nil, false
		}

		// members of room
		var members []*model.Member
		if err := g.db.Where("room_id = ?", room.ID).Find(&members).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"event":  "repo",
				"target": "room",
				"action": "find same gender",
			}).Error(err)
			return nil, false
		}

		// check gender of mebers
		for _, member := range members {
			var user model.User
			if err := g.db.Where(&model.User{ID: member.UserID}).Error; err != nil {
				logrus.WithFields(logrus.Fields{
					"event":  "repo",
					"target": "room",
					"action": "find same gender",
				}).Error(err)
				return nil, false
			}

			if user.Gender == gender {
				return room, true
			}
		}
	}
	return nil, false
}

func (g *roomGorm) FindSameBirthYear(old int, year int) (*model.Room, bool) {
	rooms, ok := g.fetchAll()
	if !ok {
		return nil, false
	}

	// find empty room and same birth year
	for _, room := range rooms {
		count, ok := countByRoom(g.db, room.ID)
		if !ok {
			return nil, false
		}

		if count >= variable.LimitRoom {
			return nil, false
		}

		if room.ID == old {
			return nil, false
		}

		// members of room
		var members []*model.Member
		if err := g.db.Where("room_id = ?", room.ID).Find(&members).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"event":  "repo",
				"target": "room",
				"action": "find same birth year",
			}).Error(err)
			return nil, false
		}

		// check gender of mebers
		for _, member := range members {
			var user model.User
			if err := g.db.Where(&model.User{ID: member.UserID}).Error; err != nil {
				logrus.WithFields(logrus.Fields{
					"event":  "repo",
					"target": "room",
					"action": "find same birth year",
				}).Error(err)
				return nil, false
			}

			if user.BirthYear == year {
				return room, true
			}
		}
	}
	return nil, false
}

func (g *roomGorm) FindByUser(userID int) (*model.Room, bool) {
	// find members
	var members []*model.Member
	if err := g.db.Where("user_id = ?", userID).
		Find(&members).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "find by user",
		}).Error(err)
		return nil, false
	}

	// make sure user in 1 room
	if len(members) != 1 {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "find by user",
		}).Error("user not in 1 room")
		return nil, false
	}

	// find room
	var room model.Room
	if err := g.db.Where("id = ?", members[0].RoomID).
		First(&room).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "find by user",
		}).Error(err)
		return nil, false
	}
	return &room, true
}

func (g *roomGorm) Create() (*model.Room, bool) {
	var room model.Room
	if err := g.db.Create(&room).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "create",
		}).Error(err)
		return nil, false
	}
	return &room, true
}
