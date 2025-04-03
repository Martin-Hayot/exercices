package main

import (
	"fmt"
	"sort"
)

// func anagrams(stringA string, stringB string) bool {
// 	var mapStringA map[string]int
// 	var mapStringB map[string]int

// 	mapStringA = make(map[string]int)
// 	mapStringB = make(map[string]int)

// 	for i := 0; i < len(stringA); i++ {
// 		if mapStringA[string(stringA[i])] > 0 {
// 			mapStringA[string(stringA[i])] = mapStringA[string(stringA[i])] + 1
// 		}
// 		if mapStringA[string(stringA[i])] == 0 {
// 			mapStringA[string(stringA[i])] = 1
// 		}
// 	}

// 	for i := 0; i < len(stringB); i++ {
// 		if mapStringB[string(stringB[i])] > 0 {
// 			mapStringB[string(stringB[i])] = mapStringB[string(stringB[i])] + 1
// 		}
// 		if mapStringB[string(stringB[i])] == 0 {
// 			mapStringB[string(stringB[i])] = 1
// 		}
// 	}
// 	fmt.Println("map of string A", mapStringA)
// 	fmt.Println("map of string B", mapStringB)

// 	return compareMaps(mapStringA, mapStringB)
// }

// func compareMaps(map1, map2 map[string]int) bool {
// 	if len(map1) != len(map2) {
// 		return false
// 	}
// 	for key, value := range map1 {
// 		if map2[key] != value {
// 			return false
// 		}
// 	}
// 	return true
// }

func anagrams(stringA string, stringB string) bool {
	// Convert strings to rune slices
	runesA := []rune(stringA)
	runesB := []rune(stringB)

	// Sort the rune slices
	sort.Slice(runesA, func(i, j int) bool { return runesA[i] < runesA[j] })
	sort.Slice(runesB, func(i, j int) bool { return runesB[i] < runesB[j] })

	// Compare the sorted strings
	return string(runesA) == string(runesB)
}

func main() {
	res := anagrams("rail safety", "fairy tales")
	fmt.Println(res)
}
