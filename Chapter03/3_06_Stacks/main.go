package main

import (
	"fmt"
	"strconv"
)

// Element class
type Element struct {
	elementValue int
}

// String method on Element class
func (element *Element) String() string {
	return strconv.Itoa(element.elementValue)
}

// Stack is a basic LIFO stack that resizes as needed
type Stack struct {
	elements     []*Element
	elementCount int
}

// New returns a new stack
func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

// Push adds a node to the stack
func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount = stack.elementCount + 1
}

// Pop removes  and returns a node form the stack in last to first order
func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}

	length := len(stack.elements)
	element := stack.elements[length-1]
	//stack.elementCount = stack.elementCount - 1

	if length > 1 {
		stack.elements = stack.elements[:length-1]
	} else {
		stack.elements = stack.elements[0:]
	}
	stack.elementCount = len(stack.elements)
	return element
}

func main() {
	stack := Stack{}
	stack.New()

	element1 := Element{3}
	element2 := Element{5}
	element3 := Element{7}
	element4 := Element{9}
	stack.Push(&element1)
	stack.Push(&element2)
	stack.Push(&element3)
	stack.Push(&element4)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}
