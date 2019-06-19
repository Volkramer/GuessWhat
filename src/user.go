package main

import (
	"github.com/gorilla/websocket"
)

//User Object
type User struct {
	name   string
	room   *Room
	socket *websocket.Conn
	output chan []byte
}

func newUser(name string, room *Room, socket *websocket.Conn) *User {
	return &User{
		name:   name,
		room:   room,
		socket: socket,
		output: make(chan []byte),
	}
}

func (user *User) read() {
	defer func() {
		user.room.leave <- user
	}()
	for {
		_, data, err := user.socket.ReadMessage()
		if err != nil {
			break
		}
		user.room.onMessage(data, user)
	}
}

func (user *User) write() {
	for {
		select {
		case data, ok := <-user.output:
			if !ok {
				user.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			user.socket.WriteMessage(websocket.TextMessage, data)
		}
	}
}

func (user *User) run() {
	go user.read()
	go user.write()
}

func (user *User) close() {
	user.socket.Close()
	close(user.output)
}
