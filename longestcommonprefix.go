package main

func longestCommonPrefix(strs []string) string {
	prefix := ""
	looping := true
	i := 0
	for looping {
		letter := ""
		for index, str := range strs {
			if len(str) <= i {
				looping = false
				letter = ""
				break
			}
			if index == 0 {
				letter = string(str[i])
			}
			if letter != string(str[i]) {
				letter = ""
				looping = false
				break
			}
		}
		if letter != "" {
			prefix += letter
		}
		i++
	}
	return prefix
}
