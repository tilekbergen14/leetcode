package main

func lengthOfLongestSubstring(s string) int {
	longest := 0
	str := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if inString(s[j], str) {
				str = ""
				break
			} else {
				str += string(s[j])
			}
			if longest < len(str) {
				longest = len(str)
			}
		}
	}
	return longest
}

func inString(letter byte, str string) bool {
	for _, current := range str {
		if current == rune(letter) {
			return true
		}
	}
	return false
}
