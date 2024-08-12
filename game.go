package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/rinjui/evildungeon/view"
)

func NewGame() *game {
	return &game{
		views: make(map[view.ViewID]view.View),
	}
}

type game struct {
	currentView view.View
	views       map[view.ViewID]view.View
}

func (g *game) RegisterView(id view.ViewID, v view.View) {
	if _, ok := g.views[id]; ok {
		panic("duplicated view")
	}

	g.views[id] = v
}

func (g *game) Redirect(id view.ViewID) error {
	v, ok := g.views[id]
	if !ok {
		return fmt.Errorf("unknown view")
	}

	g.currentView = v
	return nil
}

func (g *game) Update() error {
	return g.currentView.Update()
}

func (g *game) Draw(screen *ebiten.Image) {
	g.currentView.Draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// return 640, 480
	// fmt.Println(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}

func (g *game) SetCurrentView(v view.View) {
	g.currentView = v
}
