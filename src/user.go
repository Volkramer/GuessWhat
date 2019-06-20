package main

import (
	"log"

	"github.com/gorilla/websocket"
)

//User Object
type User struct {
	username string
	room     *Room
	socket   *websocket.Conn
	output   chan []byte
}

func newUser(name string, room *Room, socket *websocket.Conn) *User {
	return &User{
		username: name,
		room:     room,
		socket:   socket,
		output:   make(chan []byte),
	}
}

func (user *User) read() {
	defer func() {
		user.room.leave <- user
	}()
	for {
		msgT, data, err := user.socket.ReadMessage()
		if err != nil {
			log.Println(msgT, ":", err)
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
