package combat

import (
    "fmt"
    "dungeon/personnage"
	"dungeon/mimic"
)

func Attack(player *personnage.Character, enemy *mimic.Monster) {
    totalDamage := 0
    for _, attack := range player.Attacks {
        totalDamage += attack.Damage
    }
    if totalDamage < 0 {
        totalDamage = 0
    }
    enemy.CurrentHP -= totalDamage
    fmt.Printf("%s attaque %s et inflige %d dégâts !\n", player.Name, enemy.Name, totalDamage)
}

func enAttack(enemy *mimic.Monster, player *personnage.Character) {
    totalDamage := 0
    for _, attack := range enemy.Attacks {
        totalDamage += attack.Damage
    }
    if totalDamage < 0 {
        totalDamage = 0
    }
    player.CurrentHP -= totalDamage
    fmt.Printf("%s attaque %s et inflige %d dégâts !\n", enemy.Name, player.Name, totalDamage)
}


func Battle(player *personnage.Character, enemy *mimic.Monster) {
    for player.CurrentHP > 0 && enemy.CurrentHP > 0 {
        fmt.Println("\n--- Tour du joueur ---")
        Attack(player, enemy)
        if enemy.CurrentHP <= 0 {
            fmt.Println(enemy.Name, "est vaincu !")
            break
        }

        fmt.Println("\n--- Tour de l'ennemi ---")
        enAttack(enemy, player)
        if player.CurrentHP <= 0 {
            fmt.Println(player.Name, "est vaincu !")
            break
        }
    }
}
