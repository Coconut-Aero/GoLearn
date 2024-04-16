package main

import (
	"fmt"
	"math/rand"
)

func PrintArr(arr []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("%3d", arr[i])
	}
	fmt.Println()
}
func main() {
	var size = 6
	array := make([]int, size)
	for i := 0; i < 6; i++ {
		_, err := fmt.Scan(&array[i])
		if err != nil {
			return
		}
	}
	PrintArr(array, size)
	var a = array[0]
	for i := 0; i < 5; i++ {
		array[i] = array[i+1]
	}
	array[5] = a
	PrintArr(array, size)
	for i := 0; i < 3; i++ {
		var m = array[i]
		var n = array[5-i]
		array[i] = n
		array[5-i] = m
	}
	PrintArr(array, size)
	for i := 0; i < 3; i++ {
		var q = rand.Intn(3)
		var p = rand.Intn(3) + 3
		var r = array[q]
		var s = array[p]
		array[q] = s
		array[p] = r
	}
	PrintArr(array, size)
}
