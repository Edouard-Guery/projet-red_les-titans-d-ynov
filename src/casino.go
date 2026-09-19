package main

import (
	"fmt"
	"math/rand/v2"
)

func casino(p *Personnage) {
	var paris int
	var couleur string

	fmt.Println("Bienvenue au casino !")
	fmt.Println("Vous pouvez jouer à la roulette pour tenter de doubler votre argent.")
	fmt.Println("Entrez le montant que vous souhaitez miser (ou tapez 0 pour quitter) :")
	fmt.Scan(&paris)
	fmt.Println("Entrez la couleur sur laquelle vous souhaitez miser (rouge, noir, vert) :")
	fmt.Scan(&couleur)

	if paris <= 0 {
		fmt.Println("Vous quittez le casino.")
		return
	} else if paris > p.Argent {
		fmt.Println("Vous n'avez pas assez d'argent pour miser cette somme.")
		return
	} else {
		if couleur == "vert" {
			if vert() {
				p.Argent += paris * 35
				fmt.Printf("Félicitations ! Vous avez gagné %d oboles !\n", paris*35)
			} else {
				p.Argent -= paris
				fmt.Printf("Désolé, vous avez perdu %d oboles.\n", paris)
			}
		} else if couleur == "rouge" {
			if rouge() {
				p.Argent += paris
				fmt.Printf("Félicitations ! Vous avez gagné %d oboles !\n", paris)
			} else {
				p.Argent -= paris
				fmt.Printf("Désolé, vous avez perdu %d oboles.\n", paris)
			}
		} else if couleur == "noir" {
			if noir() {
				p.Argent += paris
				fmt.Printf("Félicitations ! Vous avez gagné %d oboles !\n", paris)
			} else {
				p.Argent -= paris
				fmt.Printf("Désolé, vous avez perdu %d oboles.\n", paris)
			}
		} else {
			fmt.Println("Couleur invalide.")
		}
	}
}

func vert() bool {
	if rand.N(37) == 0 {
		return true
	}
	return false
}

func rouge() bool {
	if rand.N(37) < 18 {
		return true
	}
	return false
}

func noir() bool {
	if rand.N(37) < 18 {
		return true
	}
	return false
}

