package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("🧙‍♂️ Une porte devant vous barre votre chemin.")
	fmt.Println("Elle ne présente ni serrure ni poignée, mais une tête de sphinx.")
	fmt.Println("Ses yeux rouges s'allument et une voix résonne alors :")
	fmt.Println()
	fmt.Println("« Jamais je ne suis loin de mon autre jumelle,")
	fmt.Println("on m'associe souvent au parfum vomitif d'une partie du corps")
	fmt.Println("qui n'est pas vraiment belle, localisée fort loin de l'organe olfactif. »")
	fmt.Println()
	fmt.Print("👉 Quelle est votre réponse ? ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	answer := strings.TrimSpace(strings.ToLower(input))

	if answer == "chaussette" {
		fmt.Println("✅ La porte s'ouvre lentement dans un grincement sinistre...")
		fmt.Println("Vous pouvez continuer votre aventure !")
	} else {
		fmt.Println("❌ Le sphinx reste silencieux. La porte ne bouge pas.")
		fmt.Println("Réfléchissez bien et réessayez...")
	}
}