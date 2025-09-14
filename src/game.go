package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bgGame   *ebiten.Image
	backBtn  *ebiten.Image
	backRect image.Rectangle
)

func init() {
	var err error
	bgGame, _, err = ebitenutil.NewImageFromFile("images/bg.png") // même fond pour l'instant
	if err != nil {
		log.Fatal(err)
	}
	backBtn, _, err = ebitenutil.NewImageFromFile("images/back.png")
	if err != nil {
		log.Fatal(err)
	}
	backRect = image.Rect(20, 20, 20+backBtn.Bounds().Dx(), 20+backBtn.Bounds().Dy())
}

func UpdateGame() {
	// clic sur bouton retour → revenir au menu
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if backRect.Min.X <= x && x <= backRect.Max.X &&
			backRect.Min.Y <= y && y <= backRect.Max.Y {
			SetState("menu")
		}
	}
	// touche Echap pour revenir aussi
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		SetState("menu")
	}
}

func DrawGame(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame.Bounds().Dx())
	scaleY := 600 / float64(bgGame.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame, op)

	// Bouton back
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(backRect.Min.X), float64(backRect.Min.Y))
	screen.DrawImage(backBtn, opts)
}
