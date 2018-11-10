package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	GameId int    `json:"gameId"`
	Action string `json:"action"`
	Data   int    `json:"data"`
}

type Game struct {
	GameId  int                `json:"gameId"`
	Players [2]*websocket.Conn `json:"-"`
	Player  string             `json:"player,omitempty"`
	// 0 = Game in play, 1 = Player One won, 2 = Player Two won, 3 = Draw
	Status int    `json:"status"`
	Turn   int    `json:"turn"`
	Board  [9]int `json:"board"`
}

var gameCounter = 0

var games = make(map[int]*Game)
var clients = make(map[*websocket.Conn]bool)

func main() {
	http.HandleFunc("/ws", handleConnections)

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Close the connection when the function returns
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
	}
}
