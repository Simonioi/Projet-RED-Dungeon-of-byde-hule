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
    Argent      = item.Item{Name: "£", Quantity: 1}
)
