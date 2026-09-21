package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ChoisirItem(hero *Personnage, monstre *Personnage, scanner *bufio.Scanner) {
		fmt.Print("Entre un chiffre : ")

	// Utilisation du scanner pour éviter le bug du saut de ligne
	if !scanner.Scan() {
		return
	}
	choixStr := strings.TrimSpace(scanner.Text())
	nombre, err := strconv.Atoi(choixStr)

	if err != nil {
		fmt.Println("Choix invalide !")
		return
	}

	if nombre == 15 {
		return
	}

	items := BoutiqueMarchand()
	if nombre < 1 || nombre > len(items) {
		fmt.Println("Choix invalide !")
		return
	}

	itemChoisi := items[nombre-1]

	// 1. Vérification de la place dans le sac à dos (Limite de 10)
	if hero.nombreObjets() >= hero.CapaciteMax {
		fmt.Println("Erreur : Votre inventaire est plein.")
		return
	}

	// 2. Vérification de la monnaie
	if hero.Argent < itemChoisi.Prix {
		fmt.Printf("Erreur : Il vous manque %d Oboles.\n", itemChoisi.Prix-hero.Argent)
		return
	}

	// 3. Achat : Déduction des Oboles et ajout dans la map de l'inventaire
	hero.Argent -= itemChoisi.Prix
	hero.Inventaire[itemChoisi.Nom]++
	fmt.Printf("Acheté : %s (-%d Oboles). Ajouté à l'inventaire !\n", itemChoisi.Nom, itemChoisi.Prix)

	// 4. Code original conservé : application de l'effet magique
	switch nombre {
	case 1:
		hero.BouclierEclair()
	case 2:
		hero.PendentifVenin()
	case 3:
		hero.FauxEternite()
	case 4:
		hero.BoitePandore(monstre)
	case 5:
		hero.CasqueSagesse()
	case 6:
		hero.ArmureHephaistos()
	case 7:
		hero.CarteArchipel()
	case 8:
		hero.FrenesieTravaux()
	case 9:
		hero.SacMagique()
	case 10:
		hero.RegardGorgone(monstre)
	case 11:
		hero.GueuleEnflammee(monstre)
	case 12:
		hero.HacheDoubleTranchant(monstre)
	case 13:
		hero.GourdeRegeneration()
	case 14:
		hero.BouclierBoisRenforce()
	}
}

func AfficherItem() {
	items := BoutiqueMarchand()
	fmt.Println("=== BOUTIQUE DES ÉQUIPEMENTS MYTHOLOGIQUES ===")
	for i, eq := range items {
		fmt.Printf("%d. %s (%s) — %d Oboles\n   └─ %s\n\n",
			i+1, eq.Nom, eq.Source, eq.Prix, eq.Description)
	}
	fmt.Println("15. annuler")
}