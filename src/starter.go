package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
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
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch GetState() {
	case "menu":
		DrawMenu(screen)
	case "dojo":
		DrawGame_dojo_before(screen)
	case "mall":
		DrawGame_mall_before(screen)
	case "place":
		DrawGame_place_before(screen)
	case "boutique_dojo":
		DrawGame_dojo_after(screen)
	case "boutique_mall":
		DrawGame_mall_after(screen)
	case "boutique_place":
		DrawGame_place_after(screen)
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
