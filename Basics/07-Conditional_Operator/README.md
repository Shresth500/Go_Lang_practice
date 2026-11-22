## CONDITIONAL OPERATORS
### If-Else Operator

main.go
```
var a int = 10
	var b int = 20
	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}
```

### switch statements

```
ch := 'A'
	switch ch {
	case 'A':
		fmt.Println("Character is A")
	case 'B':
		fmt.Println("Character is B")
	default:
		fmt.Println("Character is neither A nor B")
	}
value := 100
	switch value {
	case 50:
		fmt.Println("Value is 50")
	case 100, 200:
		fmt.Println("Value is 100 or 200")
	default:
		fmt.Println("Value is neither 50, 100 nor 200")
	}
```

Output
```
Character is A
Value is 100 or 200
```

Output

- In the above switch case it will directly compare the case and it will not go after those steps

#### Using it with fallthrough
```
// fallthrough example
	switch value {
	case 100:
		fmt.Println("Value is 100")
		fallthrough
	case 200:
		fmt.Println("Value is 200")
	default:
		fmt.Println("Value is neither 100 nor 200")
	}
```

Output
```
Value is 100
Value is 200
```