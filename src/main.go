package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Variables pour la création de personnage
	nomHero := "Héros"
	classeHero := "creature"

	// 1. Saisie du nom
	fmt.Print("Entrez le nom de votre héros : ")
	if scanner.Scan() {
		nomSaisi := strings.TrimSpace(scanner.Text())
		if nomSaisi != "" {
			nomHero = nomSaisi
		}
	}

	// 2. Saisie de la classe
	fmt.Print("Entrez la classe de votre héros (creature / demi-dieu / dieu) : ")
	if scanner.Scan() {
		classeSaisie := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if classeSaisie != "" {
			classeHero = classeSaisie
		}
	}

	// Initialisation du héros
	hero := Person(nomHero, classeHero)
	fmt.Printf("\nBienvenue, %s le %s ! Que votre quête commence.\n", hero.Nom, hero.Classe)

	index := 0

	for {
		fmt.Println("\nCommandes : start, inventaire, shop, forge, potions, casino, info, quit")
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}
		commande := strings.TrimSpace(strings.ToLower(scanner.Text()))

		switch commande {
		case "inv", "inventaire", "inv-list":
			hero.MenuInventaire(scanner)

		case "start", "fight", "tour de l'olympe", "combat":
			if index >= len(ListeMonstres) {
				fmt.Println("\nFélicitations ! Vous avez déjà vaincu tous les monstres !")
				break
			}

			monstre := ObtenirMonstre(index)
			fmt.Printf("\n--- COMBAT %d/%d : Un %s apparaît ! ---\n", index+1, len(ListeMonstres), monstre.Nom)

			combat := NouveauCombatMonstre(&hero, &monstre)
			combat.LancerDéroulement()

			// Progression si le héros survit
			if hero.PV > 0 {
				index++
				fmt.Printf("\nVictoire ! Prochain monstre débloqué.\n")
			} else {
				fmt.Println("\nGame Over... Reposez-vous avant de tenter un nouveau combat.")
			}

		case "magasin", "shop":
			AfficherItem()
			monstreActuel := Personnage{Nom: "Cible", PV: 100, Attaque: 10, Defense: 10}
			ChoisirItem(&hero, &monstreActuel, scanner)

		case "forge", "craft":
			menuForgeron(&hero, scanner)

		case "potions", "potion":
			hero.ShopPotions()

		case "stop", "end", "info":
			fmt.Printf("\nStats - Nom: %s | Classe: %s | PV: %d | Attaque: %d | Défense: %d | Oboles: %d\n",
				hero.Nom, hero.Classe, hero.PV, hero.Attaque, hero.Defense, hero.Argent)

		case "quit", "disconnect":
			os.Exit(0)

		case "casino":
			casino(&hero)

		default:
			fmt.Println("\nCommande invalide.")
		}
	}
}

func NouveauCombatMonstre(personnage *Personnage, monstre *Monstre) *Combat {
	return &Combat{
		Joueur:  personnage,
		Monstre: monstre,
		Tour:    1,
	}
}
