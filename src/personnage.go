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
	ToursBonus             bool
	EmplacementsEquipement int
	Argent                 int
	Inventaire             map[string]int
	CapaciteMax            int
	Niveau                 int
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
			joueur.Attaque = 12
			joueur.Defense = 5
			// Attaques de départ
			joueur.AttaquesApprises = []attaque{MorsureVeneuse, CoupDeBouclier}
			return joueur

		case "demi-dieu":
			joueur.Classe = "demi-dieu"
			joueur.PV = 120
			joueur.PVMax = 120
			joueur.Attaque = 20
			joueur.Defense = 12
			// Attaques de départ
			joueur.AttaquesApprises = []attaque{FureurHeroique, EgideDeBronze}
			return joueur

		case "dieu":
			joueur.Classe = "dieu"
			joueur.PV = 200
			joueur.PVMax = 200
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

// Banque de nouvelles attaques qu'on peut débloquer
var PoolAttaquesDispo = []attaque{
	{Nom: "Frappe Météore", Dommage: 55, regeneration: 0, defense: 10},
	{Nom: "Siphon d'Âme", Dommage: 30, regeneration: 25, defense: 0},
	{Nom: "Bouclier Divin", Dommage: 15, regeneration: 10, defense: 40},
	{Nom: "Éclair Foudroyant", Dommage: 70, regeneration: 0, defense: 5},
	{Nom: "Soin Sacré", Dommage: 10, regeneration: 60, defense: 15},
	{Nom: "Lame d'Ombre", Dommage: 45, regeneration: 15, defense: 0},
	{Nom: "Comète Destructrice", Dommage: 110, regeneration: 0, defense: 20},
	{Nom: "Colère Divine", Dommage: 90, regeneration: 30, defense: 30},
}

// Fonction appelée quand le joueur gagne un niveau
func (p *Personnage) GagnerNiveau(scanner *bufio.Scanner) {
	p.Niveau++
	p.PVMax += 20
	p.PV = p.PVMax // Soigne au passage
	p.Attaque += 5
	p.Defense += 3

	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Printf(Bold+Yellow+"🎉 LEVEL UP ! Vous êtes maintenant niveau %d !"+Reset+"\n", p.Niveau)
	fmt.Printf(Cyan+"📈 Stats augmentées : PV Max +20 (%d) | Attaque +5 (%d) | Défense +3 (%d)"+Reset+"\n", p.PVMax, p.Attaque, p.Defense)

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
		fmt.Println(Gray + "📚 Vous maîtrisez déjà toutes les compétences disponibles !" + Reset)
		fmt.Println(Magenta + "============================================================" + Reset)
		return
	}

	// Mélange les options et prends les 3 premières
	rand.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
	if len(options) > 3 {
		options = options[:3]
	}

	fmt.Println("\n" + Blue + "✨ Choisissez une nouvelle attaque à apprendre :" + Reset)
	for i, opt := range options {
		fmt.Printf(White+" %d. %s "+Gray+"(Dégâts: %d | Soin: %d | Def: %d)"+Reset+"\n", i+1, opt.Nom, opt.Dommage, opt.regeneration, opt.defense)
	}

	var choix int
	for {
		fmt.Print(Cyan + "👉 Votre choix (1, 2 ou 3) : " + Reset)
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