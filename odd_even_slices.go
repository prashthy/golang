package main

import "fmt"

func main() {

	oddSlc := []int{1,3,5,7}
	evenSlc := []int{2,4,6,8}
	copy(evenSlc, oddSlc[2:])
	fmt.Println("evenSlc = ", evenSlc)

	fmt.Println(find_greatest(10, 43, 64, 1, 334, 65))
	numbers := []int{6565, 232, 343, 45454, 2324}
	fmt.Println(find_greatest(numbers...))

}

// Assumes there are no negative numbers
func find_greatest(vals ...int) int {
	var max int = 0
	for _, val := range vals{
		if val > max {
			max = val
		}
	}
	return max
}
