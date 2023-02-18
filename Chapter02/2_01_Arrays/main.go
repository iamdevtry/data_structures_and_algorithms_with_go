package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 5, 7, 8}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	for i, value := range arr {
		fmt.Printf("arr[%d] = %d\n", i, value)
	}

	for _, value := range arr {
		fmt.Println("blank range", value)
	}
}
