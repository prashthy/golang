// Write a function which takes an integer. The function will have two returns.
// The first return should be the argument divided by 2. The second return
// should be a bool that let’s us know whether or not the argument was even.
// For example:
// half(1) returns (0, false)
// half(2) returns (1, true)

package main
import "fmt"
func multiple_rets(input int) (int, bool) {
	return input/2, (input%2 == 0)
}

func main() {
	half := func(input int) (int, bool) {
		return input/2, (input%2 == 0)
	}
	fmt.Println("half: ", half(5))
}

/*
func main() {
	a, b := multiple_rets(100)
	fmt.Println("100: ", a, b)
	fmt.Println(multiple_rets(100))
	a, b = multiple_rets(49)
	fmt.Println("49: ", a, b)
	a, b = multiple_rets(0)
	fmt.Println("0: ", a, b)
}
*/

// Same solution using a function expression:
/*func main() {

	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	fmt.Println(half(5))
}*/
