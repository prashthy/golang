package main
import (
	"fmt"
)

func main() {
	arr := []string{
		"1",
		"2",
		"3",
		"2",
		"3",
		"4",
		"7",
		"1",
		"2",
		"7",
	}
	m := make(map[string]bool)
	for _, val := range arr {
		_, exists := m[val]
		if !exists {
			arr[len(m)] = val
			m[val] = true
		}
	}
	arr = arr[:len(m)]
	fmt.Println("Deduplicated array: ", arr)

}



func sum(nums []int) int {
	return 0
}
