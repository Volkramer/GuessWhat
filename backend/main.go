package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//Message Object
type Message struct {
	username string
	msg      string
}

var env = "dev"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	router()
	http.ListenAndServe(":8000", nil)
}

func router() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		m := Message{}
		err := conn.ReadJSON(m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v sent: %v\n", m.username, m.msg)
		err = conn.WriteJSON(m)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if env == "dev" {
		http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
	} else if env == "prod" {
		http.ServeFile(w, r, "../frontend/dist/index.html")
	}
}
