package carte

import (
	"dungeon/coffre"
	"dungeon/combat"
	"dungeon/combatcthulhu"
	"dungeon/combatskelly"
	"dungeon/cthulhu"
	"dungeon/inventaire/stock"
	"dungeon/marchand"
	"dungeon/mimic"
	"dungeon/personnage"
	"dungeon/personnage/lvlup"
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
	"#             #  #$       #  #  #  #  #   #         #",
	"#       #        #        #  #  #  #  #   #   🛁    #",
	"#       #   ##########   ##  #  #  #  #   #         #",
	"#   ロ  #                 #     #     #___#         #",
	"#####################################################",
}

var playerX, playerY = 1, 6

func Start(p personnage.Character) {
	player = p
	coffreOuvert := false
	combatMimicFait := false
	combatSkellyFait := false
	for {
		clear()
		draw()
		var input string
		fmt.Scanln(&input)

		for _, move := range input {
			newX, newY := playerX, playerY
			switch move {
			case 'z':
				newY--
			case 's':
				newY++
			case 'q':
				newX--
			case 'd':
				newX++
			case 'x':
				fmt.Println("Meilleur qu'Ubisoft mdr")
				return

			default:
				continue
			}
			// Activation du marchand UNIQUEMENT en x=18 et y=8
			if playerX == 18 && playerY == 8 {
				marchand.ActiverMarchand(player.Inventory)
				// Déplace le joueur à la sortie du marchand
				playerX = 20
				playerY = 8
				break
			}

			if world[newY][newX] == '#' {
				// Mur, on arrête la séquence
				break
			}
			if world[newY][newX] == '-' {
				if enigmePorte() {
					playerX, playerY = newX, newY
				} else {
					playerY = 8
					playerX = 24
					fmt.Println("La porte reste fermée. Vous reculez. Appuie sur Entrée pour réessayer...")
					fmt.Scanln()
				}
				// On arrête la séquence après une énigme
				break
			}

			// Déplacement normal
			playerX, playerY = newX, newY

			if playerY == 2 && (playerX == 9 || playerX == 10) && !coffreOuvert {
				ouvert, _, _ := coffre.OuvrirCoffre([]string{"parchemin de boule de feu"}, 20)
				if ouvert {
					player.Inventory.AddMoney(20)
					player.Inventory.AddItem(stock.FireScroll)
				}
				if ouvert {
					coffreOuvert = true
				}
				// On arrête la séquence après un événement
				break
			}
			if playerY == 10 && (playerX == 38 || playerX == 39 || playerX == 40) && !coffreOuvert {
				ouvert, _, _ := coffre.Ouvrirbibliotheque([]string{"parchemin de boule de feu"}, 20)
				if ouvert {
					player.Inventory.AddItem(stock.CitrusLimus)
					player.Inventory.AddItem(stock.SolanumTuberosumPatate)
					player.Inventory.AddItem(stock.ZingiberOfficinale)
				}
				if ouvert {
					coffreOuvert = true
				}
				// On arrête la séquence après un événement
				break
			}
			// Combat qui start en x=4 et y=11 (Mimic)
			if (playerX == 4 || playerX == 5) && playerY == 11 && !combatMimicFait {
				lancerCombat()
				mimic := mimic.Mimic()
				combat.Battle(&player, &mimic)
				if player.CurrentHP > 0 && mimic.CurrentHP <= 0 {
					player.Inventory.AddMoney(20)
					fmt.Println("\033[33mDans sa gueule béante vous trouvez 20 pièces d'or !\033[0m")
				}
				fmt.Println("Le combat est terminé !")
				fmt.Println("Appuie sur Entrée pour continuer...")
				fmt.Scanln()
				combatMimicFait = true
				break
			}
			// Combat contre Skelly en x=18 et y=4
			if playerX == 18 && playerY == 4 && !combatSkellyFait {
				lancerCombatSkelly()
				skellyMonster := skelly.Skelly()
				combatskelly.Battle(&player, &skellyMonster)
				if player.CurrentHP > 0 && skellyMonster.CurrentHP <= 0 {
					playerY = 1
					playerX = 27
				}
				fmt.Println("Le combat est terminé !")
				fmt.Println("Appuie sur Entrée pour continuer...")
				fmt.Scanln()
				combatSkellyFait = true
				break
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
				break
			}
		}
	}
}

var player personnage.Character

func lancerCombatCthulhu() {
	fmt.Println("\033[33mVous entrez alors dans une longue piece éclairé par des bougies parfumé au jasmins, des accords éthéré de harpes vienne d'une sirène accompagnant le chant aigu d'une large créatures aux tentacules impresionnantes qui se savonne dans un bain d'eau fumante et de pétale de rose. \033[0m")
	fmt.Printf("«Qui ose me déranger dans mon bain ?»\n")
	fmt.Println("\033[33mLe grand Cthulhu se tourne vers vous\033[0m")
	fmt.Println("Roll for initiative")
	content, err := os.ReadFile("cltuuululuuuuu.txt")
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}
	fmt.Printf("\033[32m%s\033[0m\n", string(content))
	fmt.Println("Prépare-toi à affronter Cthulhu !")
}

func lancerCombatSkelly() {
	fmt.Println("\033[33mDevant vous une horde de squelette, un sans jambe rampe sans but, un autre sans bras cours en rond et un dernier semble plus menaçant armé d'une épée rouillé et d'un casque.\033[0m")
	fmt.Println()
	fmt.Println("\033[33mUn crâne au vous regarde et dis :\033[0m")
	fmt.Println("«Kakakaka te revoila Len..., oh! Non. Qui es tu ?»")
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
		lvlup.LevelUp(&player)
		fmt.Println("ding ding ! La caisse des donjons est heureuse de vous annoncer que vous avez gagné un niveau !")
		return true
	}
	fmt.Println("Mauvaise réponse !")
	return false
}

//ooooo
