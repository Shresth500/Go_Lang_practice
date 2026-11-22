# Go Operators

Operators are symbols used to perform operations on variables and values.

## Arithmetic Operators

| Operator | Description | Example |
| --- | --- | --- |
| `+` | Addition | `a + b` |
| `-` | Subtraction | `a - b` |
| `*` | Multiplication | `a * b` |
| `/` | Division | `a / b` |
| `%` | Modulus (remainder) | `a % b` |

```go
a := 10
b := 5

fmt.Println(a + b) // 15
fmt.Println(a - b) // 5
fmt.Println(a * b) // 50
fmt.Println(a / b) // 2
fmt.Println(a % b) // 0
```

## Comparison Operators

| Operator | Description | Example |
| --- | --- | --- |
| `==` | Equal to | `a == b` |
| `!=` | Not equal to | `a != b` |
| `>` | Greater than | `a > b` |
| `<` | Less than | `a < b` |
| `>=` | Greater than or equal to | `a >= b` |
| `<=` | Less than or equal to | `a <= b` |

```go
fmt.Println(a == b) // false
fmt.Println(a != b) // true
fmt.Println(a > b)  // true
fmt.Println(a < b)  // false
fmt.Println(a >= b) // true
fmt.Println(a <= b) // false
```

## Logical Operators

| Operator | Description | Example |
| --- | --- | --- |
| `&&` | Logical AND | `x && y` |
| `||` | Logical OR | `x || y` |
| `!` | Logical NOT | `!x` |

```go
x := true
y := false

fmt.Println(x && y) // false
fmt.Println(x || y) // true
fmt.Println(!x)      // false
```

## Assignment Operators

| Operator | Description | Example |
| --- | --- | --- |
| `=` | Assign | `value = 20` |
| `+=` | Add and assign | `value += 5` |
| `-=` | Subtract and assign | `value -= 5` |
| `*=` | Multiply and assign | `value *= 2` |
| `/=` | Divide and assign | `value /= 4` |
| `%=` | Calculate remainder and assign | `value %= 4` |

```go
value := 20
value += 5
value -= 5
value *= 2
value /= 4
value %= 4

fmt.Println(value) // 2
```

## Bitwise Operators

| Operator | Description | Example |
| --- | --- | --- |
| `&` | AND | `num1 & num2` |
| `|` | OR | `num1 | num2` |
| `^` | XOR | `num1 ^ num2` |
| `&^` | AND NOT | `num1 &^ num2` |
| `<<` | Left shift | `num1 << 1` |
| `>>` | Right shift | `num1 >> 1` |

```go
num1 := 12 // 1100
num2 := 10 // 1010

fmt.Println(num1 & num2)  // 8
fmt.Println(num1 | num2)  // 14
fmt.Println(num1 ^ num2)  // 6
fmt.Println(num1 &^ num2) // 4
fmt.Println(num1 << 1)    // 24
fmt.Println(num1 >> 1)    // 6
```