package main

import "fmt"

// IProcess2 interface
type IProcess2 interface {
	process()
}

// ProcessClass struct
type ProcessClass struct {
}

// ProcessClass method process
func (p *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

// ProcessDecorate struct
type ProcessDecorate struct {
	processInstance *ProcessClass
}

// ProcessDecorate class method process
func (pd *ProcessDecorate) process() {
	if pd.processInstance == nil {
		fmt.Println("ProcessDecorate process")
	} else {
		fmt.Print("ProcessDecorate process and ")
		pd.processInstance.process()
	}
}

// main method
func main() {
	var process = &ProcessClass{}
	var decorator = &ProcessDecorate{}
	decorator.process()
	decorator.processInstance = process
	decorator.process()
}
