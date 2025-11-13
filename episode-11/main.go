package main

import "fmt"

// Create a nested struct (struct in struct)
type Address struct {
	City string
	Street string
}

type CollegeStudents struct {
	Name string
	Age int
	Address Address
}

// Change value age from struct using pointer
func changeAge(CollegeStudent *CollegeStudents) {
	CollegeStudent.Age = 25
}

func main() {
	// Create struct for `Students` data type
	type Students struct {
		Name string
		Age int
		IsActive bool
	}

	// Create an instance `Students1` based on struct `Students`
	// then assign the value of each data
	var Student1 Students
	Student1.Name = "John Doe"
	Student1.Age = 20
	Student1.IsActive = true

	fmt.Println(Student1)

	// Create an instance `Student2` based on struct `Students`
	// but using literal method (direct assign when instance created)
	Student2 := Students{
		Name: "Jane Doe",
		Age: 22,
		IsActive: false,
	}

	// Access the field
	fmt.Println("Name:", Student2.Name)
	fmt.Println("Age:", Student2.Age)
	fmt.Println("IsActive:", Student2.IsActive)

	// Using nested struct
	CollegeStudent1 := CollegeStudents{
		Name: "FooBar",
		Age: 22,
		Address: Address{
			City: "New York",
			Street: "Wall Street",
		},
	}

	fmt.Println(CollegeStudent1.Address.City) // Output: New York\

	// Pointer to Struct
	CollegeStudent2 := CollegeStudents{
		Name: "BarFoo",
		Age: 20,
	}
	changeAge(&CollegeStudent2)
	fmt.Println(CollegeStudent2.Age) // Output: 25
}