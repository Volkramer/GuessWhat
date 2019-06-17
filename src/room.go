package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Room Object
type Room struct {
	Name  string
	Users map[string]*User
	Join  chan *User
	Leave chan *User
	Input chan *Message
}

var rooms = make(map[string]*Room)

//AddUser method forward a user to the channel join
func (room *Room) AddUser(user *User) {
	log.Println("SYSTEM:", user.username, "has joined the room", room.Name)
	room.Join <- user
	user.room = room
}

//RemoveUser Method forward a user to the channel Leave
func (room *Room) RemoveUser(user *User) {
	log.Println("SYSTEM:", user.username, "has left the room", room.Name)
	room.Leave <- user
	user.room = nil
}

//InputMessage Method forward a message to the input channel
func (room *Room) InputMessage(message *Message) {
	log.Println("SYSTEM:", message.Username, "send:", message.Text)
	room.Input <- message
}

//sendAll Method forward message to all user
func (room *Room) sendAll(message *Message) {
	for _, user := range room.Users {
		user.Write(message)
	}
}

func (room *Room) run() {
	log.Println("SYSTEM: Room", room.Name, "successfully started")
	for {
		select {
		case user := <-room.Join:
			room.Users[user.username] = user
		case user := <-room.Leave:
			delete(room.Users, user.username)
			if len(room.Users) == 0 {
				delete(rooms, room.Name)
				close(room.Join)
				close(room.Leave)
				close(room.Input)
				log.Println("SYSTEM: Room", room.Name, "closed")
			}
		case msg := <-room.Input:
			room.sendAll(msg)
		}
	}
}

//NewRoom method
func NewRoom(name string) (room *Room) {
	return &Room{
		Name:  name,
		Users: make(map[string]*User),
		Join:  make(chan *User),
		Leave: make(chan *User),
		Input: make(chan *Message),
	}
}

func getRooms() (roomsName string) {
	var roomsSlice []string
	for _, room := range rooms {
		roomsSlice = append(roomsSlice, room.Name)
	}
	roomsJSON, err := json.Marshal(roomsSlice)
	Error(err)
	roomsName = fmt.Sprint(string(roomsJSON))
	return roomsName
}
