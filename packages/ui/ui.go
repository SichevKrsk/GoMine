package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type uiEvent struct {
	targetId        string
	targetTextValue string
	consumed        bool
}

var globalEventPool = []uiEvent{}

type drawable interface {
	draw()
	getId() string
	setText(value string)
	getParent() *container
}

type interactable interface {
	drawable
	checkClick()
}

type container struct {
	position             rl.Vector2
	width                int32
	height               int32
	color                rl.Color
	staticElements       []drawable
	interactableElements []interactable
	parent               *container
	id                   string
}

func PublishEvent(targetId string, textValue string) {
	globalEventPool = append(globalEventPool, uiEvent{
		targetId:        targetId,
		targetTextValue: textValue,
		consumed:        false,
	})
}

func CreateRootContainer(position rl.Vector2, width, height int32) *container {
	return &container{
		position:             position,
		width:                width,
		height:               height,
		color:                rl.Gray,
		staticElements:       []drawable{},
		interactableElements: []interactable{},
		parent:               nil,
		id:                   "root",
	}
}

func (c *container) Draw() {
	c.draw()
}

func (c *container) EndContainerDefinition() *container {
	return c.parent
}

func (c *container) draw() {
	rl.DrawRectangle(int32(c.position.X), int32(c.position.Y), c.width, c.height, c.color)

	for _, static := range c.staticElements {
		static.draw()
		for _, event := range globalEventPool {
			if !event.consumed && event.targetId == static.getId() {
				event.consumed = true
				static.setText(event.targetTextValue)
			}
		}
	}

	for _, interactive := range c.interactableElements {
		interactive.draw()

		interactive.checkClick()

		for _, event := range globalEventPool {
			if !event.consumed && event.targetId == interactive.getId() {
				event.consumed = true
				interactive.setText(event.targetTextValue)
			}
		}
	}

}

func (c *container) AddContainer(position rl.Vector2, width, height int32, color rl.Color, id string) *container {
	container := &container{
		position:             position,
		width:                width,
		height:               height,
		color:                color,
		staticElements:       []drawable{},
		interactableElements: []interactable{},
		parent:               c,
		id:                   id,
	}

	c.staticElements = append(c.staticElements, container)

	return container
}

func (c *container) AddButton(position rl.Vector2, text string, clickHandler func(), id string) *container {
	c.interactableElements = append(c.interactableElements, &button{
		position:    position,
		width:       200,
		height:      50,
		text:        text,
		onClick:     clickHandler,
		color:       rl.DarkGray,
		borderColor: rl.Black,
		id:          id,
	})

	return c
}

func (c *container) AddLabel(position rl.Vector2, text string, id string) *container {
	c.staticElements = append(c.staticElements, &label{
		position: position,
		width:    200,
		height:   50,
		text:     text,
		color:    rl.DarkGray,
		id:       id,
	})

	return c
}

func (c *container) getParent() *container {
	return c.parent
}

func (c *container) getId() string {
	return c.id
}
func (c *container) setText(value string) {

}

type button struct {
	position    rl.Vector2
	width       int32
	height      int32
	color       rl.Color
	borderColor rl.Color
	text        string
	onClick     func()
	id          string
	parent      *container
}

func (b button) getParent() *container {
	return b.parent
}

func (b button) getId() string {
	return b.id
}

func (b *button) setText(value string) {
	b.text = value
}

func (b button) draw() {
	rl.DrawRectangle(int32(b.position.X), int32(b.position.Y), b.width, b.height, b.color)
	rl.DrawRectangleLines(int32(b.position.X), int32(b.position.Y), b.width, b.height, b.borderColor)
	rl.DrawText(b.text, int32(b.position.X), int32(b.position.Y), 12, rl.Red)
}

func (b button) checkClick() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{
		X:      b.position.X,
		Y:      b.position.Y,
		Width:  float32(b.width),
		Height: float32(b.height),
	}) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		b.onClick()
	}
}

type label struct {
	position rl.Vector2
	width    int32
	height   int32
	color    rl.Color
	text     string
	id       string
	parent   *container
}

func (b label) getParent() *container {
	return b.parent
}

func (l label) getId() string {
	return l.id
}

func (b *label) setText(value string) {
	b.text = value
}

func (l label) draw() {
	rl.DrawRectangle(int32(l.position.X), int32(l.position.Y), l.width, l.height, l.color)
	rl.DrawText(l.text, int32(l.position.X), int32(l.position.Y), 12, rl.Red)
}
