package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

//Client structure
type Client struct {
	username string
	server   *Server
	socket   *websocket.Conn
	send     chan []byte
}

func newClient(username string, server *Server, socket *websocket.Conn) *Client {
	return &Client{
		username: username,
		server:   server,
		socket:   socket,
		send:     make(chan []byte),
	}
}

func (client *Client) read() {
	defer func() {
		client.server.unregister <- client
		client.socket.Close()
	}()

	for {
		_, message, err := client.socket.ReadMessage()
		if err != nil {
			log.Println(err)
			client.server.unregister <- client
			client.socket.Close()
			break
		}
		jsonMessage, err := json.Marshal(newMessage(client.username, string(message)))
		client.server.broadcast <- jsonMessage
	}
}

func (client *Client) write() {
	defer func() {
		client.socket.Close()
	}()

	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				client.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			client.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
