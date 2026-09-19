package main

import "fmt"

func main() {

	hero := Person()
	hero.PendentifVenin()
	hero.BouclierBoisRenforce()
	fmt.Printf("PV du héros après Pendentif Venin : %d\n", hero.PV)
	fmt.Printf("Défense du héros après BouclierBoisRenforce : %d\n", hero.Defense)

}