package main

import "fmt"

func (p *Personnage) EquiperCuirasseSauvage() {
	p.Defense += 10
	p.PV += 20
}

func (p *Personnage) EquiperRegardInfernal() {
	p.Attaque += 15
}

func (p *Personnage) EquiperBrassardEnferVivant() {
	p.Defense += 5
	p.Attaque += 10
}

func (p *Personnage) EquiperLameLabyrinthe() {
	p.Attaque += 25
}

func (p *Personnage) EquiperManteauEmpoisonne() {
	p.Defense += 15
	p.Attaque += 5
}

func (p *Personnage) EquiperRegardMortel() {
	p.Attaque += 30
	p.PV += 20
}

func (p *Personnage) EquiperEgideMers() {
	p.Defense += 40
	p.PV += 50
}

func (p *Personnage) EquiperFoudreInfernale() {
	p.Attaque += 50
}

func (p *Personnage) EquiperArmureColosse() {
	p.Defense += 60
	p.PV += 100
}

func (p *Personnage) UtiliserEquipementForge(nom string) bool {
	equipped := false

	switch nom {
	case "Cuirasse Sauvage":
		p.EquiperCuirasseSauvage()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Cuirasse Sauvage " + Cyan + "[🛡️ +10 Def | ❤️ +20 PV]" + Reset)
		equipped = true
	case "Regard Infernal":
		p.EquiperRegardInfernal()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Regard Infernal " + Cyan + "[🗡️ +15 Attaque]" + Reset)
		equipped = true
	case "Brassard de l'Enfer Vivant":
		p.EquiperBrassardEnferVivant()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Brassard de l'Enfer Vivant " + Cyan + "[🗡️ +10 Attaque | 🛡️ +5 Def]" + Reset)
		equipped = true
	case "Lame du Labyrinthe":
		p.EquiperLameLabyrinthe()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Lame du Labyrinthe " + Cyan + "[🗡️ +25 Attaque]" + Reset)
		equipped = true
	case "Manteau Empoisonné":
		p.EquiperManteauEmpoisonne()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Manteau Empoisonné " + Cyan + "[🛡️ +15 Def | 🗡️ +5 Attaque]" + Reset)
		equipped = true
	case "Regard Mortel":
		p.EquiperRegardMortel()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Regard Mortel " + Cyan + "[🗡️ +30 Attaque | ❤️ +20 PV]" + Reset)
		equipped = true
	case "Égide des Mers":
		p.EquiperEgideMers()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Égide des Mers " + Cyan + "[🛡️ +40 Def | ❤️ +50 PV]" + Reset)
		equipped = true
	case "Foudre Infernale":
		p.EquiperFoudreInfernale()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Foudre Infernale " + Cyan + "[⚡ +50 Attaque]" + Reset)
		equipped = true
	case "Armure du Colosse":
		p.EquiperArmureColosse()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Armure du Colosse " + Cyan + "[🗿 +60 Def | ❤️ +100 PV]" + Reset)
		equipped = true
	}

	if equipped {
		p.Inventaire[nom]--
		if p.Inventaire[nom] <= 0 {
			delete(p.Inventaire, nom)
		}
		fmt.Println(Magenta + "============================================================" + Reset)
		return true
	}
	return false
}