package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Board struct {
	Position      rl.Vector2
	NumberOfBombs int32
	NumberOfTiles int32
	GameState     *GameState
	Tiles         [][]*Tile
}

func (b *Board) Draw() {
	for i := 0; i < len(b.Tiles); i++ {
		for j := 0; j < len(b.Tiles[i]); j++ {
			b.Tiles[i][j].Draw()
		}
	}
}

func (b *Board) Reset() {
	for i := 0; i < len(b.Tiles); i++ {
		for j := 0; j < len(b.Tiles[i]); j++ {
			b.Tiles[i][j].Status = Hidden
		}
	}
}

func (b *Board) HandleInputs() {
	X, Y, err := GetBoardCoordinates()

	if err != nil {
		return
	}

	tile := b.Tiles[X][Y]

	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		switch tile.Status {
		case Flagged:
			tile.Status = Hidden
		case Hidden:
			tile.Status = Flagged
		}
	}

	if tile.Status != Hidden {
		return
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if tile.Value == -1 {
			tile.Status = Bomb
			b.GameState.Reset()

		} else {
			tile.Status = Revealed
			b.GameState.TilesRevealed = b.GameState.TilesRevealed + 1
		}
	}
}
