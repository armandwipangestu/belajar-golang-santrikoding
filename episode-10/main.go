package main

import "fmt"

func main() {
	// Array declare first then assign the value later
	var numbers [3]int

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	fmt.Println("Value of Index 0 on Array numbers is:", numbers[0])
	fmt.Println("Value of Index 1 on Array numbers is:", numbers[1])
	fmt.Println("Value of Index 2 on Array numbers is:", numbers[2])

	// Array declare and assign the value at the same time
	number := [3]int{10, 20, 30}

	fmt.Println("Value of Index 0 on Array number is:", number[0])
	fmt.Println("Value of Index 1 on Array number is:", number[1])
	fmt.Println("Value of Index 2 on Array number is:", number[2])

	// Slice (not fixed array)
	age := []int{12, 15, 17, 20}
	fmt.Println(age[0]) // Output: 12

	age = append(age, 23)
	fmt.Println(age)

	// Create slice based on array
	tempArray := [5]int{1, 2, 3, 4, 5}
	slice := tempArray[1:4] // Get element from-1 to before-4 (1 until n-1)
	fmt.Println(slice) // [2, 3, 4]

	// Map (Key-Value or Dictionary on Python or Object on JavaScript)
	grade := map[string]string{
		"Foo": "A",
		"Bar": "B",
	}

	fmt.Println(grade["Foo"]) // Output: A
	fmt.Println(grade["Bar"]) // Output: B

	// Add data to Map
	grade["FooBar"] = "C"
	fmt.Println(grade["FooBar"])

	// Delete data from Map
	delete(grade, "Foo")

	// Check if key is exist on Map
	value, isExist := grade["JaneDoe"]

	if isExist {
		fmt.Println("Exist:", value)
	} else {
		fmt.Println("Can't find the key")
	}
}