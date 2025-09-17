package main

import (
	"dungeon/carte"
	"dungeon/personnage"
	"fmt"
	"strings"
)

func main() {
	var name, class string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&name)
	fmt.Print("Choisissez la classe (barbare/mage) : ")
	fmt.Scanln(&class)

	var player personnage.Character
	switch strings.ToLower(class) {
	case "barbare":
		player = personnage.CreateBarbarian(name)
	case "mage":
		player = personnage.CreateMage(name)
	default:
		fmt.Println("Classe inconnue, création d'un barbare par défaut.")
		player = personnage.CreateBarbarian(name)
	}

	fmt.Printf("Personnage créé : %s le %s (PV: %d/%d, PM: %d/%d)\n", player.Name, player.Class, player.CurrentHP, player.MaxHP, player.CurrentMP, player.MaxMP)
	carte.Start(player)
}
