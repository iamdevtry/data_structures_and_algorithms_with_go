package main

func main() {
	matrix := [1][3]int{
		{1, 2, 3},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			print(matrix[i][j], " ")
		}
		println()
	}
}
