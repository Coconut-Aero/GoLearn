package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var again = true
	for again {
		var words = [...]string{
			"write", "program", "that", "analysis",
			"culture", "access", "enhance", "system",
			"array", "mirror", "window", "public",
			"chinese", "physics", "chemistry", "static"}
		var chosenword = words[rand.Intn(len(words))]
		var miss = 0
		var chosenw = strings.Split(chosenword, "")
		var request string
		var guess = make([]bool, len(chosenw))
		var exit = false
		for !exit {
			fmt.Print("(Guess) Enter a letter in word ")
			for i := 0; i < len(chosenw); i++ {
				if !guess[i] {
					fmt.Print("*")
				} else {
					fmt.Print(chosenw[i])
				}
			}
			fmt.Print(" >")
			var enter string
			_, err := fmt.Scan(&enter)
			if err != nil {
				return
			}
			var notin = 0
			for i := 0; i < len(chosenw); i++ {
				if enter == chosenw[i] && !guess[i] {
					guess[i] = true
				} else if enter == chosenw[i] && guess[i] {
					fmt.Println(enter + " is already in the word")
				} else {
					notin++
				}
			}
			if notin == len(chosenw) {
				fmt.Println(enter + " is not in the word")
				miss++
			}
			var exit1 = 0
			for i := 0; i < len(chosenw); i++ {
				if guess[i] {
					exit1++
				} else {
					break
				}
			}
			if exit1 == len(chosenw) {
				exit = true
			}
		}
		fmt.Println("The word is", chosenword, ". You missed", miss, "time(s).")
		fmt.Print("Do you want to guess another time? (Y/N)>")
		_, err := fmt.Scan(&request)
		if err != nil {
			return
		}
		if request == "N" || request == "n" {
			again = false
			fmt.Print("Bye!")
		}
	}
}
