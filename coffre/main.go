package coffre

import (
	"fmt"
)

func OuvrirCoffre(objets []string, argent int) (bool, []string, int) {
	fmt.Println("\033[33mVous trouvez un coffre :\033[0m")
	fmt.Println("Que voulez-vous faire ?")
	fmt.Println("1. Ouvrir")
	fmt.Println("2. Partir")
	var choix string
	fmt.Scanln(&choix)
	if choix == "1" || choix == "Ouvrir" {
		contenu := "dedans : "
		if len(objets) > 0 {
			contenu += objets[0]
			for i := 1; i < len(objets); i++ {
				contenu += ", " + objets[i]
			}
		}
		if argent > 0 {
			if len(objets) > 0 {
				contenu += ", "
			}
			contenu += fmt.Sprintf("%d£", argent)
		}
		fmt.Println(contenu)
		fmt.Println("Appuie sur Entrée pour continuer...")
		fmt.Scanln()
		return true, objets, argent
	} else {
		fmt.Println("Vous décidez de ne pas ouvrir le coffre.")
		fmt.Println("Appuie sur Entrée pour continuer...")
		fmt.Scanln()
		return false, nil, 0
	}
}
