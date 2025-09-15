package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func UpdateMenu() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition() //capture position souris

		// si clic gauche dans playRect: go game
		if playRect.Min.X <= x && x <= playRect.Max.X &&
			playRect.Min.Y <= y && y <= playRect.Max.Y {
			SetState("game")
		}
		// si clic gauche dans exitRect: go exit
		if exitRect.Min.X <= x && x <= exitRect.Max.X &&
			exitRect.Min.Y <= y && y <= exitRect.Max.Y {
			SetState("exit")
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

	// créer une image temporaire pour le texte
	textImg := ebiten.NewImage(200, 50) // taille approximative du texte
	text.Draw(textImg, "MATRIX", basicfont.Face7x13, 0, 13, color.RGBA{108, 196, 12, 255})

	// scaler le texte
	optsText := &ebiten.DrawImageOptions{}
	optsText.GeoM.Scale(5, 5)         // agrandit 4×
	optsText.GeoM.Translate(275, 100) // position du titre
	screen.DrawImage(textImg, optsText)

	// bouton play
	opts1 := &ebiten.DrawImageOptions{}
	opts1.GeoM.Translate(float64(playRect.Min.X), float64(playRect.Min.Y))
	screen.DrawImage(playBtn, opts1)

	// bouton retour
	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Translate(float64(backRect.Min.X), float64(backRect.Min.Y))
	screen.DrawImage(backBtn, opts2)

	// bouton exit
	opts3 := &ebiten.DrawImageOptions{}
	opts3.GeoM.Translate(float64(exitRect.Min.X), float64(exitRect.Min.Y))
	screen.DrawImage(exitBtn, opts3)

	// bouton dojo
	opts4 := &ebiten.DrawImageOptions{}
	opts4.GeoM.Translate(float64(bgRect.Min.X), float64(bgRect.Min.Y))
	screen.DrawImage(bgGame, opts4)
}
