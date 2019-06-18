package service

import (
	"fmt"
	"time"

	"github.com/1612180/chat_stranger/internal/dtos"
	"github.com/1612180/chat_stranger/internal/repository"
)

type RoomService struct {
	roomRepo repository.RoomRepo
	userRepo repository.UserRepo
	Messages map[int]chan dtos.MessageResponse
}

func NewRoomService(roomRepo repository.RoomRepo, userRepo repository.UserRepo) *RoomService {
	return &RoomService{
		roomRepo: roomRepo,
		userRepo: userRepo,
		Messages: make(map[int]chan dtos.MessageResponse),
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
	id, errs := s.roomRepo.FindEmpty()
	if len(errs) != 0 {
		return id, errs
	}

	s.Messages[id] = make(chan dtos.MessageResponse)
	return id, errs
}

func (s *RoomService) Join(uid, rid int) []error {
	return s.roomRepo.Join(uid, rid)
}

func (s *RoomService) Leave(uid, rid int) []error {
	return s.roomRepo.Leave(uid, rid)
}

func (s *RoomService) Check(uid, rid int) []error {
	return s.roomRepo.Check(uid, rid)
}

func (s *RoomService) SendMsg(rid int, timeout time.Duration) (dtos.MessageResponse, error) {
	if _, ok := s.Messages[rid]; !ok {
		return dtos.MessageResponse{Sender: "Server", Body: "Something wrong :("}, fmt.Errorf("messages in room %d failed", rid)
	}

	endtime := make(chan bool)
	go func() {
		time.Sleep(timeout)
		endtime <- true
	}()

	select {
	case msgRes := <-s.Messages[rid]:
		return msgRes, nil
	case <-endtime:
		return dtos.MessageResponse{Sender: "Server", Body: "Something wrong 2 :("}, fmt.Errorf("out of time")
	}
}

func (s *RoomService) ReceiveMsg(uid int, msgReq dtos.MessageRequest) []error {
	if _, ok := s.Messages[msgReq.Rid]; !ok {
		err := fmt.Errorf("messages in room %d failed", msgReq.Rid)
		var errs []error
		errs = append(errs, err)
		return errs
	}

	user, errs := s.userRepo.Find(uid)
	if len(errs) != 0 {
		return errs
	}

	msgRes := dtos.MessageResponse{
		Sender: user.FullName,
		Body:   msgReq.Body,
	}

	s.Messages[msgReq.Rid] <- msgRes
	return nil
}
