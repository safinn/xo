package main

import (
	"testing"

	"github.com/gorilla/websocket"
)

func TestCreateNewGame(t *testing.T) {
	counter := 0
	ws := new(websocket.Conn)
	game := createNewGame(&counter, ws)

	if game.GameId != 0 {
		t.Errorf("Incorrect, got: %d, want: %d.", game.GameId, 0)
	}
}
