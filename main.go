package main

import (
	"dungeon/carte"
	"dungeon/personnage"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var musicCmd *exec.Cmd

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}

func playMusicLoop() {
	go func() {
		for {
			if runtime.GOOS == "windows" {
				musicCmd = exec.Command("powershell", "-c", "while ($true) { (New-Object Media.SoundPlayer 'Pokémon-GO-Night-Theme-_8-bit-Arrangement_.wav').PlaySync() }")
				musicCmd.Run()
			} else {
				musicCmd = exec.Command("ffplay", "-nodisp", "-autoexit", "-loop", "0", "Pokémon-GO-Night-Theme-_8-bit-Arrangement_.wav")
				musicCmd.Run()
			}
		}
	}()
}

func stopMusic() {
	if musicCmd != nil && musicCmd.Process != nil {
		musicCmd.Process.Kill()
	}
}

func main() {
	defer stopMusic()
	playMusicLoop()
	// Clear screen avant l'intro
	clearScreen()
	// Intro mission
	var missionChoice string
	fmt.Println("\033[33mByde Hule Games Présente...\033[0m")
	fmt.Println("\nAppuie sur Entrée pour continuer...")
	fmt.Scanln()
	clearScreen()
	fmt.Println("\033[33mContrat de Repérage et nettoyage de Donjon niveau 1, le Donjon de Byde Hule récompense :\033[0m")
	fmt.Println("\033[33m1000 pièces (taxe non comprise, se référer aux barème de taxation des aventuriers)\033[0m")
	fmt.Println()
	fmt.Println("Bonjour Aventurier, votre Contrat, si toutefois vous l'acceptez,")
	fmt.Println("sera de nettoyer un donjon, pour le compte de notre Contractuel le Seigneur de Maaschyn.")
	fmt.Println()
	fmt.Print("Donjon pouvant inclure : monstre mineur de type revenant ou creation nécromantique, ")
	fmt.Println("piège franchement visiblemême pour un aveugle,")
	fmt.Println("énigme ou puzzle de difficulté 3 à 6 ans (pièces manquante non incluse),")
	fmt.Println("BBEG de niveau Pathétique à Ridicule.")

	fmt.Println("\nAppuie sur Entrée pour continuer...")
	fmt.Scanln()
	clearScreen()

	fmt.Println("Information non contractuelle du donjon :")
	fmt.Println("-La caisse des donjons décline toute responsabilité en cas de mort violente, de mutilation,")
	fmt.Printf(" perte de santé mentale, téléportation dans des univers parallèle, arnaque de gobelin,")
	fmt.Println(" digestion par un cube gélatineux, sacrifice par des cultistes zélés.")
	fmt.Println()

	fmt.Println("-La Caisse des Donjons se reserve le droit de ne pas payer l'aventurier en cas ")
	fmt.Println("de non présentation du contrat et/ou preuve de mort du BBEG.")

	fmt.Println("\nAppuie sur Entrée pour continuer...")
	fmt.Scanln()
	clearScreen()

	fmt.Println("-Bob dit 'je vais vous Rosser', aventurier vétéran de niveau 20,")
	fmt.Println("directeur par interim de l'attribution des contrats aux aventurier débutants un poil suicidaire,")
	fmt.Println("désespéré et ambitieux à la Caisse des Donjons.")
	fmt.Println()

	fmt.Println("Ce message ne s'autodétruira pas mais sera probablement perdu au fond de votre sac.")
	fmt.Println("\n\033[31mAcceptez-vous la mission ? (oui/non)\033[0m")
	fmt.Print("> ")
	fmt.Scanln(&missionChoice)
	if strings.ToLower(missionChoice) != "oui" {
		fmt.Println("Tu peux aussi esquiver le scenar et aller faire un épisode plage, mais t'es venu pour jouer donc relance le jeu et tache de nous mettre une bonne note (pretty please uwu)")
		return
	}

	var name, class string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&name)
	fmt.Print("Choisissez la classe (barbare/mage) : ")
	fmt.Scanln(&class)

	var player personnage.Character
	switch strings.ToLower(class) {
	case "barbare":
		player = personnage.CreateBarbarian(name)
	case "mage":
		player = personnage.CreateMage(name)
	default:
		fmt.Println("Classe inconnue, création d'un barbare par défaut.")
		player = personnage.CreateBarbarian(name)
	}

	fmt.Printf("Personnage créé : %s le %s (PV: %d/%d, PM: %d/%d)\n", player.Name, player.Class, player.CurrentHP, player.MaxHP, player.CurrentMP, player.MaxMP)
	carte.Start(player)
}

//truce
