package main

import "fmt"

func PrintSpiral(n int) []int {
	left, top, right, bottom := 0, 0, n-1, n-1
	size := n * n

	s := make([]int, size)

	var i int
	i = 0
	for left < right {
		var c int
		for c = left; c <= right; c++ {
			s[top*n+c] = i
			i++
		}
		top++
		var r int
		for r = top; r <= bottom; r++ {
			s[r*n+right] = i
			i++
		}
		right--
		if top == bottom {
			break
		}
		for c = right; c >= left; c-- {
			s[bottom*n+c] = i
			i++
		}
		bottom--
		for r = bottom; r >= top; r-- {
			s[r*n+left] = i
			i++
		}
		left++
	}
	s[top*n+left] = i
	return s
}

func main() {
	var n int
	n = 5
	var length int
	length = 2
	var i int
	var sketch int
	for i, sketch = range PrintSpiral(n) {
		fmt.Printf("%*d ", length, sketch)
		if i%n == n-1 {
			fmt.Println("")
		}
	}

}
