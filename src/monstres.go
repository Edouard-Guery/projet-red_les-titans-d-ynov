package main

type Monstre struct {
	Nom     string
	Attaque int
	Defense int
	PV      int
}

var Orc = Monstre{
	Nom:     "Orc",
	PV:      110,
	Attaque: 2,
	Defense: 10,
}

var Minotaure = Monstre{
	Nom:     "Minotaure",
	PV:      75,
	Attaque: 10,
	Defense: 10,
}

var Meduse = Monstre{
	Nom:     "Méduse",
	PV:      50,
	Attaque: 15,
	Defense: 5,
}

var cerbere = Monstre{
	Nom:     "Cerbère",
	PV:      100,
	Attaque: 20,
	Defense: 15,
}

var Hydre_de_lerne  = Monstre{
	Nom:     "Hydre de Lerne",
	PV:      150,
	Attaque: 10,
	Defense: 15,
}

var thésée = Monstre{
	Nom:     "Thésée",
	PV:      250,
	Attaque: 15,
	Defense: 15,
}

var persée = Monstre{
	Nom:     "Persée",
	PV:      225,
	Attaque: 35,
	Defense: 15,
}

var achille = Monstre{
	Nom:     "Achille",
	PV:      275,
	Attaque: 40,
	Defense: 15,
}

var heraclès = Monstre{
	Nom:     "Héraclès",
	PV:      300,
	Attaque: 55,
	Defense: 15,
}

var athéna = Monstre{
	Nom:     "Athéna",
	PV:      400,
	Attaque: 65,
	Defense: 15,
}

var poséidon = Monstre{
	Nom:     "Poséidon",
	PV:      425,
	Attaque: 75,
	Defense: 15,
}

var hadès = Monstre{
	Nom:     "Hadès",
	PV:      440,
	Attaque: 80,
	Defense: 15,
}

var zeus = Monstre{
	Nom:     "Zeus",
	PV:      480,
	Attaque: 80,
	Defense: 15,
} 

var chronos = Monstre{
	Nom:     "Chronos",
	PV:      800,
	Attaque: 90,
	Defense: 15,
}