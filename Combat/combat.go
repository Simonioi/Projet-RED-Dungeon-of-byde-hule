package combat

import (
	"fmt"
	"math/rand"
)

type Attack struct {
	Name      string
	Damage    int
	HitChance float64
}

type Character struct {
	Name      string
	Class     string
	MaxHP     int
	CurrentHP int
	Attacks   []Attack
}

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attacks   []Attack
}

func Combat(player *Character, enemy *Monster) {
	fmt.Printf("Un %s apparaît !\n", enemy.Name)
	for player.CurrentHP > 0 && enemy.CurrentHP > 0 {
		fmt.Printf("\nVotre tour ! PV: %d/%d\n", player.CurrentHP, player.MaxHP)
		fmt.Println("Choisissez une attaque :")
		for i, atk := range player.Attacks {
			fmt.Printf("%d. %s (Dégâts: %d, Précision: %.0f%%)\n", i+1, atk.Name, atk.Damage, atk.HitChance*100)
		}
		var atkChoice int
		fmt.Print("Numéro de l'attaque : ")
		fmt.Scanln(&atkChoice)
		if atkChoice < 1 || atkChoice > len(player.Attacks) {
			fmt.Println("Choix invalide, tour perdu.")
		} else {
			atk := player.Attacks[atkChoice-1]
			if atk.HitChance > rand.Float64() {
				enemy.CurrentHP -= atk.Damage
				fmt.Printf("Vous touchez le %s pour %d dégâts !\n", enemy.Name, atk.Damage)
			} else {
				fmt.Println("Votre attaque rate !")
			}
		}
		if enemy.CurrentHP <= 0 {
			fmt.Printf("Le %s est vaincu !\n", enemy.Name)
			break
		}
		// Tour du monstre
		atkIndex := rand.Intn(len(enemy.Attacks))
		atk := enemy.Attacks[atkIndex]
		if atk.HitChance > rand.Float64() {
			player.CurrentHP -= atk.Damage
			fmt.Printf("Le %s vous attaque avec %s pour %d dégâts !\n", enemy.Name, atk.Name, atk.Damage)
		} else {
			fmt.Printf("Le %s rate son attaque !\n", enemy.Name)
		}
		if player.CurrentHP <= 0 {
			fmt.Println("Vous êtes vaincu...")
			break
		}
	}
}
