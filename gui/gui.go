package gui

import "github.com/faiface/pixel"

type Handler func()

func NewButton(idle, active *pixel.Sprite, handler Handler) *Button {
	if idle.Picture().Bounds() == active.Picture().Bounds() {
		panic("button sprites not of same size")
	}
	return &Button{
		activeSprite: active,
		idleSprite:   idle,
		handler:      handler,
		bounds:       idle.Picture().Bounds(),
	}
}

type Button struct {
	activeSprite *pixel.Sprite
	idleSprite   *pixel.Sprite

	bounds pixel.Rect

	handler Handler
}

func (b Button) Draw(t pixel.Target, matrix pixel.Matrix) {
	b.activeSprite.Draw(t, matrix)
}
