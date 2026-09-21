package main

import "fmt"

func ChoisirItem(hero *Personnage, monstre *Personnage) {
	var nombre int
	fmt.Print("Entre un chiffre : ")
	fmt.Scan(&nombre)

	fmt.Printf("Tu as tapé : %d\n", nombre)

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
		return // On quitte la fonction proprement
	default:
		fmt.Println("Choix invalide !")
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
