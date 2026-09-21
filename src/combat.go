package main

import "fmt"

type attaque struct {
	Nom          string
	Dommage      int
	regeneration int
	defense      int
}

// Combat représente un combat en cours entre un personnage et un monstre
type Combat struct {
	Joueur  *Personnage
	Monstre *Monstre
	Tour    int
}

// CRÉATURES
var MorsureVeneuse = attaque{
	Nom:          "Morsure Vénéneuse",
	Dommage:      15,
	regeneration: 0,
	defense:      0,
}

var CoupDeBouclier = attaque{
	Nom:          "Coup de Bouclier",
	Dommage:      8,
	regeneration: 0,
	defense:      12,
}

var VampirismeSauvage = attaque{
	Nom:          "Vampirisme Sauvage",
	Dommage:      12,
	regeneration: 6,
	defense:      0,
}

// DEMI-DIEUX
var FureurHeroique = attaque{
	Nom:          "Fureur Héroïque",
	Dommage:      45,
	regeneration: 10,
	defense:      15,
}

var EgideDeBronze = attaque{
	Nom:          "Égide de Bronze",
	Dommage:      20,
	regeneration: 15,
	defense:      45,
}

var EclairCeleste = attaque{
	Nom:          "Éclair Céleste",
	Dommage:      65,
	regeneration: 0,
	defense:      10,
}

// DIEUX
var CataclysmeCosmique = attaque{
	Nom:          "Cataclysme Cosmique",
	Dommage:      150,
	regeneration: 40,
	defense:      80,
}

var JugementAbsolu = attaque{
	Nom:          "Jugement Absolu",
	Dommage:      250,
	regeneration: 0,
	defense:      100,
}

var RenaissanceDivine = attaque{
	Nom:          "Renaissance Divine",
	Dommage:      50,
	regeneration: 150,
	defense:      150,
}

// NouveauCombat initialise le combat
func NouveauCombat(p *Personnage, m *Monstre) *Combat {
	return &Combat{
		Joueur:  p,
		Monstre: m,
		Tour:    1,
	}
}

// LancerDéroulement lance la boucle de combat
func (c *Combat) LancerDéroulement() {
	fmt.Printf("\n========================================\n")
	fmt.Printf("   DÉBUT DU COMBAT : %s VS %s\n", c.Joueur.Nom, c.Monstre.Nom)
	fmt.Printf("========================================\n")

	for c.Joueur.PV > 0 && c.Monstre.PV > 0 {
		fmt.Printf("\n--- Tour %d ---\n", c.Tour)

		fmt.Printf("%s : %d PV | Attaque : %d | Défense : %d\n", c.Joueur.Nom, c.Joueur.PV, c.Joueur.Attaque, c.Joueur.Defense)
		fmt.Printf("%s : %d PV | Attaque : %d | Défense : %d\n", c.Monstre.Nom, c.Monstre.PV, c.Monstre.Attaque, c.Monstre.Defense)

		// Tour du joueur
		c.TourJoueur()
		if c.Monstre.PV <= 0 {
			fmt.Printf("\n Victoire ! %s a vaincu %s !\n", c.Joueur.Nom, c.Monstre.Nom)
			break
		}

		// Tour du monstre
		c.TourMonstre()
		if c.Joueur.PV <= 0 {
			fmt.Printf("\n Défaite... %s a été terrassé par %s.\n", c.Joueur.Nom, c.Monstre.Nom)
			break
		}

		c.Tour++
	}
}

// TourJoueur propose les choix d'attaque au joueur
func (c *Combat) TourJoueur() {
	fmt.Printf("\nC'est à votre tour, %s !\n", c.Joueur.Nom)
	fmt.Println("1. Attaque de base")
	fmt.Println("2. Morsure Vénéneuse")
	fmt.Println("3. Coup de Bouclier")
	fmt.Println("4. Vampirisme Sauvage")

	var choix int
	fmt.Print("Choisissez une action : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		c.ExecuterAttaqueJoueur("Attaque de base", c.Joueur.Attaque, 0, 0)
	case 2:
		c.ExecuterAttaqueJoueur(MorsureVeneuse.Nom, MorsureVeneuse.Dommage, MorsureVeneuse.regeneration, MorsureVeneuse.defense)
	case 3:
		c.ExecuterAttaqueJoueur(CoupDeBouclier.Nom, CoupDeBouclier.Dommage, CoupDeBouclier.regeneration, CoupDeBouclier.defense)
	case 4:
		c.ExecuterAttaqueJoueur(VampirismeSauvage.Nom, VampirismeSauvage.Dommage, VampirismeSauvage.regeneration, VampirismeSauvage.defense)
	default:
		fmt.Println("Choix invalide ! Vous attaquez normalement.")
		c.ExecuterAttaqueJoueur("Attaque de base", c.Joueur.Attaque, 0, 0)
	}
}

// ExecuterAttaqueJoueur calcule les dégâts subis par le monstre et la régénération du joueur
func (c *Combat) ExecuterAttaqueJoueur(nomAttaque string, degatsBruts, regen, defBonus int) {
	fmt.Printf("\n%s utilise [%s] !\n", c.Joueur.Nom, nomAttaque)

	// Prise en compte de la défense du monstre
	degatsNets := degatsBruts - c.Monstre.Defense
	if degatsNets < 0 {
		degatsNets = 0
	}

	c.Monstre.PV -= degatsNets
	if c.Monstre.PV < 0 {
		c.Monstre.PV = 0
	}
	fmt.Printf("-> %s subit %d dégâts (PV restants : %d).\n", c.Monstre.Nom, degatsNets, c.Monstre.PV)

	// Soin si l'attaque possède de la régénération
	if regen > 0 {
		c.Joueur.PV += regen
		fmt.Printf("-> %s se régénère de %d PV (PV : %d).\n", c.Joueur.Nom, regen, c.Joueur.PV)
	}
}

// TourMonstre gère la riposte du monstre
func (c *Combat) TourMonstre() {
	fmt.Printf("\nC'est au tour de %s !\n", c.Monstre.Nom)

	degatsNets := c.Monstre.Attaque - c.Joueur.Defense
	if degatsNets < 0 {
		degatsNets = 0
	}

	c.Joueur.PV -= degatsNets
	if c.Joueur.PV < 0 {
		c.Joueur.PV = 0
	}

	fmt.Printf("%s attaque et inflige %d dégâts à %s !\n", c.Monstre.Nom, degatsNets, c.Joueur.Nom)
	fmt.Printf("-> %s a maintenant %d PV.\n", c.Joueur.Nom, c.Joueur.PV)
}