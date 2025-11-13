package main

import (
	"fmt"
	"time"
)

func exampleFizzBuzz() {
	var numbers []int = []int{31, 23, 68, 92, 123, 55, 21, 5, 95}

	for _, value := range numbers {
		if value % 2 == 0 {
			fmt.Printf("Value of %v is FizzBuzz\n", value)
		} else if value % 3 == 0 {
			fmt.Printf("Value of %v is Fizz\n", value)
		} else if value % 5 == 0 {
			fmt.Printf("Value of %v is Buzz\n", value)
		} else {
			fmt.Printf("Value of %v is nothing\n", value)
		}
	}
}

func main() {
	exampleFizzBuzz()

	// if else
	age := 22

	if age > 22 {
		fmt.Println("Tuwak luuhhh")
	} else if age == 22 {
		fmt.Println("Your prime age")
	} else if age < 22 {
		fmt.Println("Bomcil kematian")
	} else {
		fmt.Println("💀")
	}

	// switch case
	const day = "Sabtu"

	switch day {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Malasssss!")
	case "Sabtu", "Minggu":
		fmt.Println("Amjay mantap kang!")
	default:
		fmt.Println("💀")
	}

	// for loop standard (like at c/c++)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// infinite loop (while true)
	for {
		fmt.Println("Never give up coy...")
		time.Sleep(3 * time.Second) // Pause for 3 seconds
		fmt.Println("💀")
		break
	}

	// for loop using range
	numbers := []int{1, 2, 3, 4}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// without index

	for _, value := range numbers {
		fmt.Println(value)
	}
}