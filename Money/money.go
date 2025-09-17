package money

import (
	"fmt"
)

// Définition de la structure Character (money)
type money struct {
	Name   string
	Class  string
	Health int
	Gold   int // <- Attribut pour le $ (pièce d'or)
}

// Fonction pour créer un nouveau personnage avec 100 $ par défaut
func NewCharacter(name string, class string) money {
	return money{
		Name:   name,
		Class:  class,
		Health: 100, //  100 points de vie au départ
		Gold:   100, // Le joueur commence avec 100 $
	}
}

func main() {
	// Création d’un nouveau personnage
	player := NewCharacter("Arthur", "Guerrier")

	// Affichage des informations du personnage
	fmt.Println("Nom:", player.Name)
	fmt.Println("Classe:", player.Class)
	fmt.Println("Santé:", player.Health)
	fmt.Println("Or:", player.Gold)
}
