package view

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rinjui/evildungeon/models"
	"github.com/rinjui/evildungeon/util"
)

var (
	ogSkyHeight   = 3
	ogGrassHeight = 1
	ogSandHeight  = 1
	ogTotalHeight = ogSkyHeight + ogGrassHeight + ogSandHeight
	ugTotalHeight = 50
	gameHeight    = ogTotalHeight + ugTotalHeight
	gameWidth     = 50
)

func NewViewGaming(c Controller) View {
	return &viewGaming{
		ctl:     c,
		cursorX: util.ScreenWidth / 2,
		cursorY: util.ScreenHeight / 2,
		layout:  newLayout(),
	}
}

func newLayout() [][]*models.Tile {

	ogSkyImg := util.OgSkyBgTileImg.(*ebiten.Image)
	ogSkyImgW, ogSkyImgH := ogSkyImg.Bounds().Dx(), ogSkyImg.Bounds().Dy()

	ogGrassImg := util.OgGrassTileImg.(*ebiten.Image)
	ogGrassImgW, ogGrassImgH := ogGrassImg.Bounds().Dx(), ogGrassImg.Bounds().Dy()

	ogSandImg := util.OgSandTileImg.(*ebiten.Image)
	ogSandImgW, ogSandImgH := ogSandImg.Bounds().Dx(), ogSandImg.Bounds().Dy()

	ugMineImg := util.UgMineTileImg.(*ebiten.Image)
	ugMineImgW, ugMineImgH := ugMineImg.Bounds().Dx(), ugMineImg.Bounds().Dy()

	height, Y := 0, 0
	layout := make([][]*models.Tile, gameHeight)
	for i := 1; i < ogSkyHeight; i++ {
		layout[height] = make([]*models.Tile, gameWidth)
		for j := 0; j < gameWidth; j++ {
			layout[height][j] = &models.Tile{
				Bg: ogSkyImg,
				X:  j * ogSkyImgW,
				Y:  Y,
			}
		}
		height++
		Y += ogSkyImgH
	}

	for i := 0; i < ogGrassHeight; i++ {
		layout[height] = make([]*models.Tile, gameWidth)
		for j := 0; j < gameWidth; j++ {
			layout[height][j] = &models.Tile{
				Bg: ogGrassImg,
				X:  j * ogGrassImgW,
				Y:  Y,
			}
		}
		height++
		Y += ogGrassImgH
	}

	for i := 0; i < ogSandHeight; i++ {
		layout[height] = make([]*models.Tile, gameWidth)
		for j := 0; j < gameWidth; j++ {
			layout[height][j] = &models.Tile{
				Bg: ogSandImg,
				X:  j * ogSandImgW,
				Y:  Y,
			}
		}
		height++
		Y += ogSandImgH
	}

	for i := 0; i < ugTotalHeight; i++ {
		layout[height] = make([]*models.Tile, gameWidth)
		for j := 0; j < gameWidth; j++ {
			layout[height][j] = &models.Tile{
				Bg: ugMineImg,
				X:  j * ugMineImgW,
				Y:  Y,
			}
		}
		height++
		Y += ugMineImgH
	}

	return layout
}

type viewGaming struct {
	ctl Controller

	counter          int
	cursorX, cursorY int
	layout           [][]*models.Tile
}

func (v *viewGaming) Update() error {
	v.counter++
	v.counter = v.counter % ebiten.TPS()

	return nil
}

func (v *viewGaming) Draw(screen *ebiten.Image) {
	for _, l := range v.layout {
		for _, t := range l {
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(t.X), float64(t.Y))

			screen.DrawImage(t.Bg, &op)
		}
	}

	if v.counter <= util.GlitterCond {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(v.cursorX), float64(v.cursorY))
		screen.DrawImage(util.SCursorTileImg.(*ebiten.Image), &op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
