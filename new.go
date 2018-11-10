package main

import "github.com/gorilla/websocket"

func newGame(gameCounter *int, games map[int]*Game, ws *websocket.Conn) {
	*gameCounter++

	game := new(Game)
	game.GameId = *gameCounter
	game.Players[0] = ws
	games[*gameCounter] = game

	sendWsMessage(ws, game)
}
