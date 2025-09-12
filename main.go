package main

import (
	"fmt"
	"image/color"
	"log"

	"projetred/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	characters []models.Classe
	images     []*ebiten.Image
}

func loadImages() []*ebiten.Image {
	paths := []string{
		"images/neo_face.png",
		"images/trinity_face.png",
		"images/morpheus_face.png",
	}
	images := make([]*ebiten.Image, len(paths))
	for i, path := range paths {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Fatalf("Impossible de charger %s : %v", path, err)
		}
		images[i] = img
	}
	return images
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255}) // fond gris sombre

	for i, char := range g.characters {
		x := 50 + i*220
		y := 50

		// Dessiner l'image
		op := &ebiten.DrawImageOptions{}
		scaleX := 0.5
		scaleY := 0.5
		op.GeoM.Scale(scaleX, scaleY)
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(g.images[i], op)

		// Dessiner PV
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PV: %d", char.PvBase), x, y+160)

		// Dessiner description
		ebitenutil.DebugPrintAt(screen, char.Description, x, y+180)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 400
}

func main() {
	game := &Game{
		characters: []models.Classe{
			models.InitNeo(),
			models.InitTrinity(),
			models.InitMorpheus(),
		},
		images: loadImages(),
	}

	ebiten.SetWindowSize(800, 400)
	ebiten.SetWindowTitle("Matrix Characters")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
