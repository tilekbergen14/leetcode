package main

import "fmt"

func minIncrementForUnique(nums []int) int {
	arr := []int{nums[0]}
	counter := 0
	for _, current := range nums[1:] {
		for _, current1 := range arr {
			for {
				if current == current1 {
					fmt.Println(current)
					current++
					counter++
				} else {
					arr = append(arr, current)
					break
				}
			}
		}

	}
	return counter
}
