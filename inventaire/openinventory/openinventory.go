package menuinventaire

import (
    "dungeon/inventaire"
    "dungeon/personnage"
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)


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

        success, effect, message := inv.UseItem(index - 1)
        fmt.Println(message)
        if success {
            player.CurrentHP += effect
            fmt.Printf("PV actuels : %d\n", player.CurrentHP)
            return true // Objet utilisé
        }
    }
}
