package main

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

//Client structure
type Client struct {
	username string
	server   *Server
	socket   *websocket.Conn
	send     chan interface{}
	mux      sync.Mutex
}

func newClient(username string, server *Server, socket *websocket.Conn) *Client {
	return &Client{
		username: username,
		server:   server,
		socket:   socket,
		send:     make(chan interface{}),
	}
}

func (client *Client) read() {
	defer func() {
		client.server.unregister <- client
		client.socket.Close()
	}()

	for {
		var message *Message
		err := client.socket.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			client.server.unregister <- client
			client.socket.Close()
			break
		}
		client.server.broadcast <- message
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
			err := client.sendData(message)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (client *Client) sendData(data interface{}) error {
	client.mux.Lock()
	defer client.mux.Unlock()
	return client.socket.WriteJSON(data)
}
