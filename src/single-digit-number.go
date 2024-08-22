package main

import (
	"GoLearn/utils"
	"fmt"
)

func main() {
	var (
		tmp   [100]int
		count [10]int
	)
	for i := 0; i < len(tmp); i++ {
		tmp[i] = utils.GetRandomExpect(0, 10, -1, 10)
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%5d", tmp[i*10+j])
		}
		fmt.Println()
	}
	for i := 0; i < 100; i++ {
		switch tmp[i] {
		case 0:
			count[0]++
			continue
		case 1:
			count[1]++
			continue
		case 2:
			count[2]++
			continue
		case 3:
			count[3]++
			continue
		case 4:
			count[4]++
			continue
		case 5:
			count[5]++
			continue
		case 6:
			count[6]++
			continue
		case 7:
			count[7]++
			continue
		case 8:
			count[8]++
			continue
		case 9:
			count[9]++
			continue
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Number ", i, " Appears ", count[i], " Time(s).")
	}
}
