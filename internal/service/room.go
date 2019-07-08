package service

import (
	"time"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/repository"
)

type RoomService struct {
	roomRepo repository.RoomRepository
}

func NewRoomService(roomRepo repository.RoomRepository) *RoomService {
	return &RoomService{roomRepo: roomRepo}
}

func (s *RoomService) Find(userID int, status string) (*model.Room, bool) {
	if status == "empty" {
		room, ok := s.roomRepo.FindEmpty()
		if !ok {
			return s.roomRepo.Create()
		}
		return room, true
	} else if status == "next" {
		room, ok := s.roomRepo.FindByUserID(userID)
		if !ok {
			return nil, false
		}

		room, ok = s.roomRepo.FindNext(room.ID)
		if !ok {
			return s.roomRepo.Create()
		}
		return room, true
	}
	return nil, false
}

func (s *RoomService) Join(userID, roomID int) bool {
	if ok := s.roomRepo.IsUserFree(userID); !ok {
		return false
	}

	if ok := s.roomRepo.Exist(roomID); !ok {
		return false
	}

	if ok := s.roomRepo.IsEmpty(roomID); !ok {
		return false
	}

	if ok := s.roomRepo.DropMessages(roomID); !ok {
		return false
	}

	return s.roomRepo.Join(userID, roomID)
}

func (s *RoomService) Leave(userID int) bool {
	room, ok := s.roomRepo.FindByUserID(userID)
	if !ok {
		return false
	}

	if ok := s.roomRepo.DropMessages(room.ID); !ok {
		return false
	}

	return s.roomRepo.Leave(userID)
}

func (s *RoomService) SendMessage(message *model.Message) bool {
	room, ok := s.roomRepo.FindByUserID(message.UserID)
	if !ok {
		return false
	}

	message.RoomID = room.ID
	if ok := s.roomRepo.CreateMessage(message); !ok {
		return false
	}
	return true
}

func (s *RoomService) ReceiveMessage(userID int, fromTime time.Time) ([]*model.Message, bool) {
	room, ok := s.roomRepo.FindByUserID(userID)
	if !ok {
		return nil, false
	}

	messages, ok := s.roomRepo.LatestMessage(room.ID, fromTime)
	if !ok {
		return nil, false
	}
	return messages, true
}

func (s *RoomService) IsUserFree(userID int) bool {
	return s.roomRepo.IsUserFree(userID)
}

func (s *RoomService) CountMember(userID int) (int, bool) {
	room, ok := s.roomRepo.FindByUserID(userID)
	if !ok {
		return 0, false
	}

	return s.roomRepo.CountMember(room.ID)
}
