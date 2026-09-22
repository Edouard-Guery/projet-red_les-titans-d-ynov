package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	autelPVParObole   = 5
	autelSeuilMinPct  = 30
	autelFaveurGagnee = 10
)

func PrierALAutel(joueur *Personnage, scanner *bufio.Scanner) {
	if joueur.FaveurDieux == nil {
		joueur.FaveurDieux = map[string]int{}
	}
	if joueur.Benedictions == nil {
		joueur.Benedictions = map[string]bool{}
	}

	for {
		fmt.Println("\n" + Magenta + "============================================================" + Reset)
		fmt.Println(Bold + Yellow + "                    ⛩️  AUTEL DES DIEUX ⛩️" + Reset)
		fmt.Println(Magenta + "============================================================" + Reset)
		fmt.Printf(Cyan+" ❤️  PV : %d/%d  |  💰 Oboles : %d\n"+Reset, joueur.PV, joueur.PVMax, joueur.Argent)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)

		fmt.Println(Bold + Red + " [ SACRIFICES MATÉRIELS ]" + Reset)
		fmt.Printf(White+"  1. 🩸 Sacrifice de sang "+Gray+"(%d PV = 1 Obole | ⚠️ Mortel)\n"+Reset, autelPVParObole)

		fmt.Println(Bold + Green + "\n [ OFFRANDES DIVINES (Coût : 10 PV) ]" + Reset)
		fmt.Println(White + "  2. ⚡ Prier Zeus   " + Yellow + "➔ Objectif : +7 Attaque" + Reset)
		fmt.Println(White + "  3. 🛡️ Prier Athéna " + Cyan + "➔ Objectif : +7 Défense" + Reset)
		fmt.Println(White + "  4. 💀 Prier Hadès  " + Magenta + "➔ Objectif : +30 PV Max" + Reset)

		fmt.Println(Bold + Blue + "\n [ INFORMATIONS ]" + Reset)
		fmt.Println(White + "  5. ✨ Consulter la faveur des dieux" + Reset)
		fmt.Println(Gray + "  0. 🚪 Quitter l'autel" + Reset)
		fmt.Println(Blue + "------------------------------------------------------------" + Reset)
		fmt.Print(Cyan + "👉 Votre choix : " + Reset)

		if !scanner.Scan() {
			return
		}

		choix := strings.TrimSpace(scanner.Text())

		seuilMin := joueur.PVMax * autelSeuilMinPct / 100
		if joueur.PV <= seuilMin && choix != "0" && choix != "5" {
			fmt.Printf(Red+"\n❌ Les dieux refusent votre présence : vous êtes trop affaibli (%d/%d PV).\n"+Reset, joueur.PV, joueur.PVMax)
			fmt.Println(Gray + "Soignez-vous avant de revenir à l'autel." + Reset)
			continue
		}

		switch choix {
		case "1":
			fmt.Printf(White+"\nCombien de PV voulez-vous sacrifier ? "+Gray+"(%d PV = 1 Obole | max %d PV)\n"+Reset, autelPVParObole, joueur.PV)
			fmt.Print(Cyan + "👉 " + Reset)

			if !scanner.Scan() {
				return
			}

			entree := strings.TrimSpace(scanner.Text())
			sacrifice, err := strconv.Atoi(entree)

			if err != nil || sacrifice <= 0 {
				fmt.Println(Red + "❌ Offrande invalide." + Reset)
				continue
			}

			if sacrifice > joueur.PV {
				fmt.Println(Red + "❌ Vous ne possédez pas assez d'énergie vitale pour ce sacrifice." + Reset)
				continue
			}

			if sacrifice < autelPVParObole {
				fmt.Printf(Red+"❌ Un sacrifice trop faible n'attire pas l'attention. Minimum : %d PV.\n"+Reset, autelPVParObole)
				continue
			}

			obolesGagnees := sacrifice / autelPVParObole
			joueur.PV -= sacrifice
			joueur.Argent += obolesGagnees

			fmt.Printf(Bold+Green+"\n⚡ Le sang coule sur l'autel... Vous sacrifiez "+Red+"%d PV"+Green+" et recevez "+Yellow+"%d Oboles"+Green+".\n"+Reset, sacrifice, obolesGagnees)

			if joueur.PV <= 0 {
				joueur.PV = 0
				fmt.Println(Bold + Red + "\n💀 Vous avez tout sacrifié aux dieux... Votre corps s'effondre sur l'autel." + Reset)
				fmt.Println(Yellow + "Vos oboles tombent à côté de votre corps inerte. Utilisez une potion depuis votre sac si vous espérez survivre." + Reset)
				return
			}

		case "2", "3", "4":
			var dieu string
			switch choix {
			case "2":
				dieu = "Zeus"
			case "3":
				dieu = "Athéna"
			case "4":
				dieu = "Hadès"
			}

			if joueur.Benedictions[dieu] {
				fmt.Printf(Yellow+"\n✨ %s vous a déjà accordé sa bénédiction totale.\n"+Reset, dieu)
				continue
			}

			pvOffrande := 10
			if joueur.PV <= pvOffrande {
				fmt.Println(Red + "\n❌ Vous n'avez pas assez de PV pour cette offrande." + Reset)
				continue
			}

			joueur.PV -= pvOffrande
			joueur.FaveurDieux[dieu] += autelFaveurGagnee

			fmt.Printf(
				Bold+Green+"\n🙏 Vous sacrifiez %d PV en l'honneur de %s.\n✨ Faveur de %s : %d/30.\n"+Reset,
				pvOffrande, dieu, dieu, joueur.FaveurDieux[dieu],
			)

			joueur.DonnerBenediction(dieu)

		case "5":
			joueur.AfficherFaveur()

		case "0":
			fmt.Println(Gray + "\nVous quittez l'autel dans le silence." + Reset)
			return

		default:
			fmt.Println(Red + "\n❌ Choix invalide." + Reset)
		}
	}
}

func (p *Personnage) DonnerBenediction(dieu string) {
	if p.Benedictions[dieu] {
		return
	}

	if p.FaveurDieux[dieu] < 30 {
		return
	}

	fmt.Println(Magenta + "\n🌟 UNE LUMIÈRE DIVINE VOUS ENVELOPPE ! 🌟" + Reset)

	switch dieu {
	case "Zeus":
		p.Attaque += 7
		fmt.Println(Bold + Yellow + "⚡ Le Maître de l'Olympe vous bénit : votre force grandit de manière permanente ! [+7 Attaque]" + Reset)

	case "Athéna":
		p.Defense += 7
		fmt.Println(Bold + Cyan + "🛡️ La Déesse de la Stratégie vous bénit : votre peau s'endurcit ! [+7 Défense]" + Reset)

	case "Hadès":
		p.PVMax += 30
		p.PV += 30
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Println(Bold + Magenta + "💀 Le Seigneur des Enfers vous bénit : votre essence vitale s'étend ! [+30 PV Max]" + Reset)
	}

	p.Benedictions[dieu] = true
}

func (p *Personnage) AfficherFaveur() {
	fmt.Println("\n" + Bold + Blue + "✨ ===== FAVEUR DES DIEUX =====" + Reset)
	for _, dieu := range []string{"Zeus", "Athéna", "Hadès"} {
		faveur := p.FaveurDieux[dieu]
		etat := ""
		if p.Benedictions[dieu] {
			etat = Green + "(Bénédiction acquise)" + Reset
		}
		fmt.Printf(White+" - %-7s : "+Cyan+"%2d/30"+White+" %s\n"+Reset, dieu, faveur, etat)
	}
}
