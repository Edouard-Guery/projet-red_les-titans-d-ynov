package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func AfficherItem() {
	items := BoutiqueMarchand()
	
	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Println(Bold + Yellow + "          🏛️  BOUTIQUE DES ÉQUIPEMENTS MYTHOLOGIQUES 🏛️          " + Reset)
	fmt.Println(Magenta + "============================================================" + Reset)
	
	for i, eq := range items {
		fmt.Printf(White+" %2d. "+Green+"%s "+Cyan+"(%s)"+White+" — "+Yellow+"%d Oboles\n"+Reset, i+1, eq.Nom, eq.Source, eq.Prix)
		fmt.Printf(Gray+"     └─ %s\n"+Reset, eq.Description)
	}
	
	fmt.Println(Blue + "------------------------------------------------------------" + Reset)
	fmt.Println(Gray + " 16. Quitter la boutique" + Reset)
	fmt.Println(Blue + "------------------------------------------------------------" + Reset)
}

func ChoisirItem(hero *Personnage, monstre *Personnage, scanner *bufio.Scanner) {
	fmt.Print(Cyan + "👉 Entre le numéro de l'équipement souhaité : " + Reset)

	if !scanner.Scan() {
		return
	}

	fmt.Println("\n" + Magenta + "============================================================" + Reset)

	choixStr := strings.TrimSpace(scanner.Text())
	nombre, err := strconv.Atoi(choixStr)

	if err != nil {
		fmt.Println(Red + "❌ Choix invalide !" + Reset)
		return
	}

	if nombre == 16 {
		fmt.Println(Gray + "👋 Vous quittez la boutique des équipements." + Reset)
		return
	}

	items := BoutiqueMarchand()
	if nombre < 1 || nombre > len(items) {
		fmt.Println(Red + "❌ Choix invalide !" + Reset)
		return
	}

	itemChoisi := items[nombre-1]

	if hero.nombreObjets() >= hero.CapaciteMax {
		fmt.Printf(Red+"❌ Erreur : Votre inventaire est plein (%d/%d objets max).\n"+Reset, hero.nombreObjets(), hero.CapaciteMax)
		return
	}

	if hero.Argent < itemChoisi.Prix {
		fmt.Printf(Red+"❌ Erreur : Il vous manque %d Oboles.\n"+Reset, itemChoisi.Prix-hero.Argent)
		return
	}

	hero.Argent -= itemChoisi.Prix
	hero.Inventaire[itemChoisi.Nom]++
	fmt.Printf(Green+"🛍️  Acheté : %s (-%d Oboles). Ajouté à l'inventaire !\n"+Reset, itemChoisi.Nom, itemChoisi.Prix)
	fmt.Println("15. Augmentation d'inventaire (+10 places) —  30 oboles")
	

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
	case 15:
		hero.upgradeInventorySlot()
	}
}

func (p *Personnage) upgradeInventorySlot() bool {
	const prix = 30

	if p.AmeliorationsInventaire >= 3 {
		fmt.Println("❌ Vous avez déjà acheté les 3 augmentations d'inventaire.")
		return false
	}

	if p.Argent < prix {
		fmt.Println("❌ Il faut 30 pièces d'or.")
		return false
	}

	p.Argent -= prix
	p.CapaciteMax += 10
	p.AmeliorationsInventaire++

	fmt.Printf(
		"🎒 Inventaire augmenté : %d places (%d/3 achats).\n",
		p.CapaciteMax,
		p.AmeliorationsInventaire,
	)

	return true
}
