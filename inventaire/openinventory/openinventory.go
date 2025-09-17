package menuinventaire

import (
	"bufio"
    "dungeon/personnage"
	"dungeon/inventaire"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// OpenInventory permet d'utiliser un objet de l'inventaire.
// Si l'objet soigne, il soigne le joueur.
// Si l'objet inflige des dégâts, il les inflige à l'ennemi.
func OpenInventory(inv *inventaire.Inventory, player *personnage.Character) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		inv.ShowInventory()
		fmt.Print("Choix : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "fermer" {
			fmt.Println("Fermeture de l'inventaire.")
			return false // Aucun objet utilisé
		}

		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(inv.Items) {
			fmt.Println("Choix invalide.")
			continue
		}

		item := inv.Items[index-1]

		// Si l'objet soigne
		if item.HealAmount > 0 {
			inv.RemoveItem(index - 1)
			player.CurrentHP += item.HealAmount
			fmt.Printf("Vous utilisez %s et récupérez %d PV.\n", item.Name, item.HealAmount)
			fmt.Printf("PV actuels : %d\n", player.CurrentHP)
			return true
		}

		// Si l'objet inflige des dégâts
		if item.Damage > 0 {
			inv.RemoveItem(index - 1)
			player.PendingDamage = item.Damage // Stockage temporaire pour appliquer au tour suivant
			player.PendingDamageText = fmt.Sprintf("Vous utilisez %s ! Dégâts magiques : %d.", item.Name, item.Damage)
			return true
		}

		// Objet non utilisable
		fmt.Printf("%s ne peut pas être utilisé directement.\n", item.Name)
	}
}
