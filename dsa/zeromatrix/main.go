package main

import (
	"fmt"
)

func makeZeros(matrix [][]int) [][]int {
	m := len(matrix)
	i, j := 0, 0
	var row []int
	var col []int

	for i < m {
		for j < m {
			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
			j++
		}
		i++
		j = 0
	}

	i, j = 0, 0
	for i < m {
		for j < m {
			if check(i, row) || check(j, col) {
				matrix[i][j] = 0
			}
			j++
		}
		i++
		j = 0
	}

	return matrix
}

func check(i int, arr []int) bool {
	for _, v := range arr {
		if i == v {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("make Zeros matrixcd")

	matrix := [][]int{
		{1, 2, 0, 1},
		{4, 5, 6, 1},
		{7, 0, 9, 1},
		{1, 0, 1, 1},
	}

	matrix1 := [][]int{
		{0, 0, 0, 0},
		{4, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	fmt.Println(matrix)
	outMatrix := makeZeros(matrix)
	fmt.Println(outMatrix)

	fmt.Println("Expected:")
	fmt.Println(matrix1)
}
