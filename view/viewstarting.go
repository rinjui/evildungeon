package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rinjui/evildungeon/util"
)

func NewViewStarting(c Controller) View {
	return &viewStarting{
		Controller: c,
	}
}

type viewStarting struct {
	Controller

	counter int
}

func (v *viewStarting) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return v.Controller.Redirect(ViewIDGaming)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return v.Controller.Redirect(ViewIDGaming)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return v.Controller.Redirect(ViewIDGaming)
	}

	v.counter++
	v.counter = v.counter % ebiten.TPS()
	return nil
}

func (v *viewStarting) Draw(screen *ebiten.Image) {
	screen.DrawImage(v.getImg(), &ebiten.DrawImageOptions{})
}

func (v *viewStarting) getImg() *ebiten.Image {
	if v.counter <= util.GlitterCond {
		return util.BackgroundImg
	}
	return util.BackgroundImg1
}
