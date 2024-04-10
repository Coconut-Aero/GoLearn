package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var exit bool
	for !exit {
		fmt.Print("Scissor(0) , Rock (1) , Paper (2) :")
		var player int
		var computer = rand.Intn(3)
		fmt.Scan(&player)
		switch computer {
		case 0:
			fmt.Print("The computer is scissor.")
			break
		case 1:
			fmt.Print("The computer is rock.")
			break
		case 2:
			fmt.Print("The computer is paper.")
			break
		}
		switch player {
		case 0:
			fmt.Print("You are scissor.")
			break
		case 1:
			fmt.Print("You are rock.")
			break
		case 2:
			fmt.Print("You are paper.")
			break
		}
		if player != 0 && player != 1 && player != 2 {
			fmt.Print("You entered an invalid number.")
			os.Exit(1)
		} else if player < computer {
			if (player != 0) || (computer != 2) {
				fmt.Print("Computer Wins.")
			} else {
				fmt.Print("You win.")
			}
		} else if player != computer {
			if (player != 2) || (computer != 0) {
				fmt.Print("You win.")
			} else {
				fmt.Print("Computer wins.")
			}
		} else {
			fmt.Print("This round tied.")
		}
		fmt.Print("Play again?(Y/N)")
		var request string
		fmt.Scan(&request)
		if request == "N" || request == "n" {
			exit = true
		}
		fmt.Println()
	}
}
