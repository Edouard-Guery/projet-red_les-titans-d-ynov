package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/c-bata/go-prompt"
)

var Reset = "\033[0m"
var Bold = "\033[1m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Autocomplétion pour le choix de la classe
func completerClasse(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix([]prompt.Suggest{{Text: "creature"}, {Text: "demi-dieu"}, {Text: "dieu"}}, d.GetWordBeforeCursor(), true)
}

// Autocomplétion pour le menu principal
func completerMenu(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix([]prompt.Suggest{{Text: "start"}, {Text: "inv"}, {Text: "shop"}, {Text: "entrainement"}, {Text: "casino"}, {Text: "potions"}, {Text: "forge"}, {Text: "info"}, {Text: "autel"}, {Text: "quit"}, {Text: "clear"}}, d.GetWordBeforeCursor(), true)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(Yellow + "=======================================" + Reset)
	fmt.Println(Bold + Magenta + "        🔱 LES TITANS D'YNOV 🔱        " + Reset)
	fmt.Println(Yellow + "=======================================" + Reset)
	fmt.Println()

	nomHero := "Héros"
	classeHero := "creature"

	fmt.Print(Cyan + "👤 Entrez le nom de votre héros : " + Reset)
	if scanner.Scan() {
		nomSaisi := strings.TrimSpace(scanner.Text())
		if len(nomSaisi) > 0 {
			nomHero = strings.ToUpper(string(nomSaisi[0])) + strings.ToLower(nomSaisi[1:])
		}
	}

	fmt.Println(Cyan + "⚔️  Entrez la classe de votre héros (creature / demi-dieu / dieu) : " + Reset)
	classeSaisie := prompt.Input("👉 ", completerClasse)
	if strings.TrimSpace(classeSaisie) != "" {
		classeHero = strings.TrimSpace(strings.ToLower(classeSaisie))
	}

	hero := Person(nomHero, classeHero, scanner)
	// go LancerServeurWeb(&hero) -- en attente, upgrade web
	fmt.Printf("\n"+Green+"✨ Bienvenue %s, illustre %s ! Que votre quête commence."+Reset+"\n", hero.Nom, hero.Classe)

	index := 0

	for {
		fmt.Println("\n" + Blue + "-----------------------------------------------------------------------" + Reset)
		fmt.Println(Yellow + "📜 Que souhaitez-vous faire ?" + Reset)
		fmt.Println(White + " 🗡️  start      🎒 inv       🛒 shop      🎯 entrainement  🎲 casino" + Reset)
		fmt.Println(White + " 🧪 potions    ⚒️  forge     ℹ️  info      🪙  autel         ❌ quit" + Reset)
		fmt.Println(Blue + "-----------------------------------------------------------------------" + Reset)

		commandeSaisie := prompt.Input("👉 ", completerMenu)
		fmt.Println("\n" + Magenta + "============================================================" + Reset)

		commande := strings.TrimSpace(strings.ToLower(commandeSaisie))

		switch commande {
		case "clear":
			clearScreen()
			fmt.Println(Green + "✨ Le terminal a été purifié par les Dieux." + Reset)

		case "inv", "inventaire", "inv-list":
			hero.MenuInventaire(scanner)

		case "start", "fight", "combat":
			if hero.PV <= 0 {
				fmt.Println("\n" + Red + "💀 Vous êtes inconscient ! Soignez-vous avec une potion avant d'entrer en combat." + Reset)
				break
			}

			if index >= len(ListeMonstres) {
				fmt.Println("\n" + Green + "🏆 Félicitations ! Vous avez déjà vaincu tous les monstres de l'Olympe !" + Reset)
				break
			}

			monstre := ObtenirMonstre(index)
			fmt.Printf("\n"+Bold+Red+"--- ⚔️ COMBAT %d/%d : %s apparaît ! ---"+Reset+"\n", index+1, len(ListeMonstres), monstre.Nom)

			combat := NouveauCombatMonstre(&hero, &monstre)
			combat.LancerDéroulement()

			if hero.PV > 0 {
				fmt.Printf("\n"+Green+"🎉 Victoire ! %s a été vaincu."+Reset+"\n", monstre.Nom)

				if monstre.Drop != "" {
					if hero.nombreObjets() < hero.CapaciteMax {
						hero.Inventaire[monstre.Drop]++
						fmt.Printf(Yellow+"🎁 Vous avez récupéré un objet : %s !"+Reset+"\n", monstre.Drop)
					} else {
						fmt.Printf(Gray+"❌ Vous avez trouvé [%s], mais votre sac est plein !"+Reset+"\n", monstre.Drop)
					}
				}

				index++
				fmt.Printf(Cyan + "🔓 Prochain monstre débloqué." + Reset + "\n")

				if index%2 == 0 {
					hero.GagnerNiveau(scanner)
				}
			} else {
				fmt.Println("\n" + Red + "💀 Game Over... Vous vous réveillez avec 20% de vos PV." + Reset)
				hero.PV = hero.PVMax / 5
			}

		case "magasin", "shop":
			AfficherItem(&hero)
			monstreActuel := Personnage{Nom: "Cible", PV: 100, Attaque: 10, Defense: 10}
			ChoisirItem(&hero, &monstreActuel, scanner)

		case "forge", "craft":
			menuForgeron(&hero, scanner)

		case "potions", "potion":
			hero.ShopPotions(scanner)

		case "stop", "end", "info":
			fmt.Printf("\n"+Cyan+"📊 Stats - Nom: %s | Classe: %s | PV: %d/%d | ⚡ Endu: %d/%d | Attaque: %d | Défense: %d | Oboles: %d"+Reset+"\n",
				hero.Nom, hero.Classe, hero.PV, hero.PVMax, hero.Endurance, hero.EnduranceMax, hero.Attaque, hero.Defense, hero.Argent)
			fmt.Println(Magenta + "🎵 Easter Egg : Jeu certifié par ABBA et Steven Spielberg 🎬" + Reset)

		case "quit", "disconnect":
			fmt.Println("\n" + Green + "👋 Merci d'avoir joué ! À bientôt dans l'Olympe." + Reset)
			os.Exit(0)

		case "training", "train", "entrainement":
			trainingFight(&hero, scanner)

		case "casino":
			casino(&hero)

		case "autel", "faveur", "prier":
			PrierALAutel(&hero, scanner)

		default:
			fmt.Println("\n" + Red + "❌ Commande invalide. Les dieux ne comprennent pas votre requête." + Reset)
		}
	}
}

func NouveauCombatMonstre(personnage *Personnage, monstre *Monstre) *Combat {
	return &Combat{
		Joueur:  personnage,
		Monstre: monstre,
		Tour:    1,
	}
}
