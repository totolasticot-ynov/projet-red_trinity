package main

import (
	"fmt"
	"projet-red_trinity/models"
)

func main() {
	neoClasse := models.InitNeo()
	player := models.InitUserFromClasse(neoClasse, "Player1")

	fmt.Printf("Bienvenue %s (%s) avec %d PV et %d Dégâts\n",
		player.Nom, player.Classe, player.PvMax, player.Degats)
}
