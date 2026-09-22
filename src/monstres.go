package main

type Monstre struct {
	Nom      string
	PV       int
	PVMax    int
	Attaques []attaque
	Drop     string
}

// CRÉATURES
var Orc = Monstre{
	Nom:     "Orc",
	PV:      110,
	PVMax:   110,
	Drop:    "Dent d'Orc",
	Attaques: []attaque{
		{Nom: "Coup de Clavet", Dommage: 12, regeneration: 0, defense: 0},
		{Nom: "Charge Brutale", Dommage: 18, regeneration: 0, defense: 5},
		{Nom: "Coup de Pied", Dommage: 10, regeneration: 0, defense: 0},
	},
}

var Minotaure = Monstre{
	Nom:     "Minotaure",
	PV:      75,
	PVMax:   75,
	Drop:    "Corne Brisée",
	Attaques: []attaque{
		{Nom: "Coup de Corne", Dommage: 20, regeneration: 0, defense: 0},
		{Nom: "Pétinement Labyrinthe", Dommage: 25, regeneration: 0, defense: 10},
		{Nom: "Charge Brutale", Dommage: 20, regeneration: 0, defense: 5},
	},
}

var Meduse = Monstre{
	Nom:     "Méduse",
	PV:      50,
	PVMax:   50, 
	Drop:    "Œil Pétrifiant",
	Attaques: []attaque{
		MorsureVeneuse, // Fait référence à la variable globale dans combat.go
		{Nom: "Regard de Pierre", Dommage: 30, regeneration: 0, defense: 20},
		{Nom: "Morsure Vénéneuse", Dommage: 30, regeneration: 0, defense: 2},
		{Nom: "Serpents sifflants", Dommage: 35, regeneration: 0, defense: 0},
	},
}

var Cerbere = Monstre{
	Nom:     "Cerbère",
	PV:      100,
	PVMax:   100,
	Drop:    "Croc Infernal",
	Attaques: []attaque{
		VampirismeSauvage,
		{Nom: "Triple Morsure", Dommage: 35, regeneration: 5, defense: 0},
		{Nom: "Souffle infernal", Dommage: 35, regeneration: 0, defense: 5},
		{Nom: "Griffe des Enfers", Dommage: 40, regeneration: 0, defense: 0},
	},
}

var HydreDeLerne = Monstre{
	Nom:     "Hydre de Lerne",
	PV:      150,
	PVMax:   150,
	Drop:    "Écaille Régénérante",
	Attaques: []attaque{
		MorsureVeneuse,
		{Nom: "Régénération de Têtes", Dommage: 15, regeneration: 25, defense: 5},
		{Nom: "Morsure des neuf têtes", Dommage: 40, regeneration: 0, defense: 0},
		{Nom: "Venin de Lerne", Dommage: 30, regeneration: 0, defense: 0},
	},
}

// DEMI-DIEUX
var Thesee = Monstre{
	Nom:     "Thésée",
	PV:      250,
	PVMax:   250,
	Drop:    "Fil d'Ariane",
	Attaques: []attaque{
		CoupDeBouclier,
		{Nom: "Estocade du Héros", Dommage: 40, regeneration: 0, defense: 10},
		{Nom: "Coup de bouclier", Dommage: 20, regeneration: 0, defense: 30},
		{Nom: "Lame du labyrinthe", Dommage: 40, regeneration: 5, defense: 5},
		
	},
}

var Persee = Monstre{
	Nom:     "Persée",
	PV:      225,
	PVMax:   225,
	Drop:    "Fiole de Venin",
	Attaques: []attaque{
		FureurHeroique,
		{Nom: "Vol de Pégase", Dommage: 50, regeneration: 0, defense: 15},
		{Nom: "Decapitation", Dommage: 60, regeneration: 0, defense: 0},
		{Nom: "Coup d'épée", Dommage: 45, regeneration: 0, defense: 5},
	},
}

var Achille = Monstre{
	Nom:     "Achille",
	PV:      275,
	PVMax:   275,
	Drop:    "Fragment de Lame",
	Attaques: []attaque{
		EgideDeBronze,
		{Nom: "Lance Myrmidon", Dommage: 60, regeneration: 10, defense: 20},
		{Nom: "Colère d'Achille", Dommage: 65, regeneration: 0, defense: 0},
		{Nom: "Frappe du talon invincible", Dommage: 60, regeneration: 5, defense: 5},
	},
}

var Heracles = Monstre{
	Nom:     "Héraclès",
	PV:      300,
	PVMax:   300,
	Drop:    "Peau du Lion",
	Attaques: []attaque{
		FureurHeroique,
		{Nom: "Masse du Lion", Dommage: 80, regeneration: 0, defense: 10},
		{Nom: "Force des douze travaux", Dommage: 50, regeneration: 20, defense: 5},
		{Nom: "Poing de Némée", Dommage: 75, regeneration: 0, defense: 5},
	},
}

// DIEUX ET TITANS
var Athena = Monstre{
	Nom:     "Athéna",
	PV:      400,
	PVMax:   400,
	Drop:    "Éclat de Bouclier",
	Attaques: []attaque{
		EgideDeBronze,
		EclairCeleste,
		{Nom: "Stratagème Divin", Dommage: 90, regeneration: 20, defense: 20},
		{Nom: "Lance de sagesse", Dommage: 80, regeneration: 0, defense: 5},
		{Nom: "Egide divine", Dommage: 85, regeneration: 0, defense: 20},
	},
}

var Poseidon = Monstre{
	Nom:     "Poséidon",
	PV:      425,
	PVMax:   425,
	Drop:    "Trident Brisé",
	Attaques: []attaque{
		{Nom: "Coup de Trident", Dommage: 100, regeneration: 0, defense: 20},
		{Nom: "Tsunami Déchaîné", Dommage: 130, regeneration: 30, defense: 30},
		{Nom: "Vague abyssale", Dommage: 120, regeneration: 0, defense: 5},
	},
}

var Hades = Monstre{
	Nom:     "Hadès",
	PV:      440,
	PVMax:   440,
	Drop:    "Cendre des Enfers",
	Attaques: []attaque{
		CataclysmeCosmique,
		{Nom: "Rappel des Enfers", Dommage: 110, regeneration: 60, defense: 40},
		{Nom: "Flammes du Styx", Dommage: 120, regeneration: 0, defense: 30},
		{Nom: "Chaînes des morts", Dommage: 115, regeneration: 20, defense: 5},
	},
}

var Zeus = Monstre{
	Nom:     "Zeus",
	PV:      480,
	PVMax:   480,
	Drop:    "Éclair Figé",
	Attaques: []attaque{
		EclairCeleste,
		JugementAbsolu,
		{Nom: "Foudre de l'Olympe", Dommage: 200, regeneration: 0, defense: 30},
		{Nom: "Eclair céleste", Dommage: 210, regeneration: 10, defense: 5},
		{Nom: "Jugement absolu", Dommage: 300, regeneration: 0, defense: 20},
	},
}

var Chronos = Monstre{
	Nom:     "Chronos",
	PV:      800,
	PVMax:   800,
	Drop:    "",
	Attaques: []attaque{
		CataclysmeCosmique,
		RenaissanceDivine,
		{Nom: "Faux du Temps", Dommage: 320, regeneration: 50, defense: 100},
		{Nom: "Arrêt du temps", Dommage: 300, regeneration: 0, defense: 30},
		{Nom: "Renaissance divine", Dommage: 20, regeneration: 400, defense: 5},
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