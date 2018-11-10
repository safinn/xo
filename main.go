package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/games/", gameApi)

	http.HandleFunc("/ws", handleConnections)

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func gameApi(w http.ResponseWriter, r *http.Request) {
	gameIdInt, err := getParameterAsInt(r)
	if err != nil {
		http.Error(w, "Url parameter not a single int", http.StatusInternalServerError)
		return
	}

	data, err := findGame(gameIdInt, games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return
}

func getParameterAsInt(r *http.Request) (int, error) {
	gameId := strings.TrimPrefix(r.URL.Path, "/games/")
	gameIdInt, err := strconv.Atoi(gameId)
	if err != nil {
		return 0, errors.New("Url parameter not a single int")
	}
	return gameIdInt, nil
}

func findGame(gameId int, games map[int]*Game) ([]byte, error) {
	game := games[gameId]
	if game != nil {
		game.Player = ""
		data, err := json.Marshal(game)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		return data, nil
	}
	return nil, errors.New("No game found")
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

		handleAction(msg, ws)
	}
}

func handleAction(msg Message, ws *websocket.Conn) {
	switch msg.Action {
	case "NEW GAME":
		newGame(&gameCounter, games, ws)
	case "JOIN GAME":
		joinGame(games, ws, msg)
	case "MOVE":
		makeMove(games, ws, msg)
	}
}

func sendWsMessage(ws *websocket.Conn, game *Game) {
	if game.Players[0] == ws {
		game.Player = "X"
	} else {
		game.Player = "O"
	}
	err := ws.WriteJSON(game)
	if err != nil {
		log.Printf("error: %v", err)
		ws.Close()
		delete(clients, ws)
	}
}
