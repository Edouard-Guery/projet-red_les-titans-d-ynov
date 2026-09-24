# 🔱 Les Titans d'Ynov - Projet RED

Projet de fin de module développé en Go (Golang). 
Un jeu de rôle (RPG) textuel en ligne de commande (CLI) se déroulant dans la mythologie grecque.

---

## 📖 Sommaire

1. [Présentation du Projet](#1-présentation-du-projet)
2. [Fonctionnalités Principales](#2-fonctionnalités-principales)
3. [Classes Jouables](#3-classes-jouables)
4. [Prérequis et Installation](#4-prérequis-et-installation)
5. [Guide des Commandes](#5-guide-des-commandes)
6. [Architecture Technique](#6-architecture-technique) 
7. [Crédits et Remerciements](#7-crédits-et-remerciements)

---

## 1. Présentation du Projet

**Les Titans d'Ynov** est un RPG intégralement jouable dans un terminal. Réalisé dans le cadre du **Projet RED** (Bachelor Cyber - YNOV Campus), ce jeu valide les concepts fondamentaux du langage Go : 
* Manipulation de structures de données (`struct`, `map`, `slice`).
* Gestion d'états et boucles de jeu (`for`, `switch`).
* Modularité et séparation du code en plusieurs fichiers.
* Interactivité en ligne de commande (intégration du module `go-prompt`).

Le but du jeu est d'affronter 14 monstres et figures mythologiques (de l'Orc jusqu'à Chronos) en gérant son équipement, son endurance et ses Oboles (monnaie du jeu).

---

## 2. Fonctionnalités Principales

* **Combats au tour par tour :** Gestion stricte de l'endurance pour les compétences, absorption des dégâts par l'armure (Défense / 2), probabilité de coups critiques (15%) et effets de statut (poison, soin continu).
* **Économie et Inventaire :** Accumulation d'Oboles pour acheter des potions, des extensions d'inventaire ou des artefacts puissants à la boutique.
* **Système de Forge (Artisanat) :** Les monstres lâchent des composants spécifiques (ex: *Dent d'Orc*, *Fiole de Venin*). Héphaïstos permet de les combiner pour forger des équipements classés en 3 Tiers.
* **Événements interactifs :** 
  * *L'Autel des Dieux* : Convertit les PV en Oboles ou permet de prier pour obtenir des statistiques permanentes.
  * *Le Casino* : Mini-jeu de roulette pour parier ses Oboles.
* **Autocomplétion :** Interface fluide utilisant la touche TAB pour suggérer les commandes.

---

## 3. Classes Jouables

Au lancement, le joueur choisit sa classe, ce qui détermine ses statistiques de départ et son approche du combat :

| Classe | PV | Endurance | Attaque | Défense | Spécialité |
| :--- | :---: | :---: | :---: | :---: | :--- |
| **Créature** | 70 | 50 | 12 | 5 | Équilibré. Polyvalent pour les débutants. |
| **Demi-Dieu** | 80 | 70 | 16 | 4 | Offensif. Frappe fort et possède beaucoup d'endurance, mais encaisse mal. |
| **Dieu** | 110 | 40 | 8 | 8 | Défensif (Tank). Très résistant, mais s'épuise rapidement en attaquant. |

*Note : Le niveau augmente automatiquement tous les 2 monstres vaincus, améliorant les statistiques et débloquant de nouvelles compétences.*

---

## 4. Prérequis et Installation

### Prérequis Système
* **Go** : Version 1.22 ou supérieure (nécessaire pour `math/rand/v2`).
* Un terminal standard supportant les couleurs ANSI et emojis.

### Installation

Clonez ce dépôt Git sur votre machine locale et placez-vous dans le dossier source :

    git clone https://github.com/Edouard-Guery/projet-red_les-titans-d-ynov
    cd projet-red_les-titans-d-ynov/src

Le jeu requiert une bibliothèque externe officielle pour gérer l'autocomplétion des commandes. Installez-la :

    go get github.com/c-bata/go-prompt

### Lancement

Compilez et lancez le jeu avec la commande suivante :

    go run .

---

## 5. Guide des Commandes

Une fois dans le jeu, une invite de commande `👉` attend vos instructions. **Appuyez sur TAB pour afficher les suggestions.**

*utilisation de l'outil **prompt** pour l'autocomplétion*

* `start` : Lance le combat contre le boss actuel de la trame principale.
* `entrainement` : Ouvre l'arène pour combattre un monstre au choix (farming).
* `inv` : Ouvre l'inventaire pour consommer des potions ou équiper des objets.
* `shop` : Accède à la Boutique Mythologique (Artefacts et extensions de sac).
* `potions` : Accède au magasin de Dionysos (Soins, poison, endurance).
* `forge` : Permet de combiner des composants pour créer des équipements.
* `autel` : Permet d'échanger des PV contre de l'or ou d'obtenir des bénédictions.
* `casino` : Mini-jeu de roulette pour parier ses Oboles.
* `info` : Affiche les statistiques complètes de votre héros.
* `clear` : Nettoie l'affichage du terminal.
* `quit` : Quitte proprement le jeu.

---

## 6. Architecture Technique

Le code est modulaire pour respecter les bonnes pratiques de développement (Clean Code) :

* **`main.go`** : Boucle principale, initialisation de `go-prompt` et routage des commandes.
* **`personnage.go`** : Structure du Joueur, leveling et gestion des compétences.
* **`combat.go` / `IA_monstre.go` / `Trainingfight.go`** : Logique des combats, algorithmes de dégâts et IA ennemie.
* **`monstres.go`** : Base de données des 14 boss, de leurs stats et loots.
* **`inventaire.go` / `equipement.go` / `effets.go`** : Gestion du sac, équipement des objets et modificateurs de stats.
* **`marchand.go` / `shop_potions.go` / `items.go`** : Boutiques et transactions.
* **`forgeron.go`** : Système de recettes et vérification des prérequis d'artisanat.
* **`autel.go` / `casino.go`** : Logique des mini-jeux et événements extérieurs.

---

## 7. Crédits et Remerciements

* **Développement :** Alexis, Clovis et Edouard
* **Dépendance Technique :** `c-bata/go-prompt`

**Remerciements Spéciaux :**
Un grand merci à nos mentors, **Enzo** et **Nathan**, pour leur accompagnement et leur suivi.
