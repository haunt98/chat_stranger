package repository

import (
	"github.com/1612180/chat_stranger/internal/models"
	"github.com/jinzhu/gorm"
)

type MessageRepoGorm struct {
	db *gorm.DB
}

func NewMessageRepoGorm(db *gorm.DB) MessageRepo {
	db.DropTableIfExists(&models.Message{})
	db.AutoMigrate(&models.Message{})

	return &MessageRepoGorm{db: db}
}

func (g *MessageRepoGorm) FetchAll(roomid int) ([]*models.Message, []error) {
	var msgs []*models.Message

	if errs := g.db.Where("room_id = ?", roomid).Find(&msgs).GetErrors(); len(errs) != 0 {
		return nil, errs
	}

	return msgs, nil
}

func (g *MessageRepoGorm) FetchLatest(roomid, latest int) (*models.Message, int, []error) {
	var msgs []*models.Message
	var count int

	if errs := g.db.Where("room_id = ?", roomid).Find(&msgs).Count(&count).GetErrors(); len(errs) != 0 {
		return nil, -1, errs
	}

	// client already has latest msg
	if latest+1 == count {
		return nil, latest, nil
	}

	// any user join or leave
	if latest+1 > count {
		return nil, -1, nil
	}

	// sort by created time
	if errs := g.db.Where("room_id = ?", roomid).Order("created_at").Find(&msgs).GetErrors(); len(errs) != 0 {
		return nil, -1, errs
	}

	latest += 1
	return msgs[latest], latest, nil
}

func (g *MessageRepoGorm) Create(msg *models.Message) []error {
	if errs := g.db.Create(msg).GetErrors(); len(errs) != 0 {
		return errs
	}

	return nil
}
