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
	// Backgrounds
	bgMenu, bgGame_dojo, bgGame_mall, bgGame_place, bgGame_maison, bgGame_building, bgGame_forge *ebiten.Image
	round1, round2, round3, round4, round5                                                       *ebiten.Image
	bgRect_dojo, bgRect_mall, bgRect_place, bgRect_maison, bgRect_building                       image.Rectangle

	// Personnages
	neoplayer, morpheusplayer, oracleplayer, agentplayer, trinityplayer, cypherplayer, meufplayer, archplayer *ebiten.Image
	neo_playerRect, morpheusRect, oracleRect, agentRect, trinityRect, cypherRect, meufRect, archRect          image.Rectangle

	// Boutons principaux
	playBtn, exitBtn, backBtn, fightplay, menuBtn, noticeBtn      *ebiten.Image
	playRect, exitRect, backRect, fightRect, menuRect, noticeRect image.Rectangle

	// Boutons inventaire / items
	inventaireOffBtn, inventaireOnBtn, pilredBtn, pilblueBtn, cadena, forgeBtn *ebiten.Image
	inventaireOffRect, inventaireOnRect, pilredRect, pilblueRect, forgeRect    image.Rectangle

	// Techniques de combat
	boxeBtn, judoBtn, jujutsuBtn, karateBtn, lutteBtn, casqueAC, lutteAC, kimonoAC, pantAC *ebiten.Image
	boxeRect, judoRect, jujutsuRect, karateRect, lutteRect                                 image.Rectangle

	// Audio
	audioCtx                                                                                     *audio.Context
	menuPlayer, level1Player, level2Player, level3Player, level4Player, level5Player, ringplayer *audio.Player
)

func init() {
	// Backgrounds
	bgMenu, _, _ = ebitenutil.NewImageFromFile("../asset/images/background/bg.png")
	bgGame_dojo, _, _ = ebitenutil.NewImageFromFile("../asset/images/background/level1.png")
	bgGame_mall, _, _ = ebitenutil.NewImageFromFile("../asset/images/background/level2.png")
	bgGame_place, _, _ = ebitenutil.NewImageFromFile("../asset/images/background/level3.png")
	bgGame_maison, _, _ = ebitenutil.NewImageFromFile("../asset/images/background/level4.png")
	bgGame_building, _, _ = ebitenutil.NewImageFromFile("../asset/images/background/level5.png")
	bgGame_forge, _, _ = ebitenutil.NewImageFromFile("../asset/images/background/forgeron.png")
	round1, _, _ = ebitenutil.NewImageFromFile("../asset/images/round/ROUND1.png")
	round2, _, _ = ebitenutil.NewImageFromFile("../asset/images/round/ROUND2.png")
	round3, _, _ = ebitenutil.NewImageFromFile("../asset/images/round/ROUND3.png")
	round4, _, _ = ebitenutil.NewImageFromFile("../asset/images/round/ROUND4.png")
	round5, _, _ = ebitenutil.NewImageFromFile("../asset/images/round/ROUND5.png")
	bgRect_dojo = image.Rect(50, 100, 50+bgGame_dojo.Bounds().Dx(), 100+bgGame_dojo.Bounds().Dy())    //position du bouton dojo
	bgRect_mall = image.Rect(50, 200, 50+bgGame_mall.Bounds().Dx(), 200+bgGame_mall.Bounds().Dy())    //position du bouton mall
	bgRect_place = image.Rect(50, 200, 50+bgGame_place.Bounds().Dx(), 200+bgGame_place.Bounds().Dy()) //position du bouton place
	bgRect_maison = image.Rect(170, 175, 170+bgGame_place.Bounds().Dx(), 175+bgGame_place.Bounds().Dy())
	bgRect_building = image.Rect(800, 0, 800+bgGame_place.Bounds().Dx(), 0+bgGame_place.Bounds().Dy())

	// Personnages
	neoplayer, _, _ = ebitenutil.NewImageFromFile("../asset/images/personnages/neo.png")
	neo_playerRect = image.Rect(50, 300, 50+neoplayer.Bounds().Dx(), 300+neoplayer.Bounds().Dy())

	morpheusplayer, _, _ = ebitenutil.NewImageFromFile("../asset/images/personnages/morpheus.png")
	morpheusRect = image.Rect(600, 300, 600+morpheusplayer.Bounds().Dx(), 300+morpheusplayer.Bounds().Dy())

	oracleplayer, _, _ = ebitenutil.NewImageFromFile("../asset/images/personnages/oracle.png")
	oracleRect = image.Rect(200, 300, 200+oracleplayer.Bounds().Dx(), 300+oracleplayer.Bounds().Dy())

	agentplayer, _, _ = ebitenutil.NewImageFromFile("../asset/images/personnages/agentsmith.png")
	agentRect = image.Rect(550, 300, 550+agentplayer.Bounds().Dx(), 300+agentplayer.Bounds().Dy())

	trinityplayer, _, _ = ebitenutil.NewImageFromFile("../asset/images/personnages/trinity.png")
	trinityRect = image.Rect(650, 300, 650+trinityplayer.Bounds().Dx(), 300+trinityplayer.Bounds().Dy())

	cypherplayer, _, _ = ebitenutil.NewImageFromFile("../asset/images/personnages/cypher.png")
	cypherRect = image.Rect(600, 300, 650+cypherplayer.Bounds().Dx(), 300+cypherplayer.Bounds().Dy())

	archplayer, _, _ = ebitenutil.NewImageFromFile("../asset/images/personnages/artchitecte.png")
	archRect = image.Rect(600, 300, 600+archplayer.Bounds().Dx(), 300+archplayer.Bounds().Dy())

	meufplayer, _, _ = ebitenutil.NewImageFromFile("../asset/images/personnages/lafemmeenrouge.png")
	meufRect = image.Rect(600, 300, 600+meufplayer.Bounds().Dx(), 300+meufplayer.Bounds().Dy())

	// Boutons principaux
	playBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/asset/play.png")
	playRect = image.Rect(250, 450, 250+playBtn.Bounds().Dx(), 450+playBtn.Bounds().Dy())

	exitBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/asset/exit.png")
	exitRect = image.Rect(500, 450, 500+exitBtn.Bounds().Dx(), 450+exitBtn.Bounds().Dy())

	backBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/asset/back.png")
	backRect = image.Rect(50, 50, 50+backBtn.Bounds().Dx(), 50+backBtn.Bounds().Dy())

	fightplay, _, _ = ebitenutil.NewImageFromFile("../asset/images/asset/fight.png")
	fightRect = image.Rect(250, 150, 250+fightplay.Bounds().Dx(), 150+fightplay.Bounds().Dy())

	menuBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/asset/menu.png")
	menuRect = image.Rect(275, 50, 275+menuBtn.Bounds().Dx(), 50+menuBtn.Bounds().Dy())

	forgeBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/asset/forge.png")
	forgeRect = image.Rect(375, 450, 375+forgeBtn.Bounds().Dx(), 450+forgeBtn.Bounds().Dy())

	noticeBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/asset/notice.png")
	noticeRect = image.Rect(50, 450, 50+noticeBtn.Bounds().Dx(), 450+noticeBtn.Bounds().Dy())

	// Boutons inventaire / items
	inventaireOffBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/asset/inventaire_off.png")
	inventaireOffRect = image.Rect(100, 300, 100+inventaireOffBtn.Bounds().Dx(), 300+inventaireOffBtn.Bounds().Dy())

	inventaireOnBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/asset/inventaire_on.png")
	inventaireOnRect = image.Rect(100, 300, 100+inventaireOnBtn.Bounds().Dx(), 300+inventaireOnBtn.Bounds().Dy())

	pilredBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/potion (pillule)/pilule_rouge.png")
	pilredRect = image.Rect(425, 410, 425+pilredBtn.Bounds().Dx(), 410+pilredBtn.Bounds().Dy())

	pilblueBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/potion (pillule)/pilule_bleue.png")
	pilblueRect = image.Rect(575, 410, 575+pilblueBtn.Bounds().Dx(), 410+pilblueBtn.Bounds().Dy())

	cadena, _, _ = ebitenutil.NewImageFromFile("../asset/images/asset/cadena.png")

	// Techniques de combat
	boxeBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/art martiaux/boxe.png")
	boxeRect = image.Rect(175, 125, 175+boxeBtn.Bounds().Dx(), 125+boxeBtn.Bounds().Dy())

	judoBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/art martiaux/judo.png")
	judoRect = image.Rect(0, 80, 0+judoBtn.Bounds().Dx(), 80+judoBtn.Bounds().Dy())

	jujutsuBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/art martiaux/jujutsu.png")
	jujutsuRect = image.Rect(0, 75, 0+jujutsuBtn.Bounds().Dx(), 75+jujutsuBtn.Bounds().Dy())

	karateBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/art martiaux/karate.png")
	karateRect = image.Rect(0, 95, 0+karateBtn.Bounds().Dx(), 95+karateBtn.Bounds().Dy())

	lutteBtn, _, _ = ebitenutil.NewImageFromFile("../asset/images/art martiaux/lutte.png")
	lutteRect = image.Rect(0, 80, 0+lutteBtn.Bounds().Dx(), 80+lutteBtn.Bounds().Dy())

	casqueAC, _, _ = ebitenutil.NewImageFromFile("../asset/images/armes/casque.png")
	lutteAC, _, _ = ebitenutil.NewImageFromFile("../asset/images/armes/lutte.png")
	kimonoAC, _, _ = ebitenutil.NewImageFromFile("../asset/images/armes/kimono.png")
	pantAC, _, _ = ebitenutil.NewImageFromFile("../asset/images/armes/pantalon.png")

	// Audio
	audioCtx = audio.NewContext(44100)
}

// === AUDIO ===
func playMenuMusic() {
	if menuPlayer != nil && menuPlayer.IsPlaying() {
		return // La musique est déjà en cours de lecture
	}
	if level1Player != nil && level1Player.IsPlaying() {
		level1Player.Pause() // Arrêter la musique du niveau 1 si elle est en cours de lecture
	}
	if level2Player != nil && level2Player.IsPlaying() {
		level2Player.Pause() // Arrêter la musique du niveau 2 si elle est en cours de lecture
	}
	if level3Player != nil && level3Player.IsPlaying() {
		level3Player.Pause() // Arrêter la musique du niveau 3 si elle est en cours de lecture
	}
	if level4Player != nil && level4Player.IsPlaying() {
		level4Player.Pause() // Arrêter la musique du niveau 4 si elle est en cours de lecture
	}
	if level5Player != nil && level5Player.IsPlaying() {
		level5Player.Pause() // Arrêter la musique du niveau 5 si elle est en cours de lecture
	}
	f, _ := os.Open("../asset/musiques/musique du menu.mp3") // ton fichier mp3
	stream, _ := mp3.Decode(audioCtx, f)
	loop := audio.NewInfiniteLoop(stream, stream.Length())
	menuPlayer, _ = audio.NewPlayer(audioCtx, loop)
	menuPlayer.Play()
}

func playlevel1Music() {
	if level1Player != nil && level1Player.IsPlaying() {
		return // La musique est déjà en cours de lecture
	}
	if menuPlayer != nil && menuPlayer.IsPlaying() {
		menuPlayer.Pause() // Arrêter la musique du menu si elle est en cours de lecture
	}
	f, _ := os.Open("../asset/musiques/musiquelevel1.mp3") // ton fichier mp3
	stream, _ := mp3.Decode(audioCtx, f)
	loop := audio.NewInfiniteLoop(stream, stream.Length())
	level1Player, _ = audio.NewPlayer(audioCtx, loop)
	level1Player.Play()
}

func playlevel2Music() {
	if level2Player != nil && level2Player.IsPlaying() {
		return // La musique est déjà en cours de lecture
	}
	if menuPlayer != nil && menuPlayer.IsPlaying() {
		menuPlayer.Pause() // Arrêter la musique du menu si elle est en cours de lecture
	}
	f, _ := os.Open("../asset/musiques/musiquelevel2.mp3") // ton fichier mp3
	stream, _ := mp3.Decode(audioCtx, f)
	loop := audio.NewInfiniteLoop(stream, stream.Length())
	level2Player, _ = audio.NewPlayer(audioCtx, loop)
	level2Player.Play()
}

func playlevel3Music() {
	if level3Player != nil && level3Player.IsPlaying() {
		return // La musique est déjà en cours de lecture
	}
	if menuPlayer != nil && menuPlayer.IsPlaying() {
		menuPlayer.Pause() // Arrêter la musique du menu si elle est en cours de lecture
	}
	f, _ := os.Open("../asset/musiques/musiquelevel3.mp3") // ton fichier mp3
	stream, _ := mp3.Decode(audioCtx, f)
	loop := audio.NewInfiniteLoop(stream, stream.Length())
	level3Player, _ = audio.NewPlayer(audioCtx, loop)
	level3Player.Play()
}

func playlevel4Music() {
	if level4Player != nil && level4Player.IsPlaying() {
		return // La musique est déjà en cours de lecture
	}
	if menuPlayer != nil && menuPlayer.IsPlaying() {
		menuPlayer.Pause() // Arrêter la musique du menu si elle est en cours de lecture
	}
	f, _ := os.Open("../asset/musiques/musiquelevel4.mp3") // ton fichier mp3
	stream, _ := mp3.Decode(audioCtx, f)
	loop := audio.NewInfiniteLoop(stream, stream.Length())
	level4Player, _ = audio.NewPlayer(audioCtx, loop)
	level4Player.Play()
}

func playlevel5Music() {
	if level5Player != nil && level5Player.IsPlaying() {
		return // La musique est déjà en cours de lecture
	}
	if menuPlayer != nil && menuPlayer.IsPlaying() {
		menuPlayer.Pause() // Arrêter la musique du menu si elle est en cours de lecture
	}
	f, _ := os.Open("../asset/musiques/musiquelevel5.mp3") // ton fichier mp3
	stream, _ := mp3.Decode(audioCtx, f)
	loop := audio.NewInfiniteLoop(stream, stream.Length())
	level5Player, _ = audio.NewPlayer(audioCtx, loop)
	level5Player.Play()
}

func playringMusic() {
	if ringplayer != nil {
		ringplayer.Close()
	}
	f, _ := os.Open("../asset/musiques/ding.mp3")
	stream, _ := mp3.Decode(audioCtx, f)
	ringplayer, _ = audio.NewPlayer(audioCtx, stream)
	ringplayer.Play()
}
