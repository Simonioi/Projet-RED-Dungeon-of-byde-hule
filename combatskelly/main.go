package combatskelly

import (
	"bufio"
	menuinventaire "dungeon/inventaire/openinventory"
	"dungeon/personnage"
	"dungeon/skelly"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func enAttack(enemy *skelly.Monster, player *personnage.Character) {
	totalDamage := 0
	for _, attack := range enemy.Attacks {
		totalDamage += attack.Damage
	}
	if totalDamage < 0 {
		totalDamage = 0
	}
	player.CurrentHP -= totalDamage
	fmt.Printf("\033[31m%s attaque %s et inflige %d dégâts !\n\033[0m", enemy.Name, player.Name, totalDamage)
}

func ChooseAction() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nQue veux-tu faire ?")
	fmt.Println("\033[36m1. Attaque de base\033[0m")
	fmt.Println("\033[34m2. Attaque puissante\033[0m")
	fmt.Println("\033[33m3. Ouvrir l'inventaire\033[0m")
	fmt.Println("\033[31m4. Activer capacité spéciale\033[0m")
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
			fmt.Printf("\033[34m%s utilise %s et inflige %d dégâts à %s !\n\033[0m", attackerName, attack.Name, attack.Damage, defenderName)
		} else {
			fmt.Printf("\033[34m%s utilise %s mais rate son attaque !\n\033[0m", attackerName, attack.Name)
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

func Battle(player *personnage.Character, enemy *skelly.Monster) {
	for player.CurrentHP > 0 && enemy.CurrentHP > 0 {
		fmt.Println("\n--- Tour du joueur ---")
		fmt.Println("\033[34mPV Joueur:\033[0m", player.CurrentHP, "| \033[31mPV Ennemi:\033[0m", enemy.CurrentHP)

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
		}

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
			fmt.Println(enemy.Name, "\033[32mest vaincu ! \033[0m")
			break
		}

		fmt.Println("\n--- Tour de l'ennemi ---")
		enAttack(enemy, player)

		if player.CurrentHP <= 0 {
			fmt.Println(player.Name, "\033[33mest vaincu ! \033[0m")
			break
		}
	}

	fmt.Printf("PV Joueur: %d | PV Ennemi: %d\n", player.CurrentHP, enemy.CurrentHP)
	if player.CurrentHP <= 0 {
		clearScreen()
		content, err := os.ReadFile("wasted.txt")
		if err != nil {
			fmt.Println("Erreur de lecture du fichier :", err)
			return
		}
		fmt.Printf("\033[31m%s\033[0m\n", string(content))
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
