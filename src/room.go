package main

//Room object
type Room struct {
	name  string
	users []*User
	join  chan *User
	leave chan *User
}

func newRoom(name string) *Room {
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
