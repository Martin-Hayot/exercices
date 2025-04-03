package main

import "fmt"

func main() {
	fmt.Println(chunk([]int{1, 2, 3, 4, 5}, 2))
}

func chunk(array []int, size int) [][]int {
	var chunked [][]int

	for i := 0; i < len(array); i++ {
		if len(chunked) == 0 || len(chunked[len(chunked)-1]) == size {
			chunked = append(chunked, []int{})
		}
		chunked[len(chunked)-1] = append(chunked[len(chunked)-1], array[i])
	}
	return chunked
}
