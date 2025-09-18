package alchimie

import (
	"dungeon/inventaire/item"
	"dungeon/inventaire/stock"
	"fmt"
)

// Déclaration des 3 items de la potion super cool
var (
	SolanumTuberosum   = item.Item{Name: "Solanum tuberosum (patate)", Quantity: 1}
	ZingiberOfficinale = item.Item{Name: "Zingiber officinale", Quantity: 1}
	CitrusLimus        = item.Item{Name: "Citrus limus", Quantity: 1}
)

// Table d'alchimie qui combine les 3 items pour créer une Potion majeure
func UtiliserTableAlchimie(inventaire *[]item.Item) {
	hasPatate := false
	hasZingiber := false
	hasCitrus := false
	// Vérifier la présence des 3 ingrédients
	for _, it := range *inventaire {
		if it.Name == SolanumTuberosum.Name && it.Quantity > 0 {
			hasPatate = true
		}
		if it.Name == ZingiberOfficinale.Name && it.Quantity > 0 {
			hasZingiber = true
		}
		if it.Name == CitrusLimus.Name && it.Quantity > 0 {
			hasCitrus = true
		}
	}
	if hasPatate && hasZingiber && hasCitrus {
		// Retirer 1 de chaque ingrédient
		for i := range *inventaire {
			if (*inventaire)[i].Name == SolanumTuberosum.Name && (*inventaire)[i].Quantity > 0 {
				(*inventaire)[i].Quantity--
			}
			if (*inventaire)[i].Name == ZingiberOfficinale.Name && (*inventaire)[i].Quantity > 0 {
				(*inventaire)[i].Quantity--
			}
			if (*inventaire)[i].Name == CitrusLimus.Name && (*inventaire)[i].Quantity > 0 {
				(*inventaire)[i].Quantity--
			}
		}
		// Ajouter la potion majeure
		trouve := false
		for i := range *inventaire {
			if (*inventaire)[i].Name == stock.PotionMajor.Name {
				(*inventaire)[i].Quantity++
				trouve = true
			}
		}
		if !trouve {
			*inventaire = append(*inventaire, stock.PotionMajor)
		}
		fmt.Println("Vous avez créé une Potion majeure !")
	} else {
		fmt.Println("Il vous manque des ingrédients pour créer une Potion majeure.")
	}
}
