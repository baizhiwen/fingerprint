package main

import (
	"math"
)

func main() {

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ans,tail *ListNode
	for list1 != nil && list2 != nil {
		val1 := math.MaxInt32
		val2 := math.MaxInt32
		if list1 != nil {
			val1 = list1.Val
		}
		if list2 != nil {
			val2 = list2.Val
		}
		if val1 < val2 {
			list1 = list1.Next
		}else {
			list2 = list2.Next
		}
		val := min(val1,val2)
		if ans == nil {
			ans = &ListNode{}
			ans.Val = val
			tail = ans
		}else {
			next := &ListNode{}
			next.Val = val
			tail.Next = next
			tail = tail.Next
		}
	}
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}
	return ans
}
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	var ans,tail *ListNode
	for list1 != nil || list2 != nil {
		val1 := math.MaxInt32
		val2 := math.MaxInt32
		if list1 != nil {
			val1 = list1.Val
		}
		if list2 != nil {
			val2 = list2.Val
		}
		if val1 < val2 {
			list1 = list1.Next
		}else {
			list2 = list2.Next
		}
		val := min(val1,val2)
		if ans == nil {
			ans = &ListNode{}
			ans.Val = val
			tail = ans
		}else {
			next := &ListNode{}
			next.Val = val
			tail.Next = next
			tail = tail.Next
		}
	}

	return ans
}

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}