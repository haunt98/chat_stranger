package service

import (
	"log"
	"sync"
)

type Room struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	mux        sync.Mutex
}

func NewRoom() *Room {
	return &Room{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (r *Room) Run() {
	for {
		select {
		case c := <-r.Register:
			for oldC := range r.Clients {
				if err := oldC.Conn.WriteJSON(Message{
					Fullname: "Server",
					Body:     "User has joined",
				}); err != nil {
					log.Println(err)
					delete(r.Clients, oldC)
					if err := oldC.Conn.Close(); err != nil {
						log.Println(err)
					}
				}
			}
			r.Clients[c] = true
		case c := <-r.Unregister:
			delete(r.Clients, c)
			if err := c.Conn.Close(); err != nil {
				log.Println(err)
			}
			for oldC := range r.Clients {
				if err := oldC.Conn.WriteJSON(Message{
					Fullname: "Server",
					Body:     "User has leaved",
				}); err != nil {
					log.Println(err)
					delete(r.Clients, oldC)
					if err := oldC.Conn.Close(); err != nil {
						log.Println(err)
					}
				}
			}
		case message := <-r.Broadcast:
			for c := range r.Clients {
				if err := c.Conn.WriteJSON(message); err != nil {
					log.Println(err)
					delete(r.Clients, c)
					if err := c.Conn.Close(); err != nil {
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

	return len(r.Clients) >= 2
}
