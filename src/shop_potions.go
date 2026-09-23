package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func (p *Personnage) ShopPotions(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n" + Magenta + "============================================================" + Reset)
		fmt.Println(Bold + Green + "                     🍷 MAGASIN DE POTIONS DE DIONYSOS 🍷                      " + Reset)
		fmt.Println(Magenta + "============================================================" + Reset)
		fmt.Printf(Cyan+"💰 Oboles : %d | 🎒 Inventaire : %d/%d\n"+Reset, p.Argent, p.nombreObjets(), p.CapaciteMax)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)

		fmt.Println(White + "  1. Potion de vie (+)           - " + Yellow + "10 Oboles" + Reset)
		fmt.Println(White + "  2. Potion de vie (++)          - " + Yellow + "30 Oboles" + Reset)
		fmt.Println(White + "  3. Potion d'attaque (+)        - " + Yellow + "40 Oboles" + Reset)
		fmt.Println(White + "  4. Potion d'attaque (++)       - " + Yellow + "80 Oboles" + Reset)
		fmt.Println(White + "  5. Potion baisse d'attaque (-) - " + Yellow + "40 Oboles" + Reset)
		fmt.Println(White + "  6. Potion de défense (+)       - " + Yellow + "40 Oboles" + Reset)
		fmt.Println(White + "  7. Potion de défense (++)      - " + Yellow + "80 Oboles" + Reset)
		fmt.Println(White + "  8. Potion baisse défense (-)   - " + Yellow + "40 Oboles" + Reset)
		fmt.Println(White + "  9. Potion de poison (+)        - " + Yellow + "30 Oboles" + Reset)
		fmt.Println(White + " 10. Potion de poison (++)       - " + Yellow + "60 Oboles" + Reset)
		fmt.Println(White + " 11. Potion d'endurance (+)      - " + Yellow + "10 Oboles" + Reset)
		fmt.Println(White + " 12. Potion d'endurance (++)     - " + Yellow + "30 Oboles" + Reset)
		fmt.Println(Gray + " 13. Quitter la boutique" + Reset)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)

		fmt.Print(Cyan + "👉 Ton verdict : " + Reset)
		if !scanner.Scan() {
			break
		}

		fmt.Println("\n" + Magenta + "============================================================" + Reset)

		entree := strings.TrimSpace(scanner.Text())
		choix, err := strconv.Atoi(entree)
		if err != nil {
			fmt.Println(Red + "❌ Inscris un chiffre béni des dieux." + Reset)
			continue
		}

		if choix == 13 {
			fmt.Println(Gray + "👋 Vous quittez la boutique de Dionysos." + Reset)
			break
		}

		var nomPotion string
		var prix int

		switch choix {
		case 1:
			nomPotion, prix = "Potion de vie (+)", 10
		case 2:
			nomPotion, prix = "Potion de vie (++)", 30
		case 3:
			nomPotion, prix = "Potion d'attaque (+)", 40
		case 4:
			nomPotion, prix = "Potion d'attaque (++)", 80
		case 5:
			nomPotion, prix = "Potion baisse d'attaque (-)", 40
		case 6:
			nomPotion, prix = "Potion de défense (+)", 40
		case 7:
			nomPotion, prix = "Potion de défense (++)", 80
		case 8:
			nomPotion, prix = "Potion baisse défense (-)", 40
		case 9:
			nomPotion, prix = "Potion de poison (+)", 30
		case 10:
			nomPotion, prix = "Potion de poison (++)", 60
		case 11:
			nomPotion, prix = "Potion d'endurance (+)", 10
		case 12:
			nomPotion, prix = "Potion d'endurance (++)", 30
		default:
			fmt.Println(Red + "❌ Choix invalide." + Reset)
			continue
		}

		p.acheterPotion(nomPotion, prix)
	}
}

func (p *Personnage) acheterPotion(nom string, prix int) {
	if p.nombreObjets() >= p.CapaciteMax {
		fmt.Printf(Red+"❌ Erreur : Votre inventaire est plein (%d/%d objets max).\n"+Reset, p.nombreObjets(), p.CapaciteMax)
		return
	}

	if p.Argent >= prix {
		p.Argent -= prix
		if p.Inventaire == nil {
			p.Inventaire = make(map[string]int)
		}
		p.Inventaire[nom]++
		fmt.Printf(Green+"🛍️  Acheté : %s (-%d Oboles)\n"+Reset, nom, prix)
	} else {
		fmt.Printf(Red+"❌ Erreur : Il vous manque %d Oboles.\n"+Reset, prix-p.Argent)
	}
}
