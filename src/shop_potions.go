package main

import "fmt"

func (p *Personnage) ShopPotions() {
	var choix int

	for {
		fmt.Println("\n=== MAGASIN DE POTIONS (DIONYSOS) ===")
		fmt.Printf("Oboles : %d | Inventaire : %d/10\n\n", p.Argent, len(p.Inventaire))

		fmt.Println("1. Potion de vie (+)           - 5 Oboles")
		fmt.Println("2. Potion de vie (++)          - 15 Oboles")
		fmt.Println("3. Potion d'attaque (+)        - 20 Oboles")
		fmt.Println("4. Potion d'attaque (++)       - 35 Oboles")
		fmt.Println("5. Potion baisse d'attaque (-) - 25 Oboles")
		fmt.Println("6. Potion de défense (+)       - 20 Oboles")
		fmt.Println("7. Potion de défense (++)      - 35 Oboles")
		fmt.Println("8. Potion baisse défense (-)   - 25 Oboles")
		fmt.Println("9. Potion de poison (+)        - 15 Oboles")
		fmt.Println("10. Potion de poison (++)      - 30 Oboles")
		fmt.Println("11. Quitter")

		fmt.Print("\nChoix : ")
		fmt.Scan(&choix)

		if choix == 11 {
			break
		}

		nomPotion := ""
		prix := 0

		switch choix {
		case 1:
			nomPotion, prix = "Potion de vie (+)", 5
		case 2:
			nomPotion, prix = "Potion de vie (++)", 15
		case 3:
			nomPotion, prix = "Potion d'attaque (+)", 20
		case 4:
			nomPotion, prix = "Potion d'attaque (++)", 35
		case 5:
			nomPotion, prix = "Potion baisse d'attaque (-)", 25
		case 6:
			nomPotion, prix = "Potion de défense (+)", 20
		case 7:
			nomPotion, prix = "Potion de défense (++)", 35
		case 8:
			nomPotion, prix = "Potion baisse défense (-)", 25
		case 9:
			nomPotion, prix = "Potion de poison (+)", 15
		case 10:
			nomPotion, prix = "Potion de poison (++)", 30
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		p.acheterPotion(nomPotion, prix)
	}
}

func (p *Personnage) acheterPotion(nom string, prix int) {
	if len(p.Inventaire) >= 10 {
		fmt.Println("Erreur : Votre inventaire est plein (10 objets max).")
		return
	}
	if p.Argent >= prix {
		p.Argent -= prix
		p.Inventaire = append(p.Inventaire, nom)
		fmt.Printf("Acheté : %s (-%d Oboles)\n", nom, prix)
	} else {
		fmt.Println("Erreur : Vous n'avez pas assez d'Oboles.")
	}
}
