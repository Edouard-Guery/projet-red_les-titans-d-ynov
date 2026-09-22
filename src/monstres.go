package main

type Monstre struct {
	Nom           string
	PV            int
	PVMax         int
	Attaque       int
	Defense       int
	Drop          string
	EstEmpoisonne bool      // Indique si le monstre subit le poison
	DegatsPoison  int       // Dégâts infligés par tour de poison
	Attaques      []attaque // Liste des attaques propres au monstre
}

// CRÉATURES
var Orc = Monstre{
	Nom:     "Orc",
	Defense: 5,
	PV:      110,
	PVMax:   110,
	Attaque: 15,
	Drop:    "Dent d'Orc",
	Attaques: []attaque{
		{Nom: "Coup de Clava", Dommage: 12, regeneration: 0, defense: 0},
		{Nom: "Charge Brutale", Dommage: 18, regeneration: 0, defense: 5},
		{Nom: "Coup de Pied", Dommage: 10, regeneration: 0, defense: 0},
	},
}

var Minotaure = Monstre{
	Nom:     "Minotaure",
	Defense: 5,
	PV:      75,
	PVMax:   75,
	Attaque: 20,
	Drop:    "Corne Brisée",
	Attaques: []attaque{
		{Nom: "Coup de Corne", Dommage: 20, regeneration: 0, defense: 0},
		{Nom: "Pétinement Labyrinthe", Dommage: 25, regeneration: 0, defense: 10},
		{Nom: "Charge Brutale", Dommage: 20, regeneration: 0, defense: 5},
	},
}

var Meduse = Monstre{
	Nom:     "Méduse",
	Defense: 5,
	PV:      50,
	PVMax:   50,
	Attaque: 18,
	Drop:    "Œil Pétrifiant",
	Attaques: []attaque{
		MorsureVeneuse,
		{Nom: "Regard de Pierre", Dommage: 30, regeneration: 0, defense: 20},
		{Nom: "Serpents sifflants", Dommage: 35, regeneration: 0, defense: 0},
	},
}

var Cerbere = Monstre{
	Nom:     "Cerbère",
	Defense: 5,
	PV:      100,
	PVMax:   100,
	Attaque: 25,
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
	Defense: 5,
	PV:      150,
	PVMax:   150,
	Attaque: 30,
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
	Defense: 10,
	PV:      250,
	PVMax:   250,
	Attaque: 35,
	Drop:    "Fil d'Ariane",
	Attaques: []attaque{
		CoupDeBouclier,
		{Nom: "Estocade du Héros", Dommage: 40, regeneration: 0, defense: 10},
		{Nom: "Lame du labyrinthe", Dommage: 40, regeneration: 5, defense: 5},
	},
}

var Persee = Monstre{
	Nom:     "Persée",
	Defense: 10,
	PV:      225,
	PVMax:   225,
	Attaque: 40,
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
	Defense: 15,
	PV:      275,
	PVMax:   275,
	Attaque: 45,
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
	Defense: 20,
	PV:      300,
	PVMax:   300,
	Attaque: 50,
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
	Defense: 25,
	PV:      400,
	PVMax:   400,
	Attaque: 60,
	Drop:    "Éclat de Bouclier",
	Attaques: []attaque{
		EgideDeBronze,
		EclairCeleste,
		{Nom: "Stratagème Divin", Dommage: 90, regeneration: 20, defense: 20},
		{Nom: "Lance de sagesse", Dommage: 80, regeneration: 0, defense: 5},
	},
}

var Poseidon = Monstre{
	Nom:     "Poséidon",
	Defense: 30,
	PV:      425,
	PVMax:   425,
	Attaque: 70,
	Drop:    "Trident Brisé",
	Attaques: []attaque{
		{Nom: "Coup de Trident", Dommage: 100, regeneration: 0, defense: 20},
		{Nom: "Tsunami Déchaîné", Dommage: 130, regeneration: 30, defense: 30},
		{Nom: "Vague abyssale", Dommage: 120, regeneration: 0, defense: 5},
	},
}

var Hades = Monstre{
	Nom:     "Hadès",
	Defense: 35,
	PV:      440,
	PVMax:   440,
	Attaque: 75,
	Drop:    "Cendre des Enfers",
	Attaques: []attaque{
		CataclysmeCosmique,
		{Nom: "Rappel des Enfers", Dommage: 110, regeneration: 60, defense: 40},
		{Nom: "Flammes du Styx", Dommage: 120, regeneration: 0, defense: 30},
	},
}

var Zeus = Monstre{
	Nom:     "Zeus",
	Defense: 40,
	PV:      480,
	PVMax:   480,
	Attaque: 90,
	Drop:    "Éclair Figé",
	Attaques: []attaque{
		EclairCeleste,
		JugementAbsolu,
		{Nom: "Foudre de l'Olympe", Dommage: 200, regeneration: 0, defense: 30},
	},
}

var Chronos = Monstre{
	Nom:     "Chronos",
	Defense: 60,
	PV:      800,
	PVMax:   800,
	Attaque: 110,
	Drop:    "Faux Temporelle",
	Attaques: []attaque{
		CataclysmeCosmique,
		RenaissanceDivine,
		{Nom: "Faux du Temps", Dommage: 320, regeneration: 50, defense: 100},
		{Nom: "Arrêt du temps", Dommage: 300, regeneration: 0, defense: 30},
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
