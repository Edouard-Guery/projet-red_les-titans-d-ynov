// CLOVIS
package main

import (
	"bufio" //permet de lire ce que le joueur tape dans le terminal.
	"fmt"
	"sort"    //trie les objets par ordre alphabétique dans le sac à dos et les ingrédients dans les recettes
	"strconv" //transforme le texte tapé par le joueurpar exemple "2= en nombre 2.
	"strings" //enlève les espaces inutiles avec TrimSpace et rassemble une liste de textes avec Join.
)

type Recette struct {
	Nom         string
	Tier        string
	Ingredients map[string]int //map sert ici à mémoriser tous les objets du joueur et combien il en possède façon de stocker une liste avec une étiquette et une valeur
	Cout        int
}

// NOTE : La structure Joueur a été retirée car elle fait doublon avec Personnage de ton fichier personnage.go

func recettesForgeron() []Recette { //renvoie une liste de recettes. []Recette veut dire tableau
	return []Recette{
		// Tier 1
		{
			Nom:         "Cuirasse Sauvage",
			Tier:        "Tier 1",
			Ingredients: map[string]int{"Dent d'Orc": 1, "Corne Brisée": 1},
			Cout:        5, //crée une recette
			//les tiers servent à organiser l’affichage.
		},
		{
			Nom:         "Regard Infernal",
			Tier:        "Tier 1",
			Ingredients: map[string]int{"Œil Pétrifiant": 1, "Croc Infernal": 1},
			Cout:        5,
		},
		{
			Nom:         "Brassard de l'Enfer Vivant",
			Tier:        "Tier 1",
			Ingredients: map[string]int{"Croc Infernal": 1, "Écaille Régénérante": 1},
			Cout:        5,
		},

		// Tier 2
		{
			Nom:         "Lame du Labyrinthe",
			Tier:        "Tier 2",
			Ingredients: map[string]int{"Fil d'Ariane": 1, "Fragment de Lame": 1},
			Cout:        5,
		},
		{
			Nom:         "Manteau Empoisonné",
			Tier:        "Tier 2",
			Ingredients: map[string]int{"Fiole de Venin": 1, "Peau du Lion": 1},
			Cout:        5,
		},
		{
			Nom:         "Regard Mortel",
			Tier:        "Tier 2",
			Ingredients: map[string]int{"Regard Infernal": 1, "Fiole de Venin": 1},
			Cout:        5,
		},

		// Tier 3
		{
			Nom:         "Égide des Mers",
			Tier:        "Tier 3",
			Ingredients: map[string]int{"Éclat de Bouclier": 1, "Trident Brisé": 1},
			Cout:        5,
		},
		{
			Nom:         "Foudre Infernale",
			Tier:        "Tier 3",
			Ingredients: map[string]int{"Cendre des Enfers": 1, "Éclair Figé": 1},
			Cout:        5,
		},
		{
			Nom:  "Armure du Colosse",
			Tier: "Tier 3",
			Ingredients: map[string]int{
				"Brassard de l'Enfer Vivant": 1,
				"Manteau Empoisonné":         1,
			},
			Cout: 5,
		},
	}
}

func (p *Personnage) nombreObjets() int { //p *Personnage signifie que la fonction travaille sur un joueur précis. Le * est un pointeur : la fonction accède au vrai joueur, pas à une copie.
	total := 0

	for _, quantite := range p.Inventaire {
		total += quantite
	}

	return total
}

func (p *Personnage) peutForger(recette Recette) error {
	if p.Argent < recette.Cout { //si le joueur a moins d’Oboles que le prix de la recette, la forge est refusée
		return fmt.Errorf(
			"Vous n'avez pas assez d'Oboles : il faut %d Oboles, mais vous en avez %d.",
			recette.Cout,
			p.Argent,
		)
	}

	manquants := []string{} //crée une liste vide qui contiendra les ressources manquantes

	for ingredient, quantiteRequise := range recette.Ingredients { //Parcourt chaque matériau de la recette
		quantitePossedee := p.Inventaire[ingredient] //Récupère la quantité possédée dans le sac

		if quantitePossedee < quantiteRequise { //si le joueur n’a pas assez de cet ingrédient = il manque quelque chose
			manquants = append(
				manquants,
				fmt.Sprintf("%dx %s", quantiteRequise-quantitePossedee, ingredient), //ajoute un texte à la liste
			)
		}
	}

	if len(manquants) > 0 {
		return fmt.Errorf(
			"Vous n'avez pas les ressources nécessaires : %s.",
			strings.Join(manquants, ", "), //join = il rassemble plusieurs textes en un seul, séparés par "","" pour eviter de repeter toujours les memes choses
		)
	}

	nombreIngredients := 0 //compte combien d’objets seront retirés.
	for _, quantite := range recette.Ingredients {
		nombreIngredients += quantite
	}

	nombreApresFabrication := p.nombreObjets() - nombreIngredients + 1 //Calcule le contenu du sac après la fabrication

	if nombreApresFabrication > p.CapaciteMax { //Si le total prévu dépasse la capacité, la forge est refusée.
		return fmt.Errorf(
			"Votre sac a dos est plein (%d/%d objets).",
			p.nombreObjets(),
			p.CapaciteMax,
		)
	}

	return nil
}

func (p *Personnage) forger(recette Recette) error {
	if err := p.peutForger(recette); err != nil { //err := crée une variable d’erreur.
		//Si err != nil, il y a une erreur.
		//return err arrête la fonction sans modifier le joueur.
		return err
	}

	for ingredient, quantite := range recette.Ingredients { //Pour chaque ingrédient, la quantité requise est retirée du sac.
		p.Inventaire[ingredient] -= quantite

		if p.Inventaire[ingredient] <= 0 { // Si l’objet arrive à zéro, il est supprimé de la map.
			delete(p.Inventaire, ingredient)
		}
	}

	p.Argent -= recette.Cout    // Retire les x Oboles.
	p.Inventaire[recette.Nom]++ //Ajoute l’objet forgé

	return nil //La forge a réussi.
}

func texteIngredients(ingredients map[string]int) string { //fonction transforme les ingrédients d’une recette en phrase lisible
	noms := make([]string, 0, len(ingredients)) //CCrée une liste de noms vide, avec une capacité prévue égale au nombre d’ingrédien

	for nom := range ingredients { //Récupère tous les noms d’ingrédients
		noms = append(noms, nom)
	}

	sort.Strings(noms) //Les trie par ordre alphabétique.

	resultat := []string{} //Crée une liste de textes.
	for _, nom := range noms {
		resultat = append(resultat, fmt.Sprintf("%dx %s", ingredients[nom], nom)) //Transforme les noms en texte
	}

	return strings.Join(resultat, " + ")
}

func afficherInventaire(p *Personnage) { //Affiche l’or, la capacité et le contenu du sac.
	fmt.Println("\n--- Sac a dos ---")
	fmt.Printf("Argent : %d Oboles\n", p.Argent)
	fmt.Printf("Places : %d/%d\n", p.nombreObjets(), p.CapaciteMax)

	if len(p.Inventaire) == 0 { //Si la map est vide, le sac est vid
		fmt.Println("Sac a dos vide.")
		return
	}

	noms := make([]string, 0, len(p.Inventaire)) //Crée une liste pour récupérer les noms des objets.

	for nom := range p.Inventaire { //Trie puis affiche
		noms = append(noms, nom)
	}

	sort.Strings(noms)

	for _, nom := range noms {
		fmt.Printf("- %s x%d\n", nom, p.Inventaire[nom])
	}
}

func afficherRecettesTier(recettes []Recette, tier string) { //Reçoit la liste entière des recettes et le tier demandé.
	fmt.Printf("\n--- %s ---\n", tier)

	for index, recette := range recettes {
		if recette.Tier == tier { //N’affiche que les recettes du tier concerné.
			fmt.Printf(
				"%d. %s\n   Matériaux : %s\n   Prix : %d Oboles\n",
				index+1, //Sert a afficher le choix du joueur
				recette.Nom,
				texteIngredients(recette.Ingredients),
				recette.Cout,
			)
		}
	}
}

func menuForgeron(p *Personnage, scanner *bufio.Scanner) { //Cette fonction affiche le menu de forge et attend les choix du joueur.
	recettes := recettesForgeron() // Charge les 9 recettes

	for {
		fmt.Println("\n========== HEPHAISTOS LE FORGERON ==========")
		fmt.Printf("Votre argent : %d Oboles\n", p.Argent)

		afficherRecettesTier(recettes, "Tier 1")
		afficherRecettesTier(recettes, "Tier 2")
		afficherRecettesTier(recettes, "Tier 3")

		fmt.Println("\n0. Retour au menu principal")
		fmt.Print("> ")

		if !scanner.Scan() { //scanner.Scan() lit une ligne.
			//Le ! veut dire « non ».
			//Si la lecture échoue ou que l’entrée est fermée, la fonction s’arrête.
			return
		}

		choix := strings.TrimSpace(scanner.Text()) //scanner.Text() récupère la saisie.
		//TrimSpace retire les espaces et le retour à la ligne.

		if choix == "0" { //Le joueur quitte le menu forgeron.
			return
		}

		index, err := strconv.Atoi(choix)

		if err != nil || index < 1 || index > len(recettes) { //Refuse les choix non numériques, inférieurs à 1 ou supérieurs au nombre de recettes.
			fmt.Println("L'Olympe ne reconnaît pas ce geste..")
			continue
		}

		recetteChoisie := recettes[index-1]

		err = p.forger(recetteChoisie) //tente la forge

		if err != nil {
			fmt.Println("\nLes offrandes à la forge sont insuffisantes :", err)
		} else {
			fmt.Printf("\nVous avez forgé '%s' !\n", recetteChoisie.Nom)
			fmt.Println("Vos objets ont été remis à Hephaistos le forgeron et retirés de votre sac a dos..")
			fmt.Printf("Il vous reste %d d'Oboles.\n", p.Argent)
		}

		fmt.Println("\nFranchis le seuil pour poursuivre ton destin... [Entrée]")
		scanner.Scan() //met le programme en pause
	}
}
