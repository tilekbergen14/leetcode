package main

// func isValid(s string) bool {
// 	for {
// 		if len(s) == 0 {
// 			return true
// 		}
// 		if len(s) <= 1 {
// 			return false
// 		}
// 		if s[0]+1 == s[1] || s[0]+2 == s[1] {
// 			s = s[2:]
// 		} else {
// 			return false
// 		}
// 	}
// }

func isValid(s string) bool {
	for {
		if 1 >= len(s) {
			break
		}
		length := len(s)
		for i, j := range s {
			if string(j) == "]" || string(j) == "}" || string(j) == ")" {
				if i == 0 {
					return false
				}
				if string(s[i-1]) == "]" || string(s[i-1]) == "}" || string(s[i-1]) == ")" {
				} else {
					if rune(s[i-1])+1 != j && rune(s[i-1])+2 != j {
						return false
					}
				}
			}
			if rune(s[0])+1 == j || rune(s[0])+2 == j {
				s = s[1:i] + s[i+1:]
				break
			}
		}
		if length == len(s) {
			return false
		}
	}
	return len(s) == 0
}
