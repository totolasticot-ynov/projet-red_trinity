package models

type Classe struct {
	Nom            string
	Description    string
	SkillsBase     []Skill
	DegatsBase     int
	PvBase         int
	EquipementBase map[string]Equipment
}

// Neo
func InitNeo() Classe {
	return Classe{
		Nom:         "Neo",
		Description: "L'Élu, capable de manipuler la Matrice et d'esquiver les balles.",
		SkillsBase: []Skill{
			{Nom: "Coup de poing", Degats: 15, Achete: false},
			{Nom: "Esquive de balles", Degats: 0, Achete: false},
			{Nom: "Code Hack", Degats: 25, Achete: false},
		},
		DegatsBase: 20,
		PvBase:     150,
		EquipementBase: map[string]Equipment{
			"Lunettes": {Nom: "Lunettes noires", Bonus: 2},
			"Manteau":  {Nom: "Manteau long", Bonus: 5},
		},
	}
}

// Morpheus
func InitMorpheus() Classe {
	return Classe{
		Nom:         "Morpheus",
		Description: "Le mentor et guide spirituel, expert en combat rapproché.",
		SkillsBase: []Skill{
			{Nom: "Kung-Fu", Degats: 20, Achete: false},
			{Nom: "Discours motivant", Degats: 0, Achete: false},
		},
		DegatsBase: 18,
		PvBase:     140,
		EquipementBase: map[string]Equipment{
			"Sabre": {Nom: "Katana", Bonus: 10},
		},
	}
}

// Trinity
func InitTrinity() Classe {
	return Classe{
		Nom:         "Trinity",
		Description: "Rapide et agile, experte en piratage et acrobaties.",
		SkillsBase: []Skill{
			{Nom: "Kick retourné", Degats: 18, Achete: false},
			{Nom: "Piratage rapide", Degats: 22, Achete: false},
		},
		DegatsBase: 17,
		PvBase:     120,
		EquipementBase: map[string]Equipment{
			"Pistolets": {Nom: "Deux Berettas", Bonus: 8},
		},
	}
}

// Agent Smith
func InitAgentSmith() Classe {
	return Classe{
		Nom:         "Agent Smith",
		Description: "Programme corrompu, puissant et quasi immortel.",
		SkillsBase: []Skill{
			{Nom: "Cloner", Degats: 0, Achete: false},
			{Nom: "Surpuissance", Degats: 30, Achete: false},
		},
		DegatsBase: 25,
		PvBase:     200,
		EquipementBase: map[string]Equipment{
			"Costume": {Nom: "Costume noir", Bonus: 6},
		},
	}
}
