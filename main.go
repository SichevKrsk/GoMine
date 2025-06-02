package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//WIN CONDITION = all tiles revealed except bombs

type Drawable interface {
	Draw()
}

type GameState struct {
	BombsToFind   int32
	BombsFound    int32
	TilesRevealed int32
	Board         *Board
}

func (g *GameState) Reset() {
	g.Board.Reset()
	g.BombsFound = 0
	g.TilesRevealed = 0
}

func (g *GameState) Check() {
	if g.TilesRevealed == g.Board.NumberOfTiles-g.BombsToFind {
		rl.DrawText("You WON", 150, 150, 30, rl.Red)
	}
}

func main() {
	board, err := Generate(10, 10, 10)

	if err != nil {
		return
	}

	var gamestate GameState = GameState{
		BombsToFind:   3,
		BombsFound:    0,
		TilesRevealed: 0,
		Board:         board,
	}

	board.GameState = &gamestate

	for i := 0; i < len(board.Tiles); i++ {
		for j := 0; j < len(board.Tiles[i]); j++ {
			fmt.Print("[", i, ",", j, "]=", board.Tiles[i][j].Value, " ")
		}
		fmt.Println()
	}

	windowWidth, windowHeight := config.GetWindowDimensions()
	rl.InitWindow(windowWidth, windowHeight, config.GameName)

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		gamestate.Check()
		board.Draw()
		board.HandleInputs()

		rl.EndDrawing()
	}
}
