# Golang: Variables sample

![image](https://github.com/luiscoco/Golang-sample2-Variables/assets/32194879/da758e6f-766d-4d22-b89b-f2765f409122)

## 1. Source code

```go
package main

import (
    "fmt"
    "github.com/google/uuid"
)

func main() {
    var name string = "Go Programmer"
    var age int = 25
    isProgrammingFun := true // Short variable declaration

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Is Programming Fun?", isProgrammingFun)
	// Generate and print a UUID
	fmt.Println("Your unique ID is:", uuid.New().String())
}
```

## 2. How to run the application

```
go run variables.go
```

