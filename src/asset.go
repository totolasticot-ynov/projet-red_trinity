package main //creation des images

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bgMenu   *ebiten.Image
	playBtn  *ebiten.Image
	playRect image.Rectangle
	bgGame   *ebiten.Image
	backBtn  *ebiten.Image
	backRect image.Rectangle
	exitBtn  *ebiten.Image
	exitRect image.Rectangle
)

func init() {
	// backgrounds
	bgMenu, _, _ = ebitenutil.NewImageFromFile("../images/bg.png")
	bgGame, _, _ = ebitenutil.NewImageFromFile("../images/dojo.png")

	// boutons
	playBtn, _, _ = ebitenutil.NewImageFromFile("../images/play.png")
	playRect = image.Rect(500, 400, 500+playBtn.Bounds().Dx(), 400+playBtn.Bounds().Dy()) //position du bouton play

	exitBtn, _, _ = ebitenutil.NewImageFromFile("../images/exit.png")
	exitRect = image.Rect(500, 450, 500+exitBtn.Bounds().Dx(), 450+exitBtn.Bounds().Dy()) //position du bouton exit

	backBtn, _, _ = ebitenutil.NewImageFromFile("../images/back.png")
	backRect = image.Rect(50, 50, 50+backBtn.Bounds().Dx(), 50+backBtn.Bounds().Dy()) //position du bouton back
}
