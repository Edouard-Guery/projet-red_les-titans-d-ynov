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
		fmt.Println("\n" + Magenta + "============================================================" + Reset)
		fmt.Println(Bold + Yellow + "                     🎒 VOTRE INVENTAIRE 🎒                     " + Reset)
		fmt.Println(Magenta + "============================================================" + Reset)
		fmt.Printf(Cyan+"💰 Oboles : %d | 📦 Places : %d/%d\n"+Reset, p.Argent, p.nombreObjets(), p.CapaciteMax)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)

		if len(p.Inventaire) == 0 {
			fmt.Println(Gray + " Votre inventaire est vide." + Reset)
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

		// Séparation visuelle après la saisie
		fmt.Println("\n" + Magenta + "============================================================" + Reset)

		choix := strings.TrimSpace(scanner.Text())

		if choix == "0" {
			fmt.Println(Gray + "🎒 Vous fermez votre sac." + Reset)
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

// Applique l'effet de l'objet et le retire de l'inventaire
func (p *Personnage) utiliserObjet(nom string) {
	// Permet d'équiper les armures de la forge ! (Redirige vers equipement.go)
	if p.UtiliserEquipementForge(nom) {
		return
	}

	switch nom {
	case "Potion de vie (+)":
		p.PotionViePlus() // Fonction située dans effets.go
		fmt.Println(Green + "🧪 Vous avez bu une Potion de vie (+) et regagné 30 PV !" + Reset)
	case "Potion de vie (++)":
		p.PotionViePlusPlus()
		fmt.Println(Green + "🧪 Vous avez bu une Potion de vie (++) et regagné 100 PV !" + Reset)
	case "Potion d'attaque (+)":
		p.PotionAttaquePlus()
		fmt.Println(Yellow + "🧪 Vous avez bu une Potion d'attaque (+). Votre attaque augmente !" + Reset)
	case "Potion de défense (+)":
		p.PotionDefensePlus()
		fmt.Println(Blue + "🧪 Vous avez bu une Potion de défense (+). Votre défense augmente !" + Reset)
	default:
		// Si l'objet n'est pas utilisable (comme les matériaux de forge ou un objet passif)
		fmt.Printf(Red+"❌ Vous ne pouvez pas utiliser '%s' directement ici.\n"+Reset, nom)
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
		fmt.Println(Red + "❌ Votre inventaire est déjà à sa taille maximale." + Reset)
		return
	}
	p.CapaciteMax += 10
	fmt.Printf(Green+"✅ Succès ! Votre inventaire a été agrandi. Capacité actuelle : %d places.\n"+Reset, p.CapaciteMax)
}