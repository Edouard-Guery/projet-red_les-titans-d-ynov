package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	hero := Person()
	Monstre := Person() // Initialisé pour pouvoir être passé à la boutique
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
			hero.MenuInventaire(scanner) // Utilise ton menu interactif au lieu du simple affichage
		case "magasin", "shop":
			AfficherItem()
			// On passe Monstre et le scanner pour corriger le bug
			ChoisirItem(&hero, &Monstre, scanner)
		case "forge", "craft":
			menuForgeron(&hero, scanner)
		case "potions", "potion":
			hero.ShopPotions()
		case "stop", "end", "info":
			fmt.Printf("\nStats - PV: %d | Attaque: %d | Défense: %d | Oboles: %d\n", hero.PV, hero.Attaque, hero.Defense, hero.Argent)
		case "quit", "disconnect":
			os.Exit(0)
		case "casino":
			casino(&hero)
		default:
			fmt.Println("\nCommande invalide.")
		}
	}
}