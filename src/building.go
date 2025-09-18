package main

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

// Variables globales pour les pilules
func UpdateGame_building() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if backRect.Min.X <= x && x <= backRect.Max.X &&
			backRect.Min.Y <= y && y <= backRect.Max.Y {
			SetState("menu")
		}

		if fightRect.Min.X <= x && x <= fightRect.Max.X &&
			fightRect.Min.Y <= y && y <= fightRect.Max.Y {
			SetState("boutique_building")
		}

		// Achat des pilules dans la boutique (avant le combat)
		if pilredRect.Min.X <= x && x <= pilredRect.Max.X &&
			pilredRect.Min.Y <= y && y <= pilredRect.Max.Y {
			// Achat pilule rouge si pas encore possédée et assez d'argent
			if !pilredOwned && argent >= 30 {
				argent -= 30
				pilredOwned = true
				pilredUsed = false // Réinitialise pour utilisation
			}
		}

		if pilblueRect.Min.X <= x && x <= pilblueRect.Max.X &&
			pilblueRect.Min.Y <= y && y <= pilblueRect.Max.Y {
			// Achat pilule bleue si pas encore possédée et assez d'argent
			if !pilblueOwned && argent >= 30 {
				argent -= 30
				pilblueOwned = true
				pilblueUsed = false // Réinitialise pour utilisation
			}
		}
	}
}

func DrawGame_building_before(screen *ebiten.Image) {
	playlevel5Music()

	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame_building.Bounds().Dx())
	scaleY := 600 / float64(bgGame_building.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame_building, op)

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
	text.Draw(screen, "Pilule rouge: +1 win", basicfont.Face7x13, 450, 380, color.White)
	text.Draw(screen, "Pilule bleue: -1 ennemi", basicfont.Face7x13, 450, 400, color.White)
	text.Draw(screen, "Prix: 30 pieces chaque", basicfont.Face7x13, 450, 420, color.White)

	if dollarBtn != nil {
		optdollar := &ebiten.DrawImageOptions{}
		optdollar.GeoM.Scale(0.2, 0.2)
		optdollar.GeoM.Translate(float64(dollarRect.Min.X), float64(dollarRect.Min.Y))
		screen.DrawImage(dollarBtn, optdollar)
	}

	// Affichage des pilules avec indication si déjà possédées
	if pilredBtn != nil {
		optpilred := &ebiten.DrawImageOptions{}
		optpilred.GeoM.Scale(0.2, 0.2)
		optpilred.GeoM.Translate(float64(pilredRect.Min.X), float64(pilredRect.Min.Y))

		// Assombrir si déjà possédée
		if pilredOwned {
			optpilred.ColorM.Scale(0.5, 0.5, 0.5, 1.0)
		}
		screen.DrawImage(pilredBtn, optpilred)

		if pilredOwned {
			playringMusic()
			text.Draw(screen, "POSSEDEE", basicfont.Face7x13, pilredRect.Min.X+20, pilredRect.Min.Y+50, vert)
		}
	}

	if pilblueBtn != nil {
		optpilblue := &ebiten.DrawImageOptions{}
		optpilblue.GeoM.Scale(0.2, 0.2)
		optpilblue.GeoM.Translate(float64(pilblueRect.Min.X), float64(pilblueRect.Min.Y))

		// Assombrir si déjà possédée
		if pilblueOwned {
			optpilblue.ColorM.Scale(0.5, 0.5, 0.5, 1.0)
		}
		screen.DrawImage(pilblueBtn, optpilblue)

		if pilblueOwned {
			playringMusic()
			text.Draw(screen, "POSSEDEE", basicfont.Face7x13, pilblueRect.Min.X+20, pilblueRect.Min.Y+50, vert)
		}
	}

	text.Draw(screen, "argent: "+strconv.Itoa(argent), basicfont.Face7x13, 590, 327, color.White)
}

func DrawGame_building_after(screen *ebiten.Image) {
	playlevel5Music()

	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if pressed && !mouseDown && !gameFinished {
		x, y := ebiten.CursorPosition()
		// Utilisation de la pilule rouge (+1 pour le joueur)
		if pilredOwned && 70 <= x && x <= 120 &&
			220 <= y && y <= 270 {
			if score_toi < 5 {
				playringMusic()
				score_toi++
				pilredUsed = true   // Marque comme utilisée
				pilredOwned = false // Retire de l'inventaire, il faudra la racheter
			}
		}

		// Utilisation de la pilule bleue (-1 pour l'ennemi)
		if pilblueOwned && 70 <= x && x <= 120 &&
			100 <= y && y <= 180 {
			if score_enemie > 0 {
				playringMusic()
				score_enemie--
				pilblueUsed = true   // Marque comme utilisée
				pilblueOwned = false // Retire de l'inventaire, il faudra la racheter
			}
		}

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
				if karate {
					lancerCombat("Karate")
				}
			} else if 180 <= x && x <= 260 &&
				470 <= y && y <= 530 {
				resultImg = lutteBtn
				if lutte {
					lancerCombat("Lutte")
				}
			}
		}
		mouseDown = true
	} else if !pressed {
		mouseDown = false
	}

	if bgGame_building != nil {
		op := &ebiten.DrawImageOptions{}
		scaleX := 800 / float64(bgGame_building.Bounds().Dx())
		scaleY := 600 / float64(bgGame_building.Bounds().Dy())
		op.GeoM.Scale(scaleX, scaleY)
		screen.DrawImage(bgGame_building, op)
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

	switch score_toi {
	case 0:
		local_score_toi = num0
	case 1:
		local_score_toi = num1
	case 2:
		local_score_toi = num2
	case 3:
		local_score_toi = num3
	case 4:
		local_score_toi = num4
	case 5:
		local_score_toi = num5
	}

	optscore_toi := &ebiten.DrawImageOptions{}
	optscore_toi.GeoM.Scale(0.1, 0.1)
	optscore_toi.GeoM.Translate(200, 40)
	drawRoundedRect(screen, 200, 40, 50, 50, 10, color.RGBA{255, 255, 255, 255}, "")
	screen.DrawImage(local_score_toi, optscore_toi)

	// Affichage de la pilule rouge si possédée et pas encore utilisée
	if pilredOwned && !pilredUsed {
		if pilredBtn != nil {
			optpilred := &ebiten.DrawImageOptions{}
			optpilred.GeoM.Scale(0.2, 0.2)
			optpilred.GeoM.Translate(50, 200)
			screen.DrawImage(pilredBtn, optpilred)
		}
	}

	// Affichage de la pilule bleue si possédée et pas encore utilisée
	if pilblueOwned && !pilblueUsed {
		if pilblueBtn != nil {
			optpilblue := &ebiten.DrawImageOptions{}
			optpilblue.GeoM.Scale(0.2, 0.2)
			optpilblue.GeoM.Translate(50, 100)
			screen.DrawImage(pilblueBtn, optpilblue)
		}
	}

	switch score_enemie {
	case 0:
		local_score_enemie = num0
	case 1:
		local_score_enemie = num1
	case 2:
		local_score_enemie = num2
	case 3:
		local_score_enemie = num3
	case 4:
		local_score_enemie = num4
	case 5:
		local_score_enemie = num5
	}
	optscore_enemie := &ebiten.DrawImageOptions{}
	optscore_enemie.GeoM.Scale(0.1, 0.1)
	optscore_enemie.GeoM.Translate(550, 40)
	drawRoundedRect(screen, 550, 40, 50, 50, 10, color.RGBA{255, 255, 255, 255}, "")
	screen.DrawImage(local_score_enemie, optscore_enemie)

	if combatResult != "" {
		if score_toi > score_enemie {
			// Le joueur gagne
			couleur_toi = vert
			couleur_enemie = rouge
		} else if score_enemie > score_toi {
			// L'ennemi gagne
			couleur_toi = rouge
			couleur_enemie = vert
		} else {
			// Égalité
			couleur_toi = color.RGBA{255, 255, 0, 255}
			couleur_enemie = color.RGBA{255, 255, 0, 255}
		}

		drawRoundedRect(screen, 250, 300, 100, 100, 20, couleur_toi, "")
		// Carré rouge
		drawRoundedRect(screen, 500, 300, 100, 100, 20, couleur_enemie, "")

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
	if gameFinished {
		// Dessin du bouton menu
		if menuBtn != nil {
			optmenu := &ebiten.DrawImageOptions{}
			optmenu.GeoM.Scale(0.5, 0.5)
			optmenu.GeoM.Translate(float64(menuRect.Min.X), float64(menuRect.Min.Y))
			screen.DrawImage(menuBtn, optmenu)
		}

		// Détection du clic sur le bouton menu
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			if menuRect.Min.X <= x && x <= menuRect.Max.X &&
				menuRect.Min.Y <= y && y <= menuRect.Max.Y {
				SetState("menu") // retour au menu principal
				score_toi = 0
				score_enemie = 0
				currentRound = 0
				gameFinished = false
				combatResult = ""
				toi = false
				enemie = false
				// Réinitialiser les pilules pour le prochain combat
				pilredUsed = false
				pilblueUsed = false
			}
		}
	}
}
