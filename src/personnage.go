package main

type Personnage struct {
	Nom                    string
	Attaque                int
	Defense                int
	PV                     int
	ToursBonus             bool
	EmplacementsEquipement int
	Argent                 int
	Inventaire             map[string]int // Map pour gérer les quantités
	CapaciteMax            int
}

func Person() Personnage {
	hero := Personnage{
		Nom:                    "Alexis",
		PV:                     100,
		Attaque:                15,
		Defense:                10,
		ToursBonus:             false,
		EmplacementsEquipement: 3,
		Argent:                 10000,
		CapaciteMax:            10,
		Inventaire:             map[string]int{},
	}

	return hero
}
