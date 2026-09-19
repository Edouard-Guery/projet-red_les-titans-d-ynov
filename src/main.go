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
		fmt.Println("\nCommandes : start, inv, magasin, forge, potions, stop, quit")
		fmt.Print("> ")
		fmt.Scan(&commande)
		commande = strings.ToLower(commande)

		switch commande {
		case "start", "setup":
			fmt.Println("\n[Combat] - Lancement...")
		case "inv", "inventaire", "inv-list":
			fmt.Printf("\nInventaire (%d/10) : %v\n", len(hero.Inventaire), hero.Inventaire)
		case "magasin", "shop":
			AfficherItem()
			ChoisirItem(&hero, &Monstre)
		case "forge", "craft":
			fmt.Println("\n[Forge] - À venir...")
		case "potions", "potion":
			// Appel du nouveau shop -- edouardo
			hero.ShopPotions()
		case "stop", "end":
			fmt.Printf("\nStats - PV: %d | Attaque: %d | Défense: %d | Oboles: %d\n", hero.PV, hero.Attaque, hero.Defense, hero.Argent)
		case "quit", "disconnect":
			os.Exit(0)
		default:
			fmt.Println("\nCommande invalide.")
		}
	}
}
