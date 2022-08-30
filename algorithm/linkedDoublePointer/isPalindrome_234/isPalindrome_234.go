package main

import "fmt"

func main() {
	arr := []int{1,2,2,2}
	head := &ListNode{Val: 1}
	next := head
	for i := 1; i < len(arr); i++ {
		next.Next = &ListNode{Val: arr[i]}
		next = next.Next
	}
	ans := isPalindrome(head)
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

func isPalindrome(head *ListNode) bool {
	arr := make([]int,0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	l,r := 0,len(arr)-1
	for l <= r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
//链表反转：递归（测试用例，这种方法快）
func reverseListNode2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newListNode := reverseListNode2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newListNode
}

//链表反转：头插法
func reverseListNode(head *ListNode) *ListNode {
	var p, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = p
		p = head
		head = next
	}
	return p
}