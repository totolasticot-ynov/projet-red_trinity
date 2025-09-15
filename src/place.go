package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateGame_place() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// bouton retour
		if backRect.Min.X <= x && x <= backRect.Max.X &&
			backRect.Min.Y <= y && y <= backRect.Max.Y {
			SetState("menu")
		}
	}
}

func DrawGame_place(screen *ebiten.Image) {
	// background jeu
	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame_place.Bounds().Dx())
	scaleY := 600 / float64(bgGame_place.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame_place, op)

	// bouton retour
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(backRect.Min.X), float64(backRect.Min.Y))
	screen.DrawImage(backBtn, opts)

	// personnage
	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Translate(float64(neo_playerRect.Min.X), float64(neo_playerRect.Min.Y))
	screen.DrawImage(neoplayer, opts2)

	// personnage 2
	opts3 := &ebiten.DrawImageOptions{}
	opts3.GeoM.Translate(float64(morpheusRect.Min.X), float64(morpheusRect.Min.Y))
	screen.DrawImage(morpheusplayer, opts3)
}
