package main

type Monstre struct {
	Nom         string
	Attaque     int
	Defense     int
	PV          int
	PVMax       int
	SpecialTous int // le coup spécial arrive tous les X tours
	SpecialMult int // multiplicateur en % (200 = 200 %)
}

var Orc = Monstre{
	Nom:     "Orc",
	PV:      110,
	PVMax:   110,
	Attaque: 2,
	Defense: 10,
}

var Minotaure = Monstre{
	Nom:     "Minotaure",
	PV:      75,
	PVMax:   75,
	Attaque: 10,
	Defense: 10,
}

var Meduse = Monstre{
	Nom:     "Méduse",
	PV:      50,
	PVMax:   50,
	Attaque: 15,
	Defense: 5,
}

var Cerbere = Monstre{
	Nom:     "Cerbère",
	PV:      100,
	PVMax:   100,
	Attaque: 20,
	Defense: 15,
}

var HydreDeLerne = Monstre{
	Nom:     "Hydre de Lerne",
	PV:      150,
	PVMax:   150,
	Attaque: 10,
	Defense: 15,
}

var Thesee = Monstre{
	Nom:     "Thésée",
	PV:      250,
	PVMax:   250,
	Attaque: 15,
	Defense: 15,
}

var Persee = Monstre{
	Nom:     "Persée",
	PV:      225,
	PVMax:   225,
	Attaque: 35,
	Defense: 15,
}

var Achille = Monstre{
	Nom:     "Achille",
	PV:      275,
	PVMax:   275,
	Attaque: 40,
	Defense: 15,
}

var Heracles = Monstre{
	Nom:     "Héraclès",
	PV:      300,
	PVMax:   300,
	Attaque: 55,
	Defense: 15,
}

var Athena = Monstre{
	Nom:     "Athéna",
	PV:      400,
	PVMax:   400,
	Attaque: 65,
	Defense: 15,
}

var Poseidon = Monstre{
	Nom:     "Poséidon",
	PV:      425,
	PVMax:   425,
	Attaque: 75,
	Defense: 15,
}

var Hades = Monstre{
	Nom:     "Hadès",
	PV:      440,
	PVMax:   440,
	Attaque: 80,
	Defense: 15,
}

var Zeus = Monstre{
	Nom:     "Zeus",
	PV:      480,
	PVMax:   480,
	Attaque: 80,
	Defense: 15,
}

var Chronos = Monstre{
	Nom:     "Chronos",
	PV:      800,
	PVMax:   800,
	Attaque: 90,
	Defense: 15,
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