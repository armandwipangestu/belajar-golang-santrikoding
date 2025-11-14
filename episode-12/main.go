package main

import "fmt"

type Student struct {
	Name string
	Age int
}

// Create method and assign to struct `Student`
func (s Student) Greeting() {
	fmt.Printf("Hello, my name is %s and my age is %d\n", s.Name, s.Age)
}

// Create method but return the value
func (s Student) BirthYear(yearNow int) int {
	return yearNow - s.Age
}

// Create method but Value Receiver (copy only)
func (s Student) ChangeNewName(name string) {
	s.Name = name
}

// Create method but Pointer Receiver (direct override)
func (s *Student) ChangeName(name string) {
	s.Name = name
}

func main() {
	// Create an instance based on struct Student
	Student1 := Student{
		Name: "FooBar",
		Age: 22,
	}
	// Call the simple method
	Student1.Greeting()

	// Call the return value method
	fmt.Println("My born year is:", Student1.BirthYear(2025))

	// Call method but receiver value
	Student1.ChangeNewName("BarFoo")
	// This should be not override the value
	fmt.Println("New name is:", Student1.Name) // FooBar

	// Call method but pointer value
	Student1.ChangeName("FizzBuzz")
	// This should be change the value
	fmt.Println("New name is:", Student1.Name) // FizzBuzz
}