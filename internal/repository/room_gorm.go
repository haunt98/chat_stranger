package repository

import (
	"fmt"

	"github.com/1612180/chat_stranger/internal/models"
	"github.com/jinzhu/gorm"
)

type RoomRepoGorm struct {
	db *gorm.DB
}

func NewRoomRepoGorm(db *gorm.DB) RoomRepo {
	db.DropTableIfExists(&models.Room{})
	db.DropTableIfExists("room_users")
	db.AutoMigrate(&models.Room{})

	return &RoomRepoGorm{
		db: db,
	}
}

func (g *RoomRepoGorm) Limit() int {
	return 2
}

func (g *RoomRepoGorm) FetchAll() ([]*models.Room, []error) {
	var rooms []*models.Room
	if errs := g.db.Find(&rooms).GetErrors(); len(errs) != 0 {
		return nil, errs
	}

	return rooms, nil
}

func (g *RoomRepoGorm) Find(id int) (*models.Room, []error) {
	var room models.Room
	if errs := g.db.Where("id = ?", id).First(&room).GetErrors(); len(errs) != 0 {
		return nil, errs
	}

	return &room, nil
}

func (g *RoomRepoGorm) Create() (int, []error) {
	room := models.Room{}
	if errs := g.db.Create(&room).GetErrors(); len(errs) != 0 {
		return 0, nil
	}

	return room.ID, nil
}

func (g *RoomRepoGorm) Delete(id int) []error {
	tx := g.db.Begin()

	var room models.Room

	if errs := tx.Where("id = ?", id).First(&room).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	if errs := tx.Delete(&room).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	tx.Commit()
	return nil
}

func (g *RoomRepoGorm) FindEmpty() (int, []error) {
	var rooms []*models.Room
	if errs := g.db.Find(&rooms).GetErrors(); len(errs) != 0 {
		return 0, errs
	}

	for _, room := range rooms {
		var users []*models.User
		if errs := g.db.Model(room).Related(&users, "Users").GetErrors(); len(errs) != 0 {
			return 0, errs
		}
		if len(users) < g.Limit() {
			return room.ID, nil
		}
	}

	return g.Create()
}

func (g *RoomRepoGorm) NextEmpty(userid, oldroomid int) (int, []error) {
	// leave old room
	if errs := g.Leave(userid, oldroomid); len(errs) != 0 {
		return 0, nil
	}

	// get all rooms
	var rooms []*models.Room
	if errs := g.db.Find(&rooms).GetErrors(); len(errs) != 0 {
		return 0, errs
	}

	for _, room := range rooms {
		var users []*models.User
		if errs := g.db.Model(room).Related(&users, "Users").GetErrors(); len(errs) != 0 {
			return 0, errs
		}
		if len(users) < g.Limit() && room.ID != oldroomid {
			return room.ID, nil
		}
	}

	return g.Create()
}

func (g *RoomRepoGorm) Join(userid, roomid int) []error {
	// get room
	var room models.Room
	if errs := g.db.Where("id = ?", roomid).First(&room).GetErrors(); len(errs) != 0 {
		return errs
	}

	// get all users in room
	if errs := g.db.Model(&room).Related(&room.Users, "Users").GetErrors(); len(errs) != 0 {
		return errs
	}

	// if room is full
	if len(room.Users) >= g.Limit() {
		err := fmt.Errorf("room %d is full", room.ID)
		var errs []error
		errs = append(errs, err)
		return errs
	}

	// get user to join room
	var user models.User
	if errs := g.db.Where("id = ?", userid).First(&user).GetErrors(); len(errs) != 0 {
		return errs
	}

	// join
	if err := g.db.Model(&room).Association("Users").Append(&user).Error; err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	// delete all message in room
	if errs := g.db.Where("room_id = ?", roomid).Delete(models.Message{}).GetErrors(); len(errs) != 0 {
		return errs
	}

	return nil
}

func (g *RoomRepoGorm) Leave(userid, roomid int) []error {
	// get room
	var room models.Room
	if errs := g.db.Where("id = ?", roomid).First(&room).GetErrors(); len(errs) != 0 {
		return errs
	}

	// get user who leave
	var user models.User
	if errs := g.db.Where("id = ?", userid).First(&user).GetErrors(); len(errs) != 0 {
		return errs
	}

	// delete association betwwen user who leave and room
	if err := g.db.Model(&room).Association("Users").Delete(&user).Error; err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	// delete all message in room
	if errs := g.db.Where("room_id = ?", roomid).Delete(models.Message{}).GetErrors(); len(errs) != 0 {
		return errs
	}

	return nil
}

func (g *RoomRepoGorm) Check(userid, roomid int) []error {
	// get room
	var room models.Room
	if errs := g.db.Where("id = ?", roomid).First(&room).GetErrors(); len(errs) != 0 {
		return errs
	}

	// get all users in room
	if errs := g.db.Model(&room).Related(&room.Users, "Users").GetErrors(); len(errs) != 0 {
		return errs
	}

	// get user
	var user models.User
	if errs := g.db.Where("id = ?", userid).First(&user).GetErrors(); len(errs) != 0 {
		return errs
	}

	// check user in room
	for _, u := range room.Users {
		if u.ID == user.ID {
			return nil
		}
	}

	err := fmt.Errorf("user %d not in room %d", userid, roomid)
	var errs []error
	errs = append(errs, err)
	return errs
}
