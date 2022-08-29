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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	rmNode,tail := head,head
	times := 0
	for tail != nil {
		if tail.Next != nil {
			if times == n {
				rmNode = rmNode.Next
			}else {
				times++
			}
		}
		tail = tail.Next
	}
	if times < n {
		head = head.Next
	}else {
		//r :=rmNode.Next.Next
		rmNode.Next = rmNode.Next.Next
	}
	return head
}