package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rinjui/evildungeon/models"
	"github.com/rinjui/evildungeon/utils"
)

const (
	layoutX = 60
	layoutY = 35
)

func NewViewGaming(c Controller) View {
	return &viewGaming{
		ctl:    c,
		layout: newLayout(),
	}
}

type viewGaming struct {
	ctl    Controller
	layout *layout
}

func (v *viewGaming) Update() error {
	// 長按
	if d := inpututil.KeyPressDuration(ebiten.KeyUp); d > 1 && d%(ebiten.TPS()/10) == 0 {
		v.layout.cursorY--
	}

	if d := inpututil.KeyPressDuration(ebiten.KeyDown); d > 1 && d%(ebiten.TPS()/10) == 0 {
		v.layout.cursorY++
	}

	if d := inpututil.KeyPressDuration(ebiten.KeyLeft); d > 1 && d%(ebiten.TPS()/10) == 0 {
		v.layout.cursorX--
	}

	if d := inpututil.KeyPressDuration(ebiten.KeyRight); d > 1 && d%(ebiten.TPS()/10) == 0 {
		v.layout.cursorX++
	}

	if v.layout.cursorX < 0 {
		v.layout.cursorX = 0
	}

	if v.layout.cursorX >= layoutX {
		v.layout.cursorX = layoutX - 1
	}

	if v.layout.cursorY < 0 {
		v.layout.cursorY = 0
	}

	if v.layout.cursorY >= layoutY {
		v.layout.cursorY = layoutY - 1
	}

	v.layout.Update()
	return nil
}

func (v *viewGaming) Draw(screen *ebiten.Image) {
	v.drawBG(screen)
	v.drawLayout(screen)
	v.drawCursor(screen)
}

func (v *viewGaming) drawBG(screen *ebiten.Image) {
	screen.DrawImage(utils.GamingImg, nil)
}

func (v *viewGaming) drawLayout(screen *ebiten.Image) {
	screenX := layoutX / 3
	halfScreenX := layoutX / 6
	left := v.layout.cursorX - halfScreenX
	right := v.layout.cursorX + halfScreenX

	if v.layout.cursorX < halfScreenX {
		left = 0
		right = screenX
	}

	if v.layout.cursorX >= layoutX-halfScreenX {
		left = layoutX - screenX
		right = layoutX
	}

	tileX, tileY := utils.Tile0Img.Bounds().Dx(), utils.Tile0Img.Bounds().Dy()

	op := &ebiten.DrawImageOptions{}
	for i := range v.layout.tiles {
		op.GeoM.Reset()
		op.GeoM.Translate(0, float64(i*tileY))
		for j := left; j < right; j++ {
			img := v.layout.GetTileImage(i, j)
			if img == nil {
				continue
			}

			screen.DrawImage(img, op)
			op.GeoM.Translate(float64(tileX), 0)
		}
	}
}

func (v *viewGaming) drawCursor(screen *ebiten.Image) {
	wx, _ := ebiten.WindowSize()
	tileX := utils.CursorImg.Bounds().Dx()
	tileY := utils.CursorImg.Bounds().Dy()

	x := wx / 2
	y := v.layout.cursorY * tileY

	halfScreenX := layoutX / 6
	if v.layout.cursorX < halfScreenX {
		x = v.layout.cursorX * tileX
	}

	if v.layout.cursorX >= layoutX-halfScreenX {
		x = (v.layout.cursorX * tileX) % wx
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(utils.CursorImg, op)
}

func newLayout() *layout {
	tiles := make([][]*models.Tile, layoutY)
	for i := range tiles {
		tiles[i] = make([]*models.Tile, layoutX)
		for j := range tiles[i] {
			tiles[i][j] = &models.Tile{Point: utils.RInt() % 200}
		}
	}

	return &layout{
		tiles:   tiles,
		cursorX: layoutX / 2,
		cursorY: layoutY / 2,
	}
}

type layout struct {
	tiles            [][]*models.Tile
	cursorX, cursorY int
	count            int
}

func (l *layout) Update() {
	l.count++
	if l.count%ebiten.TPS() != 0 {
		return
	}

	for i := range l.tiles {
		for j := range l.tiles[i] {
			if l.tiles[i][j].Point == -1 {
				continue
			}

			l.tiles[i][j].Point += utils.RInt()%10 - 5
			if l.tiles[i][j].Point <= 0 {
				l.tiles[i][j].Point = 0
			}
		}
	}

	l.count = l.count % ebiten.TPS()
}

func (l *layout) GetTileImage(i, j int) *ebiten.Image {
	if i < 0 || i >= len(l.tiles) {
		return nil
	}

	if j < 0 || j >= len(l.tiles[0]) {
		return nil
	}

	tile := l.tiles[i][j]
	if tile.Point == -1 {
		return nil
	}

	if tile.Point >= 200 {
		return utils.Tile2Img
	}

	if tile.Point >= 100 {
		return utils.Tile1Img
	}

	return utils.Tile0Img
}
