package main

import (
	"fmt"
	"math/rand"
)

func getRandom(start int, end int, numbers ...int) int {
	var result int
	result = rand.Intn(end-start) + start
	for i := 0; i < len(numbers); i++ {
		if result == numbers[i] {
			result = rand.Intn(end-start) + start
			i = 0
		}
	}
	return result
}

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 15; j++ {
			var (
				res = getRandom(1, 100, 4, 8, 95, 93)
			)
			fmt.Printf("%4d", res)
		}
		fmt.Println()
	}
}
