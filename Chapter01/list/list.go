package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(9)
	intList.PushBack(2)
	intList.PushBack(0)
	intList.PushBack(1)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
