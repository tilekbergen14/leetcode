package main

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	sort.Ints(arr1)
	newArr := []int{}

	for _, current := range arr2 {
		for index, current1 := range arr1 {
			if current == current1 {
				newArr = append(newArr, int(current1))
				arr1[index] = -1
			}
		}
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] >= 0 {
			newArr = append(newArr, int(arr1[i]))
		}
	}

	return newArr
}
