package main

import "fmt"

func main() {
	arr := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(arr))
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
}
