package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(len(matrix[0]))
	fmt.Println(transpose(matrix))
}

func transpose(matrix [][]int) [][]int {

	// creating columns
	transposeMatrix := make([][]int, len(matrix[0]))

	// creating rows
	for i := range transposeMatrix {
		transposeMatrix[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			transposeMatrix[j][i] = matrix[i][j]
		}
	}

	return transposeMatrix
}
