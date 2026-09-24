package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func AfficherItem(hero *Personnage) {
	items := BoutiqueMarchand()

	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Println(Bold + Yellow + "          🏛️  BOUTIQUE DES ÉQUIPEMENTS MYTHOLOGIQUES 🏛️          " + Reset)
	fmt.Println(Magenta + "============================================================" + Reset)
	fmt.Printf(Cyan+"💰 Argent : %d Oboles\n"+Reset, hero.Argent)

	for i, eq := range items {
		fmt.Printf(White+" %2d. "+Green+"%s "+Cyan+"(%s)"+White+" — "+Yellow+"%d Oboles\n"+Reset, i+1, eq.Nom, eq.Source, eq.Prix)
		fmt.Printf(Gray+"     └─ %s\n"+Reset, eq.Description)
	}

	exitOption := len(items) + 1
	fmt.Println(Blue + "------------------------------------------------------------" + Reset)
	fmt.Printf(Gray+" %d. Quitter la boutique\n"+Reset, exitOption)
	fmt.Println(Blue + "------------------------------------------------------------" + Reset)
}

func ChoisirItem(hero *Personnage, monstre *Personnage, scanner *bufio.Scanner) {
	items := BoutiqueMarchand()
	exitOption := len(items) + 1

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

	if nombre == exitOption {
		fmt.Println(Gray + "👋 Vous quittez la boutique des équipements." + Reset)
		return
	}

	if nombre < 1 || nombre > len(items) {
		fmt.Println(Red + "❌ Choix invalide !" + Reset)
		return
	}

	itemChoisi := items[nombre-1]

	// Vérification de l'espace dans l'inventaire (Sauf si on achète une augmentation d'inventaire)
	if nombre == 15 {
		if hero.CapaciteMax >= 40 {
			fmt.Println(Red + "❌ Votre inventaire est déjà à sa taille maximale." + Reset)
			return
		}
	} else {
		if hero.nombreObjets() >= hero.CapaciteMax {
			fmt.Printf(Red+"❌ Erreur : Votre inventaire est plein (%d/%d objets max).\n"+Reset, hero.nombreObjets(), hero.CapaciteMax)
			return
		}
	}

	if hero.Argent < itemChoisi.Prix {
		fmt.Printf(Red+"❌ Erreur : Il vous manque %d Oboles.\n"+Reset, itemChoisi.Prix-hero.Argent)
		return
	}

	// Paiement
	hero.Argent -= itemChoisi.Prix

	// Ajout à l'inventaire (Sauf pour l'augmentation qui est un effet direct)
	if nombre != 15 {
		hero.Inventaire[itemChoisi.Nom]++
		fmt.Printf(Green+"🛍️  Acheté : %s (-%d Oboles). Ajouté à l'inventaire !\n"+Reset, itemChoisi.Nom, itemChoisi.Prix)
	} else {
		fmt.Printf(Green+"🛍️  Acheté : %s (-%d Oboles).\n"+Reset, itemChoisi.Nom, itemChoisi.Prix)
	}

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
		hero.CapaciteMax += 10
		fmt.Printf(Bold+Green+"🎒 Succès ! Votre inventaire a été agrandi. Capacité actuelle : %d places.\n"+Reset, hero.CapaciteMax)
	}
}
