## TYPE CASTING

Type Casting: 

- for conversion between integer to float
```
var i int = 90
var f float64 = float64(i)
fmt.Println("%.2f",f)
```

- Conversion to or from string

```
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

```

- Itoa(): 
    - Converts integer to string 
    - returns  one value that is string formed

- Atoi():
    - converts string to integer
    - returns two values, corresponding integer and error if any