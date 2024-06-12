package main

import "fmt"

func sortColors(nums []int) {
	sorted := false
	i := 0
	if len(nums) <= 1 {
		fmt.Println(nums)
		return
	}
	for {
		if nums[i] > nums[i+1] {
			swap := nums[i]
			nums[i] = nums[i+1]
			nums[i+1] = swap
			sorted = false
			i = 0
		} else {
			i++
			sorted = true
		}
		if sorted && len(nums)-1 == i {
			break
		}
	}
	fmt.Println(nums)
}
