package main

import "fmt"

func main() {
	var username string = "John Doe"
	fmt.Println("Username:", username)
	fmt.Printf("Variable is of Type: %T", username)

	var age int = 30
	fmt.Println("\nAge:", age)
	fmt.Printf("Variable is of Type: %T", age)

	var isActive bool = true
	fmt.Println("\nIs Active:", isActive)
	fmt.Printf("Variable is of Type: %T", isActive)

	var height float64 = 5.9
	fmt.Println("\nHeight:", height)
	fmt.Printf("Variable is of Type: %T", height)

	// default values
	var defaultString string
	var defaultInt int
	var defaultBool bool
	var defaultFloat float64

	fmt.Println("\nDefault String:", defaultString)
	fmt.Println("Default Int:", defaultInt)
	fmt.Println("Default Bool:", defaultBool)
	fmt.Println("Default Float:", defaultFloat)

	// implicit type declaration
	var implicitString = "Implicit String"
	var implicitInt = 42
	var implicitBool = false
	var implicitFloat = 3.14

	fmt.Println("\nImplicit String:", implicitString)
	fmt.Println("Implicit Int:", implicitInt)
	fmt.Println("Implicit Bool:", implicitBool)
	fmt.Println("Implicit Float:", implicitFloat)

	// no var style
	numberOfUser := 100
	fmt.Println("\nNumber of Users:", numberOfUser)
}
