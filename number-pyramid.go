package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 8; i++ {
		for j := 0; j < (20 - i + 1); j++ {
			fmt.Print("    ")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("%4d", math.Pow(2, float64(k)))
		}
		for l := i; l > -1; l++ {
			fmt.Printf("%4d", math.Pow(2, float64(l)))
		}
		fmt.Println()
	}
}
