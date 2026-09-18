package main

import "fmt"

type Personnage struct {
	Nom        string
	Attaque    int
	Defense    int
	PV         int
	ToursBonus bool
}

func Person() {
	// Création du personnage
	hero := Personnage{
		Nom:        "Alexis",
		PV:         100,
		Attaque:    15,
		Defense:    10,
		ToursBonus: false,
	}

	Monstre := Personnage{
		Nom:        "matias",
		PV:         100,
		Attaque:    15,
		Defense:    10,
		ToursBonus: false,
	}

fmt.Printf(" ton héros s'appelle %s, il a %d PV et %d dattaque et %d de point de defense, %s est nul sur poxel\n",hero.Nom,hero.PV,hero.Attaque,hero.Defense,Monstre.Nom)}