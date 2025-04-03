package main

import "fmt"

func main() {
	counter := 0
	for i := 1; i <= 100; i++ {
		counter = i + counter
	}
	fmt.Println(counter)
}
