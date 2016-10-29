package main
import (
	"fmt"
)

func main() {
	myslice := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	myslice = myslice[2:4]
	fmt.Printf("myslice = ", myslice)
	myslice = myslice[-1:3]
	fmt.Printf("myslice = ", myslice)
}


