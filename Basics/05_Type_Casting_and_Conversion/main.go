package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza Shop")
	// using bufio package
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza:")
	rating, err := reader.ReadString('\n') // reading the string until the newline character
	if err != nil {
		fmt.Println("Error:", err)
	}
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64) // removing the newline character from the string and converting to float64
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Your rating:", numRating)
	}
}
