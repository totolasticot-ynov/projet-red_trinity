package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var inventaire bool
var mouseDown bool // détecte la transition appuyé -> relâché
var arts = []string{"Boxe", "Judo", "Jujutsu", "Karate", "Lutte"}
var combatResult string // texte du dernier combat

func choixAdversaire() string {
	rand.Seed(time.Now().UnixNano())
	return arts[rand.Intn(len(arts))]
}

func resoudreCombat(joueur string, adversaire string) string {
	if joueur == adversaire {
		return fmt.Sprintf("Égalité ! Vous avez tous les deux choisi %s.", joueur)
	}

	// Règle circulaire (chifoumi à 5 coups)
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
		return "Erreur dans le choix."
	}

	// Le joueur bat les deux suivants
	if (adversaireIndex == (joueurIndex+1)%5) || (adversaireIndex == (joueurIndex+2)%5) {
		return fmt.Sprintf("Vous gagnez ! %s bat %s.", joueur, adversaire)
	}

	return fmt.Sprintf("Vous perdez... %s bat %s.", adversaire, joueur)
}

func lancerCombat(choix string) {
	adv := choixAdversaire()
	combatResult = resoudreCombat(choix, adv)
}

// -----------------------------
// Logique de jeu
// -----------------------------

func UpdateGame_dojo() {
	// Ici on ne gère que les boutons de changement d'état (menu/combat)
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
}

func DrawGame_dojo_after(screen *ebiten.Image) {
	playlevel1Music()

	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	// Détection d'un nouveau clic (edge: non-press -> press)
	if pressed && !mouseDown {
		// nouveau clic détecté : traiter l'événement une fois
		x, y := ebiten.CursorPosition()

		// Toggle inventaire si on clique sur le bouton d'inventaire
		if inventaireOffRect.Min.X <= x && x <= inventaireOffRect.Max.X &&
			inventaireOffRect.Min.Y <= y && y <= inventaireOffRect.Max.Y {
			inventaire = !inventaire
		} else if inventaire { // si inventaire ouvert, regarder les items
			// Chaque bouton lance un combat correspondant
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
		// relâchement : autorise le prochain clic
		mouseDown = false
	}

	// fond
	if bgGame_dojo != nil {
		op := &ebiten.DrawImageOptions{}
		scaleX := 800 / float64(bgGame_dojo.Bounds().Dx())
		scaleY := 600 / float64(bgGame_dojo.Bounds().Dy())
		op.GeoM.Scale(scaleX, scaleY)
		screen.DrawImage(bgGame_dojo, op)
	}

	// persos
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

	// inventaire (visuel)
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
		opts.GeoM.Scale(0.5, 0.5)

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

	// affichage du résultat (utilise basicfont.Face7x13)
	if combatResult != "" {
		fmt.Print(combatResult + "\n")
	}
}
