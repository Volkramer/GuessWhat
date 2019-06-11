package main

import (
	"encoding/json"
	"fmt"
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

//Registration Object
type registration struct {
	username string
	room     string
}

//Bus Object for System Event
type Bus struct {
	Event   string `json:"event"`
	Content string `json:"content"`
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

//Error Method
func Error(err error) {
	if err != nil {
		log.Println(err)
	}
}

//SendData Method
func SendData(event string, content string, conn *websocket.Conn) {
	jsonContent := `{"` + content + `"}`
	result, err := json.Marshal(jsonContent)
	Error(err)
	err = conn.WriteJSON(Bus{
		Event:   event,
		Content: string(result),
	})
	//log.Println("SYSTEM: event=", event, "content =", content)
	Error(err)
}

//ReceiveData Method
func ReceiveData(conn *websocket.Conn) (event string, content map[string]string) {
	var bus Bus
	var result map[string]string
	err := conn.ReadJSON(&bus)
	Error(err)
	event = bus.Event
	log.Println(bus.Content)
	json.Unmarshal([]byte(bus.Content), &result)
	log.Println("SYSTEM: content=", result)
	return event, result
}

func start(w http.ResponseWriter, r *http.Request) {
	//connection client
	conn, err := upgrader.Upgrade(w, r, nil)
	Error(err)
	log.Println(conn.RemoteAddr(), "connected")
	defer conn.Close()

	//Send list of rooms to the client
	SendData("getRooms", getRooms(), conn)
	Error(err)

	//client send his username and room name
	event, content := ReceiveData(conn)
	log.Println("SYSTEM: data received: event=", event)
	for key, value := range content {
		log.Println("SYSTEM: data received: content=", key, "=", value)
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
			user.room = foundRoom
			foundRoom.Join <- user
			defer func() {
				foundRoom.Leave <- user
			}()
		} else {
			newRoom := NewRoom(content["room"])
			if newRoom == nil {
				log.Println("SYSTEM: Creating new room failed")
			}
			rooms[newRoom.Name] = newRoom
			go newRoom.run()
			user.room = newRoom
			newRoom.Join <- user
		}
		data := fmt.Sprintln("username:", user.username, "room", user.room.Name)
		SendData("newUser", data, conn)
	}
}
