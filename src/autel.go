package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func autel(joueur *Personnage, scanner *bufio.Scanner) {
	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Println(Bold + Yellow + "                   ⚡ AUTEL DE ZEUS ⚡                   " + Reset)
	fmt.Println(Magenta + "============================================================" + Reset)
	fmt.Printf(Cyan+"❤️  PV actuels : %d/%d | 💰 Oboles : %d\n"+Reset, joueur.PV, joueur.PVMax, joueur.Argent)
	fmt.Println(Blue + "------------------------------------------------------------" + Reset)
	fmt.Println(White + "Combien de PV voulez-vous sacrifier au Dieu des Dieux ? " + Gray + "(1 PV = 2 Oboles)" + Reset)
	fmt.Print(Cyan + "👉 " + Reset)

	if !scanner.Scan() {
		return
	}

	fmt.Println("\n" + Magenta + "============================================================" + Reset)

	entree := strings.TrimSpace(scanner.Text())
	sacrifice, err := strconv.Atoi(entree)

	if err != nil || sacrifice <= 0 {
		fmt.Println(Red + "❌ Offrande invalide. L'autel exige un véritable sacrifice (au moins 1 PV)." + Reset)
		return
	}

	if sacrifice >= joueur.PV {
		joueur.PV = 0
		fmt.Println(Bold + Red + "💀 Vous avez sacrifié toute votre force vitale... Zeus emporte votre âme. GAME OVER." + Reset)
		return
	}

	joueur.PV -= sacrifice
	obolesGagnees := sacrifice * 2
	joueur.Argent += obolesGagnees

	fmt.Printf(Bold+Green+"⚡ Zeus accepte votre sang ! Vous sacrifiez "+Red+"%d PV"+Green+" et gagnez "+Yellow+"%d Oboles"+Green+".\n"+Reset, sacrifice, obolesGagnees)
	fmt.Printf(Cyan+"❤️  PV restants : %d | 💰 Nouvel équilibre : %d Oboles\n"+Reset, joueur.PV, joueur.Argent)
}