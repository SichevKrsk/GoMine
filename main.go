package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/GOsweeper/packages/ui"
)

//Add relative height based on container pos
//Add content

func main() {

	maincontainer :=
		ui.CreateRootContainer(rl.Vector2{X: 0, Y: 0}, 800, 800).
			AddContainer(rl.Vector2{X: 0, Y: 0}, 800, 200, rl.Beige, "header").
			AddLabel(rl.Vector2{X: 0, Y: 0}, "Hello world", "lbl-1").
			AddLabel(rl.Vector2{X: 0, Y: 100}, "Current score: 50", "lbl-2").
			AddButton(rl.Vector2{X: 200, Y: 10}, "Reset", func() {
				fmt.Println("Reset clicked")
				ui.PublishEvent("lbl-1", "Reset was clicked")
			}, "btn-1").
			AddButton(rl.Vector2{X: 200, Y: 110}, "Hello world button", func() {
				fmt.Println("Hello world button clicked")
				ui.PublishEvent("lbl-1", "Hello world button clicked")
			}, "btn-2").
			EndContainerDefinition().
			AddContainer(rl.Vector2{X: 0, Y: 200}, 800, 600, rl.Blue, "content").
			AddLabel(rl.Vector2{X: 0, Y: 100}, "Current score: 50", "lbl-2").
			AddButton(rl.Vector2{X: 200, Y: 110}, "Hello world button", func() {
				fmt.Println("Hello world button clicked")
				ui.PublishEvent("lbl-1", "Hello world button clicked")
			}, "btn-2").
			EndContainerDefinition()

	rl.InitWindow(800, 800, config.GameName)

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		maincontainer.Draw()

		rl.EndDrawing()
	}
}
