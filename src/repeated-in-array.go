package main

import (
	"GoLearn/utils"
	"fmt"
)

func main() {
	fmt.Print("Enter the number of integers:")
	var length int
	_, err := fmt.Scan(&length)
	if err != nil {
		return
	}
	nums := make([]int, length)
	read := make([]bool, length)
	fmt.Print("Enter the integers between 1~100:")
	for i := 0; i < length; i++ {
		_, err := fmt.Scan(&nums[i])
		if err != nil {
			return
		}
	}

	for i := 0; i < len(nums); i++ {
		if read[i] {
			continue
		}
		read[i] = true
		var tmp = nums[i]
		var time = 0
		var sortedNums = utils.Sort(nums)
		for j := i; j < len(sortedNums); j++ {
			if tmp == sortedNums[j] {
				read[j] = true
				time++
			}
		}
		if time == 0 {
			fmt.Print("No repeated integer")
		} else if time == 1 {
			fmt.Print(sortedNums[i], " occurs ", time, " time")
		} else {
			fmt.Print(sortedNums[i], " occurs ", time, " times")
		}
		fmt.Println()
	}
}
