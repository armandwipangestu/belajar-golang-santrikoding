package main

import "fmt"

// Can't change value without pointer
func changeA(a int) {
	a = 100
}

// Can change value with pointer
func changeB(b *int) {
	*b = 100
}

func main() {
	// Get Address Of Variable using `&`
	var number int = 10
	var ptrNumber *int = &number

	fmt.Println("Value number is:", number)
	fmt.Println("Address number is:", &number)
	fmt.Println("Pointer number is:", ptrNumber)

	// Get Value from Pointer using `*`
	fmt.Println("Value number from pointer is:", *ptrNumber)

	// Test change value without pointer
	a := 10
	changeA(a)
	fmt.Println("Test change value without pointer:", a) // Value still 10

	// Test change value with pointer
	b := 10
	changeB(&b)
	fmt.Println("Test change value with pointer:", b) // Value changed to 100

	// Pointer Default using `new`
	age := new(int)
	*age = 22
	fmt.Println(*age)
}