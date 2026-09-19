package main

type Equipement struct {
	Nom         string
	Source      string
	Description string
	Prix        int
}

func BoutiqueMarchand() []Equipement {
	return []Equipement{
		{
			Nom:         "Bouclier Éclair",
			Source:      "Zeus",
			Description: "Augmente massivement la défense.",
			Prix:        500,
		},
		{
			Nom:         "Pendentif des Moustiques",
			Source:      "Poséidon",
			Description: "Restaure une très grande quantité de PV à chaque tour.",
			Prix:        450,
		},
		{
			Nom:         "Faux d'Éternité",
			Source:      "Chronos / Hadès",
			Description: "Inflige de très lourds dégâts sur la durée à l'adversaire.",
			Prix:        600,
		},
		{
			Nom:         "Boîte de Pandore",
			Source:      "Pandore",
			Description: "50 % de chances de multiplier tes dégâts par 2 ; sinon, l'adversaire double ses dégâts.",
			Prix:        350,
		},
		{
			Nom:         "Casque de la Sagesse",
			Source:      "Athéna",
			Description: "20 % de chances (1 sur 5) d'esquiver totalement une attaque.",
			Prix:        400,
		},
		{
			Nom:         "Armure d'Héphaïstos",
			Source:      "Achille",
			Description: "Confère une résistance très élevée à tous les coups.",
			Prix:        300,
		},
		{
			Nom:         "Carte des Archipels",
			Source:      "Thésée",
			Description: "Augmente la puissance de toutes tes attaques.",
			Prix:        250,
		},
		{
			Nom:         "Frénésie des Travaux",
			Source:      "Hercule",
			Description: "Te donne un tour d'action supplémentaire.",
			Prix:        350,
		},
		{
			Nom:         "Sac Magique (Kibisis)",
			Source:      "Persée",
			Description: "Débloque un emplacement d'équipement tous les 5 boss vaincus.",
			Prix:        300,
		},
		{
			Nom:         "Regard de la Gorgone",
			Source:      "Méduse",
			Description: "2 % de chances (1 sur 50) de tuer instantanément le boss.",
			Prix:        500,
		},
		{
			Nom:         "Gueule Enflammée",
			Source:      "Cerbère",
			Description: "Brûle l'ennemi en lui infligeant de légers dégâts à chaque tour.",
			Prix:        200,
		},
		{
			Nom:         "Hache à Double Tranchant",
			Source:      "Le Minotaure",
			Description: "Multiplie tes dégâts par 2, mais l'adversaire attaque 2 fois par tour.",
			Prix:        250,
		},
		{
			Nom:         "Gourde de Régénération",
			Source:      "L'Hydre de Lerne",
			Description: "Restaure tes PV petit à petit à chaque tour.",
			Prix:        180,
		},
		{
			Nom:         "Bouclier en Bois Renforcé",
			Source:      "L'Orc",
			Description: "Augmente la défense globale.",
			Prix:        100,
		},
	}
}

func BouclierEclair() {

}
