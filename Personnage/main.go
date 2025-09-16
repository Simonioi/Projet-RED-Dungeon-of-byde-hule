package personnage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"dungeon/inventaire"
)

type Attack struct {
	Name      string
	Damage    int
	HitChance float64
}

type Character struct {
	Name      string
	Class     string
	MaxHP     int
	CurrentHP int
	MaxMP     int
	CurrentMP int
	Inventory *inventaire.Inventory
	Attacks1   []Attack // Attaques de base
	Attacks2   []Attack // Attaques puissante
}

func CreateBarbarian(name string) Character {
	// Créer l'inventaire du barbare avec les objets par défaut
	inventory := inventaire.CreateBarbarianInventory()

	// Récupérer les objets pour calculer les stats
	items := inventory.GetItems()
	var armor, sword inventaire.Item
	for _, item := range items {
		if item.Name == "Armure de barbare" {
			armor = item
		}
		if item.Name == "Épée" {
			sword = item
		}
	}

	attacks1 := []Attack{
		{Name: "Attaque rapide", Damage: sword.Damage, HitChance: 0.95},
	}
	attacks2 := []Attack{
		{Name: "Attaque puissante", Damage: sword.Damage + 1, HitChance: 0.80},
	}


	baseHP := 8 + armor.HealthBoost

	return Character{
		Name:      name,
		Class:     "Barbare",
		MaxHP:     baseHP,
		CurrentHP: baseHP,
		MaxMP:     4,
		CurrentMP: 4,
		Inventory: inventory,
		Attacks1:  attacks1,
		Attacks2:  attacks2,
	}
}

func CreateMage(name string) Character {
	// Créer l'inventaire du mage avec les objets par défaut
	inventory := inventaire.CreateMageInventory()

	// Récupérer les objets pour calculer les stats
	items := inventory.GetItems()
	var robe, staff inventaire.Item
	for _, item := range items {
		if item.Name == "Robe enchantée" {
			robe = item
		}
		if item.Name == "Bâton de mage" {
			staff = item
		}
	}

	attacks1 := []Attack{
		{Name: "Coup de bâton", Damage: staff.Damage, HitChance: 0.70},
	}
	attacks2 := []Attack{
	{Name: "Missile magique", Damage: 4, HitChance: 0.95},
	}

	baseHP := 5 + robe.HealthBoost

	return Character{
		Name:      name,
		Class:     "Mage",
		MaxHP:     baseHP,
		CurrentHP: baseHP,
		MaxMP:     4,
		CurrentMP: 4,
		Inventory: inventory,
		Attacks1:  attacks1,
		Attacks2:  attacks2,
	}
}

func ShowInventory(player Character) {
	player.Inventory.ShowInventory()
}

func main() {
	var name, classChoice string

	fmt.Println("Bienvenue dans le jeu !")
	fmt.Print("Choisissez votre classe (barbare/mage) : ")
	fmt.Scanln(&classChoice)

	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&name)

	var player Character
	switch strings.ToLower(classChoice) {
	case "barbare":
		player = CreateBarbarian(name)
	case "mage":
		player = CreateMage(name)
	default:
		fmt.Println("Classe inconnue.")
		return
	}

	fmt.Printf("Bienvenue %s le %s ! PV: %d/%d, PM: %d/%d\n",
		player.Name, player.Class, player.CurrentHP, player.MaxHP, player.CurrentMP, player.MaxMP)

	fmt.Println("\n=== COMMANDES DISPONIBLES ===")
	fmt.Println("1 ou 'inventaire' - Ouvrir/fermer l'inventaire")
	fmt.Println("2 ou 'aide' - Afficher l'aide")
	fmt.Println("3 ou 'quitter' - Quitter le jeu")
	fmt.Println("===============================")

	inventoryOpen := false
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		// Si l'inventaire est ouvert, gérer les commandes spéciales
		if inventoryOpen {
			switch input {
			case "fermer", "close":
				inventoryOpen = false
				fmt.Println("Inventaire fermé.")
			case "1", "2", "3", "4":
				itemIndex, err := strconv.Atoi(input)
				if err == nil {
					itemIndex-- // Convertir en index 0-based
					success, value, message := player.Inventory.UseItem(itemIndex)
					fmt.Println(message)
					if success && value > 0 {
						// Si c'est un soin
						if strings.Contains(message, "récupérez") {
							oldHP := player.CurrentHP
							player.CurrentHP += value
							if player.CurrentHP > player.MaxHP {
								player.CurrentHP = player.MaxHP
							}
							fmt.Printf("PV: %d/%d (+%d)\n", player.CurrentHP, player.MaxHP, player.CurrentHP-oldHP)
						}
						// Réafficher l'inventaire mis à jour
						fmt.Println("\nInventaire mis à jour :")
						ShowInventory(player)
					}
				}
			default:
				fmt.Println("Dans l'inventaire : tapez 1-4 pour utiliser un objet, ou 'fermer' pour fermer.")
			}
		} else {
			// Menu principal
			switch input {
			case "1", "inventaire":
				inventoryOpen = true
				fmt.Println("Inventaire ouvert.")
				ShowInventory(player)
			case "2", "aide", "help":
				fmt.Println("\n=== AIDE ===")
				fmt.Println("Commandes disponibles :")
				fmt.Println("  1 ou 'inventaire' - Afficher votre inventaire")
				fmt.Println("  2 ou 'aide' - Afficher cette aide")
				fmt.Println("  3 ou 'quitter' - Quitter le jeu")
				fmt.Printf("Votre personnage : %s le %s\n", player.Name, player.Class)
				fmt.Printf("PV: %d/%d, PM: %d/%d\n", player.CurrentHP, player.MaxHP, player.CurrentMP, player.MaxMP)
				fmt.Println("=============")
			case "3", "quitter", "quit", "exit":
				fmt.Println("Fin du jeu.")
				return
			default:
				fmt.Println("Commande inconnue. Tapez '2' ou 'aide' pour voir les commandes disponibles.")
			}
		}
	}
}
