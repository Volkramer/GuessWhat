package main

import "encoding/json"

//Room object
type Room struct {
	name  string
	users []*User
	join  chan *User
	leave chan *User
}

var rooms []string

func newRoom(name string) *Room {
	rooms = append(rooms, name)
	return &Room{
		name:  name,
		users: make([]*User, 0),
		join:  make(chan *User),
		leave: make(chan *User),
	}
}

func (room *Room) run() {
	for {
		select {
		case user := <-room.join:
			room.onConnect(user)
		case user := <-room.leave:
			room.onDisconnect(user)
		}
	}
}

func (room *Room) send(message interface{}, user *User) {
	data, _ := json.Marshal(message)
	user.output <- data
}

func (room *Room) broadcast(message interface{}, ignore *User) {
	data, _ := json.Marshal(message)
	for _, user := range room.users {
		if user != ignore {
			user.output <- data
		}
	}
}
