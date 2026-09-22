package main

import (
	"fmt"
	"math/rand/v2"
)

func (c *Combat) TourMonstre() {
	fmt.Printf("\n"+Red+"👹 C'est au tour de %s !"+Reset+"\n", c.Monstre.Nom)

	if len(c.Monstre.Attaques) == 0 {
		fmt.Printf(Gray+"⚠️ %s n'a aucune attaque disponible.\n"+Reset, c.Monstre.Nom)
		return
	}

	index := rand.N(len(c.Monstre.Attaques))
	attaqueChoisie := c.Monstre.Attaques[index]

	bonusAleatoire := 0
	if attaqueChoisie.Dommage > 0 {
		bonusAleatoire = rand.N((attaqueChoisie.Dommage / 5) + 1)
	}
	degatsBruts := attaqueChoisie.Dommage + bonusAleatoire

	degatsNets := degatsBruts - (c.Joueur.Defense / 2)

	if degatsBruts > 0 && degatsNets < 1 {
		degatsNets = 1
	} else if degatsBruts <= 0 {
		degatsNets = 0
	}
	c.Joueur.PV -= degatsNets
	if c.Joueur.PV < 0 {
		c.Joueur.PV = 0
	}

	fmt.Printf(Bold+Red+"🎲 %s utilise [%s] !\n"+Reset, c.Monstre.Nom, attaqueChoisie.Nom)
	fmt.Printf(White+"💥 %d dégâts bruts (%d de base + %d bonus)\n"+Reset, degatsBruts, attaqueChoisie.Dommage, bonusAleatoire)

	absorption := degatsBruts - degatsNets
	if absorption > 0 {
		fmt.Printf(Blue+"🛡️ Votre armure absorbe %d dégâts.\n"+Reset, absorption)
	}

	fmt.Printf(Red+"⚔️ Vous perdez %d PV.\n"+Reset, degatsNets)

	if attaqueChoisie.regeneration > 0 {
		c.Monstre.PV += attaqueChoisie.regeneration
		if c.Monstre.PV > c.Monstre.PVMax {
			c.Monstre.PV = c.Monstre.PVMax
		}
		fmt.Printf(Green+"💚 %s se régénère de %d PV (%d/%d PV).\n"+Reset, c.Monstre.Nom, attaqueChoisie.regeneration, c.Monstre.PV, c.Monstre.PVMax)
	}

	fmt.Printf(Cyan+"❤️ Vos PV actuels : %d/%d\n"+Reset, c.Joueur.PV, c.Joueur.PVMax)
}
