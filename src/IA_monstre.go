package main

import (
	"fmt"
	"math/rand"
)

func (c *Combat) TourMonstre() {
	fmt.Printf("\n👹 C'est au tour de %s !\n", c.Monstre.Nom)

	nomAttaque := "Coup de griffe"
	degatsBase := c.Monstre.Attaque

	// Si le monstre a des compétences dans sa liste, en choisir une au hasard
	if len(c.Monstre.Attaques) > 0 {
		idx := rand.Intn(len(c.Monstre.Attaques))
		att := c.Monstre.Attaques[idx]
		nomAttaque = att.Nom
		if att.Dommage > 0 {
			degatsBase = att.Dommage
		}
	}

	// Calcul des dégâts réduits par la défense du joueur
	degatsNets := degatsBase - c.Joueur.Defense
	if degatsNets < 0 {
		degatsNets = 0
	}

	c.Joueur.PV -= degatsNets
	if c.Joueur.PV < 0 {
		c.Joueur.PV = 0
	}

	// Affichage clair de l'attaque utilisée par le monstre
	fmt.Printf("   ⚔️  [MONSTRE] %s lance l'attaque [%s] !\n", c.Monstre.Nom, nomAttaque)
	fmt.Printf("   💥 %s vous inflige %d dégâts !\n", c.Monstre.Nom, degatsNets)
	fmt.Printf("   ❤️  Il vous reste %d/%d PV.\n", c.Joueur.PV, c.Joueur.PVMax)
}