package main

import (
	"fmt"
	"math/rand/v2"
)

type attaque struct {
	Nom           string
	Dommage       int
	regeneration  int
	defense       int
	CoutEndurance int // Coût en endurance de l'attaque
}

// ÉQUILIBRAGE DES COMPÉTENCES
var (
	MorsureVeneuse     = attaque{"Morsure Vénéneuse", 12, 0, 0, 5}
	CoupDeBouclier     = attaque{"Coup de Bouclier", 5, 0, 15, 8}
	VampirismeSauvage  = attaque{"Vampirisme Sauvage", 15, 10, 0, 12}
	FureurHeroique     = attaque{"Fureur Héroïque", 35, 0, -5, 15} // Baisse de def en contrepartie
	EgideDeBronze      = attaque{"Égide de Bronze", 10, 5, 30, 20}
	EclairCeleste      = attaque{"Éclair Céleste", 50, 0, 0, 25}
	CataclysmeCosmique = attaque{"Cataclysme Cosmique", 100, 0, 0, 45}
	JugementAbsolu     = attaque{"Jugement Absolu", 150, 0, 50, 60}     // Nerf des dégâts bruts (250 c'était trop)
	RenaissanceDivine  = attaque{"Renaissance Divine", 20, 100, 50, 80} // Nerf du soin abusif
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
	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Printf(Bold+Red+"          ⚔️  DÉBUT DU COMBAT : %s VS %s ⚔️\n"+Reset, c.Joueur.Nom, c.Monstre.Nom)
	fmt.Println(Magenta + "============================================================" + Reset)

	for c.Joueur.PV > 0 && c.Monstre.PV > 0 {
		fmt.Printf("\n"+Yellow+"--- ⏳ Tour %d ---"+Reset+"\n", c.Tour)

		// --- EFFET DE POISON AU DÉBUT DU TOUR ---
		if c.Monstre.EstEmpoisonne && c.Monstre.PV > 0 {
			c.Monstre.PV -= c.Monstre.DegatsPoison
			if c.Monstre.PV < 0 {
				c.Monstre.PV = 0
			}
			fmt.Printf(Red+"☠️  %s souffre du poison et subit %d dégâts ! (PV : %d/%d)\n"+Reset, c.Monstre.Nom, c.Monstre.DegatsPoison, c.Monstre.PV, c.Monstre.PVMax)

			if c.Monstre.PV <= 0 {
				break
			}
		}

		c.Joueur.Endurance += 15
		if c.Joueur.Endurance > c.Joueur.EnduranceMax {
			c.Joueur.Endurance = c.Joueur.EnduranceMax
		}

		if c.Joueur.Inventaire["GourdeRegeneration"] > 0 {
			fmt.Println(Green + "🌿 La Gourde de Régénération restaure vos PV !" + Reset)
			c.Joueur.GourdeRegeneration()
		}

		fmt.Printf(Cyan+"👤 %s "+Reset+"| ❤️  PV : %d/%d | ⚡ Endu : %d/%d | 🗡️  Attaque : %d | 🛡️  Défense : %d\n", c.Joueur.Nom, c.Joueur.PV, c.Joueur.PVMax, c.Joueur.Endurance, c.Joueur.EnduranceMax, c.Joueur.Attaque, c.Joueur.Defense)
		fmt.Printf(Red+"👹 %s "+Reset+"| ❤️  PV : %d/%d\n", c.Monstre.Nom, c.Monstre.PV, c.Monstre.PVMax)

		c.TourJoueur()

		if c.Monstre.PV <= 0 {
			// RÉCOMPENSE DYNAMIQUE : Basée sur les PV Max du monstre (ex: un monstre à 100 PV donne 10 à 20 oboles)
			recompense := (c.Monstre.PVMax / 10) + rand.N(10)
			if recompense < 5 {
				recompense = 5
			} // Minimum garanti

			c.Joueur.Argent += recompense

			fmt.Println("\n" + Magenta + "============================================================" + Reset)
			fmt.Printf(Bold+Green+"🏆 VICTOIRE ! %s a vaincu %s !\n"+Reset, c.Joueur.Nom, c.Monstre.Nom)
			fmt.Printf(Yellow+"💰 Le butin s'élève à %d oboles !\n"+Reset, recompense)
			fmt.Printf(Cyan+"💰 Oboles total : %d pièces.\n"+Reset, c.Joueur.Argent)
			fmt.Println(Magenta + "============================================================" + Reset)

			break
		}

		c.TourMonstre()

		if c.Joueur.PV <= 0 {
			break
		}

		c.Tour++
	}

	c.Monstre.EstEmpoisonne = false
	c.Monstre.DegatsPoison = 0

	if c.Monstre.PV <= 0 {
		fmt.Println("\n" + Magenta + "============================================================" + Reset)
		fmt.Printf(Bold+Green+"🏆 %s sort victorieux du combat !\n"+Reset, c.Joueur.Nom)
		fmt.Println(Magenta + "============================================================" + Reset)
	} else {
		fmt.Println("\n" + Magenta + "============================================================" + Reset)
		fmt.Printf(Bold+Red+"💀 DÉFAITE... %s a été terrassé par %s.\n"+Reset, c.Joueur.Nom, c.Monstre.Nom)
		fmt.Println(Magenta + "============================================================" + Reset)
	}
}

func (c *Combat) TourJoueur() {
	for {
		fmt.Println("\n" + Blue + "------------------------------------------------------------" + Reset)
		fmt.Printf(Cyan+"👤 C'est à votre tour, %s !"+Reset+"\n", c.Joueur.Nom)

		coutBase := 5
		fmt.Printf(White+" 1. Attaque de base (%d dégâts, %d ⚡)\n"+Reset, c.Joueur.Attaque, coutBase)

		offset := 2
		for i, att := range c.Joueur.AttaquesApprises {
			fmt.Printf(White+" %d. %s (%d dégâts base, %d soin, %d ⚡)\n"+Reset, i+offset, att.Nom, att.Dommage, att.regeneration, att.CoutEndurance)
		}

		indexObjet := len(c.Joueur.AttaquesApprises) + offset
		fmt.Printf(Yellow+" %d. Utiliser un objet (Potion)\n"+Reset, indexObjet)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)

		var choix int
		fmt.Print(Cyan + "👉 Choisissez une action : " + Reset)
		fmt.Scanln(&choix)

		fmt.Println("\n" + Magenta + "============================================================" + Reset)

		if choix == 1 {
			if c.Joueur.Endurance >= coutBase {
				c.Joueur.Endurance -= coutBase
				c.ExecuterAttaqueJoueur("Attaque de base", c.Joueur.Attaque, 0, 0)
				break
			} else {
				fmt.Println(Red + "❌ Pas assez d'endurance ! Choisissez une autre action." + Reset)
			}
		} else if choix >= offset && choix < indexObjet {
			att := c.Joueur.AttaquesApprises[choix-offset]
			if c.Joueur.Endurance >= att.CoutEndurance {
				c.Joueur.Endurance -= att.CoutEndurance
				c.ExecuterAttaqueJoueur(att.Nom, att.Dommage+c.Joueur.Attaque, att.regeneration, att.defense) // Ajoute l'attaque du joueur aux dégâts de la compétence
				break
			} else {
				fmt.Println(Red + "❌ Pas assez d'endurance pour lancer cette compétence !" + Reset)
			}
		} else if choix == indexObjet {
			if c.Joueur.utiliserObjetdurantcombat(c.Monstre) {
				break
			}
		} else {
			fmt.Println(Red + "❌ Choix invalide !" + Reset)
		}
	}
}

func (p *Personnage) utiliserObjetdurantcombat(m *Monstre) bool {
	potionsDispo := []string{}
	for nom, qte := range p.Inventaire {
		if qte > 0 && (nom == "Potion de vie (+)" || nom == "Potion de vie (++)" || nom == "Potion d'attaque (+)" || nom == "Potion de défense (+)" || nom == "Potion de poison (+)" || nom == "Potion de poison (++)") {
			potionsDispo = append(potionsDispo, nom)
		}
	}

	if len(potionsDispo) == 0 {
		fmt.Println(Red + "⚠️  Vous n'avez aucune potion utilisable en combat !" + Reset)
		return false
	}

	fmt.Println(Yellow + "🧪 --- VOS POTIONS ---" + Reset)
	for i, nom := range potionsDispo {
		fmt.Printf(White+" %d. %s (x%d)\n"+Reset, i+1, nom, p.Inventaire[nom])
	}
	fmt.Println(Gray + " 0. Annuler (retour au combat)" + Reset)

	var choix int
	fmt.Print(Cyan + "👉 Choisissez une potion à consommer : " + Reset)
	fmt.Scanln(&choix)

	fmt.Println("\n" + Magenta + "============================================================" + Reset)

	if choix <= 0 || choix > len(potionsDispo) {
		fmt.Println(Gray + "↩️  Retour au menu de combat." + Reset)
		return false
	}

	nomChoisi := potionsDispo[choix-1]

	switch nomChoisi {
	case "Potion de vie (+)":
		p.PV += 30
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Println(Green + "🧪 Vous avez bu une Potion de vie (+) (+30 PV) !" + Reset)
	case "Potion de vie (++)":
		p.PV += 80 // Légèrement nerf (100 c'était beaucoup selon les PV max probables)
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Println(Green + "🧪 Vous avez bu une Potion de vie (++) (+80 PV) !" + Reset)
	case "Potion d'attaque (+)":
		p.Attaque += 10
		fmt.Println(Yellow + "🧪 Attaque augmentée de 10 pour le reste du combat !" + Reset)
	case "Potion de défense (+)":
		p.Defense += 10
		fmt.Println(Blue + "🧪 Défense augmentée de 10 pour le reste du combat !" + Reset)
	case "Potion de poison (+)":
		p.PotionPoisonPlus(m)
	case "Potion de poison (++)":
		p.PotionPoisonPlusPlus(m)
	}

	p.Inventaire[nomChoisi]--
	if p.Inventaire[nomChoisi] <= 0 {
		delete(p.Inventaire, nomChoisi)
	}

	return true
}

func (c *Combat) ExecuterAttaqueJoueur(nomAttaque string, degatsBruts, regen, defBonus int) {
	fmt.Printf(Bold+Cyan+"⚔️  [JOUEUR] %s lance [%s] !\n"+Reset, c.Joueur.Nom, nomAttaque)

	// La défense ne devrait pas bloquer 100% des dégâts pour éviter les combats infinis
	degatsNets := degatsBruts - (c.Monstre.Defense / 2) // La def absorbe une partie des dégâts
	if degatsNets < 2 {                                 // Toujours au moins 2 de dégâts si l'attaque passe
		degatsNets = 2
	}

	c.Monstre.PV -= degatsNets
	if c.Monstre.PV < 0 {
		c.Monstre.PV = 0
	}
	fmt.Printf(Red+"   💥 %s subit %d dégâts (%d absorbés) (PV restants : %d/%d).\n"+Reset, c.Monstre.Nom, degatsNets, (degatsBruts - degatsNets), c.Monstre.PV, c.Monstre.PVMax)

	if regen > 0 {
		c.Joueur.PV += regen
		if c.Joueur.PV > c.Joueur.PVMax {
			c.Joueur.PV = c.Joueur.PVMax
		}
		fmt.Printf(Green+"   💚 %s se régénère de %d PV (PV actuels : %d/%d).\n"+Reset, c.Joueur.Nom, regen, c.Joueur.PV, c.Joueur.PVMax)
	}
}
