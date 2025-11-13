package main

import "fmt"

// Global Variable only using `var` keyword, can't use short declaration using `:=`
var version = "v1.0.0"

func main() {
	// Variable Declaration using `var` keyword and Type Declaration
	var name string = "Arman"
	var age int = 22
	var is_male bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Male:", is_male)
	
	// Variable Declaration using `var` keyword but Type Inference
	var kota = "Bandung"
	
	fmt.Println("Kota:", kota)

	// Variable Short Declaration using `:=`
	judul := "Belajar Golang"
	nilai := 90

	fmt.Println("Judul:", judul)
	fmt.Println("Nilai:", nilai)

	// Variable Multiple Declaration
	var a, b, c int = 1, 2, 3

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	x, y, z := "A", "B", "C"

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Change variable value
	var jumlah = 10
	fmt.Println("jumlah before edit:", jumlah)

	jumlah = 15
	fmt.Println("jumlah after edit:", jumlah)

	// Call global variable
	fmt.Println("Version:", version)
}