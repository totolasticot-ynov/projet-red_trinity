package main

import (
	"image/color"

	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func UpdateGame_forge() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if backRect.Min.X <= x && x <= backRect.Max.X &&
			backRect.Min.Y <= y && y <= backRect.Max.Y {
			SetState("menu")
		}
		// Achat des techniques de combat
		if !lutte && 500 <= x && x <= 550 && 300 <= y && y <= 350 && argent >= 50 {
			if casque && tenuelutte {
				argent -= 50
				lutte = true
			}

		}

		if !karate && 500 <= x && x <= 550 && 460 <= y && y <= 510 && argent >= 50 {
			if kimono && pant {
				argent -= 50
				karate = true
			}
		}
	}
}

func DrawGame_forge(screen *ebiten.Image) {
	playMenuMusic()
	op := &ebiten.DrawImageOptions{}
	scaleX := 800 / float64(bgGame_forge.Bounds().Dx())
	scaleY := 600 / float64(bgGame_forge.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(bgGame_forge, op)

	if backBtn != nil {
		optback := &ebiten.DrawImageOptions{}
		optback.GeoM.Translate(float64(backRect.Min.X), float64(backRect.Min.Y))
		screen.DrawImage(backBtn, optback)
	}

	if cypherplayer != nil {
		optmor := &ebiten.DrawImageOptions{}
		optmor.GeoM.Scale(0.5, 0.5)
		optmor.GeoM.Translate(float64(cypherRect.Min.X), float64(cypherRect.Min.Y))
		screen.DrawImage(cypherplayer, optmor)
	}

	drawRoundedRect(screen, 155, 100, 350, 150, 20, color.RGBA{0, 0, 0, 255}, "")

	text.Draw(screen, "Bienvenue dans la forge !", basicfont.Face7x13, 250, 150, color.White)
	text.Draw(screen, "Ici, vous pouvez acheter de nouvelles ", basicfont.Face7x13, 180, 180, color.White)
	text.Draw(screen, "techniques de combat.", basicfont.Face7x13, 180, 200, color.White)
	text.Draw(screen, "Cependant, cela coute de l'argent.", basicfont.Face7x13, 180, 220, color.White)

	text.Draw(screen, "Argent: "+strconv.Itoa(argent)+" $", basicfont.Face7x13, 650, 50, color.White)

	//lutte
	if casqueAC != nil {
		optcasque := &ebiten.DrawImageOptions{}
		optcasque.GeoM.Scale(0.3, 0.3)
		optcasque.GeoM.Translate(50, 250)
		if !casque {
			optcasque.ColorM.Scale(0.5, 0.5, 0.5, 1.0)
		}
		screen.DrawImage(casqueAC, optcasque)
	}
	drawRoundedRect(screen, 200, 300, 20, 60, 10, color.RGBA{255, 255, 255, 255}, "") //vertical
	drawRoundedRect(screen, 180, 320, 60, 20, 10, color.RGBA{255, 255, 255, 255}, "") //horizontal

	if lutteAC != nil {
		optlutte := &ebiten.DrawImageOptions{}
		optlutte.GeoM.Scale(0.3, 0.3)
		optlutte.GeoM.Translate(250, 250)
		if !tenuelutte {
			optlutte.ColorM.Scale(0.5, 0.5, 0.5, 1.0)
		}
		screen.DrawImage(lutteAC, optlutte)
	}

	drawRoundedRect(screen, 375, 300, 60, 20, 10, color.RGBA{255, 255, 255, 255}, "") //haut
	drawRoundedRect(screen, 375, 340, 60, 20, 10, color.RGBA{255, 255, 255, 255}, "") //bas
	drawRoundedRect(screen, 450, 250, 150, 150, 20, color.RGBA{255, 255, 255, 255}, "")

	if lutteBtn != nil {
		optlutte := &ebiten.DrawImageOptions{}
		optlutte.GeoM.Scale(0.3, 0.3)
		optlutte.GeoM.Translate(460, 250)
		screen.DrawImage(lutteBtn, optlutte)
	}
	//karate
	drawRoundedRect(screen, 190, 450, 20, 60, 10, color.RGBA{255, 255, 255, 255}, "") //vertical
	drawRoundedRect(screen, 170, 470, 60, 20, 10, color.RGBA{255, 255, 255, 255}, "") //horizontal

	if kimonoAC != nil {
		optkarate := &ebiten.DrawImageOptions{}
		optkarate.GeoM.Scale(0.3, 0.3)
		optkarate.GeoM.Translate(230, 430)
		if !kimono {
			optkarate.ColorM.Scale(0.5, 0.5, 0.5, 1.0)
		}
		screen.DrawImage(kimonoAC, optkarate)
	}

	if pantAC != nil {
		optpant := &ebiten.DrawImageOptions{}
		optpant.GeoM.Scale(0.3, 0.3)
		optpant.GeoM.Translate(50, 430)
		if !pant {
			optpant.ColorM.Scale(0.5, 0.5, 0.5, 1.0)
		}
		screen.DrawImage(pantAC, optpant)
	}

	drawRoundedRect(screen, 375, 470, 60, 20, 10, color.RGBA{255, 255, 255, 255}, "") //haut
	drawRoundedRect(screen, 375, 510, 60, 20, 10, color.RGBA{255, 255, 255, 255}, "") //bas
	drawRoundedRect(screen, 450, 420, 150, 150, 20, color.RGBA{255, 255, 255, 255}, "")

	if karateBtn != nil {
		optkarate := &ebiten.DrawImageOptions{}
		optkarate.GeoM.Scale(0.3, 0.3)
		optkarate.GeoM.Translate(460, 420)
		screen.DrawImage(karateBtn, optkarate)
	}

	// affichage mini carre noir devant les art artiaux
	if !lutte {
		drawRoundedRect(screen, 500, 300, 60, 60, 10, color.RGBA{0, 0, 0, 255}, "50 $")
	}
	if !karate {
		drawRoundedRect(screen, 500, 460, 60, 60, 10, color.RGBA{0, 0, 0, 255}, "50 $")
	}
}
