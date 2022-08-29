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
func middleNode(head *ListNode) *ListNode {
	midNode := head
	for head != nil  {
		head = head.Next
		if head != nil {
			head = head.Next
			midNode = midNode.Next
		}
	}
	return midNode
}
