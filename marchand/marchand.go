package marchand

import (
	"dungeon/inventaire/item"
	"dungeon/inventaire/stock"
	"fmt"
)

type MarchandItem struct {
	Item item.Item
	Prix int
}

var ItemsMarchand = []MarchandItem{
	{Item: stock.Armor2, Prix: 40},
	{Item: stock.Robe2, Prix: 40},
	{Item: stock.Sword2, Prix: 20},
	{Item: stock.Staff2, Prix: 20},
}

func ActiverMarchand(inv interface {
	AddItem(item.Item)
	GetItems() []item.Item
}, argent *int) {
	fmt.Println("\nBienvenue chez le Marchand mystérieux ! Voici ce qu'il propose :")
	for i, mi := range ItemsMarchand {
		fmt.Printf("  %d. %s - %d£\n", i+1, mi.Item.Name, mi.Prix)
	}
	fmt.Printf("\nVous avez %d£. Tapez le numéro de l'objet à acheter ou '0' pour quitter : ", *argent)
	var choix int
	fmt.Scanln(&choix)
	if choix < 1 || choix > len(ItemsMarchand) {
		fmt.Println("Marchand : Reviens quand tu veux !")
		fmt.Println("Appuie sur Entrée pour continuer...")
		fmt.Scanln()
		return
	}
	itemChoisi := ItemsMarchand[choix-1]
	if *argent < itemChoisi.Prix {
		fmt.Println("Marchand : Tu n'as pas assez d'argent !")
		fmt.Println("Appuie sur Entrée pour continuer...")
		fmt.Scanln()
		return
	}
	*argent -= itemChoisi.Prix
	inv.AddItem(itemChoisi.Item)
	fmt.Printf("Vous avez acheté %s !\n", itemChoisi.Item.Name)
	fmt.Println("Appuie sur Entrée pour continuer...")
	fmt.Scanln()
}

//ooo
