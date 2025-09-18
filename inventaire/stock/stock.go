package stock

import "dungeon/inventaire/item"

var (
	Armor       = item.Item{Name: "Armure de barbare", HealthBoost: 2, Quantity: 1}
	Sword       = item.Item{Name: "Épée", Damage: 3, Quantity: 1}
	PotionMinor = item.Item{Name: "Potion mineure", HealAmount: 3, Quantity: 1}
	PotionMajor = item.Item{Name: "Potion majeure", HealAmount: 5, Quantity: 1}
	Robe        = item.Item{Name: "Robe enchantée", HealthBoost: 5, Quantity: 1}
	Staff       = item.Item{Name: "Bâton de mage", Damage: 2, Quantity: 1}
	PotionMana  = item.Item{Name: "Potion de mana", HealMana: 4, Quantity: 1}
	FireScroll  = item.Item{Name: "Parchemin de boule de feu", Damage: 4, Quantity: 1}
	Argent      = item.Item{Name: "pièce d'or", Quantity: 1}
	Armor2      = item.Item{Name: "Armure de barbare lvl2", HealthBoost: 4, Quantity: 1}
	Robe2       = item.Item{Name: "Robe de mage enchanté lvl2", HealthBoost: 7, Quantity: 1}
	Sword2      = item.Item{Name: "Épée de Conan", Damage: 4, Quantity: 1}
	Staff2      = item.Item{Name: "Bâton de Elminster", Damage: 3, Quantity: 1}
)
