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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../public/index.html")
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsPage(server, w, r)
	})
	http.ListenAndServe(":8000", nil)
}

func wsPage(server *Server, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	var msgJoin *MsgJoin
	err = conn.ReadJSON(&msgJoin)
	if msgJoin.Event == "clientJoined" {
		client := newClient(msgJoin.Username, server, conn)
		server.register <- client
		go client.read()
		go client.write()
	}
}
