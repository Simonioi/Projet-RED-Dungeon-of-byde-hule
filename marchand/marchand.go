package marchand

import (
	"dungeon/inventaire"
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

func ActiverMarchand(inv *inventaire.Inventory) {
	for {
		argent := inv.GetMoney()
		fmt.Println("\nBienvenue chez le Sheik point, que souhaitez-vous ?")
		for i, mi := range ItemsMarchand {
			fmt.Printf("  %d. %s - %d£\n", i+1, mi.Item.Name, mi.Prix)
		}
		fmt.Printf("  %d. Vendre un objet\n", len(ItemsMarchand)+1)
		fmt.Printf("\nVous avez %d£. Tapez le numéro de l'objet à acheter, '%d' pour vendre, ou '0' pour quitter : ", argent, len(ItemsMarchand)+1)
		var choix int
		fmt.Scanln(&choix)
		if choix == 0 {
			fmt.Println("Reviens quand tu veux !")
			fmt.Println("Appuie sur Entrée pour continuer...")
			fmt.Scanln()
			return
		}
		if choix == len(ItemsMarchand)+1 {
			// Vente d'un objet
			vendables := []struct {
				Nom   string
				Index int
				Prix  int
			}{}
			for i, it := range inv.GetItems() {
				prix := 0
				switch it.Name {
				case "Armure de barbare":
					prix = 20
				case "Robe enchantée":
					prix = 20
				case "Épée":
					prix = 10
				case "Bâton de mage":
					prix = 10
				}
				if prix > 0 && it.Quantity > 0 {
					vendables = append(vendables, struct {
						Nom   string
						Index int
						Prix  int
					}{it.Name, i, prix})
				}
			}
			if len(vendables) == 0 {
				fmt.Println("Vous n'avez rien à vendre !")
				fmt.Println("Appuie sur Entrée pour continuer...")
				fmt.Scanln()
				continue
			}
			fmt.Println("Objets que vous pouvez vendre :")
			for i, v := range vendables {
				fmt.Printf("  %d. %s (+%d£)\n", i+1, v.Nom, v.Prix)
			}
			fmt.Printf("Choisissez l'objet à vendre (0 pour annuler) : ")
			var vchoix int
			fmt.Scanln(&vchoix)
			if vchoix < 1 || vchoix > len(vendables) {
				fmt.Println("Vente annulée.")
				fmt.Println("Appuie sur Entrée pour continuer...")
				fmt.Scanln()
				continue
			}
			v := vendables[vchoix-1]
			inv.RemoveItem(v.Index)
			inv.AddMoney(v.Prix)
			fmt.Printf("Vous avez vendu %s pour %d£ !\n", v.Nom, v.Prix)
			fmt.Println("Appuie sur Entrée pour continuer...")
			fmt.Scanln()
			continue
		}
		if choix < 1 || choix > len(ItemsMarchand) {
			fmt.Println("Choix invalide.")
			fmt.Println("Appuie sur Entrée pour continuer...")
			fmt.Scanln()
			continue
		}
		itemChoisi := ItemsMarchand[choix-1]
		if argent < itemChoisi.Prix {
			fmt.Println("Marchand : Tu n'as pas assez d'argent !")
			fmt.Println("Appuie sur Entrée pour continuer...")
			fmt.Scanln()
			continue
		}
		inv.RemoveMoney(itemChoisi.Prix)
		inv.AddItem(itemChoisi.Item)
		fmt.Printf("Vous avez acheté %s !\n", itemChoisi.Item.Name)
		fmt.Println("Appuie sur Entrée pour continuer...")
		fmt.Scanln()
	}
}

//ooo
