package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateGame() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// bouton retour
		if backRect.Min.X <= x && x <= backRect.Max.X &&
			backRect.Min.Y <= y && y <= backRect.Max.Y {
			SetState("menu")
		}
	}
}

func DrawGame(screen *ebiten.Image) {
	// background jeu
	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame.Bounds().Dx())
	scaleY := 600 / float64(bgGame.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame, op)

	// bouton retour
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(backRect.Min.X), float64(backRect.Min.Y))
	screen.DrawImage(backBtn, opts)
}
