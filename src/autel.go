package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	autelPVParObole    = 3
	autelPourcentMaxPV = 50
	autelSeuilMinPct   = 20
)

func autel(joueur *Personnage, scanner *bufio.Scanner) {
	fmt.Println("\n" + Magenta + "============================================================" + Reset)
	fmt.Println(Bold + Yellow + "                   ⚡ AUTEL DE ZEUS ⚡                   " + Reset)
	fmt.Println(Magenta + "============================================================" + Reset)
	fmt.Printf(Cyan+"❤️  PV actuels : %d/%d | 💰 Oboles : %d\n"+Reset, joueur.PV, joueur.PVMax, joueur.Argent)
	fmt.Println(Blue + "------------------------------------------------------------" + Reset)

	seuilMin := joueur.PVMax * autelSeuilMinPct / 100
	if joueur.PV <= seuilMin {
		fmt.Printf(Red+"❌ Zeus refuse votre offrande : vous êtes trop affaibli (%d/%d PV). Soignez-vous d'abord.\n"+Reset, joueur.PV, joueur.PVMax)
		fmt.Println(Magenta + "============================================================" + Reset)
		return
	}

	sacrificeMax := joueur.PV * autelPourcentMaxPV / 100
	if sacrificeMax < 1 {
		sacrificeMax = 1
	}

	fmt.Printf(White+"Combien de PV voulez-vous sacrifier au Dieu des Dieux ? "+Gray+"(%d PV = 1 Obole | max %d PV pour cette offrande)\n"+Reset, autelPVParObole, sacrificeMax)
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

	if sacrifice > sacrificeMax {
		fmt.Printf(Red+"❌ Zeus n'accepte pas plus de %d PV en une seule offrande (50%% de vos PV actuels).\n"+Reset, sacrificeMax)
		return
	}

	obolesGagnees := sacrifice / autelPVParObole
	if obolesGagnees <= 0 {
		fmt.Printf(Red+"❌ Offrande trop maigre : il faut au moins %d PV pour obtenir 1 Obole.\n"+Reset, autelPVParObole)
		return
	}

	joueur.PV -= sacrifice
	joueur.Argent += obolesGagnees

	fmt.Printf(Bold+Green+"⚡ Zeus accepte votre sang ! Vous sacrifiez "+Red+"%d PV"+Green+" et gagnez "+Yellow+"%d Oboles"+Green+".\n"+Reset, sacrifice, obolesGagnees)
	fmt.Printf(Cyan+"❤️  PV restants : %d/%d | 💰 Nouvel équilibre : %d Oboles\n"+Reset, joueur.PV, joueur.PVMax, joueur.Argent)
	fmt.Println(Magenta + "============================================================" + Reset)
}
