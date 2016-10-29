package main
import (
	"fmt"
)

func main() {
	myarr := []int {
		48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83,27,
			19, 97, 9, 17,
	}

	min := myarr[0]
	min_idx := 0
	for i, value := range myarr {
		if value < min {
			min = value
			min_idx = i
		}
	}
	fmt.Println("Minimum value: ", min, "Index: ", min_idx)
}
