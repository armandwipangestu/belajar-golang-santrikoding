package main

import "fmt"

func main() {
	// Constanta without data type declaration
	const angka = 22
	fmt.Println("angka:", angka)

	// Constanta with data type declaration
	const pi float32 = 3.14
	const aplikasi string = "Belajar Golang"
	const version string = "v1.0.0"

	fmt.Println("pi:", pi)
	fmt.Println("aplikasi:", aplikasi)
	fmt.Println("version:", version)

	// Constanta with Grouping
	const (
		satu = 1
		dua = 2
		tiga = 3
	)

	fmt.Println("satu:", satu)
	fmt.Println("dua:", dua)
	fmt.Println("tiga:", tiga)

	// Constanta `iota` keyword (for sequential value)
	const (
		A = iota 	// 0
		B 			// 1
		C			// 2
	)

	fmt.Println("A:", A)
	fmt.Println("B:", B)
	fmt.Println("C:", C)
}