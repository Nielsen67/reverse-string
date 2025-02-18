package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string = "Hello World"
	reversedString := reverseString(input)

	fmt.Println(reversedString)
}

func reverseString(input string) string {
	letters := make([]string, len(input))

	for i := len(input) - 1; i >= 0; i-- {
		letters = append(letters, string(input[i]))
	}
	return strings.Join(letters, "")
}
