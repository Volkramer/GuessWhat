package main

const (
	//EventConnected is sent when a user connect
	EventConnected = iota + 1
	//EventUserJoined is sent when a user join a room
	EventUserJoined
	//EventUserLeft is sent when a user leave a room
	EventUserLeft
	//EventStroke is sent when a stroke is drawn
	EventStroke
	//EventClear is sent when a user clear the screen
	EventClear
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type User struct {
	name string
}

type Connected struct {
	Event int    `json:"event"`
	Users []User `json:"users"`
}

func NewConnected(name string, users []User) *Connected {
	return &Connected{
		Event: EventConnected,
		Name:  name,
		Users: users,
	}
}
