package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var inventaire bool = false

func UpdateGame_dojo() {
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
			SetState("combat")
		}
		if inventaireOffRect.Min.X <= x && x <= inventaireOffRect.Max.X &&
			inventaireOffRect.Min.Y <= y && y <= inventaireOffRect.Max.Y {
			inventaire = true
		}
	}
}

func DrawGame_dojo_before(screen *ebiten.Image) {
	// background jeu
	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame_dojo.Bounds().Dx())
	scaleY := 600 / float64(bgGame_dojo.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame_dojo, op)

	// bouton retour
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(backRect.Min.X), float64(backRect.Min.Y))
	screen.DrawImage(backBtn, opts)

	// bouton fight
	opts1 := &ebiten.DrawImageOptions{}
	opts1.GeoM.Translate(float64(fightRect.Min.X), float64(fightRect.Min.Y))
	screen.DrawImage(fightplay, opts1)
}

func DrawGame_dojo_after(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame_dojo.Bounds().Dx())
	scaleY := 600 / float64(bgGame_dojo.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame_dojo, op)

	// Personnage
	opts := &ebiten.DrawImageOptions{}

	if neoplayer != nil {
		opts.GeoM.Translate(float64(neo_playerRect.Min.X), float64(neo_playerRect.Min.Y))
		screen.DrawImage(neoplayer, opts)
	}
	if morpheusplayer != nil {
		opts.GeoM.Translate(float64(morpheusRect.Min.X), float64(morpheusRect.Min.Y))
		screen.DrawImage(morpheusplayer, opts)
	}
	if !inventaire {
		opts.GeoM.Translate(float64(inventaireOffRect.Min.X), float64(inventaireOffRect.Min.Y))
		screen.DrawImage(inventaireOffBtn, opts)
	} else {
		opts.GeoM.Translate(float64(inventaireOnRect.Min.X), float64(inventaireOnRect.Min.Y))
		screen.DrawImage(inventaireOnBtn, opts)
	}

}
