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
		// bouton fight
		if fightRect.Min.X <= x && x <= fightRect.Max.X &&
			fightRect.Min.Y <= y && y <= fightRect.Max.Y {
			SetState("combat") // va au combat
		}
	}
}

func DrawGame_place_before(screen *ebiten.Image) {
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

	// bouton fight
	opts1 := &ebiten.DrawImageOptions{}
	opts1.GeoM.Translate(float64(fightRect.Min.X), float64(fightRect.Min.Y))
	screen.DrawImage(fightplay, opts1)
}

func DrawGame_place_after(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame_place.Bounds().Dx())
	scaleY := 600 / float64(bgGame_place.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame_place, op)

	// Images si elles existent
	if neoplayer != nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(neo_playerRect.Min.X), float64(neo_playerRect.Min.Y))
		screen.DrawImage(neoplayer, opts)
	}
	if morpheusplayer != nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(morpheusRect.Min.X), float64(morpheusRect.Min.Y))
		screen.DrawImage(morpheusplayer, opts)
	}
}
