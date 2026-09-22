package main

import (
	"bufio"
	"fmt"
	"math/rand/v2" // Mise à jour sur rand/v2 comme le reste du projet
	"strconv"
	"strings"
)

type Personnage struct {
	Nom                     string
	Classe                  string
	Attaque                 int
	Defense                 int
	PV                      int
	PVMax                   int
	Endurance               int
	EnduranceMax            int
	ToursBonus              bool
	EmplacementsEquipement  int
	Argent                  int
	Inventaire              map[string]int
	CapaciteMax             int
	Niveau                  int
	AmeliorationsInventaire int
	FaveurDieux             map[string]int
	Benedictions            map[string]bool
	AttaquesApprises        []attaque
}

func Person(nomSaisi string, classeSaisie string, scanner *bufio.Scanner) Personnage {
	joueur := Personnage{
		Nom:                    nomSaisi,
		ToursBonus:             false,
		EmplacementsEquipement: 3,
		Argent:                 20, // Économie fixée : on commence avec un peu de monnaie (20 oboles au lieu de 1000) pour acheter une potion de base.
		CapaciteMax:            10,
		FaveurDieux:            map[string]int{},
		Benedictions:           map[string]bool{},
		Inventaire:             map[string]int{},
		Niveau:                 1,
	}

	classe := strings.ToLower(strings.TrimSpace(classeSaisie))
	for {
		switch classe {
		// La Créature : Le personnage équilibré (Standard)
		case "creature":
			joueur.Classe = "Créature"
			joueur.PV = 80
			joueur.PVMax = 80
			joueur.Endurance = 50
			joueur.EnduranceMax = 50
			joueur.Attaque = 12
			joueur.Defense = 5
			joueur.AttaquesApprises = []attaque{MorsureVeneuse, CoupDeBouclier}
			return joueur

		// Le Demi-Dieu : Axé sur l'Attaque et l'Endurance (Glass Cannon)
		case "demi-dieu":
			joueur.Classe = "Demi-Dieu"
			joueur.PV = 70 // Moins de PV de base, mais frappe plus fort
			joueur.PVMax = 70
			joueur.Endurance = 70
			joueur.EnduranceMax = 70
			joueur.Attaque = 16
			joueur.Defense = 4
			joueur.AttaquesApprises = []attaque{FureurHeroique, EgideDeBronze} // Raccord avec le nerf des compétences
			return joueur

		// Le Dieu : Axé sur la Survie et la Défense (Tank)
		case "dieu":
			joueur.Classe = "Dieu"
			joueur.PV = 110 // Très tanky
			joueur.PVMax = 110
			joueur.Endurance = 40 // Mais peu d'endurance
			joueur.EnduranceMax = 40
			joueur.Attaque = 8 // Frappe moins fort de base
			joueur.Defense = 8
			// Retrait des attaques ultimes (CataclysmeCosmique) en début de jeu.
			joueur.AttaquesApprises = []attaque{EclairCeleste, RenaissanceDivine}
			return joueur

		default:
			fmt.Println(Red + "❌ Classe inconnue. Choisissez : creature, demi-dieu, dieu" + Reset)
			fmt.Print(Cyan + "👉 " + Reset)
			if scanner.Scan() {
				classe = strings.ToLower(strings.TrimSpace(scanner.Text()))
			}
		}
	}
}

// Les dégâts de ces compétences ont été lissés pour correspondre au nouveau système
var PoolAttaquesDispo = []attaque{
	{Nom: "Frappe Météore", Dommage: 35, regeneration: 0, defense: 5, CoutEndurance: 25},
	{Nom: "Siphon d'Âme", Dommage: 20, regeneration: 15, defense: 0, CoutEndurance: 20},
	{Nom: "Bouclier Divin", Dommage: 10, regeneration: 10, defense: 25, CoutEndurance: 15},
	{Nom: "Éclair Foudroyant", Dommage: 50, regeneration: 0, defense: 5, CoutEndurance: 30},
	{Nom: "Soin Sacré", Dommage: 5, regeneration: 40, defense: 10, CoutEndurance: 35},
	{Nom: "Lame d'Ombre", Dommage: 25, regeneration: 10, defense: 0, CoutEndurance: 20},
	{Nom: "Comète Destructrice", Dommage: 75, regeneration: 0, defense: 10, CoutEndurance: 45},
	{Nom: "Colère Divine", Dommage: 60, regeneration: 20, defense: 20, CoutEndurance: 40},
}

func (p *Personnage) GagnerNiveau(scanner *bufio.Scanner) {
	p.Niveau++
	p.PVMax += 15 // +20 c'était beaucoup tous les 2 combats
	p.PV = p.PVMax
	p.EnduranceMax += 5
	p.Endurance = p.EnduranceMax
	p.Attaque += 3
	p.Defense += 2

	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Printf(Bold+Yellow+"🎉 LA GRÂCE DIVINE VOUS FRAPPE ! Vous voilà niveau %d !"+Reset+"\n", p.Niveau)
	fmt.Printf(Cyan+"📈 Stats : PV Max +15 (%d) | Endu Max +5 (%d) | Attaque +3 (%d) | Défense +2 (%d)"+Reset+"\n", p.PVMax, p.EnduranceMax, p.Attaque, p.Defense)

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

	if len(options) == 0 {
		fmt.Println(Gray + "📚 Les dieux n'ont plus rien à t'enseigner !" + Reset)
		fmt.Println(Magenta + "============================================================" + Reset)
		return
	}

	// Utilisation de math/rand/v2
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
