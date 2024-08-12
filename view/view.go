package view

import "github.com/hajimehoshi/ebiten/v2"

type ViewID int

const (
	ViewIDStarting ViewID = iota
	ViewIDGaming
	ViewIDEnding
)

type View interface {
	Update() error

	Draw(screen *ebiten.Image)
}
