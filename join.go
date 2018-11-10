package main

import "github.com/gorilla/websocket"

func joinGame(games map[int]*Game, ws *websocket.Conn, msg Message) {
	game := games[msg.GameId]
	if game != nil && (ws == game.Players[0] || ws == game.Players[1]) {
		sendWsMessage(ws, game)
		return
	}
	if game != nil && game.Players[1] == nil {
		game.Players[1] = ws
		sendWsMessage(ws, game)
	}
}
