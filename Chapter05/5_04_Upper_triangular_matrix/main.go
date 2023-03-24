package main

func main() {
	var matrix = [3][3]int{
		{1, 2, 3},
		{0, 1, 4},
		{0, 0, 1},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			print(matrix[i][j], " ")
		}
		println()
	}
}
