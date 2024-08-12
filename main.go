package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rinjui/evildungeon/util"
	"github.com/rinjui/evildungeon/view"
)

const (
// screenWidth  = 1280
// screenHeight = 720
// screenWidth  = 640
// screenHeight = 480
)

func main() {
	ebiten.SetWindowSize(util.ScreenWidth, util.ScreenHeight)
	ebiten.SetWindowTitle("Evil Dungeon")

	g := NewGame()
	g.RegisterView(view.ViewIDStarting, view.NewViewStarting(g))
	g.RegisterView(view.ViewIDGaming, view.NewViewGaming(g))
	g.RegisterView(view.ViewIDEnding, view.NewViewEnding(g))

	g.Redirect(view.ViewIDStarting)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
