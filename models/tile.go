package models

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	Bg       *ebiten.Image
	X, Y     int
	Entities []*Entity
}
