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
