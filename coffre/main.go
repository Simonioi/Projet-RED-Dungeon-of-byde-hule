package coffre

import (
	"dungeon/inventaire"
	"dungeon/inventaire/item"
	"fmt"
)

func OuvrirCoffre(inv *inventaire.Inventory, item item.Item) (bool, []string, int) {
	fmt.Println("\033[33mVous trouvez un coffre :\033[0m")
	fmt.Println("Que voulez-vous faire ?")
	fmt.Println("1. Ouvrir")
	fmt.Println("2. Partir")
	var choix string
	fmt.Scanln(&choix)
	if choix == "1" || choix == "Ouvrir" {
		contenu := "dedans : "
		if item.Name != "" {
			contenu += item.Name
		}
		if item.Quantity > 0 {
			if item.Name != "" {
				contenu += ", "
			}
			contenu += fmt.Sprintf("%d£", item.Quantity)
		}
		fmt.Println(contenu)
		fmt.Scanln()
		fmt.Println("Vous ouvrez le coffre et récupérez son contenu.")
		fmt.Println("Appuie sur Entrée pour continuer...")
		fmt.Scanln()
		return true, []string{item.Name}, item.Quantity
	} else {
		fmt.Println("Vous décidez de ne pas ouvrir le coffre.")
		fmt.Println("Appuie sur Entrée pour continuer...")
		fmt.Scanln()
		return false, nil, 0
	}
}
