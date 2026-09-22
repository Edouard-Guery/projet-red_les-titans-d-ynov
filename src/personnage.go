package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Personnage struct {
	Nom                    string
	Classe                 string
	Attaque                int
	Defense                int
	PV                     int
	PVMax                  int
	Endurance              int
	EnduranceMax           int
	ToursBonus             bool
	EmplacementsEquipement int
	Argent                 int
	Inventaire             map[string]int
	CapaciteMax            int
	Niveau                 int
	AmeliorationsInventaire int
	AttaquesApprises       []attaque
}

func Person(nomSaisi string, classeSaisie string, scanner *bufio.Scanner) Personnage {
	joueur := Personnage{
		Nom:                    nomSaisi,
		ToursBonus:             false,
		EmplacementsEquipement: 3,
		Argent:                 1000,
		CapaciteMax:            10,
		Inventaire:             map[string]int{},
		Niveau:                 1,
	}

	classe := strings.ToLower(strings.TrimSpace(classeSaisie))
	for {
		switch classe {
		case "creature":
			joueur.Classe = "creature"
			joueur.PV = 80
			joueur.PVMax = 80
			joueur.Endurance = 50
			joueur.EnduranceMax = 50
			joueur.Attaque = 12
			joueur.Defense = 5
			// Attaques de départ
			joueur.AttaquesApprises = []attaque{MorsureVeneuse, CoupDeBouclier}
			return joueur

		case "demi-dieu":
			joueur.Classe = "demi-dieu"
			joueur.PV = 120
			joueur.PVMax = 120
			joueur.Endurance = 80
			joueur.EnduranceMax = 80
			joueur.Attaque = 20
			joueur.Defense = 12
			// Attaques de départ
			joueur.AttaquesApprises = []attaque{FureurHeroique, EgideDeBronze}
			return joueur

		case "dieu":
			joueur.Classe = "dieu"
			joueur.PV = 200
			joueur.PVMax = 200
			joueur.Endurance = 150
			joueur.EnduranceMax = 150
			joueur.Attaque = 35
			joueur.Defense = 20
			// Attaques de départ
			joueur.AttaquesApprises = []attaque{CataclysmeCosmique, JugementAbsolu}
			return joueur

		default:
			fmt.Println(Red + "❌ Classe inconnue. Choisissez une classe valide : creature, demi-dieu, dieu" + Reset)
			fmt.Print(Cyan + "👉 " + Reset)
			if scanner.Scan() {
				classe = strings.ToLower(strings.TrimSpace(scanner.Text()))
			}
		}
	}
}

// Banque de nouvelles attaques qu'on peut débloquer (avec coût en endurance)
var PoolAttaquesDispo = []attaque{
	{Nom: "Frappe Météore", Dommage: 55, regeneration: 0, defense: 10, CoutEndurance: 25},
	{Nom: "Siphon d'Âme", Dommage: 30, regeneration: 25, defense: 0, CoutEndurance: 20},
	{Nom: "Bouclier Divin", Dommage: 15, regeneration: 10, defense: 40, CoutEndurance: 15},
	{Nom: "Éclair Foudroyant", Dommage: 70, regeneration: 0, defense: 5, CoutEndurance: 30},
	{Nom: "Soin Sacré", Dommage: 10, regeneration: 60, defense: 15, CoutEndurance: 35},
	{Nom: "Lame d'Ombre", Dommage: 45, regeneration: 15, defense: 0, CoutEndurance: 20},
	{Nom: "Comète Destructrice", Dommage: 110, regeneration: 0, defense: 20, CoutEndurance: 50},
	{Nom: "Colère Divine", Dommage: 90, regeneration: 30, defense: 30, CoutEndurance: 40},
}

// Fonction appelée quand le joueur gagne un niveau
func (p *Personnage) GagnerNiveau(scanner *bufio.Scanner) {
	p.Niveau++
	p.PVMax += 20
	p.PV = p.PVMax // Soigne au passage
	p.EnduranceMax += 10
	p.Endurance = p.EnduranceMax // Restaure l'endurance
	p.Attaque += 5
	p.Defense += 3

	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Printf(Bold+Yellow+"🎉 LA GRÂCE DIVINE VOUS FRAPPE ! Vous voilà niveau %d !"+Reset+"\n", p.Niveau)
	fmt.Printf(Cyan+"📈 Stats : PV Max +20 (%d) | Endu Max +10 (%d) | Attaque +5 (%d) | Défense +3 (%d)"+Reset+"\n", p.PVMax, p.EnduranceMax, p.Attaque, p.Defense)

	// Sélection de 3 attaques inédites dans le pool
	options := []attaque{}
	for _, att := range PoolAttaquesDispo {
		dejaApprise := false
		for _, possedee := range p.AttaquesApprises {
			if possedee.Nom == att.Nom {
				dejaApprise = true
				break
			}
		}
		if !dejaApprise {
			options = append(options, att)
		}
	}

	// S'il n'y a plus d'attaques à apprendre
	if len(options) == 0 {
		fmt.Println(Gray + "📚 Les dieux n'ont plus rien à t'enseigner !" + Reset)
		fmt.Println(Magenta + "============================================================" + Reset)
		return
	}

	// Mélange les options et prends les 3 premières
	rand.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
	if len(options) > 3 {
		options = options[:3]
	}

	fmt.Println("\n" + Blue + "✨ Invoque un nouvel art du combat auprès du maître :" + Reset)
	for i, opt := range options {
		fmt.Printf(White+" %d. %s "+Gray+"(Dégâts: %d | Soin: %d | Def: %d | Endu: %d)"+Reset+"\n", i+1, opt.Nom, opt.Dommage, opt.regeneration, opt.defense, opt.CoutEndurance)
	}

	var choix int
	for {
		fmt.Print(Cyan + "👉 Quelle voie empruntes-tu ? (1, 2 ou 3) : " + Reset)
		if scanner.Scan() {
			saisie := strings.TrimSpace(scanner.Text())
			val, err := strconv.Atoi(saisie)

			if err == nil && val >= 1 && val <= len(options) {
				choix = val
				attaqueChoisie := options[choix-1]
				p.AttaquesApprises = append(p.AttaquesApprises, attaqueChoisie)
				fmt.Printf("\n"+Green+"📖 Vous avez appris [%s] !"+Reset+"\n", attaqueChoisie.Nom)
				break
			}
		}
		fmt.Println(Red + "❌ Choix invalide. Veuillez entrer un numéro valide." + Reset)
	}
	fmt.Println(Magenta + "============================================================" + Reset)
}
