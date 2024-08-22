package main

import (
	"GoLearn/utils"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Running time", time.Now().Format("2006-01-02 15:04:05"))
	var Quantity int
	var Length int
	fmt.Println("Enter the quantity of the Password: ")
	_, err := fmt.Scan(&Quantity)
	if err != nil {
		return
	}
	fmt.Println("Enter the length of each Password: ")
	_, err = fmt.Scan(&Length)
	if err != nil {
		return
	}
	Password := make([][]string, Quantity)
	for i := range Password {
		Password[i] = make([]string, Length)
	}
	var AllSymbols = initAllSymbols()
	var NoSymbols = []string{
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "0",
		"A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z"}
	fmt.Println("Enter the type of Password, 1 for all random, 2 for all random but no symbols:")
	var PasswordType int
	_, err = fmt.Scan(&PasswordType)
	if err != nil {
		return
	}
	if PasswordType == 1 {
		for i := 0; i < len(Password); i++ {
			for j := 0; j < len(Password[i]); j++ {
				Password[i][j] = AllSymbols[utils.GetRandom(0, len(AllSymbols))]
			}
		}
	} else if PasswordType == 2 {
		for i := 0; i < len(Password); i++ {
			for j := 0; j < len(Password[i]); j++ {
				Password[i][j] = NoSymbols[utils.GetRandom(0, len(NoSymbols))]
			}
		}
	} else {
		log.Fatal("Invalid PasswordType: ", PasswordType)
	}
	for i := 0; i < utils.GetRandom(100, 500); i++ {
		var a = Password[utils.GetRandom(0, len(Password))][utils.GetRandom(0, len(Password[0]))]
		var b = Password[utils.GetRandom(0, len(Password))][utils.GetRandom(0, len(Password[0]))]
		var temp = a
		a = b
		b = temp
	}
	for i := 0; i < len(Password); i++ {
		utils.PrintStringArr(Password[i])
	}
}

func initAllSymbols() []string {
	var symbols []string
	for i := 32; i < 127; i++ {
		symbols = append(symbols, string(rune(i)))
	}
	return symbols
}
