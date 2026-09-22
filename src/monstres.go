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

// ==========================================
// PALIER 1 : CRÉATURES (Tutoriel & Début de jeu)
// ==========================================
var Orc = Monstre{
	Nom:     "Orc",
	Defense: 2,
	PV:      50, // Nerf des PV de départ (110 c'était très long au niveau 1)
	PVMax:   50,
	Attaque: 8,
	Drop:    "Dent d'Orc",
	Attaques: []attaque{
		{Nom: "Coup de Clava", Dommage: 8, regeneration: 0, defense: 0},
		{Nom: "Charge Brutale", Dommage: 12, regeneration: 0, defense: 2},
		{Nom: "Coup de Pied", Dommage: 6, regeneration: 0, defense: 0},
	},
}

var Minotaure = Monstre{
	Nom:     "Minotaure",
	Defense: 5,
	PV:      80, // Repasse au-dessus de l'Orc
	PVMax:   80,
	Attaque: 12,
	Drop:    "Corne Brisée",
	Attaques: []attaque{
		{Nom: "Coup de Corne", Dommage: 15, regeneration: 0, defense: 0},
		{Nom: "Piétinement du Labyrinthe", Dommage: 18, regeneration: 0, defense: 5},
		{Nom: "Charge Enragée", Dommage: 14, regeneration: 0, defense: 2},
	},
}

var Meduse = Monstre{
	Nom:     "Méduse",
	Defense: 5,
	PV:      100, // Ajustement de la courbe de vie
	PVMax:   100,
	Attaque: 14,
	Drop:    "Œil Pétrifiant",
	Attaques: []attaque{
		MorsureVeneuse,
		{Nom: "Regard de Pierre", Dommage: 18, regeneration: 0, defense: 10},
		{Nom: "Serpents sifflants", Dommage: 20, regeneration: 0, defense: 0},
	},
}

var Cerbere = Monstre{
	Nom:     "Cerbère",
	Defense: 8,
	PV:      130,
	PVMax:   130,
	Attaque: 18,
	Drop:    "Croc Infernal",
	Attaques: []attaque{
		VampirismeSauvage,
		{Nom: "Triple Morsure", Dommage: 22, regeneration: 5, defense: 0},
		{Nom: "Souffle infernal", Dommage: 25, regeneration: 0, defense: 5},
		{Nom: "Griffe des Enfers", Dommage: 20, regeneration: 0, defense: 0},
	},
}

var HydreDeLerne = Monstre{
	Nom:     "Hydre de Lerne",
	Defense: 10,
	PV:      160,
	PVMax:   160,
	Attaque: 22,
	Drop:    "Écaille Régénérante",
	Attaques: []attaque{
		MorsureVeneuse,
		{Nom: "Régénération de Têtes", Dommage: 15, regeneration: 20, defense: 5},
		{Nom: "Morsure des neuf têtes", Dommage: 28, regeneration: 0, defense: 0},
		{Nom: "Venin de Lerne", Dommage: 25, regeneration: 0, defense: 0},
	},
}

// ==========================================
// PALIER 2 : DEMI-DIEUX (Milieu de jeu)
// ==========================================
var Thesee = Monstre{
	Nom:     "Thésée",
	Defense: 15,
	PV:      220,
	PVMax:   220,
	Attaque: 28,
	Drop:    "Fil d'Ariane",
	Attaques: []attaque{
		CoupDeBouclier,
		{Nom: "Estocade du Héros", Dommage: 32, regeneration: 0, defense: 10},
		{Nom: "Lame du labyrinthe", Dommage: 30, regeneration: 5, defense: 5},
	},
}

var Persee = Monstre{
	Nom:     "Persée",
	Defense: 18,
	PV:      260,
	PVMax:   260,
	Attaque: 32,
	Drop:    "Fiole de Venin",
	Attaques: []attaque{
		FureurHeroique,
		{Nom: "Vol de Pégase", Dommage: 38, regeneration: 0, defense: 15},
		{Nom: "Décapitation", Dommage: 45, regeneration: 0, defense: 0},
		{Nom: "Coup d'épée", Dommage: 35, regeneration: 0, defense: 5},
	},
}

var Achille = Monstre{
	Nom:     "Achille",
	Defense: 25,
	PV:      300,
	PVMax:   300,
	Attaque: 35,
	Drop:    "Fragment de Lame",
	Attaques: []attaque{
		EgideDeBronze,
		{Nom: "Lance Myrmidon", Dommage: 42, regeneration: 10, defense: 20},
		{Nom: "Colère d'Achille", Dommage: 48, regeneration: 0, defense: 0},
		{Nom: "Frappe Rapide", Dommage: 38, regeneration: 5, defense: 5},
	},
}

var Heracles = Monstre{
	Nom:     "Héraclès",
	Defense: 30,
	PV:      350,
	PVMax:   350,
	Attaque: 40,
	Drop:    "Peau du Lion",
	Attaques: []attaque{
		FureurHeroique,
		{Nom: "Masse du Lion", Dommage: 55, regeneration: 0, defense: 10},
		{Nom: "Force des douze travaux", Dommage: 45, regeneration: 20, defense: 5},
		{Nom: "Poing de Némée", Dommage: 50, regeneration: 0, defense: 5},
	},
}

// ==========================================
// PALIER 3 : DIEUX ET TITANS (Fin de jeu)
// ==========================================
var Athena = Monstre{
	Nom:     "Athéna",
	Defense: 40,
	PV:      420,
	PVMax:   420,
	Attaque: 48,
	Drop:    "Éclat de Bouclier",
	Attaques: []attaque{
		EgideDeBronze,
		EclairCeleste, // ~50 dégâts
		{Nom: "Stratagème Divin", Dommage: 60, regeneration: 20, defense: 20},
		{Nom: "Lance de Sagesse", Dommage: 55, regeneration: 0, defense: 5},
	},
}

var Poseidon = Monstre{
	Nom:     "Poséidon",
	Defense: 45,
	PV:      500,
	PVMax:   500,
	Attaque: 55,
	Drop:    "Trident Brisé",
	Attaques: []attaque{
		{Nom: "Coup de Trident", Dommage: 70, regeneration: 0, defense: 20},
		{Nom: "Tsunami Déchaîné", Dommage: 85, regeneration: 25, defense: 15},
		{Nom: "Vague Abyssale", Dommage: 75, regeneration: 0, defense: 10},
	},
}

var Hades = Monstre{
	Nom:     "Hadès",
	Defense: 50,
	PV:      580,
	PVMax:   580,
	Attaque: 60,
	Drop:    "Cendre des Enfers",
	Attaques: []attaque{
		CataclysmeCosmique, // ~100 dégâts
		{Nom: "Rappel des Enfers", Dommage: 80, regeneration: 40, defense: 20},
		{Nom: "Flammes du Styx", Dommage: 95, regeneration: 0, defense: 15},
	},
}

var Zeus = Monstre{
	Nom:     "Zeus",
	Defense: 55,
	PV:      650,
	PVMax:   650,
	Attaque: 70,
	Drop:    "Éclair Figé",
	Attaques: []attaque{
		EclairCeleste,
		JugementAbsolu, // ~150 dégâts
		{Nom: "Foudre de l'Olympe", Dommage: 110, regeneration: 0, defense: 20},
	},
}

var Chronos = Monstre{
	Nom:     "Chronos (Boss Final)",
	Defense: 70,
	PV:      900,
	PVMax:   900,
	Attaque: 85,
	Drop:    "Faux Temporelle",
	Attaques: []attaque{
		CataclysmeCosmique,
		RenaissanceDivine,
		{Nom: "Faux du Temps", Dommage: 140, regeneration: 30, defense: 30}, // Dégâts lissés (320 = one-shot du joueur)
		{Nom: "Arrêt du Temps", Dommage: 120, regeneration: 0, defense: 50},
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
