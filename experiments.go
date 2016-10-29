package main
import (
	"fmt"
)

const p string = "Test"
const q = 42

func variadic_func_sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func closure_outer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("string: ", p)
	// Variadic function
	fmt.Println(variadic_func_sum(1, 2, 3, 4))
	fmt.Println(variadic_func_sum(1,2,3))
	myarr := []int {1, 3, 5, 7, 9}
	fmt.Println(variadic_func_sum(myarr...))

	// Closure Function
	increment := closure_outer()
	fmt.Println("Printing output of closure function")
	fmt.Println(increment())
	fmt.Println(increment())
}
