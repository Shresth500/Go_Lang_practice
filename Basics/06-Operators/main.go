package main

import "fmt"

func main() {
	// Arithmetic Operators
	a := 10
	b := 5

	fmt.Println("Arithmetic Operators:")
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// Comparison Operators
	fmt.Println("\nComparison Operators:")
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)

	// Logical Operators
	fmt.Println("\nLogical Operators:")
	x := true
	y := false
	fmt.Println("x && y:", x && y)
	fmt.Println("x || y:", x || y)
	fmt.Println("!x:", !x)

	// Assignment Operators
	fmt.Println("\nAssignment Operators:")
	value := 20
	fmt.Println("Initial value:", value)

	value += 5
	fmt.Println("value += 5:", value)
	value -= 5
	fmt.Println("value -= 5:", value)
	value *= 2
	fmt.Println("value *= 2:", value)
	value /= 4
	fmt.Println("value /= 4:", value)
	value %= 4
	fmt.Println("value %= 4:", value)

	// Bitwise Operators
	fmt.Println("\nBitwise Operators:")
	num1 := 12 // 1100 in binary
	num2 := 10 // 1010 in binary
	fmt.Println("num1 & num2 =", num1&num2)
	fmt.Println("num1 | num2 =", num1|num2)
	fmt.Println("num1 ^ num2 =", num1^num2)
	fmt.Println("num1 &^ num2 =", num1&^num2)
	fmt.Println("num1 << 1 =", num1<<1)
	fmt.Println("num1 >> 1 =", num1>>1)
}
