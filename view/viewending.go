package view

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewViewEnding(c Controller) View {
	return &viewEnding{}
}

type viewEnding struct {
	ctl Controller

	counter int
}

func (v *viewEnding) Update() error {
	return nil
	// if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
	// 	return nil
	// }

	// if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
	// 	return nil
	// }

	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	// 	return nil
	// }

	// v.counter++
	// return nil
}

func (v *viewEnding) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	// screen.DrawImage(v.getImg(), &ebiten.DrawImageOptions{})
}
