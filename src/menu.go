package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bgMenu   *ebiten.Image
	playBtn  *ebiten.Image
	playRect image.Rectangle
)

func init() {
	var err error
	bgMenu, _, err = ebitenutil.NewImageFromFile("images/bg.png")
	if err != nil {
		log.Fatal(err)
	}
	playBtn, _, err = ebitenutil.NewImageFromFile("images/play.png")
	if err != nil {
		log.Fatal(err)
	}
	playRect = image.Rect(300, 200, 300+playBtn.Bounds().Dx(), 200+playBtn.Bounds().Dy())
}

func UpdateMenu() {
	// Si clic gauche dans playRect → changer de state
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if playRect.Min.X <= x && x <= playRect.Max.X &&
			playRect.Min.Y <= y && y <= playRect.Max.Y {
			SetState("game")
		}
	}
}

func DrawMenu(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgMenu.Bounds().Dx()) // largeur fenêtre / largeur image
	scaleY := 600 / float64(bgMenu.Bounds().Dy()) // hauteur fenêtre / hauteur image
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgMenu, op)

	// Bouton Play
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(playRect.Min.X), float64(playRect.Min.Y))
	screen.DrawImage(playBtn, opts)
}
