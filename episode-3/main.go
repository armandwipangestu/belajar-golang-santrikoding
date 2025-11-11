package main

import "fmt"

func printTypeAndValue(name string, v interface{}) {
	fmt.Printf("Variable Name: %s\n", name)
	fmt.Printf("Data Type: %T\n", v)
	fmt.Printf("Value: %v\n", v)
	fmt.Println()
}

func main() {
	// Int Data Type
	var age int = 22
	printTypeAndValue("age", age)

	// Float Data Type
	var tall float32 = 175.5
	printTypeAndValue("tall", tall)

	// String Data Type
	var name string = "Arman"
	printTypeAndValue("name", name)

	// Short Declaration Data Type
	message := `Hello,
	This is text
	with newline.`
	printTypeAndValue("message", message)

	// Zero Value Data Type
	var number int
	printTypeAndValue("number", number) // This should be print: 0

	// Convertion Data Type
	var x int = 10
	printTypeAndValue("x", x)
	var y float64 = float64(x)
	printTypeAndValue("y", y)
}