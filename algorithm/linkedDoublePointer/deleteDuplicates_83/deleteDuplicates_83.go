package main

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
func deleteDuplicates(head *ListNode) *ListNode {
	ans := head
	for head != nil {
		de := true
		if head.Next != nil {
			next := head.Next
			if head.Val == next.Val {
				head.Next = next.Next
				de = false
			}
		}
		if de {
			head = head.Next
		}
	}
	return ans
}