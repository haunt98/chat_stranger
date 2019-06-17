package service

import (
	"github.com/1612180/chat_stranger/internal/dtos"
	"github.com/1612180/chat_stranger/internal/repository"
)

type RoomService struct {
	roomRepo repository.RoomRepo
}

func NewRoomService(roomRepo repository.RoomRepo) *RoomService {
	return &RoomService{
		roomRepo: roomRepo,
	}
}

func (s *RoomService) FetchAll() ([]*dtos.RoomResponse, []error) {
	rooms, errs := s.roomRepo.FetchAll()
	if len(errs) != 0 {
		return nil, errs
	}

	var roomRess []*dtos.RoomResponse
	for _, room := range rooms {
		roomRess = append(roomRess, room.ToResponse())
	}

	return roomRess, nil
}

func (s *RoomService) Find(id int) (*dtos.RoomResponse, []error) {
	room, errs := s.roomRepo.Find(id)
	if len(errs) != 0 {
		return nil, errs
	}

	return room.ToResponse(), nil
}

func (s *RoomService) Create() (int, []error) {
	return s.roomRepo.Create()
}

func (s *RoomService) Delete(id int) []error {
	return s.roomRepo.Delete(id)
}

func (s *RoomService) FindEmpty() (int, []error) {
	return s.roomRepo.FindEmpty()
}
