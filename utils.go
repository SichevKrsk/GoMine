package main

import (
	"errors"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func GetBoardCoordinates(boardOffsetX, boardOffsetY int32) (int32, int32, error) {
	mousePos := rl.GetMousePosition()

	// Adjust mouse position by board offset
	adjustedX := mousePos.X - float32(boardOffsetX)
	adjustedY := mousePos.Y - float32(boardOffsetY)

	X := int32(math.Floor(float64(adjustedX / float32(config.TileDimension))))
	Y := int32(math.Floor(float64(adjustedY / float32(config.TileDimension))))

	if X >= config.ColumnNumber || Y >= config.RowNumber || X < 0 || Y < 0 {
		return 0, 0, errors.New("coordinates are out of the board bounds")
	}
	return X, Y, nil
}

func DrawTile(posX, posY, offsetX, offsetY int32, color rl.Color) {
	rl.DrawRectangle(
		offsetX+posX*config.TileDimension,
		offsetY+posY*config.TileDimension,
		config.TileDimension,
		config.TileDimension,
		color,
	)
	rl.DrawRectangleLines(
		offsetX+posX*config.TileDimension,
		offsetY+posY*config.TileDimension,
		config.TileDimension,
		config.TileDimension,
		rl.Black,
	)
}

func DrawInTile(posX, posY, width, height, offsetX, offsetY int32, callback func(x, y int32)) {
	x := offsetX + posX*config.TileDimension
	y := offsetY + posY*config.TileDimension

	textX := x + (config.TileDimension-width)/2
	textY := y + (config.TileDimension-height)/2

	callback(textX, textY)
}
