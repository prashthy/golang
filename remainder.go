package main

import (
	"fmt"
)

func main() {
	var larger, smaller, result int
	fmt.Printf("Enter the larger integer: ")
	fmt.Scan(&larger) // Alternatively, you can do fmt.Scanf("%d\n", &larger)
	fmt.Printf("Enter the smaller integer: ")
	fmt.Scan(&smaller)
	result = larger % smaller
	fmt.Printf("Remainder = %d\n", result)
	fmt.Println("Remainder = ", result)
}

