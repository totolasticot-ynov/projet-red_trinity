package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

var inventaire bool
var inventaireClicked bool

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

	playlevel1Music()
	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame_dojo.Bounds().Dx())
	scaleY := 600 / float64(bgGame_dojo.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame_dojo, op)

	if backBtn != nil {
		optback := &ebiten.DrawImageOptions{}
		optback.GeoM.Translate(float64(backRect.Min.X), float64(backRect.Min.Y))
		screen.DrawImage(backBtn, optback)
	}

	if fightplay != nil {
		optfight := &ebiten.DrawImageOptions{}
		optfight.GeoM.Translate(float64(fightRect.Min.X), float64(fightRect.Min.Y))
		screen.DrawImage(fightplay, optfight)
	}
	if oracleplayer != nil {
		optoracle := &ebiten.DrawImageOptions{}
		optoracle.GeoM.Scale(0.5, 0.5)
		optoracle.GeoM.Translate(float64(oracleRect.Min.X), float64(oracleRect.Min.Y))
		screen.DrawImage(oracleplayer, optoracle)
	}
	drawRoundedRect(screen, 400, 300, 300, 200, 20, color.RGBA{0, 0, 0, 255}, "boutique combat")

	if dollarBtn != nil {
		optdollar := &ebiten.DrawImageOptions{}
		optdollar.GeoM.Scale(0.2, 0.2)
		optdollar.GeoM.Translate(float64(dollarRect.Min.X), float64(dollarRect.Min.Y))
		screen.DrawImage(dollarBtn, optdollar)
	}
	if pilredBtn != nil {

		optpilred := &ebiten.DrawImageOptions{}
		optpilred.GeoM.Scale(0.2, 0.2)
		optpilred.GeoM.Translate(float64(pilredRect.Min.X), float64(pilredRect.Min.Y))
		screen.DrawImage(pilredBtn, optpilred)
	}
	if pilblueBtn != nil {
		optpilblue := &ebiten.DrawImageOptions{}
		optpilblue.GeoM.Scale(0.2, 0.2)
		optpilblue.GeoM.Translate(float64(pilblueRect.Min.X), float64(pilblueRect.Min.Y))
		screen.DrawImage(pilblueBtn, optpilblue)
	}
	text.Draw(screen, "la moula", basicfont.Face7x13, 590, 327, color.White)
}

func DrawGame_dojo_after(screen *ebiten.Image) {
	playlevel1Music()
	// inventaire
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if inventaireOffRect.Min.X <= x && x <= inventaireOffRect.Max.X &&
			inventaireOffRect.Min.Y <= y && y <= inventaireOffRect.Max.Y {
			if !inventaireClicked { // détecte le tout premier clic
				inventaire = !inventaire
				inventaireClicked = true
			}
		}
	} else {
		inventaireClicked = false // reset quand le clic est relâché
	}
	if bgGame_dojo != nil {
		op := &ebiten.DrawImageOptions{}
		scaleX := 800 / float64(bgGame_dojo.Bounds().Dx())
		scaleY := 600 / float64(bgGame_dojo.Bounds().Dy())
		op.GeoM.Scale(scaleX, scaleY)
		screen.DrawImage(bgGame_dojo, op)
	}

	if neoplayer != nil {
		optneo := &ebiten.DrawImageOptions{}
		optneo.GeoM.Scale(0.5, 0.5)
		optneo.GeoM.Translate(float64(neo_playerRect.Min.X), float64(neo_playerRect.Min.Y))
		screen.DrawImage(neoplayer, optneo)
	}
	if morpheusplayer != nil {
		optmor := &ebiten.DrawImageOptions{}
		optmor.GeoM.Scale(0.5, 0.5)
		optmor.GeoM.Translate(float64(morpheusRect.Min.X), float64(morpheusRect.Min.Y))
		screen.DrawImage(morpheusplayer, optmor)
	}
	optinventaire := &ebiten.DrawImageOptions{}

	if !inventaire {
		if inventaireOffBtn != nil {
			optinventaire.GeoM.Translate(float64(inventaireOffRect.Min.X), float64(inventaireOffRect.Min.Y))
			screen.DrawImage(inventaireOffBtn, optinventaire)
		}
	} else {
		if inventaireOnBtn != nil {
			optinventaire.GeoM.Translate(float64(inventaireOnRect.Min.X), float64(inventaireOnRect.Min.Y))
			screen.DrawImage(inventaireOnBtn, optinventaire)
		}
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(0.15, 0.15)

		drawRoundedRect(screen, 150, 100, 100, 450, 20, color.RGBA{255, 255, 255, 255}, "")

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
