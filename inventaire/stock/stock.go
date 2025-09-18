package stock

import "dungeon/inventaire/item"

var (
	Armor                  = item.Item{Name: "Armure de barbare", HealthBoost: 2, Quantity: 1}
	Sword                  = item.Item{Name: "Épée", BaseDamage: 3, Quantity: 1}
	PotionMinor            = item.Item{Name: "Potion mineure", HealAmount: 3, Quantity: 1}
	PotionMajor            = item.Item{Name: "Potion majeure", HealAmount: 5, Quantity: 1}
	Robe                   = item.Item{Name: "Robe enchantée", HealthBoost: 5, Quantity: 1}
	Staff                  = item.Item{Name: "Bâton de mage", BaseDamage: 2, Quantity: 1}
	PotionMana             = item.Item{Name: "Potion de mana", HealMana: 4, Quantity: 1}
	FireScroll             = item.Item{Name: "Parchemin de boule de feu", Damage: 4, Quantity: 1}
	Argent                 = item.Item{Name: "£", Quantity: 1}
	Armor2                 = item.Item{Name: "Armure de barbare lvl2", HealthBoost: 4, Quantity: 1}
	Robe2                  = item.Item{Name: "Robe de mage enchanté lvl2", HealthBoost: 7, Quantity: 1}
	Sword2                 = item.Item{Name: "Épée de Conan", BaseDamage: 4, Quantity: 1}
	Staff2                 = item.Item{Name: "Bâton de Elminster", BaseDamage: 3, Quantity: 1}
	SolanumTuberosumPatate = item.Item{Name: "Solanum tuberosum (patate)", HealAmount: 2, Quantity: 1}
	ZingiberOfficinale     = item.Item{Name: "Zingiber officinale (gingembre)", HealAmount: 3, Quantity: 1}
	CitrusLimus            = item.Item{Name: "Citrus limus (citron vert)", HealAmount: 1, Quantity: 1}
)

//bibliothèque des items dispo dans le jeu.
