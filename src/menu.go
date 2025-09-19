package main

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

var (
	selectedArena string = ""
	notice        bool   = false
	menuMouseDown bool   = false // edge detector
)

func UpdateMenu() {
	playMenuMusic()

	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if pressed && !menuMouseDown {
		x, y := ebiten.CursorPosition()

		if playRect.Min.X <= x && x <= playRect.Max.X &&
			playRect.Min.Y <= y && y <= playRect.Max.Y && selectedArena != "" {

			if menuPlayer != nil && menuPlayer.IsPlaying() {
				menuPlayer.Pause()
			}
			SetState(selectedArena)
		}

		if 0 <= x && x <= 10 && 550 <= y && y <= 600 {
			mall = true
			place = true
			maison = true
			building = true
			argent += 1000
			tenuelutte = true
			pant = true
			casque = true
			kimono = true
		}

		if exitRect.Min.X <= x && x <= exitRect.Max.X &&
			exitRect.Min.Y <= y && y <= exitRect.Max.Y {
			os.Exit(0)
		}

		if forgeRect.Min.X <= x && x <= forgeRect.Max.X &&
			forgeRect.Min.Y <= y && y <= forgeRect.Max.Y {
			SetState("forge")
		}

		if noticeRect.Min.X <= x && x <= noticeRect.Max.X &&
			noticeRect.Min.Y <= y && y <= noticeRect.Max.Y {
			notice = !notice // toggle propre
		}

		if 50 <= x && x <= 280 && 100 <= y && y <= 250 {
			selectedArena = "dojo"
		}
		if 300 <= x && x <= 530 && 100 <= y && y <= 250 && mall {
			selectedArena = "mall"
		}
		if 550 <= x && x <= 775 && 100 <= y && y <= 250 && place {
			selectedArena = "place"
		}
		if 200 <= x && x <= 430 && 250 <= y && y <= 400 && place {
			selectedArena = "maison"
		}
		if 400 <= x && x <= 630 && 250 <= y && y <= 400 && building {
			selectedArena = "building"
		}

		menuMouseDown = true
	} else if !pressed {
		menuMouseDown = false
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
	txtImg := ebiten.NewImage(200, 50) // taille de l'image contenant le texte
	text.Draw(txtImg, "MATRIX", basicfont.Face7x13, 85, 20, color.RGBA{108, 196, 12, 255})

	optitre := &ebiten.DrawImageOptions{}
	optitre.GeoM.Scale(4, 4) // double la taille
	screen.DrawImage(txtImg, optitre)

	// bouton play
	opts1 := &ebiten.DrawImageOptions{}
	opts1.GeoM.Translate(float64(playRect.Min.X), float64(playRect.Min.Y))
	screen.DrawImage(playBtn, opts1)

	// bouton exit
	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Translate(float64(exitRect.Min.X), float64(exitRect.Min.Y))
	screen.DrawImage(exitBtn, opts2)

	// bouton forge
	opts3 := &ebiten.DrawImageOptions{}
	opts3.GeoM.Scale(0.35, 0.35)
	opts3.GeoM.Translate(float64(forgeRect.Min.X), float64(forgeRect.Min.Y))
	screen.DrawImage(forgeBtn, opts3)

	// bouton notice
	drawRoundedRect(screen, noticeRect.Min.X-5, noticeRect.Min.Y-5, 80, 80, 20, color.RGBA{0, 255, 0, 255}, "")
	opts4 := &ebiten.DrawImageOptions{}
	opts4.GeoM.Scale(0.3, 0.3)
	opts4.GeoM.Translate(float64(noticeRect.Min.X), float64(noticeRect.Min.Y))
	screen.DrawImage(noticeBtn, opts4)

	// icônes des arènes
	opts_icon := &ebiten.DrawImageOptions{}
	opts_icon.GeoM.Scale(0.15, 0.15)

	// dojo
	opts_icon.GeoM.Translate(float64(bgRect_dojo.Min.X), float64(bgRect_dojo.Min.Y))
	screen.DrawImage(bgGame_dojo, opts_icon)

	// mall
	opts_icon.GeoM.Translate(float64(bgRect_mall.Min.X+200), float64(bgRect_mall.Min.Y-200))
	screen.DrawImage(bgGame_mall, opts_icon)
	if !mall {
		cadenasOpts := &ebiten.DrawImageOptions{}
		cadenasOpts.GeoM.Scale(0.5, 0.5)
		cadenasOpts.GeoM.Translate(float64(bgRect_mall.Min.X+212), float64(bgRect_mall.Min.Y-130))
		screen.DrawImage(cadena, cadenasOpts)
	}

	// place
	opts_icon.GeoM.Translate(float64(bgRect_place.Min.X+200), float64(bgRect_place.Min.Y-200))
	screen.DrawImage(bgGame_place, opts_icon)
	if !place {
		cadenasOpts := &ebiten.DrawImageOptions{}
		cadenasOpts.GeoM.Scale(0.5, 0.5)
		cadenasOpts.GeoM.Translate(float64(bgRect_place.Min.X+450), float64(bgRect_place.Min.Y-130))
		screen.DrawImage(cadena, cadenasOpts)
	}

	// maison
	opts_icon.GeoM.Translate(float64(bgRect_maison.Min.X-550), float64(bgRect_maison.Min.Y))
	screen.DrawImage(bgGame_maison, opts_icon)
	if !maison {
		cadenasOpts := &ebiten.DrawImageOptions{}
		cadenasOpts.GeoM.Scale(0.5, 0.5)
		cadenasOpts.GeoM.Translate(float64(bgRect_maison.Min.X-30), float64(bgRect_maison.Min.Y+70))
		screen.DrawImage(cadena, cadenasOpts)
	}

	// building
	opts_icon.GeoM.Translate(float64(bgRect_building.Min.X-550), float64(bgRect_building.Min.Y))
	screen.DrawImage(bgGame_building, opts_icon)
	if !maison {
		cadenasOpts := &ebiten.DrawImageOptions{}
		cadenasOpts.GeoM.Scale(0.5, 0.5)
		cadenasOpts.GeoM.Translate(390, 250)
		screen.DrawImage(cadena, cadenasOpts)
	}

	// affiche l'arène choisie
	if selectedArena != "" {
		// Crée une image temporaire pour le texte
		txtImg := ebiten.NewImage(300, 50)
		text.Draw(txtImg, "Arena: "+selectedArena, basicfont.Face7x13, 0, 20, color.RGBA{108, 196, 12, 255})

		// Options de dessin avec scale
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Scale(2, 2)         // double la taille
		opt.GeoM.Translate(340, 530) // position finale
		screen.DrawImage(txtImg, opt)
	}

	// notice activée
	if notice {
		drawRoundedRect(screen, 100, 200, 600, 200, 20, color.RGBA{0, 0, 0, 255}, "")
		text.Draw(screen, "Regles de victoire :", basicfont.Face7x13, 120, 130, color.Black)
		text.Draw(screen, "Boxe     => Victoires: Judo, Jujutsu | Defaites: Karate, Lutte", basicfont.Face7x13, 120, 260, color.RGBA{108, 196, 12, 255})
		text.Draw(screen, "Judo     => Victoires: Jujutsu, Karate | Defaites: Lutte, Boxe", basicfont.Face7x13, 120, 280, color.RGBA{108, 196, 12, 255})
		text.Draw(screen, "Jujutsu  => Victoires: Karate, Lutte | Defaites: Boxe, Judo", basicfont.Face7x13, 120, 300, color.RGBA{108, 196, 12, 255})
		text.Draw(screen, "Karate   => Victoires: Lutte, Boxe | Defaites: Judo, Jujutsu", basicfont.Face7x13, 120, 320, color.RGBA{108, 196, 12, 255})
		text.Draw(screen, "Lutte    => Victoires: Boxe, Judo | Defaites: Jujutsu, Karate", basicfont.Face7x13, 120, 340, color.RGBA{108, 196, 12, 255})
	}
}
