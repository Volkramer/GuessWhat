package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	log.Println("System: starting server")
	server := newServer()
	go server.start()
	log.Println("System: Server started successfully at address: http://localhost:8000")
	router(server)
	http.ListenAndServe(":8000", nil)
}

func router(server *Server) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		start(server, w, r)
	})
	http.Handle("/", http.FileServer(http.Dir("./public")))
}

func start(server *Server, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	log.Println("SYSTEM: client", conn.RemoteAddr(), "connected")
	var msgJoin *MsgJoin
	for {
		err = conn.ReadJSON(&msgJoin)
		if msgJoin.Event == "clientJoined" {
			client := newClient(msgJoin.Username, server, conn)
			server.register <- client
			log.Println("SYSTEM: client", client.username, "has joined the game")
			go client.read()
			go client.write()
		}
	}
}
