package main

import "slices"

func heightChecker(heights []int) int {
	sortedArr := make([]int, len(heights))
	copy(sortedArr, heights)
	slices.Sort(heights)
	counter := 0
	for j := 0; j < len(heights); j++ {
		if sortedArr[j] != heights[j] {
			counter++
		}
	}
	return counter
}
