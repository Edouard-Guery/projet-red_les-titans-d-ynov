package main

import "fmt"

func main() {

	hero := Person()
	Monstre:=Person()
	hero.PendentifVenin()
	hero.BouclierBoisRenforce()
	hero.HacheDoubleTranchant(&Monstre)
	fmt.Printf("PV du héros après Pendentif Venin : %d\n", hero.PV)
	fmt.Printf("attaque du héros après HacheDoubleTranchant : %d\n", hero.Attaque)
	fmt.Printf("attaque du monstre après HacheDoubleTranchant : %d\n", Monstre.Attaque)
	fmt.Printf("Défense du héros après BouclierBoisRenforce : %d\n", hero.Defense)

}