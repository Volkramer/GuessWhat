package main

//Message structure
type Message struct {
	Event    string `json:"event,omitempty"`
	Username string `json:"email,omitempty"`
	Message  string `json:"message,omitempty"`
}

func newMessage(username string, message string) *Message {
	return &Message{
		Event:    "Message",
		Username: username,
		Message:  message,
	}
}

//MsgSystem structure
type MsgSystem struct {
	Event      string `json:"event,omitempty"`
	SysMessage string `json:"sysmessage,omitempty"`
}

func newMsgSystem(systemMessage string) *MsgSystem {
	return &MsgSystem{
		Event:      "System",
		SysMessage: systemMessage,
	}
}

//MsgJoin structure
type MsgJoin struct {
	Event    string `json:"event,omitempty"`
	Username string `json:"username,omitempty"`
}
