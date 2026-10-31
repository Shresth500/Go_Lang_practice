## DATA TYPES
- String
- Number
    - Integer {int8, int64, uint8, uint64, intptr}
    - Float {float32, float64}
- Boolean
- Arrays
- Slices - just like array but it's more flexible than array
- Map: Key-value pair
- Structs
- Pointer

## VARIABLES

Variables :- Reference to a value 

var {variable_name} {data_type} = {value}

Syntax:- 

```
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
```

\n => to create new line

Printf:- formatting input-output
- process of inserting a custom string or variable in pre-defined text
- Format Specifiers:- 

| Format Specifiers | Uses                                  |
|:-------------------:|:---------------------------------------:|
| %v                | formats the value in a default format |
| %d                | format decimal integers               |
| %T                | Type of the value                     |
| %c                | Character                             |
| %q                | Quoted characters | string            |
| %s                | plain string                          |
| %t                | true/false                            |
| %f                | floating numbers                      |
| %.2f              | floating value upto 2 decimal places  |


## DECLARING VARIABLES

```
var user string
user = "Shresth"

var s,t = "foo", "bar"
```

Shorthand method:-
```
s := "Hello World"
```

```:=``` - Short Variable Declaration

#### Zero Values: 
- Default value given to variables according to Data Type