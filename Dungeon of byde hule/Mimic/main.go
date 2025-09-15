package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("mimic.txt")
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}

	fmt.Printf("\033[38;5;130m%s\033[0m\n", string(content))
}
