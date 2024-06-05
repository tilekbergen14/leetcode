package main

import "fmt"

func main() {
	fmt.Println(isValid("[({(())}[()])]"))
}

// "{[]}" --> true
// "([)]" --> false
// "{[]}" --> true
// [({(())}[()])] --> true
