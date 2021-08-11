package main

import (
	"fmt"

	"github.com/dusk125/pixelutils"
	"github.com/dusk125/pixelutils/gui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "GUI Example",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	})

	if err != nil {
		panic(fmt.Sprintf("error while creating window: %v", err))
	}
	buttonSprite, err := pixelutils.LoadSprite("button.png")
	if err != nil {
		panic(err)
	}
	buttonPressedSprite, err := pixelutils.LoadSprite("button_pressed.png")
	if err != nil {
		panic(err)
	}

	button := gui.NewButton(buttonSprite, buttonPressedSprite, func() {
		fmt.Println("Button Pressed!")
	})

	for !win.Closed() || win.JustPressed(pixelgl.KeyEscape) {
		button.Draw(win, pixel.IM)
		win.Update()
	}
}
