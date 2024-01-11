package main

import (
	"fmt"
	"sort"
)

func mergeIntervals(givenArray [][]int) (sortedArray [][]int) {
	sortedArray = append(sortedArray, givenArray[0])

	for i := 1; i < len(givenArray); i++ {
		lastIntervalIndexOfSortedArray := len(sortedArray) - 1
		if sortedArray[lastIntervalIndexOfSortedArray][1] < givenArray[i][0] {
			sortedArray = append(sortedArray, givenArray[i])
		}
	}
	return sortedArray
}

func main() {
	data := [][]int{
		{1, 3},
		{5, 8},
		{4, 10},
		{20, 25},
	}

	fmt.Printf("input: %+v\n", data)

	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})

	fmt.Println("sorted:", data)

	output := mergeIntervals(data)
	fmt.Println("output:", output)

}
