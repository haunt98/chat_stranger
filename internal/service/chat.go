package service

import (
	"time"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/1612180/chat_stranger/internal/repository"
	"github.com/pkg/errors"
)

//go:generate $GOPATH/bin/mockgen -destination=../mock/mock_service/mock_chat.go -source=chat.go

type ChatService interface {
	FindAnyRoom(userID int) (model.Room, error)
	FindNextRoom(userID int) (model.Room, error)
	FindSameGenderRoom(userID int) (model.Room, error)
	FindSameBirthYearRoom(userID int) (model.Room, error)
	Join(userID, roomID int) error
	Leave(userID int) error
	SendMessage(userID int, body string) error
	ReceiveMessage(userID int, from time.Time) ([]model.Message, error)
	IsUserFree(userID int) (bool, error)
	CountMembersInRoomOfUser(userID int) (int, error)
}

func NewChatService(accountRepo repository.AccountRepo, chatRepo repository.ChatRepo) ChatService {
	return &defautChatService{accountRepo: accountRepo, chatRepo: chatRepo}
}

// implement

type defautChatService struct {
	accountRepo repository.AccountRepo
	chatRepo    repository.ChatRepo
}

func (s *defautChatService) FindAnyRoom(userID int) (model.Room, error) {
	rooms, err := s.chatRepo.FindRooms()
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find any room userID=%d", userID)
	}

	for _, room := range rooms {
		count, err := s.chatRepo.CountMembersInRoom(room.ID)
		if err != nil {
			return model.Room{}, errors.Wrapf(err, "chat service: find any room userID=%d", userID)
		}
		if count < variable.LimitRoom {
			return room, nil
		}
	}

	room, err := s.chatRepo.CreateRoom()
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find any room userID=%d", userID)
	}
	return room, nil
}

func (s *defautChatService) FindNextRoom(userID int) (model.Room, error) {
	oldRoom, err := s.chatRepo.FindRoomOfUser(userID)
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find next room userID=%d", userID)
	}

	rooms, err := s.chatRepo.FindRooms()
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find next room userID=%d", userID)
	}

	for _, room := range rooms {
		if room.ID == oldRoom.ID {
			continue
		}

		count, err := s.chatRepo.CountMembersInRoom(room.ID)
		if err != nil {
			return model.Room{}, errors.Wrapf(err, "chat service: find next room userID=%d", userID)
		}
		if count < variable.LimitRoom {
			return room, nil
		}
	}

	room, err := s.chatRepo.CreateRoom()
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find next room userID=%d", userID)
	}
	return room, nil
}

func (s *defautChatService) FindSameGenderRoom(userID int) (model.Room, error) {
	curUser, _, err := s.accountRepo.FindUserCredential(userID)
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find same gender room userID=%d", userID)
	}

	oldRoom, err := s.chatRepo.FindRoomOfUser(userID)
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find same gender room userID=%d", userID)
	}

	rooms, err := s.chatRepo.FindRooms()
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find same gender room userID=%d", userID)
	}

	for _, room := range rooms {
		if room.ID == oldRoom.ID {
			continue
		}

		count, err := s.chatRepo.CountMembersInRoom(room.ID)
		if err != nil {
			return model.Room{}, errors.Wrapf(err, "chat service: find same gender room userID=%d", userID)
		}
		if count >= variable.LimitRoom {
			continue
		}

		users, err := s.chatRepo.FindUsersInRoom(room.ID)
		if err != nil {
			return model.Room{}, errors.Wrapf(err, "chat service: find same gender room userID=%d", userID)
		}
		for _, user := range users {
			if user.Gender == curUser.Gender {
				return room, nil
			}
		}
	}

	room, err := s.chatRepo.CreateRoom()
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find same gender room userID=%d", userID)
	}
	return room, nil
}

func (s *defautChatService) FindSameBirthYearRoom(userID int) (model.Room, error) {
	curUser, _, err := s.accountRepo.FindUserCredential(userID)
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find same birth year room userID=%d", userID)
	}

	oldRoom, err := s.chatRepo.FindRoomOfUser(userID)
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find same birth year room userID=%d", userID)
	}

	rooms, err := s.chatRepo.FindRooms()
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find same birth year room userID=%d", userID)
	}

	for _, room := range rooms {
		if room.ID == oldRoom.ID {
			continue
		}

		count, err := s.chatRepo.CountMembersInRoom(room.ID)
		if err != nil {
			return model.Room{}, errors.Wrapf(err, "chat service: find same birth year room userID=%d", userID)
		}
		if count >= variable.LimitRoom {
			continue
		}

		users, err := s.chatRepo.FindUsersInRoom(room.ID)
		if err != nil {
			return model.Room{}, errors.Wrapf(err, "chat service: find same birth year room userID=%d", userID)
		}
		for _, user := range users {
			if user.BirthYear == curUser.BirthYear {
				return room, nil
			}
		}
	}

	room, err := s.chatRepo.CreateRoom()
	if err != nil {
		return model.Room{}, errors.Wrapf(err, "chat service: find same birth year room userID=%d", userID)
	}
	return room, nil
}

func (s *defautChatService) Join(userID, roomID int) error {
	if err := s.chatRepo.DeleteMessagesOfRoom(roomID); err != nil {
		return errors.Wrapf(err, "chat service: join userID=%d roomID=%d failed", userID, roomID)
	}

	_, err := s.chatRepo.CreateMember(userID, roomID)
	if err != nil {
		return errors.Wrapf(err, "chat service: join userID=%d roomID=%d failed", userID, roomID)
	}
	return nil
}

func (s *defautChatService) Leave(userID int) error {
	oldRoom, err := s.chatRepo.FindRoomOfUser(userID)
	if err != nil {
		return errors.Wrapf(err, "chat service: leave userID=%d failed", userID)
	}

	if err := s.chatRepo.DeleteMessagesOfRoom(oldRoom.ID); err != nil {
		return errors.Wrapf(err, "chat service: leave userID=%d failed", userID)
	}

	if err := s.chatRepo.DeleteMember(userID, oldRoom.ID); err != nil {
		return errors.Wrapf(err, "chat service: leave userID=%d failed", userID)
	}
	return nil
}

func (s *defautChatService) SendMessage(userID int, body string) error {
	room, err := s.chatRepo.FindRoomOfUser(userID)
	if err != nil {
		return errors.Wrapf(err, "chat service: send message userID=%d failed", userID)
	}

	_, err = s.chatRepo.CreateMessage(room.ID, userID, body)
	if err != nil {
		return errors.Wrapf(err, "chat service: send message userID=%d failed", userID)
	}
	return nil
}

func (s *defautChatService) ReceiveMessage(userID int, from time.Time) ([]model.Message, error) {
	room, err := s.chatRepo.FindRoomOfUser(userID)
	if err != nil {
		return nil, errors.Wrapf(err, "chat service: receive message userID=%d failed", userID)
	}

	msgs, err := s.chatRepo.FindMessagesOfRoomFromTime(room.ID, from)
	if err != nil {
		return nil, errors.Wrapf(err, "chat service: receive message userID=%d failed", userID)
	}
	return msgs, nil
}

func (s *defautChatService) IsUserFree(userID int) (bool, error) {
	count, err := s.chatRepo.CountMembersOfUser(userID)
	if err != nil {
		return false, errors.Wrap(err, "chat service: is user free failed")
	}

	if count > 0 {
		return false, nil
	}
	return true, nil
}

func (s *defautChatService) CountMembersInRoomOfUser(userID int) (int, error) {
	room, err := s.chatRepo.FindRoomOfUser(userID)
	if err != nil {
		return 0, errors.Wrapf(err, "chat service: count members in room of user userID=%d failed", userID)
	}

	count, err := s.chatRepo.CountMembersInRoom(room.ID)
	if err != nil {
		return 0, errors.Wrapf(err, "chat service: count members in room of user userID=%d failed", userID)
	}
	return count, nil
}
