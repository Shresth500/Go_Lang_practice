package main

import "fmt"

func main() {
	// Declare an array of integers with a fixed size of 5
	var numbers [5]int
	fmt.Println(numbers)

	// Assign values to the array
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println(numbers)

	data := [...]float64{1.1, 2.2, 3.3}
	fmt.Println(data)

	// Declare and initialize an array in one line
	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println(names)

	for index, value := range names {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// Multi-dimensional array
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}
}
