// trainingfight.go
package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func trainingFight(joueur *Personnage, scanner *bufio.Scanner) {
	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Println(Bold + Yellow + "             🥊 ÉPREUVE DE L'ARÈNE : SÉLECTION DU MONSTRE 🥊             " + Reset)
	fmt.Println(Magenta + "============================================================" + Reset)

	for i, monstre := range ListeMonstres {
		fmt.Printf(White+" %2d. "+Red+"%s "+Gray+"(%d PV)\n"+Reset, i+1, monstre.Nom, monstre.PVMax)
	}

	fmt.Println(Blue + "------------------------------------------------------------" + Reset)

	for {
		fmt.Print(Cyan + "👉 Désigne le numéro de la créature à terrasser : " + Reset)

		if !scanner.Scan() {
			return
		}

		fmt.Println("\n" + Magenta + "============================================================" + Reset)

		saisie := strings.TrimSpace(scanner.Text())
		choix, err := strconv.Atoi(saisie)

		if err != nil || choix < 1 || choix > len(ListeMonstres) {
			fmt.Println(Red + "❌ Les divinités rejettent ce geste. Invoque un autre choix." + Reset)
			continue
		}

		monstreChoisi := ListeMonstres[choix-1]

		combat := NouveauCombatMonstre(joueur, &monstreChoisi)
		combat.Entrainement = true
		combat.LancerDéroulement()

		return
	}
}
