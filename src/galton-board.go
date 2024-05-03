package main

import (
	"GoLearn/utils"
	"fmt"
)

func main() {
	var numBall int
	var numSlots int
	var numPath int
	var ballWhere int = 0
	fmt.Print("Enter the number of balls to drop: ")
	_, err := fmt.Scan(&numBall)
	if err != nil {
		return
	}
	fmt.Print("Enter the number of slots in the machine: ")
	_, err = fmt.Scan(&numSlots)
	if err != nil {
		return
	}
	numPath = numSlots - 1
	numSlotBall := make([]int, numSlots)
	numPathBall := make([]bool, numPath)
	for i := 0; i < numBall; i++ {
		for j := 0; j < numPath; j++ {
			numPathBall[j] = utils.BoolRandom()
			if numPathBall[j] {
				fmt.Print("L")
			} else {
				fmt.Print("R")
				ballWhere++
			}
		}
		numSlotBall[ballWhere-1]++
		ballWhere = 0
		fmt.Println()
	}
	for i := 0; i < len(numSlotBall); i++ {
		if numSlotBall[i] != 0 {
			fmt.Println("Slot ", i, " have ", numSlotBall[i], ".")
		}
	}
}
