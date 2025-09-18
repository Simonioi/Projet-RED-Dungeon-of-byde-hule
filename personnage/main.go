package personnage

import (
	"dungeon/inventaire"
	"dungeon/inventaire/item"
	"fmt"
	"strings"
)

type Attack struct {
	Name            string
	Damage          int
	TempDamageBoost int
	TempHealthBoost int
	HitChance       float64
	Duration        int
}

// gestion de la rage avec boost tempo et durée
func (c *Character) ActivateAttackBoost(a Attack) {
	if a.TempHealthBoost > 0 {
		c.MaxHP += a.TempHealthBoost
		c.CurrentHP += a.TempHealthBoost
	}
	if a.TempDamageBoost > 0 {
		for i := range c.Attacks1 {
			c.Attacks1[i].Damage += a.TempDamageBoost
		}
		for i := range c.Attacks2 {
			c.Attacks2[i].Damage += a.TempDamageBoost
		}
	}
	c.Capacité = []Attack{a}
}

// gestion de la durée et retrait des boosts
func (c *Character) TickAttackBoost() {
	if len(c.Capacité) > 0 {
		a := &c.Capacité[0]
		if a.Duration > 0 {
			a.Duration--
			if a.Duration == 0 {
				// Retirer les boosts
				if a.TempHealthBoost > 0 {
					c.MaxHP -= a.TempHealthBoost
					if c.CurrentHP > c.MaxHP {
						c.CurrentHP = c.MaxHP
					}
				}
				if a.TempDamageBoost > 0 {
					for i := range c.Attacks1 {
						c.Attacks1[i].Damage -= a.TempDamageBoost
					}
					for i := range c.Attacks2 {
						c.Attacks2[i].Damage -= a.TempDamageBoost
					}
				}
				c.Capacité = nil
			}
		}
	}
}

type Character struct {
	Name              string
	Class             string
	MaxHP             int
	CurrentHP         int
	MaxMP             int
	CurrentMP         int 
	Inventory         *inventaire.Inventory // Inventaire du personnage
	PendingDamage     int // Dégâts en attente d’application (objets offensifs)
	PendingDamageText string // Texte descriptif des dégâts en attente
	Attacks1          []Attack // Attaques de base
	Attacks2          []Attack // Attaques puissante
	Capacité          []Attack // Attaque spéciale

}

func CreateBarbarian(name string) Character {
	// Créer l'inventaire du barbare avec les objets par défaut
	inventory := inventaire.CreateBarbarianInventory()

	// Récupérer les objets pour calculer les stats
	items := inventory.GetItems()
	var armor, sword item.Item
	for _, it := range items {
		if it.Name == "Armure de barbare" {
			armor = it
		}
		if it.Name == "Épée" {
			sword = it
		}
	}
// attaque du barbare
	attacks1 := []Attack{
		{Name: "Attaque rapide", Damage: sword.BaseDamage, HitChance: 0.95},
	}
	attacks2 := []Attack{
		{Name: "Attaque puissante", Damage: sword.BaseDamage + 1, HitChance: 0.80},
	}

	baseHP := 8 + armor.HealthBoost

	return Character{
		Name:              name,
		Class:             "Barbare",
		MaxHP:             baseHP,
		CurrentHP:         baseHP,
		MaxMP:             4,
		CurrentMP:         4,
		Inventory:         inventory,
		PendingDamage:     0,
		PendingDamageText: "",
		Attacks1:          attacks1,
		Attacks2:          attacks2,
	}
}

func CreateMage(name string) Character {
	// Créer l'inventaire du mage avec les objets par défaut
	inventory := inventaire.CreateMageInventory()

	// Récupérer les objets pour calculer les stats
	items := inventory.GetItems()
	var robe, staff item.Item
	for _, it := range items {
		if it.Name == "Robe enchantée" {
			robe = it
		}
		if it.Name == "Bâton de mage" {
			staff = it
		}
	}
// attaque du mage
	attacks1 := []Attack{
		{Name: "Coup de bâton", Damage: staff.BaseDamage, HitChance: 0.70},
	}
	attacks2 := []Attack{
		{Name: "Missile magique", Damage: 4, HitChance: 0.95},
	}

	baseHP := 5 + robe.HealthBoost

	return Character{
		Name:              name,
		Class:             "Mage",
		MaxHP:             baseHP,
		CurrentHP:         baseHP,
		MaxMP:             4,
		CurrentMP:         4,
		Inventory:         inventory,
		PendingDamage:     0,
		PendingDamageText: "",
		Attacks1:          attacks1,
		Attacks2:          attacks2,
	}
}

func ShowInventory(player Character) {
	player.Inventory.ShowInventory()
}
// on a casé ici l'intro de la mission faute de place
func main() {
	var name, classChoice, missionChoice string

	// Intro mission
	fmt.Println("Contrat de Repérage et nettoyage de Donjon niveau 1, le Donjon de Byde Hulerécompense : 1000 pièces (taxe non comprise, se référer aux barème de taxation des aventuriers)")
	fmt.Println("Bonjour Aventurier, votre Contrat, si toutefois vous l'acceptez, sera de nettoyer un donjon, pour le compte de notre Contractuel le Seigneur de Maaschyn.")
	fmt.Print("Donjon pouvant inclure : monstre mineur de type revenant ou creation nécromantique, piège franchement visible même pour un aveugle, énigme ou puzzle de difficulté 3 à 6 ans (pièces manquante non incluse), BBEG de niveau Pathétique à Ridicule.")
	fmt.Println("information non contractuelle du donjon :\n La caisse des donjons décline toute responsabilité en cas de mort violente, de mutilation, perte de santé mentale, téléportation dans des univers parallèle, arnaque de gobelin, digestion par un cube gélatineux, sacrifice par des cultistes zélés.\n La Caisse des Donjons se reserve le droit de ne pas payer l'aventurier en cas de non présentation du contrat, et/ou preuve de mort du BBEG.")
	fmt.Println("-Bob dit 'je vais vous Rosser', aventurier vétéran de niveau 20, directeur par interim de l'attribution des contrats aux aventurier débutants un poil suicidaire, désespéré et ambitieux à la Caisse des Donjons.")
	fmt.Println("Ce message ne s'autodétruira pas mais sera probablement perdu au fond de votre sac.")
	fmt.Println("\nAcceptez-vous la mission ? (oui/non)")
	fmt.Print("> ")
	fmt.Scanln(&missionChoice)
	if strings.ToLower(missionChoice) != "oui" {
		fmt.Println("Vous refusez la mission. Le jeu s'arrête ici.")
		return
	}

	fmt.Println("Bienvenue dans le jeu !")
	fmt.Print("Choisissez votre classe (barbare/mage) : ")
	fmt.Scanln(&classChoice)

	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&name)

	var player Character
	switch strings.ToLower(classChoice) {
	case "barbare":
		player = CreateBarbarian(name)
	case "mage":
		player = CreateMage(name)
	default:
		fmt.Println("Classe inconnue.")
		return
	}

	fmt.Printf("Bienvenue %s le %s ! PV: %d/%d, PM: %d/%d\n",
		player.Name, player.Class, player.CurrentHP, player.MaxHP, player.CurrentMP, player.MaxMP)
}

//truc
