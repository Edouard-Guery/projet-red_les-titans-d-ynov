package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Recette struct {
	Nom         string
	Tier        string
	Ingredients map[string]int
	Cout        int
}

func recettesForgeron() []Recette {
	return []Recette{
		// T1
		{
			Nom:         "Cuirasse Sauvage",
			Tier:        "Tier 1",
			Ingredients: map[string]int{"Dent d'Orc": 1, "Corne Brisée": 1},
			Cout:        15,
		},
		{
			Nom:         "Regard Infernal",
			Tier:        "Tier 1",
			Ingredients: map[string]int{"Œil Pétrifiant": 1, "Croc Infernal": 1},
			Cout:        15,
		},
		{
			Nom:         "Brassard de l'Enfer Vivant",
			Tier:        "Tier 1",
			Ingredients: map[string]int{"Croc Infernal": 1, "Écaille Régénérante": 1},
			Cout:        15,
		},

		// Tier 2
		{
			Nom:         "Lame du Labyrinthe",
			Tier:        "Tier 2",
			Ingredients: map[string]int{"Fil d'Ariane": 1, "Fragment de Lame": 1},
			Cout:        40,
		},
		{
			Nom:         "Manteau Empoisonné",
			Tier:        "Tier 2",
			Ingredients: map[string]int{"Fiole de Venin": 1, "Peau du Lion": 1},
			Cout:        40,
		},
		{
			Nom:         "Regard Mortel",
			Tier:        "Tier 2",
			Ingredients: map[string]int{"Regard Infernal": 1, "Fiole de Venin": 1},
			Cout:        40,
		},

		// Tier 3
		{
			Nom:         "Égide des Mers",
			Tier:        "Tier 3",
			Ingredients: map[string]int{"Éclat de Bouclier": 1, "Trident Brisé": 1},
			Cout:        80,
		},
		{
			Nom:         "Foudre Infernale",
			Tier:        "Tier 3",
			Ingredients: map[string]int{"Cendre des Enfers": 1, "Éclair Figé": 1},
			Cout:        100, // Arme surpuissante = très chère
		},
		{
			Nom:  "Armure du Colosse",
			Tier: "Tier 3",
			Ingredients: map[string]int{
				"Brassard de l'Enfer Vivant": 1,
				"Manteau Empoisonné":         1,
			},
			Cout: 120,
		},
	}
}

func (p *Personnage) nombreObjets() int {
	total := 0
	for _, quantite := range p.Inventaire {
		total += quantite
	}
	return total
}

func (p *Personnage) peutForger(recette Recette) error {
	if p.Argent < recette.Cout {
		return fmt.Errorf(
			Red+"❌ Héphaïstos secoue la tête : il demande %d Oboles, tu n'en as que %d."+Reset,
			recette.Cout,
			p.Argent,
		)
	}

	manquants := []string{}

	for ingredient, quantiteRequise := range recette.Ingredients {
		quantitePossedee := p.Inventaire[ingredient]

		if quantitePossedee < quantiteRequise {
			manquants = append(
				manquants,
				fmt.Sprintf("%dx %s", quantiteRequise-quantitePossedee, ingredient),
			)
		}
	}

	if len(manquants) > 0 {
		return fmt.Errorf(
			Red+"❌ Il te manque des matériaux : %s."+Reset,
			strings.Join(manquants, ", "),
		)
	}

	nombreIngredients := 0
	for _, quantite := range recette.Ingredients {
		nombreIngredients += quantite
	}

	nombreApresFabrication := p.nombreObjets() - nombreIngredients + 1

	if nombreApresFabrication > p.CapaciteMax {
		return fmt.Errorf(
			Red+"❌ Ton sac est trop lourd (%d/%d objets). Fais de la place d'abord."+Reset,
			p.nombreObjets(),
			p.CapaciteMax,
		)
	}

	return nil
}

func (p *Personnage) forger(recette Recette) error {
	if err := p.peutForger(recette); err != nil {
		return err
	}

	for ingredient, quantite := range recette.Ingredients {
		p.Inventaire[ingredient] -= quantite
		if p.Inventaire[ingredient] <= 0 {
			delete(p.Inventaire, ingredient)
		}
	}

	p.Argent -= recette.Cout
	p.Inventaire[recette.Nom]++

	return nil
}

func texteIngredients(ingredients map[string]int) string {
	noms := make([]string, 0, len(ingredients))

	for nom := range ingredients {
		noms = append(noms, nom)
	}

	sort.Strings(noms)

	resultat := []string{}
	for _, nom := range noms {
		resultat = append(resultat, fmt.Sprintf(Cyan+"%dx"+Reset+" %s", ingredients[nom], nom))
	}

	return strings.Join(resultat, " + ")
}

func afficherInventaire(p *Personnage) {
	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Println(Bold + Yellow + "                       🎒 VOTRE INVENTAIRE                       " + Reset)
	fmt.Println(Magenta + "============================================================" + Reset)
	fmt.Printf(Cyan+"💰 Argent : %d Oboles | 📦 Places : %d/%d\n"+Reset, p.Argent, p.nombreObjets(), p.CapaciteMax)
	fmt.Println(Blue + "------------------------------------------------------------" + Reset)

	if len(p.Inventaire) == 0 {
		fmt.Println(Gray + " Votre inventaire est vide." + Reset)
		return
	}

	noms := make([]string, 0, len(p.Inventaire))
	for nom := range p.Inventaire {
		noms = append(noms, nom)
	}
	sort.Strings(noms)

	for _, nom := range noms {
		fmt.Printf(White+" - %s "+Cyan+"x%d\n"+Reset, nom, p.Inventaire[nom])
	}
	fmt.Println(Blue + "------------------------------------------------------------" + Reset)
}

func afficherRecettesTier(recettes []Recette, tier string) {
	fmt.Printf("\n"+Bold+Yellow+"--- ⚜️  %s ⚜️  ---"+Reset+"\n", tier)

	for index, recette := range recettes {
		if recette.Tier == tier {
			fmt.Printf(
				White+" %2d. "+Green+"%s\n"+Gray+"     Matériaux : %s\n"+Gray+"     Prix : "+Yellow+"%d Oboles\n"+Reset,
				index+1,
				recette.Nom,
				texteIngredients(recette.Ingredients),
				recette.Cout,
			)
		}
	}
}

func menuForgeron(p *Personnage, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n" + Magenta + "============================================================" + Reset)
		fmt.Println(Bold + Red + "               🔥 LE VESTIBULE DE LA FORGE 🔥               " + Reset)
		fmt.Println(Magenta + "============================================================" + Reset)
		fmt.Printf(Cyan+"💰 Argent : %d Oboles\n"+Reset, p.Argent)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)
		fmt.Println(White + " 1. Parler à Héphaïstos le forgeron" + Reset)
		fmt.Println(White + " 2. Voir mon inventaire" + Reset)
		fmt.Println(Gray + " 3. Quitter la forge" + Reset)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)
		fmt.Print(Cyan + "👉 Quelle est ta volonté ? " + Reset)

		if !scanner.Scan() {
			return
		}

		fmt.Println("\n" + Magenta + "============================================================" + Reset)

		switch strings.TrimSpace(scanner.Text()) {
		case "1":
			sousMenuHephaistos(p, scanner)
		case "2":
			afficherInventaire(p)
			fmt.Print(Gray + "Prends le temps d'inspecter tes trouvailles... [Entrée]" + Reset)
			scanner.Scan()
		case "3":
			fmt.Println(Green + "👋 À bientôt, l'ami ! Que la faveur d'Héphaïstos t'accompagne !" + Reset)
			return
		default:
			fmt.Println(Red + "❌ L'Olympe ne reconnaît pas ce geste." + Reset)
		}
	}
}

func sousMenuHephaistos(p *Personnage, scanner *bufio.Scanner) {
	recettes := recettesForgeron()

	for {
		fmt.Println("\n" + Magenta + "============================================================" + Reset)
		fmt.Println(Bold + Red + "                 ⚒️  HÉPHAÏSTOS LE FORGERON ⚒️                 " + Reset)
		fmt.Println(Magenta + "============================================================" + Reset)
		fmt.Printf(Cyan+"💰 Votre argent : %d Oboles\n"+Reset, p.Argent)

		afficherRecettesTier(recettes, "Tier 1")
		afficherRecettesTier(recettes, "Tier 2")
		afficherRecettesTier(recettes, "Tier 3")

		fmt.Println(Blue + "------------------------------------------------------------" + Reset)
		fmt.Println(Gray + " 10. Retour au vestibule" + Reset)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)
		fmt.Print(Cyan + "👉 Que veux-tu forger ? : " + Reset)

		if !scanner.Scan() {
			return
		}

		fmt.Println("\n" + Magenta + "============================================================" + Reset)

		choix := strings.TrimSpace(scanner.Text())

		if choix == "10" {
			return
		}

		index, err := strconv.Atoi(choix)

		if err != nil || index < 1 || index > len(recettes) {
			fmt.Println(Red + "❌ Héphaïstos ne comprend pas ta demande..." + Reset)
			continue
		}

		recetteChoisie := recettes[index-1]
		err = p.forger(recetteChoisie)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf(Bold+Green+"🔨 INCROYABLE ! Le marteau frappe l'enclume et vous obtenez '%s' !"+Reset+"\n", recetteChoisie.Nom)
			fmt.Println(Gray + "Vos matériaux ont été fondus par Héphaïstos." + Reset)
			fmt.Printf(Cyan+"Il vous reste %d Oboles.\n"+Reset, p.Argent)
		}

		fmt.Print(Gray + "\nAppuyez sur [Entrée] pour continuer..." + Reset)
		scanner.Scan()
	}
}
