package main

import "fmt"

type attaque struct {
	Nom           string
	Dommage       int
	regeneration  int
	defense       int
	CoutEndurance int // Coût en endurance de l'attaque
}

var (
	MorsureVeneuse     = attaque{"Morsure Vénéneuse", 15, 0, 0, 8}
	CoupDeBouclier     = attaque{"Coup de Bouclier", 8, 0, 12, 5}
	VampirismeSauvage  = attaque{"Vampirisme Sauvage", 12, 6, 0, 10}
	FureurHeroique     = attaque{"Fureur Héroïque", 45, 10, 15, 20}
	EgideDeBronze      = attaque{"Égide de Bronze", 20, 15, 45, 15}
	EclairCeleste      = attaque{"Éclair Céleste", 65, 0, 10, 30}
	CataclysmeCosmique = attaque{"Cataclysme Cosmique", 150, 40, 80, 65}
	JugementAbsolu     = attaque{"Jugement Absolu", 250, 0, 100, 90}
	RenaissanceDivine  = attaque{"Renaissance Divine", 50, 150, 150, 100}
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

		// On régénère 15 d'endurance par tour
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

		// Tour du joueur
		c.TourJoueur()

		if c.Monstre.PV <= 0 {
			fmt.Println("\n" + Magenta + "============================================================" + Reset)
			fmt.Printf(Bold+Green+"🏆 VICTOIRE ! %s a vaincu %s !\n"+Reset, c.Joueur.Nom, c.Monstre.Nom)
			fmt.Println(Magenta + "============================================================" + Reset)
			break
		}

		// Tour du monstre
		c.TourMonstre()

		if c.Joueur.PV <= 0 {
			fmt.Println("\n" + Magenta + "============================================================" + Reset)
			fmt.Printf(Bold+Red+"💀 DÉFAITE... %s a été terrassé par %s.\n"+Reset, c.Joueur.Nom, c.Monstre.Nom)
			fmt.Println(Magenta + "============================================================" + Reset)
			break
		}

		c.Tour++
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
			fmt.Printf(White+" %d. %s (%d dégâts, %d soin, %d ⚡)\n"+Reset, i+offset, att.Nom, att.Dommage, att.regeneration, att.CoutEndurance)
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
				c.ExecuterAttaqueJoueur(att.Nom, att.Dommage, att.regeneration, att.defense)
				break
			} else {
				fmt.Println(Red + "❌ Pas assez d'endurance pour lancer cette compétence !" + Reset)
			}
		} else if choix == indexObjet {
			if c.Joueur.utiliserObjetdurantcombat() {
				break
			}
		} else {
			fmt.Println(Red + "❌ Choix invalide !" + Reset)
		}
	}
}

func (p *Personnage) utiliserObjetdurantcombat() bool {
	potionsDispo := []string{}
	for nom, qte := range p.Inventaire {
		if qte > 0 && (nom == "Potion de vie (+)" || nom == "Potion de vie (++)" || nom == "Potion d'attaque (+)" || nom == "Potion de défense (+)") {
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
		fmt.Println(Green + "🧪 Vous avez bu une Potion de vie (+) et regagné 30 PV !" + Reset)
	case "Potion de vie (++)":
		p.PV += 100
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Println(Green + "🧪 Vous avez bu une Potion de vie (++) et regagné 100 PV !" + Reset)
	case "Potion d'attaque (+)":
		p.Attaque += 5
		fmt.Println(Yellow + "🧪 Vous avez bu une Potion d'attaque (+). Votre attaque augmente de 5 !" + Reset)
	case "Potion de défense (+)":
		p.Defense += 5
		fmt.Println(Blue + "🧪 Vous avez bu une Potion de défense (+). Votre défense augmente de 5 !" + Reset)
	}

	p.Inventaire[nomChoisi]--
	if p.Inventaire[nomChoisi] <= 0 {
		delete(p.Inventaire, nomChoisi)
	}

	return true
}

func (c *Combat) ExecuterAttaqueJoueur(nomAttaque string, degatsBruts, regen, defBonus int) {
	fmt.Printf(Bold+Cyan+"⚔️  [JOUEUR] %s lance l'attaque [%s] !\n"+Reset, c.Joueur.Nom, nomAttaque)

	degatsNets := degatsBruts
	if degatsNets < 0 {
		degatsNets = 0
	}

	c.Monstre.PV -= degatsNets
	if c.Monstre.PV < 0 {
		c.Monstre.PV = 0
	}
	fmt.Printf(Red+"   💥 %s subit %d dégâts (PV restants : %d/%d).\n"+Reset, c.Monstre.Nom, degatsNets, c.Monstre.PV, c.Monstre.PVMax)

	if regen > 0 {
		c.Joueur.PV += regen
		if c.Joueur.PV > c.Joueur.PVMax {
			c.Joueur.PV = c.Joueur.PVMax
		}
		fmt.Printf(Green+"   💚 %s se régénère de %d PV (PV actuels : %d/%d).\n"+Reset, c.Joueur.Nom, regen, c.Joueur.PV, c.Joueur.PVMax)
	}
}