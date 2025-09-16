package main //creation des images

import (
	"image"

	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bgMenu         *ebiten.Image
	bgRect_dojo    image.Rectangle
	playBtn        *ebiten.Image
	playRect       image.Rectangle
	bgGame_dojo    *ebiten.Image
	backBtn        *ebiten.Image
	backRect       image.Rectangle
	exitBtn        *ebiten.Image
	exitRect       image.Rectangle
	neoplayer      *ebiten.Image
	neo_playerRect image.Rectangle
	bgGame_mall    *ebiten.Image
	bgRect_mall    image.Rectangle
	bgGame_place   *ebiten.Image
	bgRect_place   image.Rectangle
	morpheusplayer *ebiten.Image
	morpheusRect   image.Rectangle
	audioCtx       *audio.Context
	menuPlayer     *audio.Player
	fightplay      *ebiten.Image
	fightRect      image.Rectangle
	boxeBtn        *ebiten.Image
	boxeRect       image.Rectangle
	judoBtn        *ebiten.Image
	judoRect       image.Rectangle
	jujutsuBtn     *ebiten.Image
	jujutsuRect    image.Rectangle
	karateBtn      *ebiten.Image
	karateRect     image.Rectangle
	lutteBtn       *ebiten.Image
	lutteRect      image.Rectangle
)

func init() {
	// backgrounds
	bgMenu, _, _ = ebitenutil.NewImageFromFile("../background/images/bg.png")
	bgGame_dojo, _, _ = ebitenutil.NewImageFromFile("../background/images/level1.png")
	bgGame_mall, _, _ = ebitenutil.NewImageFromFile("../background/images/level2.png")
	bgGame_place, _, _ = ebitenutil.NewImageFromFile("../background/images/level3.png")

	// boutons
	playBtn, _, _ = ebitenutil.NewImageFromFile("../background/images/play.png")
	playRect = image.Rect(250, 400, 500+playBtn.Bounds().Dx(), 400+playBtn.Bounds().Dy()) //position du bouton play

	exitBtn, _, _ = ebitenutil.NewImageFromFile("../background/images/exit.png")
	exitRect = image.Rect(500, 400, 500+exitBtn.Bounds().Dx(), 400+exitBtn.Bounds().Dy()) //position du bouton exit

	backBtn, _, _ = ebitenutil.NewImageFromFile("../background/images/back.png")
	backRect = image.Rect(50, 50, 50+backBtn.Bounds().Dx(), 50+backBtn.Bounds().Dy()) //position du bouton back

	fightplay, _, _ = ebitenutil.NewImageFromFile("../background/images/fight.png")
	fightRect = image.Rect(250, 150, 250+fightplay.Bounds().Dx(), 150+fightplay.Bounds().Dy()) //position du bouton fight

	bgRect_dojo = image.Rect(50, 200, 50+bgGame_dojo.Bounds().Dx(), 200+bgGame_dojo.Bounds().Dy()) //position du bouton dojo

	bgRect_mall = image.Rect(50, 200, 50+bgGame_mall.Bounds().Dx(), 200+bgGame_mall.Bounds().Dy()) //position du bouton mall

	bgRect_place = image.Rect(50, 200, 50+bgGame_place.Bounds().Dx(), 200+bgGame_place.Bounds().Dy()) //position du bouton place

	// personnage
	neoplayer, _, _ = ebitenutil.NewImageFromFile("../background/images/neo.png")
	neo_playerRect = image.Rect(50, 300, 50+neoplayer.Bounds().Dx(), 300+neoplayer.Bounds().Dy()) //position de neo

	morpheusplayer, _, _ = ebitenutil.NewImageFromFile("../background/images/morpheus.png")
	morpheusRect = image.Rect(600, 300, 600+morpheusplayer.Bounds().Dx(), 300+morpheusplayer.Bounds().Dy()) //position de morpheus

	// audio
	audioCtx = audio.NewContext(44100)

	// technique de combat
	boxeBtn, _, _ = ebitenutil.NewImageFromFile("../background/images/art martiaux/boxe.png")
	boxeRect = image.Rect(50, 150, 50+boxeBtn.Bounds().Dx(), 150+boxeBtn.Bounds().Dy()) //position du bouton boxe

	judoBtn, _, _ = ebitenutil.NewImageFromFile("../background/images/art martiaux/judo.png")
	judoRect = image.Rect(200, 150, 200+judoBtn.Bounds().Dx(), 150+judoBtn.Bounds().Dy()) //position du bouton judo

	jujutsuBtn, _, _ = ebitenutil.NewImageFromFile("../background/images/art martiaux/jujutsu.png")
	jujutsuRect = image.Rect(350, 150, 350+jujutsuBtn.Bounds().Dx(), 150+jujutsuBtn.Bounds().Dy()) //position du bouton jujutsu

	karateBtn, _, _ = ebitenutil.NewImageFromFile("../background/images/art martiaux/karate.png")
	karateRect = image.Rect(500, 150, 500+karateBtn.Bounds().Dx(), 150+karateBtn.Bounds().Dy()) //position du bouton karate

	lutteBtn, _, _ = ebitenutil.NewImageFromFile("../background/images/art martiaux/lutte.png")
	lutteRect = image.Rect(650, 150, 650+lutteBtn.Bounds().Dx(), 150+lutteBtn.Bounds().Dy()) //position du bouton lutte
}

// === AUDIO ===
func playMenuMusic() {
	if menuPlayer != nil && menuPlayer.IsPlaying() {
		return // La musique est déjà en cours de lecture
	}
	f, _ := os.Open("../musiques/musique du menu.mp3") // ton fichier mp3
	stream, _ := mp3.Decode(audioCtx, f)
	loop := audio.NewInfiniteLoop(stream, stream.Length())
	menuPlayer, _ = audio.NewPlayer(audioCtx, loop)
	menuPlayer.Play()
}
