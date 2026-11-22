## LOOP

There is only one loop operator in Go and that is for loop

### Different types of for loop format
1. Format-1:
    ```
    for initilization;condition;increment {

    }
    ```
    Example
    ```
    for i:=1;i<10;i++{

    }

    for i int = 0; i < 10;i++{

    }
    ```
2. Format-2

    ```
    initilization

    for condition {
        incrementation
    }
    ```

    ```
    // Another way to write a for loop
	j := 0
	for j < 5 {
		fmt.Println("Hello, Go!")
		j++
	}
    ```

### Infinite loop
```
for {
    fmt.Println("Hello, Go!")
}

// Infinite loop with a break condition
i := 0
for {
    fmt.Println("Hello")
    i++
    if i == 10 {
        break
    }
}
```