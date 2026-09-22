package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func casino(p *Personnage) {
	// Utilisation d'un scanner local pour éviter les bugs de frappe
	scanner := bufio.NewScanner(os.Stdin)

	// En-tête stylisée du Casino
	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Println(Bold + Yellow + "                 🎰 CASINO DE L'OLYMPE 🎰                 " + Reset)
	fmt.Println(Magenta + "============================================================" + Reset)
	fmt.Printf(Cyan+"💰 Vos Oboles : %d\n"+Reset, p.Argent)
	fmt.Println(White + "Règles : " + Red + "Rouge" + White + " / " + Gray + "Noir" + White + " (Gain x2) | " + Green + "Vert" + White + " (Gain x35)" + Reset)
	fmt.Println(Magenta + "------------------------------------------------------------" + Reset)

	// Saisie de la mise
	fmt.Print(Cyan + "Combien d'Oboles souhaitez-vous miser ? (0 pour quitter) : " + Reset)
	if !scanner.Scan() {
		return
	}
	parisStr := strings.TrimSpace(scanner.Text())
	paris, err := strconv.Atoi(parisStr)

	if err != nil || paris <= 0 {
		fmt.Println(Yellow + "\n🚶 Vous quittez le casino." + Reset)
		return
	}

	if paris > p.Argent {
		fmt.Println(Red + "\n❌ Vous n'avez pas assez d'Oboles pour miser cette somme." + Reset)
		return
	}

	// Saisie de la couleur
	fmt.Print(Cyan + "Sur quelle couleur ? (rouge, noir, vert) : " + Reset)
	if !scanner.Scan() {
		return
	}
	couleur := strings.ToLower(strings.TrimSpace(scanner.Text()))

	if couleur != "rouge" && couleur != "noir" && couleur != "vert" {
		fmt.Println(Red + "\n❌ Couleur invalide. Vous quittez la table." + Reset)
		return
	}

	// Lancer de la roulette
	fmt.Println(Yellow + "\n🎡 La roulette tourne..." + Reset)
	
	// Un seul tirage entre 0 et 36
	resultat := rand.N(37) 
	couleurTiree := ""
	couleurAffichage := ""

	// Définition de la couleur gagnante
	if resultat == 0 {
		couleurTiree = "vert"
		couleurAffichage = Green + "Vert 🟢" + Reset
	} else if resultat <= 18 {
		couleurTiree = "rouge"
		couleurAffichage = Red + "Rouge 🔴" + Reset
	} else {
		couleurTiree = "noir"
		couleurAffichage = Gray + "Noir ⚫" + Reset
	}

	fmt.Printf(White+"Le numéro gagnant est le %d... C'est le %s !\n"+Reset, resultat, couleurAffichage)

	// Résultat des gains/pertes
	if couleur == couleurTiree {
		if couleur == "vert" {
			gain := paris * 35
			p.Argent += gain
			fmt.Printf(Bold+Green+"🎉 JACKPOT ! Vous avez gagné %d Oboles !\n"+Reset, gain)
		} else {
			p.Argent += paris
			fmt.Printf(Bold+Green+"🎉 Félicitations ! Vous avez gagné %d Oboles !\n"+Reset, paris)
		}
	} else {
		p.Argent -= paris
		fmt.Printf(Red+"💸 Désolé, vous avez perdu %d Oboles.\n"+Reset, paris)
	}

	fmt.Printf(Cyan+"💰 Solde actuel : %d Oboles\n"+Reset, p.Argent)
	fmt.Println(Magenta + "============================================================" + Reset)

	// Pause avant de retourner au menu principal
	fmt.Print(Gray + "Appuyez sur [Entrée] pour retourner à l'Olympe..." + Reset)
	scanner.Scan()
}