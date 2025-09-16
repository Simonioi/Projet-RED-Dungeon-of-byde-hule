package inventaire

import "fmt"

// Item représente un objet dans l'inventaire
type Item struct {
	Name        string
	HealthBoost int
	Damage      int
	HealAmount  int
}

// Inventory représente l'inventaire d'un personnage
type Inventory struct {
	Items []Item
}

// NewInventory crée un nouvel inventaire vide
func NewInventory() *Inventory {
	return &Inventory{
		Items: make([]Item, 0),
	}
}

// AddItem ajoute un objet à l'inventaire
func (inv *Inventory) AddItem(item Item) {
	inv.Items = append(inv.Items, item)
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
func (inv *Inventory) GetItems() []Item {
	return inv.Items
}

// ShowInventory affiche le contenu de l'inventaire
func (inv *Inventory) ShowInventory() {
	if len(inv.Items) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	fmt.Println("Inventaire :")
	for i, item := range inv.Items {
		fmt.Printf("  %d. %s", i+1, item.Name)
		if item.HealthBoost > 0 {
			fmt.Printf(" (+%d PV)", item.HealthBoost)
		}
		if item.Damage > 0 {
			fmt.Printf(" (+%d Dégâts)", item.Damage)
		}
		if item.HealAmount > 0 {
			fmt.Printf(" (Soigne %d PV)", item.HealAmount)
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
	
	item := inv.Items[index]
	
	// Si c'est un objet utilisable (potion, parchemin)
	if item.HealAmount > 0 {
		// Supprimer l'objet de l'inventaire après utilisation
		inv.RemoveItem(index)
		return true, item.HealAmount, fmt.Sprintf("Vous utilisez %s et récupérez %d PV.", item.Name, item.HealAmount)
	} else if item.Name == "Parchemin de boule de feu" {
		// Supprimer le parchemin après utilisation
		inv.RemoveItem(index)
		return true, item.Damage, fmt.Sprintf("Vous utilisez %s ! Dégâts magiques : %d.", item.Name, item.Damage)
	} else {
		// Objets non utilisables (armure, armes)
		return false, 0, fmt.Sprintf("%s ne peut pas être utilisé directement.", item.Name)
	}
}

// CreateBarbarianInventory crée l'inventaire par défaut pour un barbare
func CreateBarbarianInventory() *Inventory {
	inv := NewInventory()

	armor := Item{Name: "Armure de barbare", HealthBoost: 2}
	sword := Item{Name: "Épée", Damage: 3}
	potionMinor := Item{Name: "Potion mineure", HealAmount: 3}
	potionMajor := Item{Name: "Potion majeure", HealAmount: 5}

	inv.AddItem(armor)
	inv.AddItem(sword)
	inv.AddItem(potionMinor)
	inv.AddItem(potionMajor)

	return inv
}

// CreateMageInventory crée l'inventaire par défaut pour un mage
func CreateMageInventory() *Inventory {
	inv := NewInventory()
	
	robe := Item{Name: "Robe enchantée", HealthBoost: 5}
	staff := Item{Name: "Bâton de mage", Damage: 2}
	scrollFireball := Item{Name: "Parchemin de boule de feu", Damage: 6}
	potionMana := Item{Name: "Potion de mana", HealAmount: 4}
	
	inv.AddItem(robe)
	inv.AddItem(staff)
	inv.AddItem(scrollFireball)
	inv.AddItem(potionMana)
	
	return inv
}