package models

func InitUserFromClasse(classe Classe, nom string) User {
	return User{
		Nom:                   nom,
		Classe:                classe.Nom,
		Niveau:                1,
		Or:                    0,
		PvMax:                 classe.PvBase,
		PvActuel:              classe.PvBase,
		Inventaire:            make(map[string]int),
		ResourcesJ:            make(map[string]int),
		Skills:                classe.SkillsBase,
		Equipement:            classe.EquipementBase,
		Equipements:           Equipment{}, // si tu veux gérer un équipement actif
		hasReceivedFreePotion: false,
		Maxinventaire:         10,
		Degats:                classe.DegatsBase,
	}
}
