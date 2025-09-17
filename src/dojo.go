package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func UpdateGame_dojo() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if backRect.Min.X <= x && x <= backRect.Max.X &&
			backRect.Min.Y <= y && y <= backRect.Max.Y {
			SetState("menu")
		}

		if fightRect.Min.X <= x && x <= fightRect.Max.X &&
			fightRect.Min.Y <= y && y <= fightRect.Max.Y {
			SetState("boutique_dojo")
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
			if 180 <= x && x <= 260 &&
				130 <= y && y <= 200 {
				resultImg = boxeBtn
				lancerCombat("Boxe")
			} else if 180 <= x && x <= 230 &&
				210 <= y && y <= 270 {
				resultImg = judoBtn
				lancerCombat("Judo")
			} else if 180 <= x && x <= 260 &&
				290 <= y && y <= 350 {
				resultImg = jujutsuBtn
				lancerCombat("Jujutsu")
			} else if 180 <= x && x <= 260 &&
				380 <= y && y <= 440 {
				resultImg = karateBtn
				lancerCombat("Karate")
			} else if 180 <= x && x <= 260 &&
				470 <= y && y <= 530 {
				resultImg = lutteBtn
				lancerCombat("Lutte")
			} else {
				lancerCombat("miss")
			}
		}

		mouseDown = true
	} else if !pressed {
		mouseDown = false
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

	DrawInventaire(screen)
	DrawRounds(screen)

	if combatResult != "" {
		if toi {
			couleur_toi = vert
			couleur_enemie = rouge
		} else if enemie {
			couleur_toi = rouge
			couleur_enemie = vert
		} else { // égalité
			couleur_toi = color.RGBA{200, 200, 0, 255}
			couleur_enemie = color.RGBA{200, 200, 0, 255}
		}
		drawRoundedRect(screen, 250, 300, 100, 100, 20, couleur_toi, "")
		// Carré rouge
		drawRoundedRect(screen, 500, 300, 100, 100, 20, couleur_enemie, "")

		// Affichage du choix du joueur dans le carré vert
		text.Draw(screen, combatResult, basicfont.Face7x13, 275, 500, color.Black)

		// Optionnel : dessiner l'image du choix
		optresult := &ebiten.DrawImageOptions{}
		optresult.GeoM.Scale(0.2, 0.2)
		optresult.GeoM.Translate(250, 300) // carré rouge
		screen.DrawImage(resultImg, optresult)

		// Affichage du choix de l'adversaire dans le carré rouge
		optresult_red := &ebiten.DrawImageOptions{}
		optresult_red.GeoM.Scale(0.2, 0.2)
		optresult_red.GeoM.Translate(500, 300) // carré rouge
		screen.DrawImage(resultimg_red, optresult_red)
	}
}
