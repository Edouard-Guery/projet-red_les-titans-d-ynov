package main

import (
	"fmt"
	"math/rand/v2"
)

func (p *Personnage) BouclierEclair() {
	p.Defense += 50
	fmt.Println(Bold + Cyan + "⚡ Le Bouclier Éclair renforce votre posture ! " + Green + "[🛡️ +50 Défense]" + Reset)
}

func (p *Personnage) PendentifVenin() {
	p.PV += 20
	fmt.Println(Bold + Green + "🦟 Le Pendentif des Moustiques pulse... " + Cyan + "[❤️ +20 PV]" + Reset)
}

func (p *Personnage) FauxEternite() {
	p.Attaque += 30
	fmt.Println(Bold + Red + "⏳ La Faux d'Éternité s'imprègne de puissance ! " + Yellow + "[🗡️ +30 Attaque]" + Reset)
}

func (p *Personnage) BoitePandore(m *Personnage) {
	fmt.Println(Bold + Magenta + "📦 Vous ouvrez la redoutable Boîte de Pandore..." + Reset)
	if rand.N(3) < 2 {
		p.Attaque *= 2
		fmt.Println(Green + "✨ Bénédiction ! Votre puissance explose ! " + Cyan + "[🗡️ Attaque x2]" + Reset)
	} else {
		m.Attaque *= 2
		fmt.Println(Red + "💀 Malédiction ! La créature ciblée gagne en puissance ! " + Yellow + "[👹 Attaque Monstre x2]" + Reset)
	}
}

func (p *Personnage) ArmureHephaistos() {
	p.Defense *= 2
	fmt.Println(Bold + Yellow + "🛡️ L'Armure d'Héphaïstos s'ajuste à votre taille. " + Cyan + "[🛡️ Défense x2]" + Reset)
}

func (p *Personnage) CarteArchipel() {
	p.Attaque += 10
	fmt.Println(Bold + Blue + "🗺️ La Carte des Archipels révèle des points faibles. " + Yellow + "[🗡️ +10 Attaque]" + Reset)
}

func (p *Personnage) FrenesieTravaux() {
	p.ToursBonus = true
	fmt.Println(Bold + Red + "🔥 Frénésie des Travaux activée ! " + Cyan + "[⏱️ Tours Bonus activés]" + Reset)
}

func (p *Personnage) SacMagique() {
	p.EmplacementsEquipement += 1
	fmt.Println(Bold + Yellow + "🎒 Le Sac Magique s'agrandit... " + Cyan + "[📦 +1 Emplacement d'équipement]" + Reset)
}

func (p *Personnage) RegardGorgone(m *Personnage) bool {
	fmt.Println(Bold + Green + "🐍 Vous utilisez le Regard de la Gorgone..." + Reset)
	if rand.N(100) < 20 {
		m.PV = 0
		fmt.Println(Red + "🗿 PÉTRIFICATION ! L'ennemi est changé en pierre instantanément !" + Reset)
		return true
	}
	fmt.Println(Gray + "💨 L'ennemi a détourné le regard à temps..." + Reset)
	return false
}

func (p *Personnage) GueuleEnflammee(m *Personnage) {
	degatsBrulure := 15
	m.PV -= degatsBrulure
	fmt.Printf(Bold+Red+"🔥 La Gueule Enflammée crache le feu ! "+Yellow+"[-15 PV à %s]\n"+Reset, m.Nom)
}

func (p *Personnage) HacheDoubleTranchant(m *Personnage) {
	p.Attaque *= 2
	m.Attaque *= 2
	fmt.Println(Bold + Red + "🪓 Hache à Double Tranchant ! Le sang appelle le sang ! " + Cyan + "[🗡️ Attaque x2 pour vous ET l'ennemi]" + Reset)
}

func (p *Personnage) GourdeRegeneration() {
	if p.PV > 0 && p.PV < p.PVMax {
		p.PV += 10
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Printf(Green+"💧 [Gourde] %s se régénère doucement ! "+Cyan+"(%d/%d PV)\n"+Reset, p.Nom, p.PV, p.PVMax)
	}
}

func (p *Personnage) BouclierBoisRenforce() {
	p.Defense += 15
	fmt.Println(Bold + Yellow + "🛡️ Bouclier en Bois Renforcé équipé. " + Cyan + "[🛡️ +15 Défense]" + Reset)
}

func (p *Personnage) CasqueSagesse() {
	fmt.Println(Bold + Blue + "🦉 Casque de la Sagesse activé..." + Reset)
	if rand.N(100) < 20 {
		p.Defense += 20
		fmt.Println(Green + "✨ La sagesse d'Athéna vous protège ! " + Cyan + "[🛡️ +20 Défense]" + Reset)
	} else {
		fmt.Println(Gray + "💨 Le casque reste silencieux..." + Reset)
	}
}

func (p *Personnage) PotionViePlus() {
	p.PV += 30
}

func (p *Personnage) PotionViePlusPlus() {
	p.PV += 100
}

func (p *Personnage) PotionAttaquePlus() {
	p.Attaque += 10
}

func (p *Personnage) PotionAttaquePlusPlus() {
	p.Attaque += 25
}

func (p *Personnage) PotionDefensePlus() {
	p.Defense += 10
}

func (p *Personnage) PotionBaisseAttaque(m *Personnage) {
	m.Attaque -= 10
}

func (p *Personnage) PotionBaisseDefense(m *Personnage) {
	m.Defense -= 10
}