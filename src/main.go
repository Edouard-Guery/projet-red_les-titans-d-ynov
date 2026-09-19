package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	hero := Person()
	Monstre := Person()
	var commande string

	for {
		fmt.Println("\nCommandes : start, inventaire, L'Échoppe de l'Olympe, forge de Héphaïstos, les potions de Dionysos, The Midas Club, info, stop, quit")
		fmt.Print("> ")
		fmt.Scan(&commande)
		commande = strings.ToLower(commande)

		switch commande {
		case "start", "setup":
			fmt.Println("\n[Combat] - Lancement...")
		case "inv", "inventaire", "inv-list":
			fmt.Printf("\nInventaire (%d/10) : %v\n", len(hero.Inventaire), hero.Inventaire)
		case "magasin", "shop","L'Échoppe de l'Olympe":
			AfficherItem()
			ChoisirItem(&hero, &Monstre)
		case "forge", "craft","forge de Héphaïstos":
			fmt.Println("\n[Forge] - À venir...")
		case "potions", "potion","les potions de Dionysos":
			// Appel du nouveau shop -- edouardo
			hero.ShopPotions()
		case "stop", "end":
			fmt.Printf("\nStats - PV: %d | Attaque: %d | Défense: %d | Oboles: %d\n", hero.PV, hero.Attaque, hero.Defense, hero.Argent)
		case "quit", "disconnect":
			os.Exit(0)
		case "casino", "The Midas Club":
			casino(&hero)
		case "info":
			fmt.Printf("\nStats - PV: %d | Attaque: %d | Défense: %d | Oboles: %d\n", hero.PV, hero.Attaque, hero.Defense, hero.Argent)
		default:
			fmt.Println("\nCommande invalide.")
		}
	}
}
