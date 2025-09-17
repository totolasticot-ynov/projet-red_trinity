package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func drawRoundedRect(screen *ebiten.Image, x, y, w, h, r int, col color.Color, txt string) {

	// coins = cercles plus clairs
	for _, cx := range []int{x + r, x + w - r} {
		for _, cy := range []int{y + r, y + h - r} {
			drawCircle(screen, cx, cy, r, col)
		}
	}

	// rectangle central qui relie les cercles
	ebitenutil.DrawRect(screen, float64(x+r), float64(y+r), float64(w-2*r), float64(h-2*r), col)

	// rectangles pour relier les côtés
	ebitenutil.DrawRect(screen, float64(x+r), float64(y), float64(w-2*r), float64(r), col)     // haut
	ebitenutil.DrawRect(screen, float64(x+r), float64(y+h-r), float64(w-2*r), float64(r), col) // bas
	ebitenutil.DrawRect(screen, float64(x), float64(y+r), float64(r), float64(h-2*r), col)     // gauche
	ebitenutil.DrawRect(screen, float64(x+w-r), float64(y+r), float64(r), float64(h-2*r), col) // droite

	// texte centré approximatif
	text.Draw(screen, txt, basicfont.Face7x13, x+w/3, y+h/4, color.White)
}

func drawCircle(screen *ebiten.Image, cx, cy, r int, col color.Color) {
	img := ebiten.NewImage(2*r, 2*r)
	img.Fill(color.RGBA{0, 0, 0, 0}) // transparent
	for dy := -r; dy <= r; dy++ {
		for dx := -r; dx <= r; dx++ {
			if dx*dx+dy*dy <= r*r {
				img.Set(r+dx, r+dy, col) // col doit avoir A < 255 si tu veux semi-transparent
			}
		}
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cx-r), float64(cy-r))
	screen.DrawImage(img, op) // ici la transparence est respectée
}
