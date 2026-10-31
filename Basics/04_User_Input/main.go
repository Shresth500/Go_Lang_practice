package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	// using bufio package
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza:")
	review, err := reader.ReadString('\n') // reading the string until the newline character
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Your review:", review)

}
