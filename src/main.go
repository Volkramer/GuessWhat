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

//Main function
func main() {
	router()
	log.Println("http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

//Router function to handle connection and websocket
func router() {
	http.HandleFunc("/ws", start)
	http.Handle("/", http.FileServer(http.Dir("./public")))
}

func start(w http.ResponseWriter, r *http.Request) {
	//connection client
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "could not connect", http.StatusInternalServerError)
		return
	}
	log.Println(socket.RemoteAddr(), "connected")

	/* //Send list of rooms to the client
	SendData("getRooms", getRooms(), conn)
	Error(err)

	//client send his username and room name
	event, content := ReceiveData(conn)
	if event == "registration" {
		log.Println("SYSTEM: data received: event =", event)
		for key, value := range content {
			log.Println("SYSTEM: data received: content =", key, "=", value)
		}

		if content["room"] == "" {
			SendData("badRoom", "", conn)
		} else if content["username"] == "" {
			SendData("badUsername", "", conn)
		} else {

			//Creation of the user from form
			user := NewUser(content["username"], conn)
			if user == nil {
				log.Println("SYSTEM: Creating new user failed")
			}

			//Verification of existing room, if not create it
			foundRoom, exist := rooms[content["room"]]
			if exist {
				foundRoom.AddUser(user)
			} else {
				newRoom := NewRoom(content["room"])
				if newRoom == nil {
					log.Println("SYSTEM: Creating new room failed")
				}
				rooms[newRoom.name] = newRoom
				go newRoom.run()
				newRoom.AddUser(user)
			}
			data := Join{
				Username: user.username,
				Room:     user.room.name,
			}
			SendData("newUser", data, conn)
			user.Listen()
		}
	} */
}
