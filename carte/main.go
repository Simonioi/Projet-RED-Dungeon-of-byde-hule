package carte

import (
	"dungeon/combat"
	"dungeon/combatcthulhu"
	"dungeon/combatskelly"
	"dungeon/cthulhu"
	"dungeon/mimic"
	"dungeon/personnage"
	"dungeon/skelly"
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

func Start(p personnage.Character) {
	player = p
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
			if world[newY][newX] == '-' {
				if enigmePorte() {
					// Bonne réponse : avancer
					playerX, playerY = newX, newY
				} else {
					// Mauvaise réponse : reculer à Y=8, X=24
					playerY = 8
					playerX = 24
					fmt.Println("La porte reste fermée. Vous reculez. Appuie sur Entrée pour réessayer...")
					fmt.Scanln()
				}
			} else {
				// Déplacement normal
				playerX, playerY = newX, newY
				// Combat qui start en x=4 et y=11
				if playerX == 4 && playerY == 11 {
					lancerCombat()
					mimic := mimic.Mimic()
					combat.Battle(&player, &mimic)
					fmt.Println("Le combat est terminé !")
					fmt.Println("Appuie sur Entrée pour continuer...")
					fmt.Scanln()
				}
				// Combat contre Skelly en x=18 et y=4
				if playerX == 18 && playerY == 4 {
					lancerCombatSkelly()
					skellyMonster := skelly.Skelly()
					combatskelly.Battle(&player, &skellyMonster)
					if player.CurrentHP > 0 && skellyMonster.CurrentHP <= 0 {
						playerY = 1
						playerX = 27
					}
				}
				// Combat contre Cthulhu en x=43 et y=1
				if playerX == 43 && playerY == 1 {
					lancerCombatCthulhu()
					cthulhuMonster := cthulhu.Cthulhu()
					combatcthulhu.Battle(&player, &cthulhuMonster)
					if player.CurrentHP > 0 && cthulhuMonster.CurrentHP <= 0 {
						playerY = 1
						playerX = 43
					}
				}

			}
		}

	}
}

var player personnage.Character

func lancerCombatCthulhu() {
	fmt.Println("\033[37mdans son bain ya l autre gland\033[0m")
	content, err := os.ReadFile("cltuuululuuuuu.txt")
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}
	fmt.Printf("\033[32m%s\033[0m\n", string(content))
	fmt.Println("Prépare-toi à affronter Cthulhu !")
}

func lancerCombatSkelly() {
	fmt.Println("\033[33mDevant vous une horde de squelette. Un sans jambe rampe sans but, un autre sans bras cours en rond et un dernier semble plus menaçant armé d'une épée rouillé et d'un casque.\033[0m")
	fmt.Println()
	fmt.Println("\033[33mUn crâne au vous regarde et dis :\033[0m")
	fmt.Println("«Kakakaka te revoila Len..., oh! Non. Qui es tu ? »")
	fmt.Println("\033[31mATTAQUE LE !!!\033[0m")
	content, err := os.ReadFile("skelly.txt")
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}
	fmt.Printf("\033[97m%s\033[0m\n", string(content))

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
