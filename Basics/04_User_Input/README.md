## USER INPUT

user input

```
package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello, " + name + "!")

	var username string
	var isStudent bool

	fmt.Println("Enter username and Student: ")
	count, err := fmt.Scanf("%s %t", &username, &isStudent)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Number of inputs successfully scanned:", count)
		fmt.Scanf("%s %t", &username, &isStudent)
		fmt.Println("Username:", username)
		fmt.Println("Is Student:", isStudent)
	}
}


```

```&name``` - Similar to reference and pointers

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

Use the access specifier while taking the input and output

fmt.Scanf() return values: 
- count - the no of arguments the function writes to
- err - any error thrown during execution of function

### Find the Type of Variable
- %T format specifier
- reflect.TypeOf function from the reflect package
	- reflect.TypeOf()


### Using the bufio for getting the user input
```
reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza:")
	review, err := reader.ReadString('\n') // reading the string until the newline character
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Your review:", review)
```