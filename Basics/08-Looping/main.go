package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, World!")
	}
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Another way to write a for loop
	j := 0
	for j < 5 {
		fmt.Println("Hello, Go!")
		j++
	}

	// Infinite loop with a break condition
	i := 0
	for {
		fmt.Println("Hello, Go!")
		i++
		if i == 3 {
			break
		}
	}
}
