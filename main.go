package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	//"os"
	"strings"
	"math/rand"
	"time"
)

var filename string = "dict.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func dummy() {
	mySl := make([]int, 3, 9)
	fmt.Println("Len: ", len(mySl))

	//b := [] string{"Penn", "Foo"}
	//fmt.Println(len(b))

}

func main() {
	content, err := ioutil.ReadFile(filename)
	check(err)
	lines := strings.Split(string(content), "\n")

	var spanWords = make([]string, len(lines))
	var engWords = make([]string, len(lines))

	//spanToEng = make(map[string]string)
	//engToSpan = make(map[string]string)

	for i, val := range lines {
		tmp := strings.Split(val, "=")
		if tmp[0] != "" {
			spanWords[i] = tmp[0]
			engWords[i] = tmp[1]
		}
		i++
	}
	fmt.Println("spanWords")
	for _, value := range spanWords {
		fmt.Println(value);
	}
	fmt.Println("engWords")
	for _, value := range engWords {
		fmt.Println(value);
	}
	dummy()
	numWords := len(spanWords)
	rand.Seed(time.Now().UnixNano())
	for {
		idx := rand.Intn(numWords)
		fmt.Println(spanWords[idx])
		var i byte
		fmt.Scanf("Press any key to continue: %c", &i)
		fmt.Println(engWords[idx])
		fmt.Println()
		fmt.Scanf("Press any key to continue: %c", &i)
	}
}
