package main

import "fmt"

// twiceValue method given slice of int type
func twiceValue(slice []int) {
	for i, value := range slice {
		slice[i] = 2 * value
	}
}

func main() {
	var slice = []int{1, 3, 5, 6}

	slice = append(slice, 8)
	fmt.Println("Capacity : ", cap(slice))
	fmt.Println("Length : ", len(slice))

	fmt.Println("==================================================")
	fmt.Println("Twice Slice")
	twiceValue(slice)
	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Println("==================================================")
	fmt.Println("=============Two-dimensional slices===============")
	//twoDArray := [8][8]int{}
	//twoDArray[3][6] = 18
	//twoDArray[7][4] = 3
	//for i := 0; i < len(twoDArray[i]); i++ {
	//	for j := 0; j < len(twoDArray[j]); j++ {
	//		fmt.Printf("%d ", twoDArray[i][j])
	//	}
	//	fmt.Println()
	//
	//}

	//var rows int
	//var cols int
	//rows = 7
	//cols = 9
	//var twodslices = make([][]int, rows)
	//var i int
	//for i = range twodslices {
	//	twodslices[i] = make([]int, cols)
	//}
	//fmt.Println(twodslices)

	var arr = []int{5, 6, 7, 8, 9}
	var slic1 = arr[:3]
	fmt.Println("slice 1", slic1)
	var slic2 = arr[1:5]
	fmt.Println("slice 2", slic2)
	var slic3 = append(slic2, 12)
	fmt.Println("slice3", slic3)

	fmt.Printf("%p\n", arr)
	fmt.Printf("%p\n", slic1)
	fmt.Printf("%p\n", slic2)
	fmt.Printf("%p\n", slic3)
}
