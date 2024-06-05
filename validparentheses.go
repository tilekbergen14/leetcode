package main

import "fmt"

func isValid(s string) bool {
	for {
		length := len(s)
		if len(s)%2 != 0 || len(s) <= 1 {
			return false
		}
		fmt.Println(s, len(s))

		for index, current := range s {
			if index+1 < len(s) {
				if current+1 == rune(s[index+1]) || current+2 == rune(s[index+1]) {
					s = s[:index] + s[index+2:]
					break
				}
			}
		}

		if len(s) == 0 {
			return true
		}
		if length == len(s) {
			return false
		}
	}
}
