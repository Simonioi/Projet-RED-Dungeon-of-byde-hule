package inventaire

import (
	"dungeon/inventaire/stock"
	"dungeon/inventaire/item"
	"fmt"
)

// Inventory représente l'inventaire d'un personnage
type Inventory struct {
	Items []item.Item
}

// NewInventory crée un nouvel inventaire vide
func NewInventory() *Inventory {
	return &Inventory{
		Items: make([]item.Item, 0),
	}
}

// AddItem ajoute un objet à l'inventaire (empile si déjà présent)
func (inv *Inventory) AddItem(it item.Item) {
	for i, existing := range inv.Items {
		if existing.Name == it.Name {
			inv.Items[i].Quantity += it.Quantity
			return
		}
	}
	inv.Items = append(inv.Items, it)
}

// RemoveItem retire un objet de l'inventaire par index
func (inv *Inventory) RemoveItem(index int) bool {
	if index < 0 || index >= len(inv.Items) {
		return false
	}
	inv.Items = append(inv.Items[:index], inv.Items[index+1:]...)
	return true
}

// GetItems retourne tous les objets de l'inventaire
func (inv *Inventory) GetItems() []item.Item {
	return inv.Items
}

// ShowInventory affiche le contenu de l'inventaire
func (inv *Inventory) ShowInventory() {
	if len(inv.Items) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	fmt.Println("Inventaire :")
	for i, it := range inv.Items {
		fmt.Printf("  %d. %s", i+1, it.Name)
		if it.Quantity > 1 {
			fmt.Printf(" x%d", it.Quantity)
		}
		if it.HealthBoost > 0 {
			fmt.Printf(" (+%d PV)", it.HealthBoost)
		}
		if it.Damage > 0 {
			fmt.Printf(" (+%d Dégâts)", it.Damage)
		}
		if it.HealAmount > 0 {
			fmt.Printf(" (Soigne %d PV)", it.HealAmount)
		}
		if it.HealMana > 0 {
			fmt.Printf(" (Restaure %d Mana)", it.HealMana)
		}
		fmt.Println()
	}
	fmt.Println("Tapez le numéro de l'objet pour l'utiliser, ou 'fermer' pour fermer l'inventaire.")
}

// UseItem utilise un objet de l'inventaire par index et retourne ses effets
func (inv *Inventory) UseItem(index int) (bool, int, string) {
	if index < 0 || index >= len(inv.Items) {
		return false, 0, "Objet inexistant."
	}

	it := inv.Items[index]

	// Si c'est un objet utilisable
	if it.HealAmount > 0 {
		inv.RemoveItem(index)
		return true, it.HealAmount, fmt.Sprintf("Vous utilisez %s et récupérez %d PV.", it.Name, it.HealAmount)
	} else if it.Name == "Parchemin de boule de feu" {
		inv.RemoveItem(index)
		return true, it.Damage, fmt.Sprintf("Vous utilisez %s ! Dégâts magiques : %d.", it.Name, it.Damage)
	} else if it.HealMana > 0 {
		inv.RemoveItem(index)
		return true, it.HealMana, fmt.Sprintf("Vous utilisez %s et récupérez %d Mana.", it.Name, it.HealMana)
	} else {
		return false, 0, fmt.Sprintf("%s ne peut pas être utilisé directement.", it.Name)
	}
}

// CreateBarbarianInventory crée l'inventaire par défaut pour un barbare
func CreateBarbarianInventory() *Inventory {
	inv := NewInventory()
	inv.AddItem(stock.Armor)
	inv.AddItem(stock.Sword)
	inv.AddItem(stock.PotionMinor)
	inv.AddItem(stock.PotionMajor)
	return inv
}

// CreateMageInventory crée l'inventaire par défaut pour un mage
func CreateMageInventory() *Inventory {
	inv := NewInventory()
	inv.AddItem(stock.Robe)
	inv.AddItem(stock.Staff)
	inv.AddItem(stock.PotionMana)
	return inv
}
