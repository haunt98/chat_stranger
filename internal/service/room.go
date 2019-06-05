package service

import (
	"fmt"
	"log"
	"sync"

	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/repository"
)

type Room struct {
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
	broadcast  chan Message
	mux        sync.Mutex
	repo       repository.UserRepo
}

func NewRoom(repo repository.UserRepo) *Room {
	r := &Room{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message),
		repo:       repo,
	}

	go r.Run()

	return r
}

type RoomService struct {
	rooms map[int]*Room
	mux   sync.Mutex
	repo  repository.UserRepo
}

func NewRoomService(repo repository.UserRepo) *RoomService {
	return &RoomService{
		rooms: make(map[int]*Room),
		repo:  repo,
	}
}

func (r *Room) Run() {
	for {
		select {
		case c := <-r.register:
			u, err := r.repo.Find(c.id)
			if err != nil {
				log.Println(err)
				return
			}

			for oldC := range r.clients {
				if err := oldC.conn.WriteJSON(Message{
					Fullname: u.FullName,
					Body:     "has joined",
				}); err != nil {
					log.Println(err)
					delete(r.clients, oldC)
					if err := oldC.conn.Close(); err != nil {
						log.Println(err)
					}
				}
			}
			r.clients[c] = true
		case c := <-r.unregister:
			u, err := r.repo.Find(c.id)
			if err != nil {
				log.Println(err)
				return
			}

			delete(r.clients, c)
			if err := c.conn.Close(); err != nil {
				log.Println(err)
			}

			for oldC := range r.clients {
				if err := oldC.conn.WriteJSON(Message{
					Fullname: u.FullName,
					Body:     "has leaved",
				}); err != nil {
					log.Println(err)
					delete(r.clients, oldC)
					if err := oldC.conn.Close(); err != nil {
						log.Println(err)
					}
				}
			}
		case message := <-r.broadcast:
			for c := range r.clients {
				if err := c.conn.WriteJSON(message); err != nil {
					log.Println(err)
					delete(r.clients, c)
					if err := c.conn.Close(); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}

func (r *Room) IsFull() bool {
	r.mux.Lock()
	defer r.mux.Unlock()

	return len(r.clients) >= 2
}

func (s *RoomService) NextRoom(cur int) int {
	s.mux.Lock()
	defer s.mux.Unlock()

	for id := range s.rooms {
		if !s.rooms[id].IsFull() && id != cur {
			return id
		}
	}

	id := len(s.rooms) + 1
	s.rooms[id] = NewRoom(s.repo)

	return id
}

func (s *RoomService) JoinRoom() int {
	s.mux.Lock()
	defer s.mux.Unlock()

	for id := range s.rooms {
		if !s.rooms[id].IsFull() {
			return id
		}
	}

	id := len(s.rooms) + 1
	s.rooms[id] = NewRoom(s.repo)

	return id
}

func (s *RoomService) CheckRoom(id int) error {
	if _, ok := s.rooms[id]; !ok {
		return fmt.Errorf(response.ResponseCode[411])
	}

	if s.rooms[id].IsFull() {
		return fmt.Errorf(response.ResponseCode[410])
	}

	return nil
}

func (s *RoomService) GetRoom(id int) *Room {
	return s.rooms[id]
}
