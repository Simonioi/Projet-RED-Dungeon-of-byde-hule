package combat

import (
    "fmt"
    "dungeon/personnage"
	"dungeon/mimic"
)

func Attack(attacker *personnage.Character, defender *mimic.Monster) {
    totalDamage := 0
    for _, attack := range attacker.Attacks {
        totalDamage += attack.Damage
    }
    if totalDamage < 0 {
        totalDamage = 0
    }
    defender.CurrentHP -= totalDamage
    fmt.Printf("%s attaque %s et inflige %d dégâts !\n", attacker.Name, defender.Name, totalDamage)
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
        Attack(player, enemy)
        if player.CurrentHP <= 0 {
            fmt.Println(player.Name, "est vaincu !")
            break
        }
    }
}
