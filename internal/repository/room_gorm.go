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
	db.AutoMigrate(&models.Room{})

	return &RoomRepoGorm{
		db: db,
	}
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
		if len(users) < 2 {
			return room.ID, nil
		}
	}

	return g.Create()
}

func (g *RoomRepoGorm) Join(uid, rid int) []error {
	var room models.Room
	if errs := g.db.Where("id = ?", rid).First(&room).GetErrors(); len(errs) != 0 {
		return errs
	}

	if errs := g.db.Model(&room).Related(&room.Users, "Users").GetErrors(); len(errs) != 0 {
		return errs
	}

	if len(room.Users) >= 2 {
		err := fmt.Errorf("room %d is full", room.ID)
		var errs []error
		errs = append(errs, err)
		return errs
	}

	var user models.User
	if errs := g.db.Where("id = ?", uid).First(&user).GetErrors(); len(errs) != 0 {
		return errs
	}

	if err := g.db.Model(&room).Association("Users").Append(&user).Error; err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	return nil
}

func (g *RoomRepoGorm) Leave(uid, rid int) []error {
	var room models.Room
	if errs := g.db.Where("id = ?", rid).First(&room).GetErrors(); len(errs) != 0 {
		return errs
	}

	if errs := g.db.Model(&room).Related(&room.Users, "Users").GetErrors(); len(errs) != 0 {
		return errs
	}

	if len(room.Users) >= 2 {
		err := fmt.Errorf("room %d is full", room.ID)
		var errs []error
		errs = append(errs, err)
		return errs
	}

	var user models.User
	if errs := g.db.Where("id = ?", uid).First(&user).GetErrors(); len(errs) != 0 {
		return errs
	}

	if err := g.db.Model(&room).Association("Users").Delete(&user).Error; err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	if errs := g.db.Save(&room).GetErrors(); len(errs) != 0 {
		return errs
	}

	return nil
}

func (g *RoomRepoGorm) Check(uid, rid int) []error {
	var room models.Room
	if errs := g.db.Where("id = ?", rid).First(&room).GetErrors(); len(errs) != 0 {
		return errs
	}

	if errs := g.db.Model(&room).Related(&room.Users, "Users").GetErrors(); len(errs) != 0 {
		return errs
	}

	var user models.User
	if errs := g.db.Where("id = ?", uid).First(&user).GetErrors(); len(errs) != 0 {
		return errs
	}

	for _, u := range room.Users {
		if u.ID == user.ID {
			return nil
		}
	}

	err := fmt.Errorf("user %d not in room %d", uid, rid)
	var errs []error
	errs = append(errs, err)
	return errs
}
