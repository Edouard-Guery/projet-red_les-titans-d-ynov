package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

const (
	miseMaxCasino = 100
	miseMinCasino = 1
)

func casino(p *Personnage) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Println(Bold + Yellow + "                 🎰 CASINO DE L'OLYMPE 🎰                 " + Reset)
	fmt.Println(Magenta + "============================================================" + Reset)
	fmt.Printf(Cyan+"💰 Vos Oboles : %d | ⚠️ Mise max : %d\n"+Reset, p.Argent, miseMaxCasino)
	fmt.Println(White + "Règles : " + Red + "Rouge" + White + " / " + Gray + "Noir" + White + " (Gain x2) | " + Green + "Vert" + White + " (Gain x35)" + Reset)
	fmt.Println(Magenta + "------------------------------------------------------------" + Reset)

	if p.Argent < miseMinCasino {
		fmt.Println(Red + "❌ Vous êtes à sec. Revenez avec des Oboles." + Reset)
		return
	}

	fmt.Print(Cyan + "Combien d'Oboles souhaitez-vous miser ? (0 pour quitter) : " + Reset)
	if !scanner.Scan() {
		return
	}
	parisStr := strings.TrimSpace(scanner.Text())
	paris, err := strconv.Atoi(parisStr)

	if err != nil || paris <= 0 {
		fmt.Println(Yellow + "\n🚶 Vous quittez la table." + Reset)
		return
	}

	if paris > p.Argent {
		fmt.Println(Red + "\n❌ Fonds insuffisants." + Reset)
		return
	}

	if paris > miseMaxCasino {
		fmt.Printf(Red+"\n❌ La table n'accepte pas les mises de plus de %d Oboles.\n"+Reset, miseMaxCasino)
		return
	}

	fmt.Print(Cyan + "Sur quelle couleur ? (rouge, noir, vert) : " + Reset)
	if !scanner.Scan() {
		return
	}
	couleur := strings.ToLower(strings.TrimSpace(scanner.Text()))

	if couleur != "rouge" && couleur != "noir" && couleur != "vert" {
		fmt.Println(Red + "\n❌ Choix invalide. La mise est annulée." + Reset)
		return
	}

	fmt.Println(Yellow + "\n🎡 La roulette tourne..." + Reset)

	resultat := rand.N(37)
	couleurTiree := ""
	couleurAffichage := ""

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

	if couleur == couleurTiree {
		if couleur == "vert" {
			gain := paris * 35
			p.Argent += gain
			fmt.Printf(Bold+Green+"🎉 JACKPOT ! Vous remportez %d Oboles !\n"+Reset, gain)
		} else {
			p.Argent += paris
			fmt.Printf(Bold+Green+"🎉 Bien joué ! Vous remportez %d Oboles !\n"+Reset, paris)
		}
	} else {
		p.Argent -= paris
		fmt.Printf(Red+"💸 Perdu. L'Olympe récupère vos %d Oboles.\n"+Reset, paris)
	}

	fmt.Printf(Cyan+"💰 Solde actuel : %d Oboles\n"+Reset, p.Argent)
	fmt.Println(Magenta + "============================================================" + Reset)

	fmt.Print(Gray + "Appuyez sur [Entrée] pour continuer..." + Reset)
	scanner.Scan()
}
