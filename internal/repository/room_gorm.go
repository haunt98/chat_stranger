package repository

import (
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
		if errs := g.db.Model(room).Related(&users).GetErrors(); len(errs) != 0 {
			return 0, errs
		}
		if len(users) < 2 {
			return room.ID, nil
		}
	}

	return g.Create()
}
