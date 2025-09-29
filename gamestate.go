package main

import "fmt"

// GameState manages the current game state
type GameState struct {
	BombsToFind   int32
	BombsFound    int32
	TilesRevealed int32
	Board         *Board
	GameActive    bool
	GameWon       bool
	GameLost      bool
}

func (g *GameState) Reset() {
	if g.Board != nil {
		g.Board.Reset()
	}
	g.BombsFound = 0
	g.TilesRevealed = 0
	g.GameActive = false
	g.GameWon = false
	g.GameLost = false
}

func (g *GameState) Check() {
	if g.GameActive && g.TilesRevealed == g.Board.NumberOfTiles-g.BombsToFind {
		g.GameWon = true
		g.GameActive = false
	}
}

func (g *GameState) StartNewGame() {
	// Generate new board with 10 bombs on 10x10 grid
	board, err := Generate(10, 10, 10)
	if err != nil {
		fmt.Println("Error generating board:", err)
		return
	}

	g.Board = board
	g.BombsToFind = 10
	g.BombsFound = 0
	g.TilesRevealed = 0
	g.GameActive = true
	g.GameWon = false
	g.GameLost = false

	// Set the game state reference in the board
	board.GameState = g
}
