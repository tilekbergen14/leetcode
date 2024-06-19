package main

import (
	"slices"
)

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	day := slices.Min(bloomDay)
	for {
		bloomed := true
		for index, current := range bloomDay {
			if int(current) == day {
				bloomDay[index] = 0
			}
			if bloomDay[index] != 0 {
				bloomed = false
			}
		}
		if checkFlowers(m, k, bloomDay) {
			break
		}
		if bloomed {
			break
		}
		day++

	}
	return day
}

func checkFlowers(m, k int, bloomDay []int) bool {
	i := 0
	counter := 0
	for {
		if k+i <= len(bloomDay) {
			enough := true
			for _, current := range bloomDay[i : i+k] {
				if current != 0 {
					enough = false
					i++
					break
				}
			}
			if enough {
				i = i + k
				counter++
			}
		} else {
			break
		}
	}
	return counter == m
}
