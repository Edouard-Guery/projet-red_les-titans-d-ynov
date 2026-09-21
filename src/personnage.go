package main

import "fmt"

type Personnage struct {
	Nom                    string
	Classe                 string
	Attaque                int
	Defense                int
	PV                     int
	PVMax				   int
	ToursBonus             bool
	EmplacementsEquipement int
	Argent                 int
	Inventaire             map[string]int // Map pour gérer les qte
	CapaciteMax            int
	Niveau                 int
}

func Person(nomSaisi string, classeSaisie string) Personnage {

	joueur := Personnage{
		Nom:                    nomSaisi,
		Classe:                 classeSaisie,
		ToursBonus:             false,
		EmplacementsEquipement: 3,
		Argent:                 1000,
		CapaciteMax:            10,
		Inventaire:             map[string]int{},
		Niveau:                 1,
	}

	switch classeSaisie {
	case "creature":
		joueur.PV = 80
		joueur.PVMax = 80
		joueur.Attaque = 12
		joueur.Defense = 5
	case "demi-dieu":
		joueur.PV = 120
		joueur.PVMax = 120
		joueur.Attaque = 20
		joueur.Defense = 12
	case "dieu":
		joueur.PV = 200
		joueur.PVMax = 200
		joueur.Attaque = 35
		joueur.Defense = 20
	default:
		fmt.Println("Classe inconnue, choissiser une classe valide : creature, demi-dieu, dieu")
		return Person(nomSaisi, classeSaisie)
	}

	return joueur
}