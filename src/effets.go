package main

import (
	"fmt"
	"math/rand/v2"
)

func (p *Personnage) BouclierEclair() {
	p.Defense += 50
}

func (p *Personnage) PendentifVenin() {
	p.PV += 20
}

func (p *Personnage) FauxEternite() {
	p.Attaque += 30
}

// m utilise la structure Personnage
func (p *Personnage) BoitePandore(m *Personnage) {
	if rand.N(3) < 2 {
		p.Attaque *= 2
	} else {
		m.Attaque *= 2
	}
}

func (p *Personnage) ArmureHephaistos() {
	p.Defense *= 2
}

func (p *Personnage) CarteArchipel() {
	p.Attaque += 10
}

func (p *Personnage) FrenesieTravaux() {
	p.ToursBonus = true
}

func (p *Personnage) SacMagique() {
	p.EmplacementsEquipement += 1
}

// m utilise la structure Personnage
func (p *Personnage) RegardGorgone(m *Personnage) bool {
	if rand.N(100) < 20 {
		m.PV = 0
		return true
	}
	return false
}

// m utilise la structure Personnage
func (p *Personnage) GueuleEnflammee(m *Personnage) {
	degatsBrulure := 15
	m.PV -= degatsBrulure
}

// m utilise la structure Personnage
func (p *Personnage) HacheDoubleTranchant(m *Personnage) {
	p.Attaque *= 2
	m.Attaque *= 2
}

func (p *Personnage) GourdeRegeneration() {
	// S'exécute seulement si le joueur n'est pas mort et n'a pas déjà sa vie au max
	if p.PV > 0 && p.PV < p.PVMax {
		p.PV += 10
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Printf("💧 [Gourde] %s se régénère de 10 PV ! (%d/%d PV)\n", p.Nom, p.PV, p.PVMax)
	}
}

func (p *Personnage) BouclierBoisRenforce() {
	p.Defense += 15
}

func (p *Personnage) CasqueSagesse() {
	if rand.N(100) < 20 {
		p.Defense += 20
	}
}

// --- EFFETS DES POTIONS ---- //

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

// m utilise la structure Personnage
func (p *Personnage) PotionBaisseAttaque(m *Personnage) {
	m.Attaque -= 10
}

// m utilise la structure Personnage
func (p *Personnage) PotionBaisseDefense(m *Personnage) {
	m.Defense -= 10
}