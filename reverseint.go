package main

func reverseInt(x int) int {
	digit := 10
	newnum := 0
	operator := 1
	if x < 0 {
		operator = -1
		x *= -1
	}
	for x != 0 {
		newnum = newnum*digit + x%10
		x = x / 10
	}
	if newnum > 2147483647 || newnum < -2147483648 {
		return 0
	}
	return newnum * operator
}
