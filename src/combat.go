package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	inventaire         bool
	mouseDown          bool
	arts               = []string{"Boxe", "Judo", "Jujutsu", "Karate", "Lutte"}
	combatResult       string
	currentRound       int  = 1
	gameFinished       bool = false
	playerChoice       string
	adversaireChoice   string
	resultImg          *ebiten.Image
	resultimg_red      *ebiten.Image
	vert               = color.RGBA{0, 255, 0, 255}
	rouge              = color.RGBA{255, 0, 0, 255}
	toi                bool
	enemie             bool
	couleur_toi        color.RGBA
	couleur_enemie     color.RGBA
	score_toi          int
	local_score_toi    *ebiten.Image
	local_score_enemie *ebiten.Image
	score_enemie       int
	mall               bool = false
	place              bool = false
	maison             bool = false
	argent             int
	karate             bool = false
	lutte              bool = false
	pilred             bool
	pilblue            bool
)

func choixAdversaire() string {
	rand.Seed(time.Now().UnixNano())
	return arts[rand.Intn(len(arts))]
}

func lancerCombat(choix string) {
	if gameFinished {
		return
	}

	// reset couleurs à chaque round
	toi = false
	enemie = false
	couleur_toi = color.RGBA{0, 0, 0, 0}
	couleur_enemie = color.RGBA{0, 0, 0, 0}

	if choix == "miss" {
		combatResult = fmt.Sprintf("Round %d - Vous avez manqué votre attaque !", currentRound)
	}

	playerChoice = choix
	adversaireChoice = choixAdversaire()

	switch adversaireChoice {
	case "Boxe":
		resultimg_red = boxeBtn
	case "Judo":
		resultimg_red = judoBtn
	case "Jujutsu":
		resultimg_red = jujutsuBtn
	case "Karate":
		resultimg_red = karateBtn
	case "Lutte":
		resultimg_red = lutteBtn
	}

	combatResult = resoudreCombat(playerChoice, adversaireChoice)

	currentRound++
	if currentRound > 5 {
		gameFinished = true
		combatResult += "\nCombat terminé ! 5 rounds effectués."
	}
}

func resoudreCombat(joueur string, adversaire string) string {
	if joueur == adversaire {
		toi = false
		enemie = false
		return fmt.Sprintf("Round %d - Égalité ! Vous avez tous les deux choisi %s.", currentRound, joueur)
	}

	joueurIndex := -1
	adversaireIndex := -1

	for i, v := range arts {
		if v == joueur {
			joueurIndex = i
		}
		if v == adversaire {
			adversaireIndex = i
		}
	}

	if joueurIndex == -1 || adversaireIndex == -1 {
		return fmt.Sprintf("Round %d - Erreur dans le choix.", currentRound)
	}

	if (adversaireIndex == (joueurIndex+1)%5) || (adversaireIndex == (joueurIndex+2)%5) {
		toi = true
		score_toi++
		argent += 10
		return fmt.Sprintf("Round %d - Vous gagnez ! %s bat %s.", currentRound, joueur, adversaire)
	}
	toi = false
	enemie = false
	score_enemie++
	return fmt.Sprintf("Round %d - Vous perdez... %s bat %s.", currentRound, adversaire, joueur)
}

func DrawInventaire(screen *ebiten.Image) {
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

		drawRoundedRect(screen, 150, 100, 100, 450, 20, color.RGBA{255, 255, 255, 255}, "")

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(0.15, 0.15)

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
			if !karate {
				// dessine un cadenas si le niveau n'est pas débloqué
				cadenasOpts := &ebiten.DrawImageOptions{}
				cadenasOpts.GeoM.Scale(0.2, 0.2)
				cadenasOpts.GeoM.Translate(float64(bgRect_mall.Min.X+90), float64(bgRect_mall.Min.Y+175))
				screen.DrawImage(cadena, cadenasOpts)
			}
		}
		if lutteBtn != nil {
			opts.GeoM.Translate(float64(lutteRect.Min.X), float64(lutteRect.Min.Y))
			screen.DrawImage(lutteBtn, opts)
			if !lutte {
				// dessine un cadenas si le niveau n'est pas débloqué
				cadenasOpts := &ebiten.DrawImageOptions{}
				cadenasOpts.GeoM.Scale(0.2, 0.2)
				cadenasOpts.GeoM.Translate(float64(bgRect_mall.Min.X+90), float64(bgRect_mall.Min.Y+250))
				screen.DrawImage(cadena, cadenasOpts)
			}
		}
	}
}

func DrawRounds(screen *ebiten.Image) {
	var roundImg *ebiten.Image

	switch currentRound {
	case 1:
		roundImg = round1
	case 2:
		roundImg = round2
	case 3:
		roundImg = round3
	case 4:
		roundImg = round4
	case 5:
		roundImg = round5
	}

	if roundImg != nil {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(100, 30)
		screen.DrawImage(roundImg, opt)
	}
}
