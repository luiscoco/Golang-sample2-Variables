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

