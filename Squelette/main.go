package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("skelly.txt")
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}

	fmt.Printf("\033[97m%s\033[0m\n", string(content))
}
