package main

import "fmt"

func main() {

	hero := Person()
	Monstre:=Person()
	AfficherItem()
	ChoisirItem(&hero,&Monstre)
	fmt.Printf("attaque du héros : %d\n", hero.Attaque)
	fmt.Printf("attaque du monstre : %d\n", Monstre.Attaque)
}