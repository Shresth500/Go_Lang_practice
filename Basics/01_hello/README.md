## Basics

main.go file

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

```

- Every go file should start with package declaration
- To write package (develop package to store):
    ```
    package <package_name>
    ```
- import: To improt packages from other files
- fmt: package used for input/output

- // is for comment
- main(): main functions are special functions for executing program

- fmt.Println() - function inside fmt for printing
- /**/ -> multiline Comment
- To run a go file
```
go run <go_file.go>
```

go.mod file

- Core configuration file for dependency management in Go.
- acts as project blueprint that tracks your application module name, the requirement Go lang version and exact version of external libraries

- To make a go.mod file
```
go mod init <package_name>
```