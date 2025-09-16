package main

import (
	"fmt"
	"dungeon/personnage"
	"dungeon/mimic"
	"dungeon/combat"
)

func main() {
	var name, class string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&name)
	fmt.Print("Choisissez la classe (barbare/mage) : ")
	fmt.Scanln(&class)

	var player *personnage.Character
	switch class {
	case "barbare":
		p := personnage.CreateBarbarian(name)
		player = &p
	case "mage":
		p := personnage.CreateMage(name)
		player = &p
	default:
		fmt.Println("Classe inconnue, création d'un barbare par défaut.")
		p := personnage.CreateBarbarian(name)
		player = &p
	}


	var enemy *mimic.Monster
	fmt.Printf("Un combat commence \n")
	combat.Battle(player, enemy)
}
