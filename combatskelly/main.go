package combatskelly

import (
    "bufio"
    "dungeon/personnage"
    "dungeon/skelly"
    "dungeon/inventaire/openinventory"
    "fmt"
    "math/rand"
    "os"
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

func Battle(player *personnage.Character, enemy *skelly.Monster) {
    for player.CurrentHP > 0 && enemy.CurrentHP > 0 {
        fmt.Println("\n--- Tour du joueur ---")
        fmt.Println("\033[34mPV Joueur:\033[0m", player.CurrentHP, "| \033[31mPV Ennemi:\033[0m", enemy.CurrentHP)

        choice := ChooseAction()
        switch choice {
        case 1:
            ExecuteAttack(player.Name, player.Attacks1, enemy.Name, &enemy.CurrentHP)
        case 2:
            ExecuteAttack(player.Name, player.Attacks2, enemy.Name, &enemy.CurrentHP)
        case 3:
            used := menuinventaire.OpenInventory(player.Inventory, player)
            if !used {
                // Si aucun objet utilisé, on recommence le tour du joueur
                continue
            }
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

    fmt.Printf("\033[34mPV Joueur: %d\033[0m | \033[31mPV Ennemi: %d\033[0m\n", player.CurrentHP, enemy.CurrentHP)
}