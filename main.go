package main

import "fmt"

// func main() {
// 	node3 := ListNode{
// 		Val:  3,
// 		Next: nil,
// 	}

// 	node2 := ListNode{
// 		Val:  4,
// 		Next: &node3,
// 	}
// 	node1 := ListNode{
// 		Val:  2,
// 		Next: &node2,
// 	}

// 	node6 := ListNode{
// 		Val:  4,
// 		Next: nil,
// 	}

// 	node5 := ListNode{
// 		Val:  6,
// 		Next: &node6,
// 	}
// 	node4 := ListNode{
// 		Val:  5,
// 		Next: &node5,
// 	}

// 	current := addTwoNumbers(&node1, &node4)
// 	for current != nil {
// 		fmt.Println(current.Val)
// 		current = current.Next
// 	}
// }

func main() {
	// arr := [][]int{{1, 2}, {2, 3}, {4, 2}}
	arr1 := [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}

	fmt.Println(findCenter(arr1))
}
