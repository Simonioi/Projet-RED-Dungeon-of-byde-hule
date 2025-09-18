package lvlup

import (
	"dungeon/personnage"
	"fmt"
)

func LevelUp(p *personnage.Character) {
	p.MaxHP += 8
	p.CurrentHP = p.MaxHP
	fmt.Println("\033[33mLa Caisse des Donjons est heureuse de vous annoncer que vous avez gagné un niveau !\033[0m")
	fmt.Printf("\033[35mVotre vie maximale a augmenté de 5. PV actuels restaurés à %d/%d.\033[0m\n", p.CurrentHP, p.MaxHP)
	switch p.Class {
	case "Mage":
		p.MaxMP += 4
		p.CurrentMP = p.MaxMP
		fmt.Printf("\033[35mVotre mana maximal a augmenté de 4. PM actuels restaurés à %d/%d.\033[0m\n", p.CurrentMP, p.MaxMP)
		p.Capacité = append(p.Capacité, personnage.Attack{Name: "Boule de feu", Damage: 8, HitChance: 1.0})
		fmt.Println("\033[35mNouvelle capacité apprise : Boule de feu (Dégâts : 8)\033[0m")
	case "Barbare":
		p.Capacité = append(p.Capacité, personnage.Attack{Name: "I WOULD LIKE TO RAGE", TempDamageBoost: 4, TempHealthBoost: 2, Duration: 6})
		fmt.Println("\033[35mNouvelle capacité apprise : I WOULD LIKE TO RAGE\033[0m")
	}
	fmt.Println("Appuyez sur Entrée pour continuer...")
	fmt.Scanln()
}

// gestion du LVL UP
