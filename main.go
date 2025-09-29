package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/sichevkrsk/GoMine/packages/ui"
)

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
			AddLabel(rl.Vector2{X: 10, Y: 10}, "GoMine", "title").
			AddLabel(rl.Vector2{X: 10, Y: 40}, "Flags: 0/10", "flag-counter").
			AddLabel(rl.Vector2{X: 10, Y: 70}, "Status: Menu", "game-status").
			AddButton(rl.Vector2{X: 600, Y: 20}, "New Game", func() {
				gameState.StartNewGame()
				ui.PublishEvent("game-status", "Playing")
				ui.PublishEvent("flag-counter", fmt.Sprintf("Flags: %d/%d", gameState.BombsFound, gameState.BombsToFind))
			}, "new-game-btn").
			EndContainerDefinition().
			AddContainer(rl.Vector2{X: 0, Y: 100}, 800, 500, rl.White, "game-area").EndContainerDefinition()

	rl.InitWindow(800, 600, "GoMine")

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
