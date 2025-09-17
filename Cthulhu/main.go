package cthulhu

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("cltuuululuuuuu.txt")
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}

	fmt.Printf("\033[32m%s\033[0m\n", string(content))
}
