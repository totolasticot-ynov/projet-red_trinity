package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func UpdateGame_place() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if backRect.Min.X <= x && x <= backRect.Max.X &&
			backRect.Min.Y <= y && y <= backRect.Max.Y {
			SetState("menu")
		}

		if fightRect.Min.X <= x && x <= fightRect.Max.X &&
			fightRect.Min.Y <= y && y <= fightRect.Max.Y {
			SetState("boutique_place")
		}
	}
}

func DrawGame_place_before(screen *ebiten.Image) {
	playlevel2Music()

	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame_place.Bounds().Dx())
	scaleY := 600 / float64(bgGame_place.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame_place, op)

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

func DrawGame_place_after(screen *ebiten.Image) {
	playlevel2Music()

	// Affichage du round actuel
	roundText := fmt.Sprintf("Round %d/5", currentRound)
	text.Draw(screen, roundText, basicfont.Face7x13, 350, 30, color.White)

	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if pressed && !mouseDown && !gameFinished {
		x, y := ebiten.CursorPosition()

		if inventaireOffRect.Min.X <= x && x <= inventaireOffRect.Max.X &&
			inventaireOffRect.Min.Y <= y && y <= inventaireOffRect.Max.Y {
			inventaire = !inventaire
		} else if inventaire {
			if boxeRect.Min.X <= x && x <= boxeRect.Max.X &&
				boxeRect.Min.Y <= y && y <= boxeRect.Max.Y {
				lancerCombat("Boxe")
			} else if judoRect.Min.X <= x && x <= judoRect.Max.X &&
				judoRect.Min.Y <= y && y <= judoRect.Max.Y {
				lancerCombat("Judo")
			} else if jujutsuRect.Min.X <= x && x <= jujutsuRect.Max.X &&
				jujutsuRect.Min.Y <= y && y <= jujutsuRect.Max.Y {
				lancerCombat("Jujutsu")
			} else if karateRect.Min.X <= x && x <= karateRect.Max.X &&
				karateRect.Min.Y <= y && y <= karateRect.Max.Y {
				lancerCombat("Karate")
			} else if lutteRect.Min.X <= x && x <= lutteRect.Max.X &&
				lutteRect.Min.Y <= y && y <= lutteRect.Max.Y {
				lancerCombat("Lutte")
			}
		}

		mouseDown = true
	} else if !pressed {
		mouseDown = false
	}

	if bgGame_place != nil {
		op := &ebiten.DrawImageOptions{}
		scaleX := 800 / float64(bgGame_place.Bounds().Dx())
		scaleY := 600 / float64(bgGame_place.Bounds().Dy())
		op.GeoM.Scale(scaleX, scaleY)
		screen.DrawImage(bgGame_place, op)
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

	DrawInventaire(screen)
	DrawRounds(screen)

	if combatResult != "" {
		drawRoundedRect(screen, 250, 300, 100, 100, 20, color.RGBA{0, 255, 0, 255}, "")
		drawRoundedRect(screen, 500, 300, 100, 100, 20, color.RGBA{255, 0, 0, 255}, "")
		text.Draw(screen, combatResult, basicfont.Face7x13, 300, 550, color.White)
	}
}
