package main

import "fmt"

type Attaque struct {
	Nom         string
	Dommage     int
	regeneration int
	defense     int
}

var (
	morsureVeneuse      = Attaque{Nom: "Morsure Vénéneuse", Dommage: 18, regeneration: 7, defense: 0}
	coupDeBouclier      = Attaque{Nom: "Coup de Bouclier", Dommage: 14, regeneration: 0, defense: 4}
	vampirismeSauvage   = Attaque{Nom: "Vampirisme Sauvage", Dommage: 16, regeneration: 6, defense: 0}
	fureurHeroique      = Attaque{Nom: "Fureur Héroïque", Dommage: 20, regeneration: 0, defense: 0}
	egideDeBronze       = Attaque{Nom: "Égide de Bronze", Dommage: 0, regeneration: 10, defense: 6}
	eclairCeleste       = Attaque{Nom: "Éclair Céleste", Dommage: 22, regeneration: 0, defense: 0}
	cataclysmeCosmique  = Attaque{Nom: "Cataclysme Cosmique", Dommage: 28, regeneration: 0, defense: 0}
	jugementAbsolu      = Attaque{Nom: "Jugement Absolu", Dommage: 24, regeneration: 0, defense: 0}
	nenaissanceDivine    = Attaque{Nom: "Renaissance Divine", Dommage: 0, regeneration: 15, defense: 3}
)

// Structure représentant le combat
type Combat struct {
	Joueur       *Personnage
	Monstre      *Monstre
	Tour         int
	JoueurPVMax  int
	MonstreMaxPV int
}

// NouveauCombatMonstre initialise le combat (mémorise les PV de départ comme PV max)
func NouveauCombatMonstre(p *Personnage, m *Monstre) *Combat {
	return &Combat{
		Joueur:       p,
		Monstre:      m,
		Tour:         1,
		JoueurPVMax:  p.PV,
		MonstreMaxPV: m.PV,
	}
}

// LancerDéroulement lance la boucle de combat
func (c *Combat) LancerDéroulement() {
	fmt.Printf("\n========================================\n")
	fmt.Printf("   DÉBUT DU COMBAT : %s VS %s\n", c.Joueur.Nom, c.Monstre.Nom)
	fmt.Printf("========================================\n")

	for c.Joueur.PV > 0 && c.Monstre.PV > 0 {
		fmt.Printf("\n--- Tour %d ---\n", c.Tour)

		fmt.Printf("%s : %d/%d PV | Attaque : %d | Défense : %d\n",
			c.Joueur.Nom, c.Joueur.PV, c.Joueur.PVMax, c.Joueur.Attaque, c.Joueur.Defense)
		fmt.Printf("%s : %d/%d PV | Attaque : %d | Défense : %d\n",
			c.Monstre.Nom, c.Monstre.PV, c.MonstreMaxPV, c.Monstre.Attaque, c.Monstre.Defense)

		// Tour du joueur
		c.TourJoueur()
		if c.Monstre.PV <= 0 {
			fmt.Printf("\n Victoire ! %s a vaincu %s !\n", c.Joueur.Nom, c.Monstre.Nom)
			break
		}

		// Tour du monstre (pattern défini dans IA_monstre.go)
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
		c.ExecuterAttaqueJoueur(morsureVeneuse.Nom, morsureVeneuse.Dommage, morsureVeneuse.regeneration, morsureVeneuse.defense)
	case 3:
		c.ExecuterAttaqueJoueur(coupDeBouclier.Nom, coupDeBouclier.Dommage, coupDeBouclier.regeneration, coupDeBouclier.defense)
	case 4:
		c.ExecuterAttaqueJoueur(vampirismeSauvage.Nom, vampirismeSauvage.Dommage, vampirismeSauvage.regeneration, vampirismeSauvage.defense)
	default:
		fmt.Println("Choix invalide ! Vous attaquez normalement.")
		c.ExecuterAttaqueJoueur("Attaque de base", c.Joueur.Attaque, 0, 0)
	}
}

// ExecuterAttaqueJoueur calcule les dégâts subis par le monstre et la régénération du joueur
func (c *Combat) ExecuterAttaqueJoueur(nomAttaque string, degatsBruts, regen, defBonus int) {
	fmt.Printf("\n%s utilise [%s] !\n", c.Joueur.Nom, nomAttaque)

	degatsNets := degatsBruts - c.Monstre.Defense
	if degatsNets < 0 {
		degatsNets = 0
	}

	c.Monstre.PV -= degatsNets
	if c.Monstre.PV < 0 {
		c.Monstre.PV = 0
	}
	fmt.Printf("-> %s subit %d dégâts (PV : %d/%d).\n", c.Monstre.Nom, degatsNets, c.Monstre.PV, c.MonstreMaxPV)

	if regen > 0 {
		c.Joueur.PV += regen
		if c.Joueur.PV > c.Joueur.PVMax {
			c.Joueur.PV = c.Joueur.PVMax
		}
		fmt.Printf("-> %s se régénère de %d PV (PV : %d/%d).\n", c.Joueur.Nom, regen, c.Joueur.PV, c.Joueur.PVMax)
	}
}

// TourMonstre délègue l'attaque au pattern du monstre
// monstrePattern applique l'IA simple du monstre : attaque de base, avec bonus à certains tours.
func monstrePattern(monstre *Monstre, joueur *Personnage, tour, joueurPVMax int) {
	if monstre == nil || joueur == nil {
		return
	}

	fmt.Printf("\n%s utilise son attaque !\n", monstre.Nom)

	degats := monstre.Attaque
	if tour%3 == 0 {
		degats += 4
		fmt.Printf("-> %s est motivé et frappe plus fort !\n", monstre.Nom)
	}

	degatsNets := degats - joueur.Defense
	if degatsNets < 0 {
		degatsNets = 0
	}

	joueur.PV -= degatsNets
	if joueur.PV < 0 {
		joueur.PV = 0
	}

	fmt.Printf("-> %s subit %d dégâts (PV : %d/%d).\n", joueur.Nom, degatsNets, joueur.PV, joueurPVMax)
}

func (c *Combat) TourMonstre() {
	fmt.Printf("\nC'est au tour de %s !\n", c.Monstre.Nom)
	monstrePattern(c.Monstre, c.Joueur, c.Tour, c.Joueur.PVMax)
}