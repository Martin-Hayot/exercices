package main

import "fmt"

func main() {
	pyramids(6)
}

func pyramids(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < (n * 2); j++ {
			if j >= n-i && j <= n+i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()

	}
}
