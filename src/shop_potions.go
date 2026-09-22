package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func (p *Personnage) ShopPotions(scanner *bufio.Scanner) {
	for {
		// En-tête stylisée
		fmt.Println("\n" + Magenta + "============================================================" + Reset)
		fmt.Println(Bold + Green + "              🍷 MAGASIN DE POTIONS DE DIONYSOS 🍷              " + Reset)
		fmt.Println(Magenta + "============================================================" + Reset)
		fmt.Printf(Cyan+"💰 Oboles : %d | 🎒 Inventaire : %d/%d\n"+Reset, p.Argent, p.nombreObjets(), p.CapaciteMax)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)

		// Liste des potions alignée
		fmt.Println(White + "  1. Potion de vie (+)           - " + Yellow + "5 Oboles" + Reset)
		fmt.Println(White + "  2. Potion de vie (++)          - " + Yellow + "15 Oboles" + Reset)
		fmt.Println(White + "  3. Potion d'attaque (+)        - " + Yellow + "20 Oboles" + Reset)
		fmt.Println(White + "  4. Potion d'attaque (++)       - " + Yellow + "35 Oboles" + Reset)
		fmt.Println(White + "  5. Potion baisse d'attaque (-) - " + Yellow + "25 Oboles" + Reset)
		fmt.Println(White + "  6. Potion de défense (+)       - " + Yellow + "20 Oboles" + Reset)
		fmt.Println(White + "  7. Potion de défense (++)      - " + Yellow + "35 Oboles" + Reset)
		fmt.Println(White + "  8. Potion baisse défense (-)   - " + Yellow + "25 Oboles" + Reset)
		fmt.Println(White + "  9. Potion de poison (+)        - " + Yellow + "15 Oboles" + Reset)
		fmt.Println(White + " 10. Potion de poison (++)       - " + Yellow + "30 Oboles" + Reset)
		fmt.Println(Gray + " 11. Quitter la boutique" + Reset)
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

		if choix == 11 {
			fmt.Println(Gray + "👋 Vous quittez la boutique de Dionysos." + Reset)
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