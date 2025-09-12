package models

type User struct {
	Nom                   string
	Classe                string
	Niveau                int
	Or                    int
	PvMax                 int
	PvActuel              int
	Inventaire            map[string]int
	ResourcesJ            map[string]int
	Skills                []Skill
	hasReceivedFreePotion bool
	Maxinventaire         int
	Equipement            map[string]Equipment
	Equipements           Equipment
	Degats                int
}

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
		Equipements:           Equipment{},
		hasReceivedFreePotion: false,
		Maxinventaire:         10,
		Degats:                classe.DegatsBase,
	}
}
