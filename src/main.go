package main

import (
	"log"

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
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch GetState() {
	case "menu":
		DrawMenu(screen)
	case "dojo":
		DrawGame_dojo(screen)
	case "mall":
		DrawGame_mall(screen)
	case "place":
		DrawGame_place(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Projet Matrix - Menu")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
