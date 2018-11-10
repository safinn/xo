package main

import "github.com/gorilla/websocket"

func newGame(gameCounter *int, games map[int]*Game, ws *websocket.Conn) {
	*gameCounter++

	game := createNewGame(gameCounter, ws)
	games[*gameCounter] = game

	sendWsMessage(ws, game)
}

func createNewGame(gameCounter *int, ws *websocket.Conn) *Game {
	game := new(Game)
	game.GameId = *gameCounter
	game.Players[0] = ws

	return game
}
