package main

import (
	"fmt"
	"math"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	moves := 0
	// for i := range students {
	// 	for j := range seats {
	// 		if students[i] == seats[j] {
	// 			students[i] = -1
	// 			seats[j] = -1
	// 		}
	// 	}
	// }
	sort.Sort(sort.Reverse(sort.IntSlice(students)))
	for i := range students {
		if students[i] < 0 {
			continue
		}
		closest := -1
		for j := range seats {
			if seats[j] < 0 {
				continue
			}

			if closest < 0 {
				closest = j
				continue
			}
			if math.Abs(float64(students[i]-seats[j])) < math.Abs(float64(students[i]-seats[closest])) {
				closest = j
			}
			if math.Abs(float64(students[i]-seats[j])) == math.Abs(float64(students[i]-seats[closest])) {
				if students[i]-seats[j] > 0 {
					closest = j
				}
			}
		}
		moves += int(math.Abs(float64(students[i] - seats[closest])))
		seats[closest] = -1
		fmt.Println(seats, moves)
	}
	return moves
}
