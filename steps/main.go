package main

import "fmt"

func main() {

	steps(5)
}

func steps(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j <= i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
