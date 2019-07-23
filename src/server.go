package main

import (
	"fmt"
)

//Server structure
type Server struct {
	clients    map[*Client]bool
	broadcast  chan interface{}
	register   chan *Client
	unregister chan *Client
}

func newServer() *Server {
	return &Server{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan interface{}),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (server *Server) start() {
	for {
		select {
		case client := <-server.register:
			server.clients[client] = true
			msgSystem := newMsgSystem(fmt.Sprintln(client.username, "has joined the game"))
			server.send(msgSystem, client)
		case client := <-server.unregister:
			if _, ok := server.clients[client]; ok {
				close(client.send)
				delete(server.clients, client)
				msgSystem := newMsgSystem(fmt.Sprintln(client.username, "has left the game"))
				server.send(msgSystem, client)
			}
		case message := <-server.broadcast:
			for client := range server.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(server.clients, client)
				}
			}
		}
	}
}

func (server *Server) send(message interface{}, ignore *Client) {
	for client := range server.clients {
		if client != ignore {
			client.send <- message
		}
	}
}
