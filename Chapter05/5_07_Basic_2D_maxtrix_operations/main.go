package main

import "fmt"

func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			matrix1[i][j] += matrix2[i][j]
		}
	}

	return matrix1
}

func subtract(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			matrix1[i][j] -= matrix2[i][j]
		}
	}

	return matrix1
}

func multiply(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var product [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				product[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return product
}

func transpose(matrix1 [2][2]int) [2][2]int {
	var m int
	var l int
	var transMatrix [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			transMatrix[l][m] = matrix1[m][l]
		}
	}
	return transMatrix
}

func determinant(matrix1 [2][2]int) float32 {
	var det float32
	det = float32(matrix1[0][0]*matrix1[1][1] - matrix1[0][1]*matrix1[1][0])
	return det
}

func inverse(matrix1 [2][2]int) [2][2]int {
	var invMatrix [2][2]int
	var det float32
	det = determinant(matrix1)
	invMatrix[0][0] = matrix1[1][1] / int(det)
	invMatrix[0][1] = -matrix1[0][1] / int(det)
	invMatrix[1][0] = -matrix1[1][0] / int(det)
	invMatrix[1][1] = matrix1[0][0] / int(det)
	return invMatrix
}

func main() {
	matrix1 := [2][2]int{
		{4, 5},
		{1, 2},
	}
	matrix2 := [2][2]int{
		{6, 7},
		{3, 4},
	}
	var sum [2][2]int
	sum = add(matrix1, matrix2)
	fmt.Println(sum)

	sum = subtract(matrix1, matrix2)
	fmt.Println(sum)

	sum = multiply(matrix1, matrix2)
	fmt.Println(sum)

	sum = transpose(matrix1)
	fmt.Println(sum)

	fmt.Println(determinant(matrix1))
	fmt.Println(inverse(matrix1))

}
