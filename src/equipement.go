package main

import "fmt"

// LISSAGE DES STATS D'ÉQUIPEMENT
func (p *Personnage) EquiperCuirasseSauvage() {
	p.Defense += 5
	p.PVMax += 20 // Il vaut mieux augmenter les PV Max que juste soigner
	p.PV += 20
}

func (p *Personnage) EquiperRegardInfernal() {
	p.Attaque += 8
}

func (p *Personnage) EquiperBrassardEnferVivant() {
	p.Defense += 5
	p.Attaque += 5
}

func (p *Personnage) EquiperLameLabyrinthe() {
	p.Attaque += 15
}

func (p *Personnage) EquiperManteauEmpoisonne() {
	p.Defense += 12
	p.Attaque += 3
}

func (p *Personnage) EquiperRegardMortel() {
	p.Attaque += 22
	p.PVMax += 15
	p.PV += 15
}

func (p *Personnage) EquiperEgideMers() {
	p.Defense += 25
	p.PVMax += 40
	p.PV += 40
}

func (p *Personnage) EquiperFoudreInfernale() {
	p.Attaque += 35
}

func (p *Personnage) EquiperArmureColosse() {
	p.Defense += 45
	p.PVMax += 80
	p.PV += 80
}

func (p *Personnage) UtiliserEquipementForge(nom string) bool {
	equipped := false

	// Note: Idéalement, il faudrait vérifier ici si le joueur n'est pas DÉJÀ équipé
	// d'une arme/armure pour éviter le cumul infini.

	switch nom {
	case "Cuirasse Sauvage":
		p.EquiperCuirasseSauvage()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Cuirasse Sauvage " + Cyan + "[🛡️ +5 Def | ❤️ +20 PV Max]" + Reset)
		equipped = true
	case "Regard Infernal":
		p.EquiperRegardInfernal()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Regard Infernal " + Cyan + "[🗡️ +8 Attaque]" + Reset)
		equipped = true
	case "Brassard de l'Enfer Vivant":
		p.EquiperBrassardEnferVivant()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Brassard de l'Enfer Vivant " + Cyan + "[🗡️ +5 Attaque | 🛡️ +5 Def]" + Reset)
		equipped = true
	case "Lame du Labyrinthe":
		p.EquiperLameLabyrinthe()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Lame du Labyrinthe " + Cyan + "[🗡️ +15 Attaque]" + Reset)
		equipped = true
	case "Manteau Empoisonné":
		p.EquiperManteauEmpoisonne()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Manteau Empoisonné " + Cyan + "[🛡️ +12 Def | 🗡️ +3 Attaque]" + Reset)
		equipped = true
	case "Regard Mortel":
		p.EquiperRegardMortel()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Regard Mortel " + Cyan + "[🗡️ +22 Attaque | ❤️ +15 PV Max]" + Reset)
		equipped = true
	case "Égide des Mers":
		p.EquiperEgideMers()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Égide des Mers " + Cyan + "[🛡️ +25 Def | ❤️ +40 PV Max]" + Reset)
		equipped = true
	case "Foudre Infernale":
		p.EquiperFoudreInfernale()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Foudre Infernale " + Cyan + "[⚡ +35 Attaque]" + Reset)
		equipped = true
	case "Armure du Colosse":
		p.EquiperArmureColosse()
		fmt.Println(Bold + Yellow + "✨ Équipement revêtu : " + White + "Armure du Colosse " + Cyan + "[🗿 +45 Def | ❤️ +80 PV Max]" + Reset)
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
