package main

import (
	"github.com/gorilla/websocket"
)

//User Object
type User struct {
	username string
	output   chan *Message
	room     *Room
	conn     *websocket.Conn
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
