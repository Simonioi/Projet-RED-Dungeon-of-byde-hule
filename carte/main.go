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
			playerX, playerY = newX, newY

			// Déclencheur d'événement : combat si le joueur est sur une case spéciale
			if world[playerY][playerX] == '$' {
				lancerCombat()
			}
		}
	}
}

// Fonction simulant le lancement d'un combat
func lancerCombat() {
	fmt.Println("Un combat commence ! Prépare-toi !")
	// Ici, tu pourrais appeler la logique de combat réelle
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

//systeme d evenement
