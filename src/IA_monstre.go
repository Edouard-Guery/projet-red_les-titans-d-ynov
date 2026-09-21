package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// // orcPattern choisit au hasard l'une des deux attaques de l'Orc.
// func orcPattern(orc *Monstre, joueur *Personnage, tour int) {
// 	var nomAttaque string
// 	var degats int

// 	if rand.Intn(2) == 0 {
// 		nomAttaque = "Coup de hache"
// 		degats = orc.Attaque // 100 % de l'attaque
// 	} else {
// 		nomAttaque = "Frappe sauvage"
// 		degats = orc.Attaque * 3 / 2 // 150 % de l'attaque
// 	}

// 	fmt.Printf("%s utilise %s !\n", orc.Nom, nomAttaque)
// 	infligerDegats(orc.Nom, joueur, degats)
// }

// // monstreIA choisit le schéma correspondant au monstre affronté.
// func monstreIA(monstre *Monstre, joueur *Personnage, tour int) {
// 	switch monstre.Nom {
// 	case "Orc":
// 		orcPattern(monstre, joueur, tour)
// 	default:
// 		infligerDegats(monstre.Nom, joueur, monstre.Attaque)
// 	}
// }

// // infligerDegats enlève les PV, puis affiche le résultat demandé.
// func infligerDegats(nomAttaquant string, joueur *Personnage, degats int) {
// 	joueur.PV -= degats
// 	if joueur.PV < 0 {
// 		joueur.PV = 0
// 	}

// 	fmt.Printf("%s inflige à %s %d de dégâts\n", nomAttaquant, joueur.Nom, degats)
// 	fmt.Printf("%s : %d/%d PV\n", joueur.Nom, joueur.PV, joueur.PVMax)
// }

// // Exemple d'utilisation du Gobelin créé dans ton fichier de monstres :
// // gobelin := InitGoblin()
// // goblinPattern(gobelin, joueur, tour)
// //
// // Exemple d'utilisation de l'Orc créé dans ton fichier de monstres :
// // orc := InitOrc()
// // orcPattern(orc, joueur, tour)
// //
// // Dans TourMonstre, remplace l'appel à monstrePattern(...) par :
// // monstreIA(c.Monstre, c.Joueur, c.Tour)
// func (c *Combat) LancerDéroulement() {
// 	fmt.Printf("\n========================================\n")
// 	fmt.Printf("   DÉBUT DU COMBAT : %s VS %s\n", c.Joueur.Nom, c.Monstre.Nom)
// 	fmt.Printf("========================================\n")

// 	for c.Joueur.PV > 0 && c.Monstre.PV > 0 {
// 		fmt.Printf("\n--- Tour %d ---\n", c.Tour)
		
// 		fmt.Printf("%s : %d PV | Attaque : %d | Défense : %d\n", c.Joueur.Nom, c.Joueur.PV, c.Joueur.Attaque, c.Joueur.Defense)
// 		fmt.Printf("%s : %d PV | Attaque : %d | Défense : %d\n", c.Monstre.Nom, c.Monstre.PV, c.Monstre.Attaque, c.Monstre.Defense)

// 		// Tour du joueur
// 		c.TourJoueur()
// 		if c.Monstre.PV <= 0 {
// 			fmt.Printf("\n Victoire ! %s a vaincu %s !\n", c.Joueur.Nom, c.Monstre.Nom)
// 			break
// 		}

// 		// Tour du monstre
// 		c.TourMonstre()
// 		if c.Joueur.PV <= 0 {
// 			fmt.Printf("\n Défaite... %s a été terrassé par %s.\n", c.Joueur.Nom, c.Monstre.Nom)
// 			break
// 		}

// 		c.Tour++
// 	}
// }
