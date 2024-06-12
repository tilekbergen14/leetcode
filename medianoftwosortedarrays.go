package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1) == 1 {
		return float64(nums1[0])
	}
	median := 0.0
	if len(nums1)%2 == 0 {
		median = float64(nums1[len(nums1)/2]+(nums1[len(nums1)/2-1])) / 2
	} else {
		return float64(nums1[len(nums1)/2])
	}

	return median
}
