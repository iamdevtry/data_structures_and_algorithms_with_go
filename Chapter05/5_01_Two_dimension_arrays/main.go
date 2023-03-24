package main

func main() {
	var a [3][4]int
	a[1][2] = 1
	a[2][3] = 2
	a[1][1] = 3
	a[0][3] = 4
	a[2][2] = 5
	a[0][0] = 6
	a[2][0] = 7
	a[0][1] = 8
	a[1][3] = 9
	a[0][2] = 10

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			print(a[i][j], " ")
		}
		println()
	}
}
