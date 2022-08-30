package main

import "fmt"

func main() {
	head := &ListNode{Val: 5}
	next := head
	for i := 1; i <= 4; i++ {
		next.Next = &ListNode{Val: i}
		next = next.Next
	}
	ans := reverseBetween(head,2,4)
	fmt.Println(ans)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	temp := &ListNode{0,head}
	ans := temp
	n := right - left +1
	for i := 1; i < left; i++ {
		if temp == nil {
			break
		}
		temp = temp.Next
	}
	if temp == nil {
		return head
	}

	newListNode := reverseListNode(temp.Next,n)
	temp.Next = newListNode

	return ans.Next
}

//链表反转：头插法
func reverseListNode(head *ListNode,n int) *ListNode {
	var p, next *ListNode
	tail := head
	i := 0
	for head != nil {
		next = head.Next
		head.Next = p
		p = head
		head = next
		i++
		if i == n {
			tail.Next = head
			break
		}
	}
	return p
}
