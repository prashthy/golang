package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &a)
	fmt.Println("Hello, ", a)
}
