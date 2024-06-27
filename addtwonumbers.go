package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res1 := 0
	multiplier := 1
	for l1 != nil {
		res1 += l1.Val * multiplier
		l1 = l1.Next
		multiplier *= 10
	}

	res2 := 0
	multiplier = 1
	for l2 != nil {
		res1 += l2.Val * multiplier
		l2 = l2.Next
		multiplier *= 10
	}
	res := res1 + res2
	list := LinkedList{head: nil}

	for res != 0 {
		val := res % 10
		res = res / 10
		list.insertAtBack(val)
	}
	node := list.head
	for node != nil {
		fmt.Println(node)
	}

	return l1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (list *LinkedList) insertAtBack(data int) {
	newNode := &ListNode{Val: data, Next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}
