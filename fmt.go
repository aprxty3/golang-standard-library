package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	firstName := "John"
	lastName := "Doe"
	age := 30

	fmt.Println("Hello '", firstName, lastName, "' you are", age, "years old")
	fmt.Printf("Hello '%s %s' you are %d years old", firstName, lastName, age)
}
