package main

import (
	"GoLearn/utils"
	"fmt"
)

func main() {
	var size = 3
	matrix1 := make([][]float64, size)
	for i := range matrix1 {
		matrix1[i] = make([]float64, size)
	}

	matrix2 := make([][]float64, size)
	for i := range matrix2 {
		matrix2[i] = make([]float64, size)
	}

	fmt.Println("Enter Matrix 1:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			_, err := fmt.Scan(&matrix1[i][j])
			if err != nil {
				return
			}
		}
	}
	fmt.Println("Enter Matrix 2:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			_, err := fmt.Scan(&matrix2[i][j])
			if err != nil {
				return
			}
		}
	}
	var res = utils.AddMatrix(matrix1, matrix2)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%5.1f", matrix1[i][j])
		}
		if i == 1 {
			fmt.Print("  +  ")
		} else {
			fmt.Print("     ")
		}
		for j := 0; j < 3; j++ {
			fmt.Printf("%5.1f", matrix2[i][j])
		}
		if i == 1 {
			fmt.Print("  =  ")
		} else {
			fmt.Print("     ")
		}
		for j := 0; j < 3; j++ {
			fmt.Printf("%5.1f", res[i][j])
		}
		fmt.Println()
	}
}

func addMatrix(matrix1 [][]float64, matrix2 [][]float64) [][]float64 {
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
