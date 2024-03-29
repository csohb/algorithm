package main

import (
	"fmt"
	"sort"
)

func mergeIntervals(givenArray [][]int) (sortedArray [][]int) {
	sortedArray = append(sortedArray, givenArray[0])

	for i := 1; i < len(givenArray); i++ {
		lastIndex := len(sortedArray) - 1
		if sortedArray[lastIndex][1] < givenArray[i][0] {
			sortedArray = append(sortedArray, givenArray[i])
		}
	}
	return sortedArray
}

func main() {
	data := [][]int{
		{1, 3},
		{30, 8},
		{4, 10},
		{20, 25},
	}

	fmt.Printf("input: %+v\n", data)

	for _, v := range data {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}

	fmt.Println("sorted1:", data)
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})

	fmt.Println("sorted2:", data)

	output := mergeIntervals(data)
	fmt.Println("output:", output)

}
