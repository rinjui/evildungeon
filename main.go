package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rinjui/evildungeon/utils"
	"github.com/rinjui/evildungeon/view"
)

const (
	// screenWidth  = 720
	// screenHeight = 1280
	screenWidth  = 360
	screenHeight = 640
)

func main() {
	fmt.Println("MAX TPS: ", ebiten.TPS())
	fmt.Println("SCREEN: ", screenWidth, screenHeight)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Evil Dungeon")

	utils.InitResource()

	g := NewGame()
	g.RegisterView(view.ViewIDStarting, view.NewViewStarting(g))
	g.RegisterView(view.ViewIDGaming, view.NewViewGaming(g))
	g.RegisterView(view.ViewIDEnding, view.NewViewEnding(g))

	g.Redirect(view.ViewIDStarting)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
