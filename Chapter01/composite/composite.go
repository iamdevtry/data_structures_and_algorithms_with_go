package main

import "fmt"

// IComposite interface
type IComposite interface {
	perform()
}

// Leaflet struct
type Leaflet struct {
	name string
}

// Leaflet class method perform
func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}

// Branch struct
type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

// Branch class method perform
func (b *Branch) perform() {
	fmt.Println("Branch: " + b.name)

	for _, leaf := range b.leafs {
		leaf.perform()
	}

	for _, branch := range b.branches {
		branch.perform()
	}
}

// Branch class method add leaflet
func (b *Branch) add(l Leaflet) {
	b.leafs = append(b.leafs, l)
}

// Branch class method add branch
func (b *Branch) addBranch(newBranch Branch) {
	b.branches = append(b.branches, newBranch)
}

// Branch class method get getLeaflets
func (b *Branch) getLeaflets() []Leaflet {
	return b.leafs
}

func main() {
	var branch = Branch{name: "branch 1"}
	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}

	var branch2 = Branch{name: "branch 2"}
	branch.add(leaf1)
	branch.add(leaf2)
	branch.addBranch(branch2)
	branch.perform()
}
