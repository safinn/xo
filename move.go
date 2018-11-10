package main

import (
	"errors"

	"github.com/gorilla/websocket"
)

func makeMove(games map[int]*Game, ws *websocket.Conn, msg Message) {
	game := games[msg.GameId]
	if game != nil && game.Players[game.Turn] == ws && game.Status == 0 {

		err := setMove(game, msg)
		if err != nil {
			return
		}

		checkWinOrDraw(game)
		changeTurn(game)
		updateGameStateForPlayers(game)
	}
}

func setMove(game *Game, msg Message) error {
	if game.Board[msg.Data] != 0 {
		return errors.New("space already taken")
	}
	// 0 = blank, 1 = player one, 2 = player two
	game.Board[msg.Data] = game.Turn + 1
	return nil
}

func checkWinOrDraw(game *Game) {
	win := false
	board := game.Board
	winConditions := [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}
	for _, condition := range winConditions {
		if board[condition[0]] != 0 && board[condition[0]] == board[condition[1]] && board[condition[0]] == board[condition[2]] {
			game.Status = board[condition[0]]
			win = true
			break
		}
	}
	if !win {
		draw := true
		for i := 0; i < len(game.Board); i++ {
			if game.Board[i] == 0 {
				draw = false
			}
		}
		if draw {
			game.Status = 3
		}
	}
}

func updateGameStateForPlayers(game *Game) {
	for i := 0; i < len(game.Players); i++ {
		client := game.Players[i]
		if client != nil {
			sendWsMessage(client, game)
		}
	}
}

func changeTurn(game *Game) {
	if game.Turn == 0 {
		game.Turn = 1
	} else {
		game.Turn = 0
	}
}
