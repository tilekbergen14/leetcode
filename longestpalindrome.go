package main

func longestPalindrome(s string) string {
	longest := string(s[0])
	for j := range s {
		for i := j + 1; i <= len(s); i++ {
			reversed := reverse(s[j:i])
			if reversed == s[j:i] {
				if len(reversed) > len(longest) {
					longest = reversed
				}
			}
		}
	}
	return longest
}

func reverse(s string) string {
	newstr := ""
	for _, current := range s {
		newstr = string(current) + newstr
	}
	return newstr
}
