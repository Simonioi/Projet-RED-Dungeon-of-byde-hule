package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Name        string
	HealthBoost int
	Damage      int
	HealAmount  int
}

type Attack struct {
	Name      string
	Damage    int
	HitChance float64
}

type Character struct {
	Name      string
	Class     string
	MaxHP     int
	CurrentHP int
	MaxMP     int
	CurrentMP int
	Inventory []Item
	Attacks   []Attack
}

func CreateBarbarian(name string) Character {
	armor := Item{Name: "Armure de barbare", HealthBoost: 2}
	sword := Item{Name: "Épée", Damage: 3}
	potionMinor := Item{Name: "Potion mineure", HealAmount: 3}
	potionMajor := Item{Name: "Potion majeure", HealAmount: 5}

	attacks := []Attack{
		{Name: "Attaque rapide", Damage: sword.Damage, HitChance: 0.95},
		{Name: "Attaque puissante", Damage: sword.Damage + 1, HitChance: 0.80},
	}

	baseHP := 8 + armor.HealthBoost

	return Character{
		Name:      name,
		Class:     "Barbare",
		MaxHP:     baseHP,
		CurrentHP: baseHP,
		MaxMP:     4,
		CurrentMP: 4,
		Inventory: []Item{armor, sword, potionMinor, potionMajor},
		Attacks:   attacks,
	}
}

func CreateMage(name string) Character {
	robe := Item{Name: "Robe enchantée", HealthBoost: 5}
	staff := Item{Name: "Bâton de mage", Damage: 2}
	potionMinor := Item{Name: "Potion mineure", HealAmount: 3}
	potionMajor := Item{Name: "Potion majeure", HealAmount: 5}

	attacks := []Attack{
		{Name: "Coup de bâton", Damage: staff.Damage, HitChance: 0.70},
		{Name: "Missile magique", Damage: 4, HitChance: 0.95},
	}

	baseHP := 5 + robe.HealthBoost

	return Character{
		Name:      name,
		Class:     "Mage",
		MaxHP:     baseHP,
		CurrentHP: baseHP,
		MaxMP:     4,
		CurrentMP: 4,
		Inventory: []Item{robe, staff, potionMinor, potionMajor},
		Attacks:   attacks,
	}
}

func ShowInventory(player Character) {
	if len(player.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	fmt.Println("Inventaire :")
	for i, item := range player.Inventory {
		fmt.Printf("  %d. %s", i+1, item.Name)
		if item.HealthBoost > 0 {
			fmt.Printf(" (+%d PV)", item.HealthBoost)
		}
		if item.Damage > 0 {
			fmt.Printf(" (+%d Dégâts)", item.Damage)
		}
		if item.HealAmount > 0 {
			fmt.Printf(" (Soigne %d PV)", item.HealAmount)
		}
		fmt.Println()
	}
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

	inventoryOpen := false
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "inventaire":
			inventoryOpen = !inventoryOpen
			if inventoryOpen {
				fmt.Println("Inventaire ouvert.")
				ShowInventory(player)
			} else {
				fmt.Println("Inventaire fermé.")
			}
		case "quitter":
			fmt.Println("Fin du jeu.")
			return
		default:
			fmt.Println("Commande inconnue.")
		}
	}
}

//salut
