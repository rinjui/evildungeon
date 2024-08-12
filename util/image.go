package util

import (
	"bytes"
	_ "embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/background.png
var backgroundImgByte []byte
var BackgroundImg *ebiten.Image

//go:embed assets/background-1.png
var backgroundImg1Byte []byte
var BackgroundImg1 *ebiten.Image

//go:embed assets/tiles.png
var tilesByte []byte
var TilesImg *ebiten.Image

// //go:embed assets/ground.png
// var groundByte []byte
// var GroundImg *ebiten.Image

// on the ground ui
var OgSkyBgTileImg image.Image // 32x32
var OgGrassTileImg image.Image // 32x16
var OgSandTileImg image.Image  // 32x16

// underground ui
var UgSandTileImg image.Image // 32x32
var UgMineTileImg image.Image // 32x32

// special ui
var SCursorTileImg image.Image // 32x32

func init() {
	img, _, err := image.Decode(bytes.NewReader(backgroundImgByte))
	if err != nil {
		panic(err)
	}
	BackgroundImg = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(backgroundImg1Byte))
	if err != nil {
		panic(err)
	}
	BackgroundImg1 = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(tilesByte))
	if err != nil {
		panic(err)
	}
	TilesImg = ebiten.NewImageFromImage(img)

	// img, _, err = image.Decode(bytes.NewReader(groundByte))
	// if err != nil {
	// 	panic(err)
	// }
	// GroundImg = ebiten.NewImageFromImage(img)

	OgSkyBgTileImg = TilesImg.SubImage(image.Rect(32*7, 32*5, 32*8, 32*6))
	OgGrassTileImg = TilesImg.SubImage(image.Rect(32*2, 32*3, 32*3, 32*4-16))
	OgSandTileImg = TilesImg.SubImage(image.Rect(32*7, 32*1, 32*8, 32*2-24))

	UgSandTileImg = TilesImg.SubImage(image.Rect(32*0, 32*5, 32*1, 32*6))
	UgMineTileImg = TilesImg.SubImage(image.Rect(32*4, 32*4, 32*5, 32*5))
	SCursorTileImg = TilesImg.SubImage(image.Rect(32*0, 32*6, 32*1, 32*7))
}
