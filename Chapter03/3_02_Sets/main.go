package main

import "fmt"

// Set class
type Set struct {
	integerMap map[int]bool
}

// New create the map of integer and bool
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// ContainsElement checks if element is in the set
func (set *Set) ContainsElement(element int) bool {
	exists := set.integerMap[element]

	return exists
}

// AddElement adds the element to the set
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

// DeleteElement deletes the element from the set
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

// Intersect  method returns the set which intersects with anotherSet
func (set *Set) Intersect(anotherSet *Set) *Set {
	intersectSet := Set{}
	intersectSet.New()
	for value := range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return &intersectSet
}

// Union method return the set which i union of the set with anotherSet
func (set *Set) Union(another *Set) *Set {
	unionSet := Set{}
	unionSet.New()

	for value := range set.integerMap {
		unionSet.AddElement(value)
	}

	for value := range another.integerMap {
		unionSet.AddElement(value)
	}
	return &unionSet
}

func main() {
	set := Set{}
	set.New()

	set.AddElement(1)
	set.AddElement(2)

	fmt.Println("initial set: ", set)
	fmt.Println(set.ContainsElement(1))
	fmt.Println(set.ContainsElement(3))

	anotherSet := Set{}
	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)

	fmt.Println("intersect")
	fmt.Printf("%p\n : ", &anotherSet)
	fmt.Println(set.Intersect(&anotherSet))
	fmt.Println("union")
	fmt.Printf("%p\n : ", &anotherSet)
	fmt.Println(set.Union(&anotherSet))
}
