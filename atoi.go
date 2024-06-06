package main

import "fmt"

func myAtoi(s string) int {
	started := false
	result := float64(0)
	operator := float64(1)
	for _, current := range s {
		if current == '-' && !started {
			started = true
			operator = -1
			continue
		}
		if current == '+' && !started {
			started = true
			operator = 1
			continue
		}
		if current == ' ' && !started {
			continue
		}
		if current < 48 || current > 57 {
			break
		}
		result = result*10 + float64(current-48)
		started = true
	}
	if result*operator > 2147483647 {
		return 2147483647
	}
	if result*operator < -2147483648 {
		fmt.Println("smaller")
		return -2147483648

	}
	return int(result * operator)
}
