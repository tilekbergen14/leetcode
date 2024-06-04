package main

func isPalindrome(x int) bool {
	initialX := x
	newNum := 0
	for x > 0 {
		if x/10 > 0 {
			if x == initialX {
				newNum = x % 10
			} else {
				newNum = (newNum * 10) + x%10
			}
			x = x / 10
		} else {
			if x != 0 {
				newNum = (newNum * 10) + x
			}
			x = 0
		}
	}
	return newNum == initialX
}
