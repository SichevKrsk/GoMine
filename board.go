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
			b.revealTileAndAdjacent(int32(X), int32(Y))
		}
	}
}

// revealTileAndAdjacent reveals a tile and recursively reveals adjacent empty tiles
func (b *Board) revealTileAndAdjacent(x, y int32) {
	// Check bounds
	if x < 0 || x >= int32(len(b.Tiles)) || y < 0 || y >= int32(len(b.Tiles[0])) {
		return
	}

	tile := b.Tiles[x][y]

	// If tile is already revealed or flagged, don't do anything
	if tile.Status != Hidden {
		return
	}

	// Reveal the current tile
	tile.Status = Revealed
	b.GameState.TilesRevealed++

	// If this tile has a value (number of adjacent bombs), don't reveal adjacent tiles
	if tile.Value > 0 {
		return
	}

	// If this tile is empty (value == 0), reveal all adjacent tiles
	for dx := int32(-1); dx <= 1; dx++ {
		for dy := int32(-1); dy <= 1; dy++ {
			// Skip the current tile itself
			if dx == 0 && dy == 0 {
				continue
			}

			newX := x + dx
			newY := y + dy

			// Recursively reveal adjacent tiles
			b.revealTileAndAdjacent(newX, newY)
		}
	}
}
