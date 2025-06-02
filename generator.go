package main

import (
	"errors"
	"math/rand"
)

func generateTiles(columns, rows int32) [][]*Tile {
	tiles := make([][]*Tile, rows)

	for i := 0; i < int(rows); i++ {
		tiles[i] = make([]*Tile, columns)
		for j := 0; j < int(columns); j++ {
			tiles[i][j] = &Tile{
				X:      int32(i),
				Y:      int32(j),
				Value:  0,
				Status: Hidden,
			}
		}
	}

	return tiles
}

func Generate(numberOfBombs, columns, rows int32) (*Board, error) {
	if columns <= 0 {
		return nil, errors.New("column size should be bigger than 0")
	}

	if rows <= 0 {
		return nil, errors.New("row size should be bigger than 0")
	}

	board := Board{
		NumberOfTiles: columns * rows,
		NumberOfBombs: numberOfBombs,
		Tiles:         generateTiles(columns, rows),
	}

	for i := 0; i < int(numberOfBombs); i++ {
		x := rand.Intn(int(rows))
		y := rand.Intn(int(columns))
		board.Tiles[x][y].Value = -1
	}

	for i := 0; i < int(rows); i++ {
		for j := 0; j < int(columns); j++ {
			current := board.Tiles[i][j]

			if current.Value == -1 {
				continue
			}

			for ofX := -1; ofX <= 1; ofX++ {
				for ofY := -1; ofY <= 1; ofY++ {
					evaluatedX := i + ofX
					evaluatedY := j + ofY

					//if out of bounds, do nothing
					if evaluatedX < 0 || evaluatedX == int(rows) || evaluatedY < 0 || evaluatedY == int(columns) || (ofX == i && ofY == j) {
						continue
					}

					evaluated := board.Tiles[evaluatedX][evaluatedY]

					if evaluated.Value == -1 {
						current.Value = current.Value + 1
					}

				}
			}
		}
	}

	return &board, nil
}
