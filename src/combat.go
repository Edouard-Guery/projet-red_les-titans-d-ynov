package main

import "fmt"

type attaque struct {
	Nom          string
	Dommage      int
	regeneration int
	defense      int
}

// Déclaration des attaques du jeu
var (
	// CRÉATURES
	MorsureVeneuse    = attaque{Nom: "Morsure Vénéneuse", Dommage: 15, regeneration: 0, defense: 0}
	CoupDeBouclier    = attaque{Nom: "Coup de Bouclier", Dommage: 8, regeneration: 0, defense: 12}
	VampirismeSauvage = attaque{Nom: "Vampirisme Sauvage", Dommage: 12, regeneration: 6, defense: 0}

	// DEMI-DIEUX
	FureurHeroique = attaque{Nom: "Fureur Héroïque", Dommage: 45, regeneration: 10, defense: 15}
	EgideDeBronze  = attaque{Nom: "Égide de Bronze", Dommage: 20, regeneration: 15, defense: 45}
	EclairCeleste  = attaque{Nom: "Éclair Céleste", Dommage: 65, regeneration: 0, defense: 10}

	// DIEUX
	CataclysmeCosmique = attaque{Nom: "Cataclysme Cosmique", Dommage: 150, regeneration: 40, defense: 80}
	JugementAbsolu     = attaque{Nom: "Jugement Absolu", Dommage: 250, regeneration: 0, defense: 100}
	RenaissanceDivine  = attaque{Nom: "Renaissance Divine", Dommage: 50, regeneration: 150, defense: 150}
)

type Combat struct {
	Joueur  *Personnage
	Monstre *Monstre
	Tour    int
}

func NouveauCombat(p *Personnage, m *Monstre) *Combat {
	return &Combat{
		Joueur:  p,
		Monstre: m,
		Tour:    1,
	}
}

func (c *Combat) LancerDéroulement() {
	fmt.Printf("\n========================================\n")
	fmt.Printf("   DÉBUT DU COMBAT : %s VS %s\n", c.Joueur.Nom, c.Monstre.Nom)
	fmt.Printf("========================================\n")

	for c.Joueur.PV > 0 && c.Monstre.PV > 0 {
		fmt.Printf("\n--- Tour %d ---\n", c.Tour)
		if c.Joueur.Inventaire["GourdeRegeneration"] > 0 {
    		fmt.Println("Le joueur possède l'objet !")
		c.Joueur.GourdeRegeneration()
		}
		fmt.Printf("%s : %d/%d PV | Attaque : %d | Défense : %d\n", c.Joueur.Nom, c.Joueur.PV, c.Joueur.PVMax, c.Joueur.Attaque, c.Joueur.Defense)
		fmt.Printf("%s : %d/%d PV | Attaque : %d | Défense : %d\n", c.Monstre.Nom, c.Monstre.PV, c.Monstre.PVMax, c.Monstre.Attaque, c.Monstre.Defense)

		// Tour du joueur
		c.TourJoueur()
		if c.Monstre.PV <= 0 {
			fmt.Printf("\n🏆 Victoire ! %s a vaincu %s !\n", c.Joueur.Nom, c.Monstre.Nom)
			break
		}

		// Tour du monstre
		c.TourMonstre()
		if c.Joueur.PV <= 0 {
			fmt.Printf("\n💀 Défaite... %s a été terrassé par %s.\n", c.Joueur.Nom, c.Monstre.Nom)
			break
		}

		c.Tour++
	}
}

// TourJoueur propose les choix d'attaque ou l'utilisation d'objets
func (c *Combat) TourJoueur() {
	for {
		fmt.Printf("\nC'est à votre tour, %s !\n", c.Joueur.Nom)

		// 1. Attaque de base
		fmt.Printf("1. Attaque de base (%d dégâts)\n", c.Joueur.Attaque)

		// 2. Affichage dynamique des attaques apprises
		offset := 2
		for i, att := range c.Joueur.AttaquesApprises {
			fmt.Printf("%d. %s (%d dégâts, %d soin)\n", i+offset, att.Nom, att.Dommage, att.regeneration)
		}

		// Option pour utiliser un objet
		indexObjet := len(c.Joueur.AttaquesApprises) + offset
		fmt.Printf("%d. Utiliser un objet (Potion)\n", indexObjet)

		var choix int
		fmt.Print("Choisissez une action : ")
		fmt.Scanln(&choix)

		if choix == 1 {
			c.ExecuterAttaqueJoueur("Attaque de base", c.Joueur.Attaque, 0, 0)
			break // Fin du tour joueur -> tour du monstre
		} else if choix >= offset && choix < indexObjet {
			att := c.Joueur.AttaquesApprises[choix-offset]
			c.ExecuterAttaqueJoueur(att.Nom, att.Dommage, att.regeneration, att.defense)
			break // Fin du tour joueur -> tour du monstre
		} else if choix == indexObjet {
			// Si un objet est bien utilisé, on termine le tour
			if c.Joueur.utiliserObjetdurantcombat() {
				break // Objet utilisé -> tour consommé
			}
			// Si aucun objet n'a été utilisé ou annulation, la boucle continue
		} else {
			fmt.Println("Choix invalide ! Vous effectuez une attaque normale.")
			c.ExecuterAttaqueJoueur("Attaque de base", c.Joueur.Attaque, 0, 0)
			break // Fin du tour
		}
	}
}

// utiliserObjetdurantcombat permet de choisir une potion et renvoie true si un objet a été consommé
func (p *Personnage) utiliserObjetdurantcombat() bool {
	// Filtrer les objets utilisables en combat
	potionsDispo := []string{}
	for nom, qte := range p.Inventaire {
		if qte > 0 && (nom == "Potion de vie (+)" || nom == "Potion de vie (++)" || nom == "Potion d'attaque (+)" || nom == "Potion de défense (+)") {
			potionsDispo = append(potionsDispo, nom)
		}
	}

	if len(potionsDispo) == 0 {
		fmt.Println("⚠️  Vous n'avez aucune potion utilisable en combat !")
		return false // Le tour n'est pas consommé
	}

	fmt.Println("\n--- VOS POTIONS ---")
	for i, nom := range potionsDispo {
		fmt.Printf("%d. %s (x%d)\n", i+1, nom, p.Inventaire[nom])
	}
	fmt.Println("0. Annuler (retour au combat)")

	var choix int
	fmt.Print("Choisissez une potion à consommer : ")
	fmt.Scanln(&choix)

	if choix <= 0 || choix > len(potionsDispo) {
		fmt.Println("Retour au menu de combat.")
		return false // Annulé, ne consomme pas le tour
	}

	nomChoisi := potionsDispo[choix-1]

	switch nomChoisi {
	case "Potion de vie (+)":
		p.PV += 30
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Println("🧪 Vous avez bu une Potion de vie (+) et regagné 30 PV !")
	case "Potion de vie (++)":
		p.PV += 100
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Println("🧪 Vous avez bu une Potion de vie (++) et regagné 100 PV !")
	case "Potion d'attaque (+)":
		p.Attaque += 5
		fmt.Println("🧪 Vous avez bu une Potion d'attaque (+). Votre attaque augmente de 5 !")
	case "Potion de défense (+)":
		p.Defense += 5
		fmt.Println("🧪 Vous avez bu une Potion de défense (+). Votre défense augmente de 5 !")
	}

	p.Inventaire[nomChoisi]--
	if p.Inventaire[nomChoisi] <= 0 {
		delete(p.Inventaire, nomChoisi)
	}

	return true // Tour consommé
}
func (c *Combat) ExecuterAttaqueJoueur(nomAttaque string, degatsBruts, regen, defBonus int) {
	fmt.Printf("\n⚔️  [JOUEUR] %s lance l'attaque [%s] !\n", c.Joueur.Nom, nomAttaque)

	degatsNets := degatsBruts - c.Monstre.Defense
	if degatsNets < 0 {
		degatsNets = 0
	}

	c.Monstre.PV -= degatsNets
	if c.Monstre.PV < 0 {
		c.Monstre.PV = 0
	}
	fmt.Printf("   💥 %s subit %d dégâts (PV restants : %d/%d).\n", c.Monstre.Nom, degatsNets, c.Monstre.PV, c.Monstre.PVMax)

	if regen > 0 {
		c.Joueur.PV += regen
		if c.Joueur.PV > c.Joueur.PVMax {
			c.Joueur.PV = c.Joueur.PVMax
		}
		fmt.Printf("   💚 %s se régénère de %d PV (PV actuels : %d/%d).\n", c.Joueur.Nom, regen, c.Joueur.PV, c.Joueur.PVMax)
	}
}