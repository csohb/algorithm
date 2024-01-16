package main

import "fmt"

func rotateListByK(arr []int, k int) {
	n := len(arr)

	k = k % n

	reverse(arr, 0, n-k-1)
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	k := 2

	rotateListByK(arr, k)
	fmt.Println("Rotated array by", k, "elements:", arr)
}
