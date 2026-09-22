package main

// import (
// 	"bufio"
// 	"fmt"
// 	"strings"
// )

// func (p *Personnage) PrierALAutel(scanner *bufio.Scanner) {
// 	if p.FaveurDieux == nil {
// 		p.FaveurDieux = map[string]int{}
// 	}

// 	if p.Benedictions == nil {
// 		p.Benedictions = map[string]bool{}
// 	}

// 	fmt.Println("\n⛩️ ===== AUTEL DES DIEUX =====")
// 	fmt.Println("Offrir 10 PV donne 20 pièces d'or et 10 points de faveur.")
// 	fmt.Println("1. Zeus   — bénédiction : +7 attaque à 30 faveur")
// 	fmt.Println("2. Athéna — bénédiction : +7 défense à 30 faveur")
// 	fmt.Println("3. Hadès  — bénédiction : +30 PV maximum à 30 faveur")
// 	fmt.Println("0. Retour")
// 	fmt.Print("👉 Choisissez un dieu : ")

// 	if !scanner.Scan() {
// 		return
// 	}

// 	choix := strings.TrimSpace(scanner.Text())

// 	dieu := ""

// 	switch choix {
// 	case "1", "zeus":
// 		dieu = "Zeus"
// 	case "2", "athena", "athéna":
// 		dieu = "Athéna"
// 	case "3", "hades", "hadès":
// 		dieu = "Hadès"
// 	case "0":
// 		return
// 	default:
// 		fmt.Println("❌ Choix invalide.")
// 		return
// 	}

// 	if p.PV <= 10 {
// 		fmt.Println("❌ Vous n'avez pas assez de PV pour faire une offrande.")
// 		return
// 	}

// 	p.PV -= 10
// 	p.Argent += 20
// 	p.FaveurDieux[dieu] += 10

// 	fmt.Printf(
// 		"🙏 Vous offrez 10 PV à %s.\n💰 Vous recevez 20 pièces d'or.\n✨ Faveur de %s : %d/30.\n",
// 		dieu,
// 		dieu,
// 		p.FaveurDieux[dieu],
// 	)

// 	p.DonnerBenediction(dieu)
// }

// func (p *Personnage) DonnerBenediction(dieu string) {
// 	// Une seule bénédiction par dieu.
// 	if p.Benedictions[dieu] {
// 		return
// 	}

// 	if p.FaveurDieux[dieu] < 30 {
// 		return
// 	}

// 	switch dieu {
// 	case "Zeus":
// 		p.Attaque += 7
// 		fmt.Println("⚡ Zeus vous bénit : +7 attaque !")

// 	case "Athéna":
// 		p.Defense += 7
// 		fmt.Println("🛡️ Athéna vous bénit : +7 défense !")

// 	case "Hadès":
// 		p.PVMax += 30
// 		p.PV += 30

// 		if p.PV > p.PVMax {
// 			p.PV = p.PVMax
// 		}

// 		fmt.Println("💀 Hadès vous bénit : +30 PV maximum !")
// 	}

// 	p.Benedictions[dieu] = true
// }

// func (p *Personnage) AfficherFaveur() {
// 	fmt.Println("\n✨ ===== FAVEUR DES DIEUX =====")

// 	for _, dieu := range []string{"Zeus", "Athéna", "Hadès"} {
// 		fmt.Printf("%s : %d/30 faveur\n", dieu, p.FaveurDieux[dieu])
// 	}
// }
