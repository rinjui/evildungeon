package models

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rinjui/evildungeon/utils"
)

type Tile struct {
	Point int
}

func (t Tile) GetImage() *ebiten.Image {
	if t.Point == -1 {
		return nil
	}

	if t.Point >= 200 {
		return utils.Tile2Img
	}

	if t.Point >= 100 {
		return utils.Tile1Img
	}

	return utils.Tile0Img
}
