package main

import (
	"encoding/json"
	"log"

	"github.com/tidwall/gjson"
)

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

func (room *Room) onConnect(user *User) {
	log.Println("client connected: ", user.socket.RemoteAddr())
	users := []UserJSON{}
	for _, u := range room.users {
		users = append(users, UserJSON{Username: u.username})
	}
	room.send(NewConnected(user.username, users), user)
	room.broadcast(NewUserJoined(user.username), user)
}

func (room *Room) onDisconnect(user *User) {
	log.Println("client disconnected: ", user.socket.RemoteAddr())
	user.close()
	i := -1
	for j, c := range room.users {
		if c.username == user.username {
			i = j
			break
		}
	}
	copy(room.users[i:], room.users[i+1:])
	room.users[len(room.users)-1] = nil
	room.users = room.users[:len(room.users)-1]
	room.broadcast(NewUserLeft(user.username), nil)
}

func (room *Room) onMessage(data []byte, user *User) {
	event := gjson.GetBytes(data, "event").Int()
	if event == EventStroke {
		var msg StrokeJSON
		if json.Unmarshal(data, &msg) != nil {
			return
		}
		msg.Username = user.username
		room.broadcast(msg, user)
	} else if event == EventClear {
		var msg ClearJSON
		if json.Unmarshal(data, &msg) != nil {
			return
		}
		msg.Username = user.username
		room.broadcast(msg, user)
	}
}
