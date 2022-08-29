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
func hasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]struct{},0)
	for head != nil  {
		if _,ok := nodeMap[head];ok {
			return true
		}
		nodeMap[head]= struct{}{}
		head = head.Next
	}
	return false
}
