package main

import (
	"log"

	"github.com/gorilla/websocket"
)

const channelBufSize = 100

//User Object
type User struct {
	username string
	output   chan *Message
	room     *Room
	conn     *websocket.Conn
}

//Join Object
type Join struct {
	Username string `json:"username"`
	Room     string `json:"room"`
}

//NewUser Method
func NewUser(username string, conn *websocket.Conn) (user *User) {
	return &User{
		username: username,
		output:   make(chan *Message),
		room:     nil,
		conn:     conn,
	}
}

//Write Method
func (user *User) Write(message *Message) {
	select {
	case user.output <- message:
	default:
		user.room.RemoveUser(user)
		log.Println("SYSTEM: User", user.username, "is disconnected.")
	}
}

//Listen Method
func (user *User) Listen() {
	go user.listenWrite()
	user.listenRead()
}

//ListenWrite Method
func (user *User) listenWrite() {
	for {
		select {}
	}
}
