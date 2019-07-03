package storage

import "github.com/1612180/chat_stranger/internal/model"

type User interface {
	Find(id int) (*model.User, bool)
	Save(user *model.User) bool
}
