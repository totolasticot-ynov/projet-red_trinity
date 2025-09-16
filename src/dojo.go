package main

import (
	"log"

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
	}
}

func DrawGame_dojo_before(screen *ebiten.Image) {
	if bgGame_dojo == nil {
		log.Println("bgGame_dojo est nil")
		return
	}
	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame_dojo.Bounds().Dx())
	scaleY := 600 / float64(bgGame_dojo.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame_dojo, op)

	if backBtn != nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(backRect.Min.X), float64(backRect.Min.Y))
		screen.DrawImage(backBtn, opts)
	}

	if fightplay != nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(fightRect.Min.X), float64(fightRect.Min.Y))
		screen.DrawImage(fightplay, opts)
	}
}

func DrawGame_dojo_after(screen *ebiten.Image) {
	// inventaire
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if inventaireOffRect.Min.X <= x && x <= inventaireOffRect.Max.X &&
			inventaireOffRect.Min.Y <= y && y <= inventaireOffRect.Max.Y {
			if !inventaire {
				inventaire = true
			} else {
				inventaire = false
			}
		}
	}
	if bgGame_dojo != nil {
		op := &ebiten.DrawImageOptions{}
		scaleX := 800 / float64(bgGame_dojo.Bounds().Dx())
		scaleY := 600 / float64(bgGame_dojo.Bounds().Dy())
		op.GeoM.Scale(scaleX, scaleY)
		screen.DrawImage(bgGame_dojo, op)
	}

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
	opts := &ebiten.DrawImageOptions{}

	if !inventaire {
		if inventaireOffBtn != nil {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(inventaireOffRect.Min.X), float64(inventaireOffRect.Min.Y))
			screen.DrawImage(inventaireOffBtn, opts)
		}
	} else {
		if inventaireOnBtn != nil {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(inventaireOnRect.Min.X), float64(inventaireOnRect.Min.Y))
			screen.DrawImage(inventaireOnBtn, opts)
		}
		if boxeBtn != nil {
			opts.GeoM.Translate(float64(boxeRect.Min.X), float64(boxeRect.Min.Y))
			screen.DrawImage(boxeBtn, opts)
		}
		if judoBtn != nil {
			opts.GeoM.Translate(float64(judoRect.Min.X), float64(judoRect.Min.Y))
			screen.DrawImage(judoBtn, opts)
		}
		if jujutsuBtn != nil {
			opts.GeoM.Translate(float64(jujutsuRect.Min.X), float64(jujutsuRect.Min.Y))
			screen.DrawImage(jujutsuBtn, opts)
		}
		if karateBtn != nil {
			opts.GeoM.Translate(float64(karateRect.Min.X), float64(karateRect.Min.Y))
			screen.DrawImage(karateBtn, opts)
		}
		if lutteBtn != nil {
			opts.GeoM.Translate(float64(lutteRect.Min.X), float64(lutteRect.Min.Y))
			screen.DrawImage(lutteBtn, opts)
		}
	}
}
