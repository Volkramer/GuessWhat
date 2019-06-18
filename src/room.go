package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Room Object
type Room struct {
	name  string
	users map[string]*User
	join  chan *User
	leave chan *User
	input chan *Message
}

var rooms = make(map[string]*Room)

//AddUser method forward a user to the channel join
func (room *Room) AddUser(user *User) {
	log.Println("SYSTEM:", user.username, "has joined the room", room.name)
	room.join <- user
	user.room = room
}

//RemoveUser Method forward a user to the channel Leave
func (room *Room) RemoveUser(user *User) {
	log.Println("SYSTEM:", user.username, "has left the room", room.name)
	room.leave <- user
	user.room = nil
}

//InputMessage Method forward a message to the input channel
func (room *Room) InputMessage(message *Message) {
	log.Println("SYSTEM:", message.Username, "send:", message.Text)
	room.input <- message
}

//sendAll Method forward message to all user
func (room *Room) sendAll(message *Message) {
	for _, user := range room.users {
		user.Write(message)
	}
}

func (room *Room) run() {
	log.Println("SYSTEM: Room", room.name, "successfully started")
	for {
		select {
		case user := <-room.join:
			room.users[user.username] = user
			SendData("userJoin", user.username, user.conn)
		case user := <-room.leave:
			delete(room.users, user.username)
			SendData("userLeave", user.username, user.conn)
			if len(room.users) == 0 {
				delete(rooms, room.name)
				close(room.join)
				close(room.leave)
				close(room.input)
				log.Println("SYSTEM: Room", room.name, "closed")
			}
		case msg := <-room.input:
			room.sendAll(msg)
		}
	}
}

//NewRoom method
func NewRoom(name string) (room *Room) {
	return &Room{
		name:  name,
		users: make(map[string]*User),
		join:  make(chan *User),
		leave: make(chan *User),
		input: make(chan *Message),
	}
}

func getRooms() (roomsName string) {
	var roomsSlice []string
	for _, room := range rooms {
		roomsSlice = append(roomsSlice, room.name)
	}
	roomsJSON, err := json.Marshal(roomsSlice)
	Error(err)
	roomsName = fmt.Sprint(string(roomsJSON))
	return roomsName
}
