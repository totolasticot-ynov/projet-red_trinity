package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	lastF11 bool
}

func (g *Game) Update() error {
	// Toggle plein écran si on appuie sur F11
	f11 := ebiten.IsKeyPressed(ebiten.KeyF11)
	if f11 && !g.lastF11 {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	g.lastF11 = f11

	switch GetState() {
	case "menu":
		UpdateMenu()
	case "dojo":
		UpdateGame_dojo()
	case "mall":
		UpdateGame_mall()
	case "place":
		UpdateGame_place()
	case "maison":
		UpdateGame_maison()
	case "building":
		UpdateGame_building()
	case "forge":
		UpdateGame_forge()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch GetState() {
	case "menu":
		DrawMenu(screen)
	case "forge":
		DrawGame_forge(screen)
	case "dojo":
		DrawGame_dojo_before(screen)
	case "mall":
		DrawGame_mall_before(screen)
	case "place":
		DrawGame_place_before(screen)
	case "maison":
		DrawGame_maison_before(screen)
	case "building":
		DrawGame_building_before(screen)
	case "boutique_dojo":
		DrawGame_dojo_after(screen)
	case "boutique_mall":
		DrawGame_mall_after(screen)
	case "boutique_place":
		DrawGame_place_after(screen)
	case "boutique_maison":
		DrawGame_maison_after(screen)
	case "boutique_building":
		DrawGame_building_after(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Projet Matrix - Menu")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
