package main

import (
	"fmt"
	"math"
)

func binarySearch(array []int, value int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		middlePoint := int(math.Floor(float64(left+right) / 2))
		middleValue := array[middlePoint]

		if middleValue > value {
			right = middlePoint - 1
		} else if middleValue < value {
			left = middlePoint + 1
		} else {
			return middlePoint
		}
	}
	return -1

}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(array, 5)
	fmt.Println("result", result)

}
