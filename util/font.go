package util

import (
	"bytes"
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed assets/font.ttf
var fontEmbed []byte
var FontFaceSource *text.GoTextFaceSource

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fontEmbed))
	if err != nil {
		panic(err)
	}
	FontFaceSource = s
}
