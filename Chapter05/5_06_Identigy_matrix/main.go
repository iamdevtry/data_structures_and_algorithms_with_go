package main

import "fmt"

func Identity(order int) [][]float64 {
	matrix := make([][]float64, order)

	for i := 0; i < order; i++ {
		matrix[i] = make([]float64, order)
		matrix[i][i] = 1
	}

	return matrix
}

func main() {
	fmt.Println(Identity(4))
}
