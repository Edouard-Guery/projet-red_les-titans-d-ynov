package main
import "math/rand/v2"
func (p *Personnage)BouclierEclair()  {
	p.Defense+=50
}

func (p *Personnage)PendentifVenin(){
	p.PV+=20
}

func (p *personnage) FauxEternite()  {
	p.Attaque+=30
}

func  BoitePandore(p*Personnage,m*Monstre)  {
	if nombre := rand.N(3) <2{
		p*Attaque= Attaque*2
	} else {
		m* Attaque=Attaque*2
	}
}

func (p *Personnage) ArmureHephaistos() {
	p*Defense+=Defense
}

func (p * Personnage)CarteArchipel(){
	p*Attaque+=10
}

func (p *Personnage)FrenesieTravaux(){

}


Frénésie des Travaux : Donne un tour supplémentaire.

Sac Magique (Kibisis) : Débloque un emplacement d'équipement.

Regard de la Gorgone : Chance d'élimination instantanée.

Gueule Enflammée : Brûlure et dégâts continus.

Hache à Double Tranchant : Dégâts x2, attaques ennemie x2.

Gourde de Régénération : Soins progressifs chaque tour.

Bouclier en Bois Renforcé : Augmente la défense globale