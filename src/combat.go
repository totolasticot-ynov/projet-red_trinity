package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	inventaire       bool
	mouseDown        bool
	arts             = []string{"Boxe", "Judo", "Jujutsu", "Karate", "Lutte"}
	combatResult     string
	currentRound     int  = 1
	gameFinished     bool = false
	playerChoice     string
	adversaireChoice string
	resultImg        *ebiten.Image
	resultimg_red    *ebiten.Image
	vert             = color.RGBA{0, 255, 0, 255}
	rouge            = color.RGBA{255, 0, 0, 255}
	toi              bool
	enemie           bool
	couleur_toi      color.RGBA
	couleur_enemie   color.RGBA
	score_toi        int
	score_enemie     int
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
		couleur_toi = color.RGBA{200, 200, 0, 255} // jaune pour égalité
		couleur_enemie = color.RGBA{200, 200, 0, 255}
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

	if adversaireIndex == (joueurIndex+1)%5 || adversaireIndex == (joueurIndex+2)%5 {
		// joueur gagne
		toi = true
		enemie = false
		couleur_toi = vert
		couleur_enemie = rouge
		score_toi++ // ✅ ajoute un point au joueur
		return fmt.Sprintf("Round %d - Vous gagnez ! %s bat %s.", currentRound, joueur, adversaire)
	}

	// sinon joueur perd
	toi = false
	enemie = true
	couleur_toi = rouge
	couleur_enemie = vert
	score_enemie++ // ✅ ajoute un point à l'adversaire
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
		}
		if lutteBtn != nil {
			opts.GeoM.Translate(float64(lutteRect.Min.X), float64(lutteRect.Min.Y))
			screen.DrawImage(lutteBtn, opts)
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

func DrawScore(screen *ebiten.Image) {
	// Scores en haut à droite
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Score Vous: %d", score_toi), 400, 30)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Score Ennemi: %d", score_enemie), 400, 50)

	// Message final
	if gameFinished {
		if score_toi > score_enemie {
			ebitenutil.DebugPrintAt(screen, "🏆 Vous remportez le combat !", 350, 80)
		} else if score_enemie > score_toi {
			ebitenutil.DebugPrintAt(screen, "❌ L'ennemi remporte le combat...", 350, 80)
		} else {
			ebitenutil.DebugPrintAt(screen, "🤝 Match nul !", 350, 80)
		}
	}
}

// ---- Fonction Draw principale ----
func (g *Game) Draw(screen *ebiten.Image) {
	DrawInventaire(screen)
	DrawRounds(screen)
	DrawScore(screen)
}
