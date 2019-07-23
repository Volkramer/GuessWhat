package main

//Message structure
type Message struct {
	Event    string `json:"event,omitempty"`
	Username string `json:"username,omitempty"`
	Message  string `json:"message,omitempty"`
}

func newMessage(username string, message string) *Message {
	return &Message{
		Event:    "message",
		Username: username,
		Message:  message,
	}
}

//MsgSystem structure
type MsgSystem struct {
	Event      string `json:"event,omitempty"`
	SysMessage string `json:"systemMessage,omitempty"`
}

func newMsgSystem(systemMessage string) *MsgSystem {
	return &MsgSystem{
		Event:      "system",
		SysMessage: systemMessage,
	}
}

//MsgJoin structure
type MsgJoin struct {
	Event    string `json:"event,omitempty"`
	Username string `json:"username,omitempty"`
}

func newMsgJoin(username string) *MsgJoin {
	return &MsgJoin{
		Event:    "clientJoined",
		Username: username,
	}
}

//MsgClientList structure
type MsgClientList struct {
	Event   string   `json:"event,omitempty"`
	Clients []string `json:"clients,omitempty"`
}

func newMsgClientList(server *Server) *MsgClientList {
	clients := make([]string, 0, len(server.clients))
	for val := range server.clients {
		clients = append(clients, val.username)
	}
	return &MsgClientList{
		Event:   "clientList",
		Clients: clients,
	}
}

//Point structure
type Point struct {
	X int `json:"x"`
	Y int `json:"Y"`
}

//Stroke structure
type Stroke struct {
	Event    string  `json:"event"`
	Username string  `json:"username"`
	Points   []Point `json:"points"`
	Finish   bool    `json:"finish"`
}

//Clear structure
type Clear struct {
	Event    string `json:"event"`
	Username string `json:"username"`
}
