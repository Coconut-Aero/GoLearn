package main

import (
	"GoLearn/utils"
	"fmt"
)

func main() {
	fmt.Println("Enter some number and I will tell you whether is it shorted.")
	fmt.Println("Enter the size of the list:")
	var size int
	_, err := fmt.Scan(&size)
	if err != nil {
		return
	}
	list := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Print("Enter the number #", i+1, ":")
		_, err := fmt.Scan(&list[i])
		if err != nil {
			return
		}
	}
	var IsShorted = IsSorted(list)
	fmt.Print("The list have ", size, " integers:")
	utils.PrintArr(list, size)
	if !IsShorted {
		fmt.Println("The list is not sorted.")
	} else {
		fmt.Println("The list is sorted.")
	}
}

func IsSorted(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}
