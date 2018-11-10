package main

import (
	"errors"
	"testing"
)

func TestChangeTurn(t *testing.T) {
	tables := []struct {
		turn   int
		should int
	}{
		{0, 1},
		{1, 0},
	}

	for _, table := range tables {
		game := new(Game)
		game.Turn = table.turn
		changeTurn(game)

		if game.Turn != table.should {
			t.Errorf("Incorrect, got: %d, want: %d.", game.Turn, table.should)
		}
	}
}

func TestCheckWinOrDraw(t *testing.T) {
	tables := []struct {
		board  [9]int
		status int
	}{
		{[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{[9]int{1, 1, 1, 0, 0, 0, 0, 0, 0}, 1},
		{[9]int{0, 0, 0, 2, 2, 2, 0, 0, 0}, 2},
		{[9]int{1, 2, 2, 1, 1, 2, 1, 2, 1}, 1},
		{[9]int{0, 1, 0, 1, 0, 0, 1, 2, 2}, 0},
		{[9]int{1, 1, 2, 2, 1, 1, 1, 2, 2}, 3},
	}

	for _, table := range tables {
		game := new(Game)
		game.Board = table.board

		checkWinOrDraw(game)

		if game.Status != table.status {
			t.Errorf("Incorrect, got: %d, want: %d.", game.Status, table.status)
		}
	}
}

func TestSetMove(t *testing.T) {
	tables := []struct {
		board [9]int
		turn  int
		data  int
		err   error
		move  int
	}{
		{[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 0, nil, 1},
		{[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}, 1, 5, nil, 2},
		{[9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 0, errors.New(""), 1},
		{[9]int{0, 2, 0, 0, 0, 0, 0, 0, 0}, 1, 1, errors.New(""), 2},
	}

	for _, table := range tables {
		game := new(Game)
		msg := new(Message)
		game.Board = table.board
		game.Turn = table.turn
		msg.Data = table.data

		err := setMove(game, *msg)

		_, tableErrOk := table.err.(error)
		_, errOk := err.(error)
		if tableErrOk != errOk {
			t.Errorf("Incorrect, got: %e, want: %e.", err, table.err)
		}

		if game.Board[msg.Data] != table.move {
			t.Errorf("Incorrect, got: %d, want: %d.", game.Board[msg.Data], table.move)
		}
	}
}
