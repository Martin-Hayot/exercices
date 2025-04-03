package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func main() {
	fmt.Println(vowels("hEllo world iii"))
	fmt.Println(vowelsViaRegex("hEllo world iii"))
}

func vowels(str string) int {
	count := 0
	for _, char := range str {
		char = unicode.ToLower(char)
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
		}
	}
	return count
}

func vowelsViaRegex(str string) int {
	re := regexp.MustCompile("[aeiouAEIOU]")
	return len(re.FindAllString(str, -1))
}
