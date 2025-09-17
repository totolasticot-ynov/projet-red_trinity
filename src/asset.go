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
	bgMenu            *ebiten.Image
	bgRect_dojo       image.Rectangle
	playBtn           *ebiten.Image
	playRect          image.Rectangle
	bgGame_dojo       *ebiten.Image
	backBtn           *ebiten.Image
	backRect          image.Rectangle
	exitBtn           *ebiten.Image
	exitRect          image.Rectangle
	neoplayer         *ebiten.Image
	neo_playerRect    image.Rectangle
	bgGame_mall       *ebiten.Image
	bgRect_mall       image.Rectangle
	bgGame_place      *ebiten.Image
	bgRect_place      image.Rectangle
	morpheusplayer    *ebiten.Image
	morpheusRect      image.Rectangle
	audioCtx          *audio.Context
	menuPlayer        *audio.Player
	fightplay         *ebiten.Image
	fightRect         image.Rectangle
	boxeBtn           *ebiten.Image
	boxeRect          image.Rectangle
	judoBtn           *ebiten.Image
	judoRect          image.Rectangle
	jujutsuBtn        *ebiten.Image
	jujutsuRect       image.Rectangle
	karateBtn         *ebiten.Image
	karateRect        image.Rectangle
	lutteBtn          *ebiten.Image
	lutteRect         image.Rectangle
	inventaireOffBtn  *ebiten.Image
	inventaireOffRect image.Rectangle
	inventaireOnBtn   *ebiten.Image
	inventaireOnRect  image.Rectangle
	oracleplayer      *ebiten.Image
	oracleRect        image.Rectangle
	level1Player      *audio.Player
	dollarBtn         *ebiten.Image
	dollarRect        image.Rectangle
	pilredBtn         *ebiten.Image
	pilredRect        image.Rectangle
	pilblueBtn        *ebiten.Image
	pilblueRect       image.Rectangle
)

func init() {
	// backgrounds
	bgMenu, _, _ = ebitenutil.NewImageFromFile("../images/background/bg.png")
	bgGame_dojo, _, _ = ebitenutil.NewImageFromFile("../images/background/level1.png")
	bgGame_mall, _, _ = ebitenutil.NewImageFromFile("../images/background/level2.png")
	bgGame_place, _, _ = ebitenutil.NewImageFromFile("../images/background/level3.png")

	// boutons
	playBtn, _, _ = ebitenutil.NewImageFromFile("../images/asset/play.png")
	playRect = image.Rect(250, 400, 500+playBtn.Bounds().Dx(), 400+playBtn.Bounds().Dy()) //position du bouton play

	exitBtn, _, _ = ebitenutil.NewImageFromFile("../images/asset/exit.png")
	exitRect = image.Rect(500, 400, 500+exitBtn.Bounds().Dx(), 400+exitBtn.Bounds().Dy()) //position du bouton exit

	backBtn, _, _ = ebitenutil.NewImageFromFile("../images/asset/back.png")
	backRect = image.Rect(50, 50, 50+backBtn.Bounds().Dx(), 50+backBtn.Bounds().Dy()) //position du bouton back

	fightplay, _, _ = ebitenutil.NewImageFromFile("../images/asset/fight.png")
	fightRect = image.Rect(250, 150, 250+fightplay.Bounds().Dx(), 150+fightplay.Bounds().Dy()) //position du bouton fight

	inventaireOffBtn, _, _ = ebitenutil.NewImageFromFile("../images/asset/inventaire_off.png")
	inventaireOffRect = image.Rect(100, 300, 100+inventaireOffBtn.Bounds().Dx(), 300+inventaireOffBtn.Bounds().Dy())

	inventaireOnBtn, _, _ = ebitenutil.NewImageFromFile("../images/asset/inventaire_on.png")
	inventaireOnRect = image.Rect(100, 300, 100+inventaireOnBtn.Bounds().Dx(), 300+inventaireOnBtn.Bounds().Dy())

	dollarBtn, _, _ = ebitenutil.NewImageFromFile("../images/asset/dollar.png")
	dollarRect = image.Rect(650, 300, 650+dollarBtn.Bounds().Dx(), 300+dollarBtn.Bounds().Dy())

	pilredBtn, _, _ = ebitenutil.NewImageFromFile("../images/potion (pillule)/pilule_rouge.png")
	pilredRect = image.Rect(425, 350, 425+pilredBtn.Bounds().Dx(), 350+pilredBtn.Bounds().Dy())

	pilblueBtn, _, _ = ebitenutil.NewImageFromFile("../images/potion (pillule)/pilule_bleue.png")
	pilblueRect = image.Rect(575, 350, 575+pilblueBtn.Bounds().Dx(), 550+pilblueBtn.Bounds().Dy())

	bgRect_dojo = image.Rect(50, 200, 50+bgGame_dojo.Bounds().Dx(), 200+bgGame_dojo.Bounds().Dy()) //position du bouton dojo

	bgRect_mall = image.Rect(50, 200, 50+bgGame_mall.Bounds().Dx(), 200+bgGame_mall.Bounds().Dy()) //position du bouton mall

	bgRect_place = image.Rect(50, 200, 50+bgGame_place.Bounds().Dx(), 200+bgGame_place.Bounds().Dy()) //position du bouton place

	// personnage
	neoplayer, _, _ = ebitenutil.NewImageFromFile("../images/personnages/neo.png")
	neo_playerRect = image.Rect(50, 300, 50+neoplayer.Bounds().Dx(), 300+neoplayer.Bounds().Dy()) //position de neo

	morpheusplayer, _, _ = ebitenutil.NewImageFromFile("../images/personnages/morpheus.png")
	morpheusRect = image.Rect(600, 300, 600+morpheusplayer.Bounds().Dx(), 300+morpheusplayer.Bounds().Dy()) //position de morpheus

	oracleplayer, _, _ = ebitenutil.NewImageFromFile("../images/personnages/oracle.png")
	oracleRect = image.Rect(200, 300, 200+oracleplayer.Bounds().Dx(), 300+oracleplayer.Bounds().Dy()) //position de l'oracle

	// audio
	audioCtx = audio.NewContext(44100)

	// technique de combat
	boxeBtn, _, _ = ebitenutil.NewImageFromFile("../images/art martiaux/boxe.png")
	boxeRect = image.Rect(175, 125, 175+boxeBtn.Bounds().Dx(), 125+boxeBtn.Bounds().Dy()) //position du bouton boxe

	judoBtn, _, _ = ebitenutil.NewImageFromFile("../images/art martiaux/judo.png")
	judoRect = image.Rect(0, 80, 0+judoBtn.Bounds().Dx(), 80+judoBtn.Bounds().Dy()) //position du bouton judo*/

	jujutsuBtn, _, _ = ebitenutil.NewImageFromFile("../images/art martiaux/jujutsu.png")
	jujutsuRect = image.Rect(0, 75, 0+jujutsuBtn.Bounds().Dx(), 75+jujutsuBtn.Bounds().Dy()) //position du bouton jujutsu

	karateBtn, _, _ = ebitenutil.NewImageFromFile("../images/art martiaux/karate.png")
	karateRect = image.Rect(0, 95, 0+karateBtn.Bounds().Dx(), 95+karateBtn.Bounds().Dy()) //position du bouton karate

	lutteBtn, _, _ = ebitenutil.NewImageFromFile("../images/art martiaux/lutte.png")
	lutteRect = image.Rect(0, 80, 0+lutteBtn.Bounds().Dx(), 80+lutteBtn.Bounds().Dy()) //position du bouton lutte
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

func playlevel1Music() {
	if level1Player != nil && level1Player.IsPlaying() {
		return // La musique est déjà en cours de lecture
	}
	f, _ := os.Open("../musiques/musiquelevel1.mp3") // ton fichier mp3
	stream, _ := mp3.Decode(audioCtx, f)
	loop := audio.NewInfiniteLoop(stream, stream.Length())
	level1Player, _ = audio.NewPlayer(audioCtx, loop)
	level1Player.Play()
}

func playlevel2Music() {
	if level1Player != nil && level1Player.IsPlaying() {
		return // La musique est déjà en cours de lecture
	}
	f, _ := os.Open("../musiques/musiquelevel2.mp3") // ton fichier mp3
	stream, _ := mp3.Decode(audioCtx, f)
	loop := audio.NewInfiniteLoop(stream, stream.Length())
	level1Player, _ = audio.NewPlayer(audioCtx, loop)
	level1Player.Play()
}
