package main

import (
	"fmt"
	"math/rand"
)

func printArr(arr [6]int) {
	for i := 0; i < 6; i++ {
		fmt.Printf("%3d", arr[i])
	}
	fmt.Println()
}
func main() {
	var array [6]int
	for i := 0; i < 6; i++ {
		fmt.Scan(&array[i])
	}
	printArr(array)
	var a = array[0]
	for i := 0; i < 5; i++ {
		array[i] = array[i+1]
	}
	array[5] = a
	printArr(array)
	for i := 0; i < 3; i++ {
		var m = array[i]
		var n = array[5-i]
		array[i] = n
		array[5-i] = m
	}
	printArr(array)
	for i := 0; i < 3; i++ {
		var q = rand.Intn(3)
		var p = rand.Intn(3) + 3
		var r = array[q]
		var s = array[p]
		array[q] = s
		array[p] = r
	}
	printArr(array)
}
