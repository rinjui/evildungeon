package view

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rinjui/evildungeon/models"
	"github.com/rinjui/evildungeon/utils"
)

const (
	tileXCnt = 25
	tileYCnt = 25
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
	windowX, windoxY := ebiten.WindowSize()
	tileX := utils.CursorImg.Bounds().Dx()
	tileY := utils.CursorImg.Bounds().Dy()
	layoutX, layoutY := tileXCnt*tileX, tileYCnt*tileY

	halfWindowX := windowX / 2
	halfWindowY := windoxY / 2

	// calcuate the real position of cursor
	cursorX := v.layout.tileCursorX * tileX
	cursorY := v.layout.tileCursorY * tileY

	leftX := cursorX - halfWindowX
	rightX := cursorX + halfWindowX
	upY := cursorY - halfWindowY
	downY := cursorY + halfWindowY

	if leftX <= 0 {
		leftX = 0
		rightX = windowX
	}

	if rightX >= layoutX {
		leftX = layoutX - windowX
		rightX = layoutX
	}

	if upY <= 0 {
		upY = 0
		downY = windoxY
	}

	if downY >= layoutY {
		upY = layoutY - windoxY
		downY = layoutY
	}

	leftTile := leftX / tileX
	rightTile := rightX / tileX
	upTile := upY / tileY
	downTile := downY / tileY

	fmt.Println(v.layout.tileCursorX, v.layout.tileCursorY, cursorX, cursorY)
	fmt.Println(leftX, rightX, upY, downY, halfWindowX, halfWindowY, layoutX, layoutY)
	fmt.Println(leftTile, rightTile, upTile, downTile)
	op := &ebiten.DrawImageOptions{}
	for i := upTile; i < downTile; i++ {
		for j := leftTile; j < rightTile; j++ {
			img := v.layout.tiles[i][j].GetImage()
			if img == nil {
				continue
			}

			op.GeoM.Reset()
			op.GeoM.Translate(float64((j-leftTile)*tileX), float64((i-upTile)*tileY))
			screen.DrawImage(img, op)
		}
	}
}

func (v *viewGaming) drawCursor(screen *ebiten.Image) {
	windowX, windoxY := ebiten.WindowSize()
	tileX := utils.CursorImg.Bounds().Dx()
	tileY := utils.CursorImg.Bounds().Dy()
	layoutX, layoutY := tileXCnt*tileX, tileYCnt*tileY

	halfWindowX := windowX / 2
	halfWindowY := windoxY / 2

	// calcuate the real position of cursor
	cursorX := v.layout.tileCursorX * tileX
	cursorY := v.layout.tileCursorY * tileY

	// put in "middle" position if it's "far away" the boundary
	if cursorX >= halfWindowX && cursorX <= layoutX-halfWindowX {
		cursorX = halfWindowX
	}

	if cursorY >= halfWindowY && cursorY <= layoutY-halfWindowY {
		cursorY = halfWindowY
	}

	// adjust the deviation
	if cursorX%tileX != 0 {
		cursorX = int(cursorX/tileX) * tileX
	}

	if cursorY%tileY != 0 {
		cursorY = int(cursorY/tileY) * tileY
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cursorX), float64(cursorY))
	screen.DrawImage(utils.CursorImg, op)
}

func newLayout() *layout {
	tiles := make([][]*models.Tile, tileYCnt)
	for i := range tiles {
		tiles[i] = make([]*models.Tile, tileXCnt)
		for j := range tiles[i] {
			tiles[i][j] = &models.Tile{Point: utils.RInt() % 200}
		}
	}

	return &layout{
		tiles:       tiles,
		tileCursorX: tileXCnt / 2,
		tileCursorY: tileYCnt / 2,
	}
}

type layout struct {
	tiles                    [][]*models.Tile
	tileCursorX, tileCursorY int
	count                    int
}

func (l *layout) Update() {
	l.DigUpdate()
	l.CursorUpdate()
	l.TilePointUpdate()
}

func (l *layout) CursorUpdate() {
	// keep pressing
	if d := inpututil.KeyPressDuration(ebiten.KeyUp); d > 1 && d%(ebiten.TPS()/10) == 0 {
		l.tileCursorY--
	}

	if d := inpututil.KeyPressDuration(ebiten.KeyDown); d > 1 && d%(ebiten.TPS()/10) == 0 {
		l.tileCursorY++
	}

	if d := inpututil.KeyPressDuration(ebiten.KeyLeft); d > 1 && d%(ebiten.TPS()/10) == 0 {
		l.tileCursorX--
	}

	if d := inpututil.KeyPressDuration(ebiten.KeyRight); d > 1 && d%(ebiten.TPS()/10) == 0 {
		l.tileCursorX++
	}

	if l.tileCursorX < 0 {
		l.tileCursorX = 0
	}

	if l.tileCursorX >= tileXCnt {
		l.tileCursorX = tileXCnt - 1
	}

	if l.tileCursorY < 0 {
		l.tileCursorY = 0
	}

	if l.tileCursorY >= tileYCnt {
		l.tileCursorY = tileYCnt - 1
	}

	// 滑鼠操控
	// x, y := ebiten.WindowSize()
	// cursorX, cursorY := ebiten.CursorPosition()
	// if cursorX <= 0 || cursorX >= x {
	// 	return
	// }

	// if cursorY <= 0 || cursorY >= y {
	// 	return
	// }

	// tileX, tileY := utils.Tile0Img.Bounds().Dx(), utils.Tile0Img.Bounds().Dy()
	// l.cursorX = cursorX / tileX
	// l.cursorY = cursorY / tileY
	// x, y := ebiten.WindowPosition()
	// fmt.Println(ebiten.CursorPosition())
	// fmt.Println(ebiten.WindowPosition())
}

func (l *layout) DigUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		fmt.Println(l.tileCursorX, l.tileCursorY)
		l.tiles[l.tileCursorY][l.tileCursorX].Point = -1
	}
}

func (l *layout) TilePointUpdate() {
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
