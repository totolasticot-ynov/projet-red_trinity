package main

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

var selectedArena string = ""

func UpdateMenu() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// bouton play
		if playRect.Min.X <= x && x <= playRect.Max.X &&
			playRect.Min.Y <= y && y <= playRect.Max.Y && selectedArena != "" {
			SetState(selectedArena)
		}

		// bouton exit
		if exitRect.Min.X <= x && x <= exitRect.Max.X &&
			exitRect.Min.Y <= y && y <= exitRect.Max.Y {
			os.Exit(0)
		}

		// choisir une arène
		if 50 <= x && x <= 280 && 200 <= y && y <= 350 {
			selectedArena = "dojo"
		}
		if 300 <= x && x <= 530 && 200 <= y && y <= 350 {
			selectedArena = "mall"
		}
		if 550 <= x && x <= 775 && 200 <= y && y <= 350 {
			selectedArena = "place"
		}
	}
}

func DrawMenu(screen *ebiten.Image) {
	// background menu
	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgMenu.Bounds().Dx())
	scaleY := 600 / float64(bgMenu.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgMenu, op)

	// titre
	text.Draw(screen, "MATRIX", basicfont.Face7x13, 400, 100, color.RGBA{108, 196, 12, 255})

	// bouton play
	opts1 := &ebiten.DrawImageOptions{}
	opts1.GeoM.Translate(float64(playRect.Min.X), float64(playRect.Min.Y))
	screen.DrawImage(playBtn, opts1)

	// bouton exit
	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Translate(float64(exitRect.Min.X), float64(exitRect.Min.Y))
	screen.DrawImage(exitBtn, opts2)

	opts_icon := &ebiten.DrawImageOptions{}
	opts_icon.GeoM.Scale(0.15, 0.15)

	// dojo
	opts_icon.GeoM.Translate(float64(bgRect_dojo.Min.X), float64(bgRect_dojo.Min.Y))
	screen.DrawImage(bgGame_dojo, opts_icon)

	// mall
	opts_icon.GeoM.Translate(float64(bgRect_mall.Min.X+200), float64(bgRect_mall.Min.Y-200))
	screen.DrawImage(bgGame_mall, opts_icon)

	// place
	opts_icon.GeoM.Translate(float64(bgRect_place.Min.X+200), float64(bgRect_place.Min.Y-200))
	screen.DrawImage(bgGame_place, opts_icon)

	// affiche l'arène choisie
	if selectedArena != "" {
		text.Draw(screen, "Arena: "+selectedArena, basicfont.Face7x13, 300, 550, color.White)
	}
}
