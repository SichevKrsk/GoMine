package main

import (
	"errors"
	"math"
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

type Drawable interface {
	Draw()
}

type GameState struct {
	BombsToFind int32
	BombsFound  int32
}

type Config struct {
	GameName      string
	TileDimension int32
	RowNumber     int32
	ColumnNumber  int32
}

func (c *Config) GetWindowDimensions() (int32, int32) {
	return c.TileDimension * c.ColumnNumber, c.TileDimension * c.RowNumber
}

var config = Config{
	GameName:      "Minesweeper",
	TileDimension: 32,
	RowNumber:     10,
	ColumnNumber:  10,
}

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
		break
	case Flagged:
		color = rl.Blue
		break
	case Revealed:
		color = rl.Gray
		break
	case Bomb:
		color = rl.Red
		break
	default:
		color = rl.Beige
	}

	rl.DrawRectangle(t.X*config.TileDimension, t.Y*config.TileDimension, config.TileDimension, config.TileDimension, color)
	rl.DrawRectangleLines(t.X*config.TileDimension, t.Y*config.TileDimension, config.TileDimension, config.TileDimension, rl.Black)

	// Draw number if revealed and not a bomb and value > 0
	if t.Status == Revealed && t.Value > 0 {
		x := t.X * config.TileDimension
		y := t.Y * config.TileDimension
		size := config.TileDimension

		text := strconv.Itoa(int(t.Value))
		fontSize := int32(20)

		textWidth := rl.MeasureText(text, fontSize)
		textX := x + (size-textWidth)/2
		textY := y + (size-fontSize)/2

		rl.DrawText(text, textX, textY, fontSize, rl.Green)
	}
}

func GetBoardCoordinates() (int32, int32, error) {
	mousePos := rl.GetMousePosition()

	X := int32(math.Floor(float64(mousePos.X / float32(config.TileDimension))))
	Y := int32(math.Floor(float64(mousePos.Y / float32(config.TileDimension))))

	if X >= config.ColumnNumber || Y >= config.RowNumber || X < 0 || Y < 0 {
		return 0, 0, errors.New("coordinates are out of the board bounds")
	}
	return X, Y, nil
}

func (b *Board) HandleInputs() {
	X, Y, err := GetBoardCoordinates()

	if err != nil {
		return
	}

	tile := b.Tiles[Y][X]

	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		if tile.Status == Flagged {
			tile.Status = Hidden
		} else if tile.Status == Hidden {
			tile.Status = Flagged
		}
	}

	if tile.Status != Hidden {
		return
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {

		if tile.Value == -1 {
			tile.Status = Bomb
		} else {
			tile.Status = Revealed
		}
	}
}

type Board struct {
	BombsToFind int32
	Tiles       [][]*Tile
}

func (b *Board) Draw() {
	for i := 0; i < len(b.Tiles); i++ {
		for j := 0; j < len(b.Tiles[i]); j++ {
			b.Tiles[i][j].Draw()
		}
	}
}

func main() {
	board := Board{
		BombsToFind: 10,
		Tiles:       testTiles,
	}

	windowWidth, windowHeight := config.GetWindowDimensions()
	rl.InitWindow(windowWidth, windowHeight, config.GameName)

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		board.Draw()
		board.HandleInputs()

		rl.EndDrawing()
	}
}
