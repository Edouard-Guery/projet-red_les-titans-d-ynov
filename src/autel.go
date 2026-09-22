package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func autel(joueur *Personnage, scanner *bufio.Scanner) {
	fmt.Println("\n=== AUTEL DE ZEUS ===")
	fmt.Printf("PV actuels : %d/%d | Oboles : %d\n", joueur.PV, joueur.PVMax, joueur.Argent)
	fmt.Println("Combien de PV voulez-vous sacrifier ? (1 PV = 2 Oboles)")
	fmt.Print("> ")

	if !scanner.Scan() {
		return
	}

	entree := strings.TrimSpace(scanner.Text())
	sacrifice, err := strconv.Atoi(entree)

	// Empêche les entrées invalides ou les nombres négatifs
	if err != nil || sacrifice <= 0 {
		fmt.Println("Offrande invalide. Vous devez sacrifier au moins 1 PV.")
		return
	}

	// Vérifie si le sacrifice tue le personnage
	if sacrifice >= joueur.PV {
		joueur.PV = 0
		fmt.Println("\nVous avez sacrifié toute votre force vitale... Zeus accepte votre âme. Game Over.")
		return
	}

	// Application du sacrifice
	joueur.PV -= sacrifice
	obolesGagnees := sacrifice * 2
	joueur.Argent += obolesGagnees

	fmt.Printf("\nZeus accepte votre sang ! Vous sacrifiez %d PV et gagnez %d Oboles.\n", sacrifice, obolesGagnees)
	fmt.Printf("PV restants : %d | Oboles : %d\n", joueur.PV, joueur.Argent)
}