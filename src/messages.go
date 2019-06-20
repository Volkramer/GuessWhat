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

//PointJSON Object for json through websocket
type PointJSON struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//UserJSON Object for json through websocket
type UserJSON struct {
	Username string `json:"username"`
}

//ConnectedJSON Object for json through websocket
type ConnectedJSON struct {
	Event int        `json:"event"`
	Users []UserJSON `json:"users"`
}

//UserJoinedJSON Object for json through websocket
type UserJoinedJSON struct {
	Event int      `json:"event"`
	User  UserJSON `json:"user"`
}

//UserLeftJSON Object for json through websocket
type UserLeftJSON struct {
	Event    int    `json:"event"`
	Username string `json:"username"`
}

//StrokeJSON Object for json through websocket
type StrokeJSON struct {
	Event    int         `json:"event"`
	Username string      `json:"username"`
	Points   []PointJSON `json:"points"`
	Finish   bool        `json:"finish"`
}

//ClearJSON Object for json through websocket
type ClearJSON struct {
	Event    int    `json:"event"`
	Username string `json:"username"`
}

//NewConnected constructor method
func NewConnected(name string, users []UserJSON) *ConnectedJSON {
	return &ConnectedJSON{
		Event: EventConnected,
		Users: users,
	}
}

//NewUserLeft constructor method
func NewUserLeft(username string) *UserLeftJSON {
	return &UserLeftJSON{
		Event:    EventUserLeft,
		Username: username,
	}
}

//NewUserJoined constructor method
func NewUserJoined(username string) *UserJoinedJSON {
	return &UserJoinedJSON{
		Event: EventUserJoined,
		User:  UserJSON{Username: username},
	}
}
