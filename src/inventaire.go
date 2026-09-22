package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func (p *Personnage) MenuInventaire(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n" + Magenta + "============================================================" + Reset)
		fmt.Println(Bold + Yellow + "                     🎒 VOTRE INVENTAIRE 🎒                     " + Reset)
		fmt.Println(Magenta + "============================================================" + Reset)
		fmt.Printf(Cyan+"💰 Oboles : %d | 📦 Places : %d/%d\n"+Reset, p.Argent, p.nombreObjets(), p.CapaciteMax)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)

		if len(p.Inventaire) == 0 {
			fmt.Println(Gray + " Votre sac est vide." + Reset)
			fmt.Println(Blue + "------------------------------------------------------------" + Reset)
			return
		}

		// Trie les objets par ordre alphabétique
		noms := make([]string, 0, len(p.Inventaire))
		for nom := range p.Inventaire {
			noms = append(noms, nom)
		}
		sort.Strings(noms)

		fmt.Println(White + " Objets disponibles :" + Reset)
		for i, nom := range noms {
			fmt.Printf(White+" %2d. %s "+Cyan+"(x%d)\n"+Reset, i+1, nom, p.Inventaire[nom])
		}

		fmt.Println(Blue + "------------------------------------------------------------" + Reset)
		fmt.Println(Gray + "  0. Fermer l'inventaire" + Reset)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)
		fmt.Print(Cyan + "👉 Entrez le numéro de l'objet à utiliser : " + Reset)

		if !scanner.Scan() {
			return
		}

		fmt.Println("\n" + Magenta + "============================================================" + Reset)

		choix := strings.TrimSpace(scanner.Text())

		if choix == "0" {
			fmt.Println(Gray + "🎒 Vous refermez votre sac." + Reset)
			return
		}

		index, err := strconv.Atoi(choix)
		if err != nil || index < 1 || index > len(noms) {
			fmt.Println(Red + "❌ Choix invalide." + Reset)
			continue
		}

		objetChoisi := noms[index-1]
		p.utiliserObjet(objetChoisi)
	}
}

func (p *Personnage) utiliserObjet(nom string) {
	if p.UtiliserEquipementForge(nom) {
		return
	}

	switch nom {
	case "Potion de vie (+)":
		if p.PV == p.PVMax {
			fmt.Println(Yellow + "⚠️ Vos PV sont déjà au maximum. Inutile de gâcher une potion." + Reset)
			return
		}
		p.PotionViePlus()
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Printf(Green+"🧪 Vous avez bu une Potion de vie (+). PV actuels : %d/%d\n"+Reset, p.PV, p.PVMax)
	case "Potion de vie (++)":
		if p.PV == p.PVMax {
			fmt.Println(Yellow + "⚠️ Vos PV sont déjà au maximum. Inutile de gâcher une potion." + Reset)
			return
		}
		p.PotionViePlusPlus()
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Printf(Green+"🧪 Vous avez bu une Potion de vie (++). PV actuels : %d/%d\n"+Reset, p.PV, p.PVMax)
	case "Potion d'attaque (+)":
		p.PotionAttaquePlus()
		fmt.Println(Yellow + "🧪 Vous avez bu une Potion d'attaque (+). Votre puissance brute augmente de manière permanente !" + Reset)
	case "Potion de défense (+)":
		p.PotionDefensePlus()
		fmt.Println(Blue + "🧪 Vous avez bu une Potion de défense (+). Votre peau s'endurcit de manière permanente !" + Reset)
	default:
		fmt.Printf(Red+"❌ Cet objet ('%s') ne peut pas être consommé ou équipé depuis ce menu.\n"+Reset, nom)
		return
	}

	p.Inventaire[nom]--
	if p.Inventaire[nom] <= 0 {
		delete(p.Inventaire, nom)
	}
}

func (p *Personnage) UpgradeInventorySlot() {
	if p.CapaciteMax >= 40 {
		fmt.Println(Red + "❌ Votre sac est déjà à sa taille maximale. Les dieux ne peuvent pas l'agrandir davantage." + Reset)
		return
	}
	p.CapaciteMax += 10
	fmt.Printf(Green+"✅ Succès ! Votre inventaire a été agrandi. Capacité actuelle : %d places.\n"+Reset, p.CapaciteMax)
}
