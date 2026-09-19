package main

import "math/rand/v2"

func (p *Personnage)BouclierEclair()  {
	p.Defense+=50
}

func (p *Personnage)PendentifVenin(){
	p.PV+=20
}

func (p *Personnage) FauxEternite()  {
	p.Attaque+=30
}

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

func (p *Personnage) RegardGorgone(m *Personnage) bool {
	if rand.N(100) < 20 {
		m.PV = 0
		return true
	}
	return false
}

func (p *Personnage) GueuleEnflammee(m *Personnage) {
	degatsBrulure := 15
	m.PV -= degatsBrulure
}

func (p *Personnage) HacheDoubleTranchant(m *Personnage) {
	p.Attaque *= 2
	m.Attaque *= 2
}

func (p *Personnage) GourdeRegeneration() {
	p.PV += 10
}

func (p *Personnage) BouclierBoisRenforce() {
	p.Defense += 15
}
func (p *Personnage) CasqueSagesse() {
	if rand.N(100) < 20 {
		p.Defense += 20
	}
}