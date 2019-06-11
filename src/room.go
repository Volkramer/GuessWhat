package main

import "fmt"

//Room Object
type Room struct {
	Name  string
	Users map[*User]bool
	Join  chan *User
	Leave chan *User
	Input chan *Message
}

var rooms = make(map[string]*Room)

func (r *Room) run() {
flag:
	for {
		select {
		case user := <-r.Join:
			r.Users[user] = true
			go func() {
				r.Input <- &Message{
					Username: "SYSTEM",
					Text:     fmt.Sprintln(user.username, "joined"),
				}
			}()
		case user := <-r.Leave:
			delete(r.Users, user)
			go func() {
				r.Input <- &Message{
					Username: "SYSTEM",
					Text:     fmt.Sprintln(user.username, "left"),
				}
			}()
			if len(r.Users) == 0 {
				delete(rooms, r.Name)
				close(r.Join)
				close(r.Leave)
				close(r.Input)
				break flag
			}
		case msg := <-r.Input:
			for user := range r.Users {
				user.output <- msg
			}
		}
	}
}

//NewRoom method
func NewRoom(name string) (room *Room) {
	return &Room{
		Name:  name,
		Users: make(map[*User]bool),
		Join:  make(chan *User),
		Leave: make(chan *User),
		Input: make(chan *Message),
	}
}

func getRooms() (roomsName string) {
	roomsName = "'rooms': {"
	for room := range rooms {
		roomsName = fmt.Sprint(roomsName + "'" + room + "', ")
	}
	roomsName = fmt.Sprint(roomsName + "}")
	return roomsName
}
