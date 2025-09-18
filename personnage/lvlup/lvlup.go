package lvlup

import (
	"dungeon/personnage"
	"fmt"
)
func LevelUp(p *personnage.Character) {
	p.MaxHP += 8
	p.CurrentHP = p.MaxHP
	fmt.Println("La Caisse des Donjons est heureuse de vous annoncer que vous avez gagné un niveau !")
	fmt.Printf("Votre vie maximale a augmenté de 5. PV actuels restaurés à %d/%d.\n", p.CurrentHP, p.MaxHP)
	switch p.Class {
		case "Mage":
			p.MaxMP += 4
			p.CurrentMP = p.MaxMP
			fmt.Printf("Votre mana maximal a augmenté de 4. PM actuels restaurés à %d/%d.\n", p.CurrentMP, p.MaxMP)
			p.Capacité = append(p.Capacité, personnage.Attack{Name: "Boule de feu", Damage: 8, HitChance: 1.0})
			fmt.Println("Nouvelle capacité apprise : Boule de feu (Dégâts : 8)")
		case "Barbare":
			p.Capacité = append(p.Capacité, personnage.Attack{Name: "I WOULD LIKE TO RAGE", TempDamageBoost: 4, TempHealthBoost: 2, Duration: 5})
			fmt.Println("Nouvelle capacité apprise : I WOULD LIKE TO RAGE")
	}
	fmt.Println("Appuyez sur Entrée pour continuer...")
	fmt.Scanln()
}
