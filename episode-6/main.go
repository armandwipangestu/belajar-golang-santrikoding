package main

import "fmt"

func main() {
	// Operator Arithmetic
	const a float32 = 3.14
	const b float32 = 22.7
	result := a + b
	fmt.Println("result of a+b is:", result)

	// Operator Comparison
	x := 5
	y := 7
	result2 := x < y
	result3 := x == y
	fmt.Println("is x < y:", result2)
	fmt.Println("is x == y:", result3)

	// Operator Logical
	var j bool = true
	var k bool = false
	result4 := j && k
	result5 := j || k
	result6 := !j
	fmt.Println("result of j&&k is:", result4)
	fmt.Println("result of j||k is:", result5)
	fmt.Println("result of !j is:", result6)
}