package main

import (
	"fmt"
	"math/rand"
)

func (c *Combat) TourMonstre() {
	fmt.Printf("\n👹 C'est au tour de %s !\n", c.Monstre.Nom)

	// Chaque monstre doit avoir au moins une attaque.
	if len(c.Monstre.Attaques) == 0 {
		fmt.Printf("⚠️ %s n'a aucune attaque.\n", c.Monstre.Nom)
		return
	}

	// Choisit une attaque au hasard dans la liste du monstre.
	// S'il y a 3 attaques, chacune a 1 chance sur 3.
	index := rand.Intn(len(c.Monstre.Attaques))
	attaqueChoisie := c.Monstre.Attaques[index]

	// Petit bonus aléatoire : les dégâts changent à chaque tour.
	bonusMaximum := attaqueChoisie.Dommage / 5
	bonusAleatoire := rand.Intn(bonusMaximum + 1)
	degatsBruts := attaqueChoisie.Dommage + bonusAleatoire

	// La défense du joueur réduit les dégâts.
	degatsNets := degatsBruts - c.Joueur.Defense
	if degatsNets < 0 {
		degatsNets = 0
	}

	// Le joueur perd ses PV.
	c.Joueur.PV -= degatsNets
	if c.Joueur.PV < 0 {
		c.Joueur.PV = 0
	}

	// Régénération éventuelle du monstre.
	if attaqueChoisie.regeneration > 0 {
		c.Monstre.PV += attaqueChoisie.regeneration

		if c.Monstre.PV > c.Monstre.PVMax {
			c.Monstre.PV = c.Monstre.PVMax
		}

		fmt.Printf(
			"💚 %s récupère %d PV ! (%d/%d PV)\n",
			c.Monstre.Nom,
			attaqueChoisie.regeneration,
			c.Monstre.PV,
			c.Monstre.PVMax,
		)
	}

	// Affichage du tour.
	fmt.Printf("🎲 %s utilise [%s] !\n", c.Monstre.Nom, attaqueChoisie.Nom)
	fmt.Printf(
		"💥 %d dégâts + %d bonus = %d dégâts bruts.\n",
		attaqueChoisie.Dommage,
		bonusAleatoire,
		degatsBruts,
	)
	fmt.Printf("⚔️ Vous recevez %d dégâts.\n", degatsNets)
	fmt.Printf("❤️ Vos PV : %d/%d\n", c.Joueur.PV, c.Joueur.PVMax)
}