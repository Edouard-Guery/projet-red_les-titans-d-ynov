package main

import "fmt"

type Personnage struct {
	Nom                    string
	Attaque                int
	Defense                int
	PV                     int
	ToursBonus             bool
	EmplacementsEquipement int
}

func Person() Personnage {
	hero := Personnage{
		Nom:                    "Alexis",
		PV:                     100,
		Attaque:                15,
		Defense:                10,
		ToursBonus:             false,
		EmplacementsEquipement: 3,
	}

	Monstre := Personnage{
		Nom:                    "matias",
		PV:                     100,
		Attaque:                15,
		Defense:                10,
		ToursBonus:             false,
		EmplacementsEquipement: 0,
	}

	fmt.Printf("Ton héros s'appelle %s, il a %d PV, %d d'attaque et %d de défense. %s est nul sur Pixel.\n", hero.Nom, hero.PV, hero.Attaque, hero.Defense, Monstre.Nom)

	return hero
}