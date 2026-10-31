package main

import "fmt"

var a int = 10 // Global level variable

func main() {
	fmt.Println("Variable Scopes")
	fmt.Println("Global level variable 'a':", a)

	// Function level variable / Local level variable
	var b int = 20
	fmt.Println("Function level variable 'b':", b)

	// Block level variable
	{
		var c int = 30
		fmt.Println("Block level variable 'c':", c)
		{
			var d int = 40
			fmt.Println("Nested Block level variable 'd':", d)
			fmt.Println("Accessing Global level variable 'a' inside nested block:", a)
			fmt.Println("Accessing Function level variable 'b' inside nested block:", b)
			fmt.Println("Accessing Block level variable 'c' inside nested block:", c)
		}
	}
}
