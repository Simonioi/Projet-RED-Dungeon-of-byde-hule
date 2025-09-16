package main

import (
	"fmt"
	"os"
	"os/exec"
)

var world = []string{
	"#####################################################",
	"#        ロ      #           #     #                #",
	"#                #        #  #  #  #      #         #",
	"#  ############  #        #  #  #  #  #####         #",
	"#             #  # ########  #  #  #   ⚗️  #         #",
	"##########    #  #        #  #  #  #      #         #",
	"#             #  #######-##  #  #  #####  #         #",
	"##########    #  #        #  #  #  #      #         #",
	"#             #  # $      #  #  #  #  #   #         #",
	"#       #        #        #  #  #  #  #   #   🛁    #",
	"#       #   ##########   ##  #  #  #  #   #         #",
	"#   ロ  #                 #     #     #   #         #",
	"#####################################################",
}

var playerX, playerY = 1, 6

func main() {
	for {
		clear()
		draw()
		var input string
		fmt.Scanln(&input)

		newX, newY := playerX, playerY
		switch input {
		case "z":
			newY--
		case "s":
			newY++
		case "q":
			newX--
		case "d":
			newX++
		case "x":
			fmt.Println("Meilleur qu'Ubisoft mdr")
			return
		}

		if world[newY][newX] != '#' {
			// Système d'énigme sur la porte '-'
			if world[newY][newX] == '-' {
				if enigmePorte() {
					playerX, playerY = newX, newY
				} else {
					fmt.Println("La porte reste fermée. Appuie sur Entrée pour réessayer...")
					fmt.Scanln()
				}
			} else {
				playerX, playerY = newX, newY
				// Déclencheur d'événement : combat uniquement aux coordonnées (x=4, y=11)
				if playerX == 4 && playerY == 11 {
					lancerCombat()
				}
			}
			playerX, playerY = newX, newY
			// Déclencheur d'événement : combat uniquement aux coordonnées (x=4, y=11)
			if playerX == 18 && playerY == 4 {
				lancerCombatSkelly()
			}
		}

	}
}

func lancerCombatSkelly() {
	fmt.Println("skelly se reveil pour te casser la gueule")
	fmt.Println("il est en fasse de toi pret a te malaxer bien comme il faut")
	content, err := os.ReadFile("skelly.txt")
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}
	fmt.Printf("\033[97m%s\033[0m\n", string(content))

	// ici logique reel de combat tah pokemon
	fmt.Println("(Simulation du combat...)")
	fmt.Println("Le combat est terminé !")
	fmt.Println("Appuie sur Entrée pour continuer...")
	playerY = 1
	playerX = 27
	fmt.Scanln()
}

// Fonction simulant le lancement d'un combat
func lancerCombat() {
	fmt.Println("\033[33mVous trouvez un coffre :\033[0m")
	fmt.Println("Le coffre s'agite, et s'ouvre revelant une gueule béante et des dents acérés, s'ouvre alors des yeux mauvais")
	fmt.Println("Un mimic dardant sa langue s'apprete a vous dévorer.")
	content, err := os.ReadFile("mimic.txt")
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}
	fmt.Printf("\033[38;5;130m%s\033[0m\n", string(content))

	fmt.Println("Roll for initiative")
	// ici logique reel de combat tah pokemon
	fmt.Println("(Simulation du combat...)")
	fmt.Println("Le combat est terminé !")
	fmt.Println("Appuie sur Entrée pour continuer...")
	fmt.Scanln()
}
func draw() {
	for y, line := range world {
		for x, char := range line {
			if x == playerX && y == playerY {
				fmt.Print("\033[34m¤\033[0m")
			} else if char == '#' {
				fmt.Print("\033[30m#\033[0m")
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func enigmePorte() bool {
	fmt.Println("Une porte devant vous barre votre chemin.")
	fmt.Println("Elle ne présente ni serrure ni poignée, mais une tête de sphinx.")
	fmt.Println("Ses yeux rouges s'allument et une voix résonne alors :")
	fmt.Println()
	fmt.Println("« Jamais je ne suis loin de mon autre jumelle,\non m'associe souvent au parfum vomitif d'une partie du corps\nqui n'est pas vraiment belle, localisée fort loin de l'organe olfactif. »")
	fmt.Println()
	fmt.Print("Quelle est votre réponse ? ")
	var reponse string
	fmt.Scanln(&reponse)
	if reponse == "chaussette" || reponse == "Chaussette" {
		fmt.Println("La porte s'ouvre lentement dans un grincement sinistre...")
		return true
	}
	fmt.Println("Mauvaise réponse !")
	return false
}
