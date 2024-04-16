package main

import (
	"fmt"
	"math"
)

func pi(i int) float64 {
	var result float64
	for j := 1; j < i; j++ {
		result += math.Pow(-1, float64(j+1)) / (float64(2.0*j - 1))
	}
	result *= 4
	return result
}

func main() {
	fmt.Printf("%9s%15s", "i", "m(i)")
	fmt.Println()
	for i := 1; i < 10; i++ {
		var res = pi(i)
		fmt.Printf("%9d%15.9f", i, res)
		fmt.Println()
	}
}
