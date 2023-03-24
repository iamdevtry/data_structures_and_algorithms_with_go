package main

func main() {
	var matrix = [3][3]int{
		{1, 0, 0},
		{1, 1, 0},
		{2, 1, 1},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			print(matrix[i][j], " ")
		}
		println()
	}
}
