package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileStatus string

const (
	Hidden   TileStatus = "hidden"
	Flagged  TileStatus = "flagged"
	Revealed TileStatus = "revealed"
	Bomb     TileStatus = "bomb"
)

type Tile struct {
	X, Y   int32
	Value  int32 // -1 for mine, 0-8 for number of adjacent mines
	Status TileStatus
}

func (t *Tile) Draw() {
	var color rl.Color
	switch t.Status {
	case Hidden:
		color = rl.Gray
	case Flagged:
		color = rl.Blue
	case Revealed:
		color = rl.White
	case Bomb:
		color = rl.Red
	default:
		color = rl.Beige
	}

	DrawTile(t.X, t.Y, color)

	// Draw number if revealed and not a bomb and value > 0
	if t.Status == Revealed {
		if t.Value > 0 {
			text := strconv.Itoa(int(t.Value))
			fontSize := int32(20)
			textWidth := rl.MeasureText(text, fontSize)
			DrawInTile(t.X, t.Y, textWidth, fontSize, func(x, y int32) {
				rl.DrawText(text, x, y, fontSize, rl.Green)
			})
		}
	}
}
