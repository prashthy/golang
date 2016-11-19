package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"os"
	"strings"
	"bufio"
	"math/rand"
	"time"
	"flag"
)

var filename string = "dict.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var eng_to_span = flag.Bool("eng_to_span", false, "Set to false for Spanish to English")
	var fname = flag.String("filename", "all.txt", "Name of the file that has the words")
	flag.Parse()
	var content []byte
	var err error

	if (*fname == "") {
		content, err = ioutil.ReadFile(filename)
	} else {
		content, err = ioutil.ReadFile(*fname)
	}
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
	numWords := len(spanWords)
	rand.Seed(time.Now().UnixNano())
	for {
		idx := rand.Intn(numWords)
		if (*eng_to_span) {
			fmt.Print(engWords[idx])
			scan := bufio.NewScanner(os.Stdin)
			scan.Scan()
			fmt.Printf("%s\n", spanWords[idx])
		} else {
			fmt.Print(spanWords[idx])
			scan := bufio.NewScanner(os.Stdin)
			scan.Scan()
			fmt.Printf("%s\n", engWords[idx])
		}
		//fmt.Print(engWords[idx])
	}
}
