package combatcthulhu

import (
	"bufio"
	"dungeon/cthulhu"
	"dungeon/personnage"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"dungeon/inventaire/openinventory"
)

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
		selectedAttack = enemy.Attack1 // Revenir à l'attaque 1 si le tour dépasse 5
	}

	totalDamage := 0
	for _, attack := range selectedAttack {
		if rand.Float64() <= attack.HitChance {
			totalDamage += attack.Damage
			fmt.Printf("\033[31m%s utilise %s et inflige %d dégâts à %s !\n\033[0m", enemy.Name, attack.Name, attack.Damage, player.Name)
		} else {
			fmt.Printf("\033[31m%s utilise %s mais rate son attaque !\n\033[0m", enemy.Name, attack.Name)
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

func ChooseAction() int {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("\nQue veux-tu faire ?")
    fmt.Println("1. Attaque de base")
    fmt.Println("2. Attaque puissante")
    fmt.Println("3. Ouvrir l'inventaire")
    fmt.Print("Choix : ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    choice, err := strconv.Atoi(input)
    if err != nil || (choice < 1 || choice > 3) {
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

func Battle(player *personnage.Character, enemy *cthulhu.Monster) {
    turn := 1

    for player.CurrentHP > 0 && enemy.CurrentHP > 0 {
        fmt.Println("\n--- Tour du joueur ---")
        fmt.Printf("PV Joueur: %d | PV Ennemi: %d\n", player.CurrentHP, enemy.CurrentHP)

        choice := ChooseAction()
        switch choice {
        case 1:
            ExecuteAttack(player.Name, player.Attacks1, enemy.Name, &enemy.CurrentHP)
        case 2:
            ExecuteAttack(player.Name, player.Attacks2, enemy.Name, &enemy.CurrentHP)
        case 3:
            menuinventaire.OpenInventory(player.Inventory, player)
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
}
