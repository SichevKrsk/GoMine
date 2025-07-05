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
	offsetX := int32(b.Position.X)
	offsetY := int32(b.Position.Y)
	for i := 0; i < len(b.Tiles); i++ {
		for j := 0; j < len(b.Tiles[i]); j++ {
			b.Tiles[i][j].Draw(offsetX, offsetY)
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
	offsetX := int32(b.Position.X)
	offsetY := int32(b.Position.Y)
	X, Y, err := GetBoardCoordinates(offsetX, offsetY)

	if err != nil {
		return
	}

	tile := b.Tiles[X][Y]

	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		switch tile.Status {
		case Flagged:
			tile.Status = Hidden
			b.GameState.BombsFound--
		case Hidden:
			tile.Status = Flagged
			b.GameState.BombsFound++
		}
	}

	if tile.Status != Hidden {
		return
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if tile.Value == -1 {
			tile.Status = Bomb
			b.GameState.GameLost = true
			b.GameState.GameActive = false
		} else {
			tile.Status = Revealed
			b.GameState.TilesRevealed = b.GameState.TilesRevealed + 1
		}
	}
}
