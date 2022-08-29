package main

import "math"

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
//二分法合并
func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	ans = merge(lists,0, len(lists)-1)
	return ans
}
func merge(lists []*ListNode, l,r int) *ListNode {
	var ans *ListNode
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l+r)/2
	return mergeTwoLists(merge(lists,l,mid),merge(lists,mid+1,r))
	return ans
}

//顺序合并
func mergeKLists1(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i:=0;i<len(lists);i++ {
		ans = mergeTwoLists(ans,lists[i])
	}
	return ans
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}
	return ans
}

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}