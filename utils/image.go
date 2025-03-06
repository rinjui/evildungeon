package utils

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nfnt/resize"
)

//go:embed assets/gamestartingbg.png
var gameStartingBGImgByte []byte
var GameStartingBGImg *ebiten.Image

//go:embed assets/gamestarting.png
var gameStartingImgByte []byte
var GameStartingImg *ebiten.Image

//go:embed assets/gamingbg.png
var gamingImgByte []byte
var GamingImg *ebiten.Image

//go:embed assets/gamingboundarytile.png
var gamingBoundaryTileImgByte []byte
var GamingBoundaryTileImg *ebiten.Image

//go:embed assets/tile0.png
var tile0ImgByte []byte
var Tile0Img *ebiten.Image

//go:embed assets/tile1.png
var tile1ImgByte []byte
var Tile1Img *ebiten.Image

//go:embed assets/tile2.png
var tile2ImgByte []byte
var Tile2Img *ebiten.Image

//go:embed assets/door.png
var doorImgByte []byte
var DoorImg *ebiten.Image

//go:embed assets/cursor.png
var cursorImgByte []byte
var CursorImg *ebiten.Image

func InitResource() {
	img, _, err := image.Decode(bytes.NewReader(gameStartingBGImgByte))
	if err != nil {
		panic(err)
	}

	x, y := ebiten.WindowSize()
	ratioX, ratioY := calcFactors(img, x, y)

	GameStartingBGImg = DecodeImage(gameStartingBGImgByte, ratioX, ratioY)
	GameStartingImg = DecodeImage(gameStartingImgByte, ratioX, ratioY)
	GamingImg = DecodeImage(gamingImgByte, ratioX, ratioY)
	GamingBoundaryTileImg = DecodeImage(gamingBoundaryTileImgByte, ratioX, ratioY)
	Tile0Img = DecodeImage(tile0ImgByte, ratioX, ratioY)
	Tile1Img = DecodeImage(tile1ImgByte, ratioX, ratioY)
	Tile2Img = DecodeImage(tile2ImgByte, ratioX, ratioY)
	DoorImg = DecodeImage(doorImgByte, ratioX, ratioY)
	CursorImg = DecodeImage(cursorImgByte, ratioX, ratioY)
}

func DecodeImage(b []byte, ratioX, ratioY float64) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	newX := uint(float64(img.Bounds().Dx()) * ratioX)
	newY := uint(float64(img.Bounds().Dy()) * ratioY)
	return ebiten.NewImageFromImage(resize.Resize(newX, newY, img, resize.Lanczos3))
}

func calcFactors(img image.Image, newX, newY int) (float64, float64) {
	scaleX := float64(newX) / float64(img.Bounds().Dx())
	scaleY := float64(newY) / float64(img.Bounds().Dy())

	return scaleX, scaleY
}
