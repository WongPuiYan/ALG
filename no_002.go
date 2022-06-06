package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		root, current_node *ListNode
		temp               int
	)
	for l1 != nil || l2 != nil || temp != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		temp = n1 + n2 + temp
		if root == nil {
			root = &ListNode{Val: temp % 10}
			current_node = root
		} else {
			current_node.Next = &ListNode{Val: temp % 10}
			current_node = current_node.Next
		}
		temp = temp / 10
	}
	return root
}

func main() {
	var (
		l1 = []int{9, 9, 9, 9, 9, 9, 9}
		l2 = []int{9, 9, 9, 9}
		// l1           = []int{2, 4, 3}
		// l2           = []int{5, 6, 4}
		n1, n2, temp *ListNode
	)
	for i := 0; i < len(l1); i++ {
		if n1 == nil {
			n1 = &ListNode{Val: l1[i]}
			temp = n1
		} else {
			temp.Next = &ListNode{Val: l1[i]}
			temp = temp.Next
		}
	}
	// for n1 != nil {
	// 	fmt.Println(n1.Val)
	// 	n1 = n1.Next
	// }
	for i := 0; i < len(l2); i++ {
		if n2 == nil {
			n2 = &ListNode{Val: l2[i]}
			temp = n2
		} else {
			temp.Next = &ListNode{Val: l2[i]}
			temp = temp.Next
		}
	}
	// for n2 != nil {
	// 	fmt.Println(n2.Val)
	// 	n2 = n2.Next
	// }
	// for i := 0; i < len(l2); i++ {
	// 	n2 = &ListNode{Val: l2[i]}
	// 	n2
	// }
	results := addTwoNumbers(n1, n2)
	// fmt.Println(results)
	for results != nil {
		fmt.Println(results.Val)
		results = results.Next
	}
	// addTwoNumbers(n1, temp)
	// fmt.Println(results)
}
