package main

import (
	"fmt"
	"strings"
)

func capitalize(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(capitalize("hello world"))
}
