package main

import "fmt"

// Define interface
type CanWalk interface {
	Walk()
}

// Define struct
type Human struct {
	Name string
}

type Robot struct {
	Code string
}

// Binding method to the struct, with the method name same as interface name
func (h Human) Walk() {
	fmt.Println(h.Name, "on walk")
}

func (r Robot) Walk() {
	fmt.Println("Robot", r.Code, "on automatic walk")
}

// Create function receive an instance of struct / object
func Move(c CanWalk) {
	c.Walk()
}

// Create an empty interface (can receive any data type)
func PrintData(data interface{}) {
	fmt.Println(data)
}

func main() {
	// Create an instance of struct
	h := Human{Name: "Foo"}
	r := Robot{Code: "Apollo"}

	Move(h) // Output: Foo on walk
	Move(r) // Output: Robot Apollo on automatic walk

	PrintData(h) // Output: {Foo}
}