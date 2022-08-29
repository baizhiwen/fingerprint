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
func detectCycle(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]struct{},0)
	for head != nil  {
		if _,ok := nodeMap[head];ok {
			return head
		}
		nodeMap[head]= struct{}{}
		head = head.Next
	}
	return nil
}
