package main

import "fmt"

func main() {
	matrix := matrix(3)
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func matrix(n int) [][]int {
	matrix := make([][]int, n)
	count := 1
	startCol := 0
	startRow := 0
	endCol := n - 1
	endRow := n - 1

	// Initialize the matrix
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for count <= n*n {
		// Fill top row left to right
		for i := startCol; i <= endCol && count <= n*n; i++ {
			matrix[startRow][i] = count
			count++
		}
		startRow++

		// Fill right column top to bottom
		for i := startRow; i <= endRow && count <= n*n; i++ {
			matrix[i][endCol] = count
			count++
		}
		endCol--

		// Fill bottom row right to left
		for i := endCol; i >= startCol && count <= n*n; i-- {
			matrix[endRow][i] = count
			count++
		}
		endRow--

		// Fill left column bottom to top
		for i := endRow; i >= startRow && count <= n*n; i-- {
			matrix[i][startCol] = count
			count++
		}
		startCol++
	}

	return matrix
}
