package main

import "fmt"

// h function which returns the product of parameters x and y
func h(x, y int) int {
	return x * y
}

// g function which returns x and y parameters after modification
func g(l, m int) (x, y int) {
	return 2 * l, 4 * m
}

func main() {
	fmt.Println(h(g(2, 1)))
}
