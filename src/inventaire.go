package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Menu interactif pour voir et utiliser les objets de l'inventaire
func (p *Personnage) MenuInventaire(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n=== VOTRE INVENTAIRE ===")
		fmt.Printf("Oboles : %d | Places : %d/%d\n", p.Argent, p.nombreObjets(), p.CapaciteMax)

		if len(p.Inventaire) == 0 {
			fmt.Println("Votre inventaire est vide.")
			return
		}

		// Trie les objets par ordre alphabétique
		noms := make([]string, 0, len(p.Inventaire))
		for nom := range p.Inventaire {
			noms = append(noms, nom)
		}
		sort.Strings(noms)

		fmt.Println("Objets disponibles :")
		for i, nom := range noms {
			fmt.Printf("%d. %s (x%d)\n", i+1, nom, p.Inventaire[nom])
		}

		fmt.Println("\n0. Fermer l'inventaire")
		fmt.Print("> Entrez le numéro de l'objet à utiliser : ")

		if !scanner.Scan() {
			return
		}
		choix := strings.TrimSpace(scanner.Text())

		if choix == "0" {
			return
		}

		index, err := strconv.Atoi(choix)
		if err != nil || index < 1 || index > len(noms) {
			fmt.Println("Choix invalide.")
			continue
		}

		objetChoisi := noms[index-1]
		p.utiliserObjet(objetChoisi)
	}
}

// Applique l'effet de l'objet et le retire de l'inventaire
func (p *Personnage) utiliserObjet(nom string) {
	switch nom {
	case "Potion de vie (+)":
		p.PotionViePlus() // Fonction située dans effets.go
		fmt.Println("Vous avez bu une Potion de vie (+) et regagné 30 PV !")
	case "Potion de vie (++)":
		p.PotionViePlusPlus()
		fmt.Println("Vous avez bu une Potion de vie (++) et regagné 100 PV !")
	case "Potion d'attaque (+)":
		p.PotionAttaquePlus()
		fmt.Println("Vous avez bu une Potion d'attaque (+). Votre attaque augmente !")
	case "Potion de défense (+)":
		p.PotionDefensePlus()
		fmt.Println("Vous avez bu une Potion de défense (+). Votre défense augmente !")
	default:
		// Si l'objet n'est pas utilisable (comme les matériaux de forge)
		fmt.Printf("Vous ne pouvez pas utiliser '%s' directement ici.\n", nom)
		return
	}

	// Retire 1 objet du sac après utilisation
	p.Inventaire[nom]--
	if p.Inventaire[nom] <= 0 {
		delete(p.Inventaire, nom)
	}
}

func (p *Personnage) UpgradeInventorySlot() {
	if p.CapaciteMax >= 40 {
		fmt.Println("Votre inventaire est déjà à sa taille maximale.")
		return
	}
	p.CapaciteMax += 10
	fmt.Printf("Succès ! Votre inventaire a été agrandi. Capacité actuelle : %d places.\n", p.CapaciteMax)
}
