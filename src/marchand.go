package main

import "fmt"

func ChoisirItem() {
	var nombre int
	fmt.Print("Entre un chiffre : ")
	fmt.Scan(&nombre)

	fmt.Printf("Tu as tapé : %d\n", nombre)
}

func AfficherItem() {
	items := BoutiqueMarchand()
	fmt.Println("=== BOUTIQUE DES ÉQUIPEMENTS MYTHOLOGIQUES ===")
	for i, eq := range items {
		fmt.Printf("%d. %s (%s) — %d Or\n   └─ %s\n\n",
			i+1, eq.Nom, eq.Source, eq.Prix, eq.Description)
	}
	fmt.Println("15. annuler")
}
