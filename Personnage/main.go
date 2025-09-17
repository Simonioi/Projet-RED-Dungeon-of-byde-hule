package personnage

import (
	"fmt"
	"strings"

	"dungeon/inventaire"
)

type Attack struct {
	Name      string
	Damage    int
	HitChance float64
}

type Character struct {
	Name              string
	Class             string
	MaxHP             int
	CurrentHP         int
	MaxMP             int
	CurrentMP         int
	Inventory         *inventaire.Inventory
	Attacks1          []Attack // Attaques de base
	Attacks2          []Attack // Attaques puissante
	PendingDamage     int
	PendingDamageText string
}

func CreateBarbarian(name string) Character {
	// Créer l'inventaire du barbare avec les objets par défaut
	inventory := inventaire.CreateBarbarianInventory()

	// Récupérer les objets pour calculer les stats
	items := inventory.GetItems()
	var armor, sword inventaire.Item
	for _, item := range items {
		if item.Name == "Armure de barbare" {
			armor = item
		}
		if item.Name == "Épée" {
			sword = item
		}
	}

	attacks1 := []Attack{
		{Name: "Attaque rapide", Damage: sword.Damage, HitChance: 0.95},
	}
	attacks2 := []Attack{
		{Name: "Attaque puissante", Damage: sword.Damage + 1, HitChance: 0.80},
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
		Attacks1:          attacks1,
		Attacks2:          attacks2,
		PendingDamage:     0,
		PendingDamageText: "",
	}
}

func CreateMage(name string) Character {
	// Créer l'inventaire du mage avec les objets par défaut
	inventory := inventaire.CreateMageInventory()

	// Récupérer les objets pour calculer les stats
	items := inventory.GetItems()
	var robe, staff inventaire.Item
	for _, item := range items {
		if item.Name == "Robe enchantée" {
			robe = item
		}
		if item.Name == "Bâton de mage" {
			staff = item
		}
	}

	attacks1 := []Attack{
		{Name: "Coup de bâton", Damage: staff.Damage, HitChance: 0.70},
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
		Attacks1:          attacks1,
		Attacks2:          attacks2,
		PendingDamage:     0,
		PendingDamageText: "",
	}
}

func ShowInventory(player Character) {
	player.Inventory.ShowInventory()
}

func main() {
	var name, classChoice string

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
