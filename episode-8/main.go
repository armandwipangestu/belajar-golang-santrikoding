package main

import "fmt"

// Example function
func say() {
	fmt.Println("Hello from function!")
}

// Example function with parameter
func greeting(name string) {
	fmt.Println("Hello", name)
}

// Example function with multiple parameter
func adds(a int, b int) {
	fmt.Println("Result:", a + b)
}

// Example function with return value
func times(a, b int) int {
	return a * b
}

// Example function with multiple return value
func count(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	say()
	greeting("king")
	adds(3, 6)

	result := times(3, 3)
	fmt.Println(result) // Output: 9

	add, time := count(3, 3)
	fmt.Println("Add:", add)
	fmt.Println("Time:", time)

	// Example function on Variable (Function as Value)
	f := func(name string) {
		fmt.Println("Hello,", name)
	}
	f("Golang")
}