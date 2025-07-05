package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/GOsweeper/packages/ui"
)

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

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create game state
	gameState := &GameState{}
	gameState.Reset()

	// Create main UI container
	maincontainer :=
		ui.CreateRootContainer(rl.Vector2{X: 0, Y: 0}, 800, 600).
			AddContainer(rl.Vector2{X: 0, Y: 0}, 800, 100, rl.LightGray, "header").
			AddLabel(rl.Vector2{X: 10, Y: 10}, "GOsweeper", "title").
			AddLabel(rl.Vector2{X: 10, Y: 40}, "Flags: 0/10", "flag-counter").
			AddLabel(rl.Vector2{X: 10, Y: 70}, "Status: Menu", "game-status").
			AddButton(rl.Vector2{X: 600, Y: 20}, "New Game", func() {
				gameState.StartNewGame()
				ui.PublishEvent("game-status", "Playing")
				ui.PublishEvent("flag-counter", fmt.Sprintf("Flags: %d/%d", gameState.BombsFound, gameState.BombsToFind))
			}, "new-game-btn").
			EndContainerDefinition().
			AddContainer(rl.Vector2{X: 0, Y: 100}, 800, 500, rl.White, "game-area").EndContainerDefinition()

	rl.InitWindow(800, 600, "GOsweeper")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Draw UI
		maincontainer.Draw()

		// Draw game board if active
		if gameState.GameActive && gameState.Board != nil {
			// Calculate board dimensions
			boardWidth := config.ColumnNumber * config.TileDimension

			// Position board below header, centered horizontally
			boardOffsetX := int32((800 - boardWidth) / 2) // Center horizontally
			boardOffsetY := int32(120)                    // Position below header with some padding

			// Set board position for drawing
			gameState.Board.Position = rl.Vector2{X: float32(boardOffsetX), Y: float32(boardOffsetY)}

			gameState.Board.Draw()
			gameState.Board.HandleInputs()

			// Check win condition
			gameState.Check()

			// Update flag counter
			ui.PublishEvent("flag-counter", fmt.Sprintf("Flags: %d/%d", gameState.BombsFound, gameState.BombsToFind))
		}

		// Draw game over messages
		if gameState.GameWon {
			rl.DrawText("YOU WON!", 300, 250, 40, rl.Green)
			ui.PublishEvent("game-status", "Won!")
		} else if gameState.GameLost {
			rl.DrawText("GAME OVER!", 300, 250, 40, rl.Red)
			ui.PublishEvent("game-status", "Lost!")
		}

		rl.EndDrawing()
	}
}
