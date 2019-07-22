package main

import (
	"encoding/json"
	"log"
)

//Server structure
type Server struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func newServer() *Server {
	return &Server{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (server *Server) start() {
	for {
		select {
		case client := <-server.register:
			server.clients[client] = true
			jsonMessage, err := json.Marshal(newMsgSystem("/A new socket has connected."))
			if err != nil {
				log.Println(err)
			}
			server.send(jsonMessage, client)
		case client := <-server.unregister:
			if _, ok := server.clients[client]; ok {
				close(client.send)
				delete(server.clients, client)
				jsonMessage, err := json.Marshal(newMsgSystem("/A socket has disconnected."))
				if err != nil {
					log.Println(err)
				}
				server.send(jsonMessage, client)
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

func (server *Server) send(message []byte, ignore *Client) {
	for client := range server.clients {
		if client != ignore {
			client.send <- message
		}
	}
}
