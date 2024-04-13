package main

import (
	"fmt"
	"math/rand"
)

func printArr(arr []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("%3d", arr[i])
	}
	fmt.Println()
}
func main() {
	var array []int
	var size = 6
	for i := 0; i < 6; i++ {
		_, err := fmt.Scan(&array[i])
		if err != nil {
			return
		}
	}
	printArr(array, size)
	var a = array[0]
	for i := 0; i < 5; i++ {
		array[i] = array[i+1]
	}
	array[5] = a
	printArr(array, size)
	for i := 0; i < 3; i++ {
		var m = array[i]
		var n = array[5-i]
		array[i] = n
		array[5-i] = m
	}
	printArr(array, size)
	for i := 0; i < 3; i++ {
		var q = rand.Intn(3)
		var p = rand.Intn(3) + 3
		var r = array[q]
		var s = array[p]
		array[q] = s
		array[p] = r
	}
	printArr(array, size)
}
