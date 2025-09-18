package combatcthulhu

import (
	"bufio"
	"dungeon/cthulhu"
	menuinventaire "dungeon/inventaire/openinventory"
	"dungeon/personnage"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)
// script de combat contre Cthulhu, (oui, c'est pas possible de gagner, deso :/)
func enAttack(enemy *cthulhu.Monster, player *personnage.Character, turn int) {
	var selectedAttack []cthulhu.Attack

	switch turn {
	case 1:
		selectedAttack = enemy.Attack1
	case 2:
		selectedAttack = enemy.Attack2
	case 3:
		selectedAttack = enemy.Attack3
	case 4:
		selectedAttack = enemy.Attack4
	case 5:
		selectedAttack = enemy.Attack5
	default:
		selectedAttack = enemy.Attack1
	}

	totalDamage := 0
	for _, attack := range selectedAttack {
		if rand.Float64() <= attack.HitChance {
			totalDamage += attack.Damage
			fmt.Printf("%s utilise %s et inflige %d dégâts à %s !\n", enemy.Name, attack.Name, attack.Damage, player.Name)
		} else {
			fmt.Printf("%s utilise %s mais rate son attaque !\n", enemy.Name, attack.Name)
		}
	}

	if totalDamage < 0 {
		totalDamage = 0
	}
	player.CurrentHP -= totalDamage
	if player.CurrentHP < 0 {
		player.CurrentHP = 0
	}
}
// fonction de choix d'action du joueur
func ChooseAction() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nQue veux-tu faire ?")
	fmt.Println("1. Attaque de base")
	fmt.Println("2. Attaque puissante")
	fmt.Println("3. Ouvrir l'inventaire")
	fmt.Println("4. Activer capacité spéciale")
	fmt.Print("Choix : ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || (choice < 1 || choice > 4) {
		fmt.Println("Choix invalide, attaque de base utilisée.")
		return 1
	}
	return choice
}

func ExecuteAttack(attackerName string, attacks []personnage.Attack, defenderName string, defenderHP *int) {
	totalDamage := 0
	for _, attack := range attacks {
		if rand.Float64() <= attack.HitChance {
			totalDamage += attack.Damage
			fmt.Printf("%s utilise %s et inflige %d dégâts à %s !\n", attackerName, attack.Name, attack.Damage, defenderName)
		} else {
			fmt.Printf("%s utilise %s mais rate son attaque !\n", attackerName, attack.Name)
		}
	}
	if totalDamage < 0 {
		totalDamage = 0
	}
	*defenderHP -= totalDamage
	if *defenderHP < 0 {
		*defenderHP = 0
	}
}
// fonction de combat contre cthulhu qui implémente un script de tour afin d'avoir des attaques précises a chaque tour.
func Battle(player *personnage.Character, enemy *cthulhu.Monster) {
	turn := 1

	for player.CurrentHP > 0 && enemy.CurrentHP > 0 {
		fmt.Println("\n--- Tour du joueur ---")
		fmt.Printf("PV Joueur: %d | PV Ennemi: %d\n", player.CurrentHP, enemy.CurrentHP)

		player.TickAttackBoost()

		if len(player.Capacité) > 0 && player.Capacité[0].Duration > 0 {
			fmt.Printf("\033[35mBoost actif: %s (%d tours restants)\033[0m\n", player.Capacité[0].Name, player.Capacité[0].Duration)
		}

		
        choice := ChooseAction()
		switch choice {
		case 1:
			ExecuteAttack(player.Name, player.Attacks1, enemy.Name, &enemy.CurrentHP)
		case 2:
			ExecuteAttack(player.Name, player.Attacks2, enemy.Name, &enemy.CurrentHP)
		case 3:
			used := menuinventaire.OpenInventory(player.Inventory, player)
			if !used {
				continue
			}
		case 4:
			ExecuteAttack(player.Name, player.Capacité, enemy.Name, &enemy.CurrentHP)
			if player.Class == "Barbare" {
				player.ActivateAttackBoost(personnage.Attack{Name: "I WOULD LIKE TO RAGE", TempDamageBoost: 4, TempHealthBoost: 2, Duration: 6})
				continue
			}
// cas spéciale de la capacité spé. le barbare a un boost, et le mage une boule de feu.
		}
		// Appliquer les dégâts d’un objet offensif si présent
		if player.PendingDamage > 0 {
			fmt.Println(player.PendingDamageText)
			enemy.CurrentHP -= player.PendingDamage
			if enemy.CurrentHP < 0 {
				enemy.CurrentHP = 0
			}
			player.PendingDamage = 0
			player.PendingDamageText = ""
		}

		if enemy.CurrentHP <= 0 {
			fmt.Println(enemy.Name, "est vaincu !")
			break
		}

		fmt.Println("\n--- Tour de l'ennemi ---")
		enAttack(enemy, player, turn)

		if player.CurrentHP <= 0 {
			fmt.Println(player.Name, "est vaincu !")
			break
		}

		turn++
	}

	fmt.Printf("PV Joueur: %d | PV Ennemi: %d\n", player.CurrentHP, enemy.CurrentHP)
	if player.CurrentHP <= 0 {
		clearScreen()
		fmt.Println("Tu es mort.")
		fmt.Println("Appuie sur Entrée pour quitter...")
		fmt.Scanln()
		os.Exit(0)
	}
}

// Fonction utilitaire pour clear l'écran
func clearScreen() {
	if os.Getenv("OS") == "Windows_NT" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}
