package main

import "fmt"

func rotate(matrix [][]int) [][]int {
	m := len(matrix)
	i, j := 0, 0

	newmAtrix := make([][]int, m)
	for l := range newmAtrix {
		newmAtrix[l] = make([]int, m)
	}

	for i < m {
		for j < m {
			newmAtrix[j][m-i-1] = matrix[i][j]
			j++
		}
		i++
		j = 0
	}
	return newmAtrix
}

func main() {
	fmt.Println("Rotate matrixcd")

	matrix := [][]int{
		{1, 2, 3, 1},
		{4, 5, 6, 1},
		{7, 8, 9, 1},
		{1, 1, 1, 1},
	}

	matrix1 := [][]int{
		{1, 7, 4, 1},
		{1, 8, 5, 2},
		{1, 9, 6, 3},
		{1, 1, 1, 1},
	}

	fmt.Println(matrix)
	outMatrix := rotate(matrix)
	fmt.Println(outMatrix)

	fmt.Println("Expected:")
	fmt.Println(matrix1)
}
