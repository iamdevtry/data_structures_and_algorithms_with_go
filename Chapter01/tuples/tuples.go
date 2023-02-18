package main

import "fmt"

func powerSeries(a int) (int, int, error) {
	return a * a, a * a * a, nil
}

func main() {
	square, cube, err := powerSeries(2)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Square : ", square, " Cube : ", cube)
}
