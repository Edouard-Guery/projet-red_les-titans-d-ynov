package main

type Monstre struct {
	Nom      string
	PV       int
	PVMax    int
	Attaque  int
	Defense  int
	Attaques []attaque
}

// CRÉATURES
var Orc = Monstre{
	Nom:     "Orc",
	PV:      110,
	PVMax:   110,
	Attaque: 2,
	Defense: 10,
	Attaques: []attaque{
		{Nom: "Coup de Clavet", Dommage: 12, regeneration: 0, defense: 0},
		{Nom: "Charge Brutale", Dommage: 18, regeneration: 0, defense: 5},
	},
}

var Minotaure = Monstre{
	Nom:     "Minotaure",
	PV:      75,
	PVMax:   75,
	Attaque: 10,
	Defense: 10,
	Attaques: []attaque{
		{Nom: "Coup de Corne", Dommage: 20, regeneration: 0, defense: 0},
		{Nom: "Pétinement Labyrinthe", Dommage: 25, regeneration: 0, defense: 10},
	},
}

var Meduse = Monstre{
	Nom:     "Méduse",
	PV:      50,
	PVMax:   50,
	Attaque: 15,
	Defense: 5,
	Attaques: []attaque{
		MorsureVeneuse,
		{Nom: "Regard de Pierre", Dommage: 30, regeneration: 0, defense: 20},
	},
}

var Cerbere = Monstre{
	Nom:     "Cerbère",
	PV:      100,
	PVMax:   100,
	Attaque: 20,
	Defense: 15,
	Attaques: []attaque{
		VampirismeSauvage,
		{Nom: "Triple Morsure", Dommage: 35, regeneration: 5, defense: 0},
	},
}

var HydreDeLerne = Monstre{
	Nom:     "Hydre de Lerne",
	PV:      150,
	PVMax:   150,
	Attaque: 10,
	Defense: 15,
	Attaques: []attaque{
		MorsureVeneuse,
		{Nom: "Régénérationde Têtes", Dommage: 15, regeneration: 25, defense: 5},
	},
}

// DEMI-DIEUX
var Thesee = Monstre{
	Nom:     "Thésée",
	PV:      250,
	PVMax:   250,
	Attaque: 15,
	Defense: 15,
	Attaques: []attaque{
		CoupDeBouclier,
		{Nom: "Estocade du Héros", Dommage: 40, regeneration: 0, defense: 10},
	},
}

var Persee = Monstre{
	Nom:     "Persée",
	PV:      225,
	PVMax:   225,
	Attaque: 35,
	Defense: 15,
	Attaques: []attaque{
		FureurHeroique,
		{Nom: "Vol de Pégase", Dommage: 50, regeneration: 0, defense: 15},
	},
}

var Achille = Monstre{
	Nom:     "Achille",
	PV:      275,
	PVMax:   275,
	Attaque: 40,
	Defense: 15,
	Attaques: []attaque{
		EgideDeBronze,
		{Nom: "Lance Myrmidon", Dommage: 60, regeneration: 10, defense: 20},
	},
}

var Heracles = Monstre{
	Nom:     "Héraclès",
	PV:      300,
	PVMax:   300,
	Attaque: 55,
	Defense: 15,
	Attaques: []attaque{
		FureurHeroique,
		{Nom: "Masse du Lion", Dommage: 80, regeneration: 0, defense: 10},
	},
}

// DIEUX ET TITANS
var Athena = Monstre{
	Nom:     "Athéna",
	PV:      400,
	PVMax:   400,
	Attaque: 65,
	Defense: 15,
	Attaques: []attaque{
		EgideDeBronze,
		EclairCeleste,
		{Nom: "Stratagème Divin", Dommage: 90, regeneration: 20, defense: 50},
	},
}

var Poseidon = Monstre{
	Nom:     "Poséidon",
	PV:      425,
	PVMax:   425,
	Attaque: 75,
	Defense: 15,
	Attaques: []attaque{
		{Nom: "Coup de Trident", Dommage: 100, regeneration: 0, defense: 20},
		{Nom: "Tsunami Déchaîné", Dommage: 130, regeneration: 30, defense: 30},
	},
}

var Hades = Monstre{
	Nom:     "Hadès",
	PV:      440,
	PVMax:   440,
	Attaque: 80,
	Defense: 15,
	Attaques: []attaque{
		CataclysmeCosmique,
		{Nom: "Rappel des Enfers", Dommage: 110, regeneration: 60, defense: 40},
	},
}

var Zeus = Monstre{
	Nom:     "Zeus",
	PV:      480,
	PVMax:   480,
	Attaque: 80,
	Defense: 15,
	Attaques: []attaque{
		EclairCeleste,
		JugementAbsolu,
		{Nom: "Foudre de l'Olympe", Dommage: 200, regeneration: 0, defense: 30},
	},
}

var Chronos = Monstre{
	Nom:     "Chronos",
	PV:      800,
	PVMax:   800,
	Attaque: 90,
	Defense: 15,
	Attaques: []attaque{
		CataclysmeCosmique,
		RenaissanceDivine,
		{Nom: "Faux du Temps", Dommage: 220, regeneration: 50, defense: 100},
	},
}

var ListeMonstres = []Monstre{
	Orc, Minotaure, Meduse, Cerbere, HydreDeLerne,
	Thesee, Persee, Achille, Heracles,
	Athena, Poseidon, Hades, Zeus, Chronos,
}

func ObtenirMonstre(index int) Monstre {
	if index >= 0 && index < len(ListeMonstres) {
		return ListeMonstres[index]
	}
	return Monstre{}
}