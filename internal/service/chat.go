package service

import (
	"time"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/1612180/chat_stranger/internal/repository"

	"github.com/sirupsen/logrus"
)

type ChatService struct {
	roomRepo    repository.RoomRepository
	memberRepo  repository.MemberRepo
	messageRepo repository.MessageRepo
}

func NewChatService(
	roomRepo repository.RoomRepository,
	memberRepo repository.MemberRepo,
	messageRepo repository.MessageRepo,
) *ChatService {
	return &ChatService{roomRepo: roomRepo, memberRepo: memberRepo, messageRepo: messageRepo}
}

func (s *ChatService) Find(userID int, status string) (*model.Room, bool) {
	if status == "empty" {
		room, ok := s.roomRepo.FindEmpty()
		if !ok {
			return s.roomRepo.Create()
		}
		return room, true
	} else if status == "next" {
		// old room
		room, ok := s.roomRepo.FindByUser(userID)
		if !ok {
			return nil, false
		}

		// next
		room, ok = s.roomRepo.FindNext(room.ID)
		if !ok {
			return s.roomRepo.Create()
		}
		return room, true
	}
	return nil, false
}

func (s *ChatService) Join(userID, roomID int) bool {
	// check user
	count, ok := s.memberRepo.CountByUser(userID)
	if !ok {
		return false
	}
	if count != 0 {
		logrus.WithFields(logrus.Fields{
			"event":  "service",
			"target": "chat",
			"action": "join",
			"userID": userID,
			"roomID": roomID,
		}).Info("user has joined another room")
		return false
	}

	// check room
	if ok := s.roomRepo.Exist(roomID); !ok {
		return false
	}

	count, ok = s.memberRepo.CountByRoom(roomID)
	if !ok {
		return false
	}
	if count >= variable.LimitRoom {
		logrus.WithFields(logrus.Fields{
			"event":  "service",
			"target": "chat",
			"action": "join",
			"userID": userID,
			"roomID": roomID,
		}).Info("room is full")
		return false
	}

	// delete old messages
	if ok := s.messageRepo.Delete(roomID); !ok {
		return false
	}
	return s.memberRepo.Create(userID, roomID)
}

func (s *ChatService) Leave(userID int) bool {
	// find room
	room, ok := s.roomRepo.FindByUser(userID)
	if !ok {
		return false
	}

	// delete old messages
	if ok := s.messageRepo.Delete(room.ID); !ok {
		return false
	}
	return s.memberRepo.Delete(userID)
}

func (s *ChatService) SendMessage(message *model.Message) bool {
	// find room
	room, ok := s.roomRepo.FindByUser(message.UserID)
	if !ok {
		return false
	}

	// create message
	message.RoomID = room.ID
	if ok := s.messageRepo.Create(message); !ok {
		return false
	}
	return true
}

func (s *ChatService) ReceiveMessage(userID int, fromTime time.Time) ([]*model.Message, bool) {
	// find room
	room, ok := s.roomRepo.FindByUser(userID)
	if !ok {
		return nil, false
	}

	// fetch messages
	messages, ok := s.messageRepo.FetchByTime(room.ID, fromTime)
	if !ok {
		return nil, false
	}
	return messages, true
}

func (s *ChatService) IsUserFree(userID int) bool {
	count, ok := s.memberRepo.CountByUser(userID)
	if !ok {
		return false
	}
	if count != 0 {
		logrus.WithFields(logrus.Fields{
			"event":  "service",
			"target": "chat",
			"action": "is user free",
		}).Info("user has joined another room")
		return false
	}
	return true
}

func (s *ChatService) CountMember(userID int) (int, bool) {
	// find room
	room, ok := s.roomRepo.FindByUser(userID)
	if !ok {
		return 0, false
	}
	return s.memberRepo.CountByRoom(room.ID)
}
