package main

func romanToInt(s string) int {
	res := 0
	for index, letter := range s {
		if len(s) > index+1 {
			if string(letter) == "I" && (string(s[index+1]) == "V" || string(s[index+1]) == "X") {
				res -= 1
				continue
			}
			if string(letter) == "X" && (string(s[index+1]) == "L" || string(s[index+1]) == "C") {
				res -= 10
				continue
			}
			if string(letter) == "C" && (string(s[index+1]) == "D" || string(s[index+1]) == "M") {
				res -= 100
				continue
			}
		}
		if string(letter) == "I" {
			res += 1
		}
		if string(letter) == "V" {
			res += 5
		}
		if string(letter) == "X" {
			res += 10
		}
		if string(letter) == "L" {
			res += 50
		}
		if string(letter) == "C" {
			res += 100
		}
		if string(letter) == "D" {
			res += 500
		}
		if string(letter) == "M" {
			res += 1000
		}
	}
	return res
}
