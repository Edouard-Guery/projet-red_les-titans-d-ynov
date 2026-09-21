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
	switch nom {
	case "Cuirasse Sauvage":
		p.EquiperCuirasseSauvage()
		fmt.Println("Équipé : Cuirasse Sauvage (+10 Défense, +20 PV)")
	case "Regard Infernal":
		p.EquiperRegardInfernal()
		fmt.Println("Équipé : Regard Infernal (+15 Attaque)")
	case "Brassard de l'Enfer Vivant":
		p.EquiperBrassardEnferVivant()
		fmt.Println("Équipé : Brassard de l'Enfer Vivant (+10 Attaque, +5 Défense)")
	case "Lame du Labyrinthe":
		p.EquiperLameLabyrinthe()
		fmt.Println("Équipé : Lame du Labyrinthe (+25 Attaque)")
	case "Manteau Empoisonné":
		p.EquiperManteauEmpoisonne()
		fmt.Println("Équipé : Manteau Empoisonné (+15 Défense, +5 Attaque)")
	case "Regard Mortel":
		p.EquiperRegardMortel()
		fmt.Println("Équipé : Regard Mortel (+30 Attaque, +20 PV)")
	case "Égide des Mers":
		p.EquiperEgideMers()
		fmt.Println("Équipé : Égide des Mers (+40 Défense, +50 PV)")
	case "Foudre Infernale":
		p.EquiperFoudreInfernale()
		fmt.Println("Équipé : Foudre Infernale (+50 Attaque)")
	case "Armure du Colosse":
		p.EquiperArmureColosse()
		fmt.Println("Équipé : Armure du Colosse (+60 Défense, +100 PV)")
	default:
		return false
	}

	p.Inventaire[nom]--
	if p.Inventaire[nom] <= 0 {
		delete(p.Inventaire, nom)
	}

	return true
}