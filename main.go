package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	everything := strings.Join(os.Args[1:], "")
	ok, index := validateInput([]rune(everything))
	if ok {
		fmt.Printf("input %s is valid\n", everything)
		return
	}
	fmt.Println("input is invalid")
	fmt.Println(everything)
	spaces := strings.Repeat(" ", index)
	fmt.Printf("%s^\n", spaces)
}
