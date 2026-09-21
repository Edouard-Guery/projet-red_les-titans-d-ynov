package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	hero := Person()
	Monstre := Person()

	// Utilisation du scanner pour tout le jeu (évite les bugs avec fmt.Scan)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nCommandes : start, inventaire, shop, forge, potions, casino, info, stop, quit")
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}
		commande := strings.TrimSpace(strings.ToLower(scanner.Text()))

		switch commande {
		case "start", "setup":
			fmt.Println("\n[Combat] - Lancement...")
		case "inv", "inventaire", "inv-list":
			afficherInventaire(&hero) // Fonction du forgeron
		case "magasin", "shop":
			AfficherItem()
			ChoisirItem(&hero, &Monstre)
		case "forge", "craft":
			menuForgeron(&hero, scanner) // Appel de la forge de Clovis
		case "potions", "potion":
			hero.ShopPotions()
		case "stop", "end", "info":
			fmt.Printf("\nStats - PV: %d | Attaque: %d | Défense: %d | Oboles: %d\n", hero.PV, hero.Attaque, hero.Defense, hero.Argent)
		case "quit", "disconnect":
			os.Exit(0)
		case "casino":
			// casino(&hero)
			fmt.Println("\n[Casino] - À venir...")
		default:
			fmt.Println("\nCommande invalide.")
		}
	}
}
