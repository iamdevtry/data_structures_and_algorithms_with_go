package main

import "fmt"

func main() {
	var m [10]int

	for k := 0; k < 10; k++ {
		m[k] = k * 200
		fmt.Printf("Element [%d] = %d\n", k, m[k])
	}

}
