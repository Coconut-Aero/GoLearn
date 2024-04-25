package utils

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

func Sort(res []int) []int {
	if !IsSorted(res) {
		for i := 0; i < len(res)-1; i++ {
			if res[i] > res[i+1] {
				var tmp int
				tmp = res[i]
				res[i] = res[i+1]
				res[i+1] = tmp
				i = 0
			}
		}
	}
	return res
}

func IsSorted(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}

func GetRandom(start int, end int, numbers ...int) int {
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

func AddMatrix(matrix1 [][]float64, matrix2 [][]float64) [][]float64 {
	res := make([][]float64, len(matrix1))
	for i := range res {
		res[i] = make([]float64, len(matrix1[i]))
	}
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[i]); j++ {
			res[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}
	return res
}
func MultiplyMatrix(matrix1 [][]float64, matrix2 [][]float64) [][]float64 {
	res := make([][]float64, len(matrix1))
	for i := range res {
		res[i] = make([]float64, len(matrix1[i]))
	}
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[i]); j++ {
			for k := 0; k < len(matrix1); k++ {
				res[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return res
}
