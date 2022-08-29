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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]struct{},0)
	for headA != nil {
		if headA != nil {
			nodeMap[headA]= struct{}{}
			headA = headA.Next
		}
	}
	for headB != nil  {
		if _,ok := nodeMap[headB];ok {
			return headB
		}
		if headB != nil {
			nodeMap[headB]= struct{}{}
			headB = headB.Next
		}
	}
	return nil
}

//错误
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]struct{},0)
	for headA != nil || headB != nil  {
		if _,ok := nodeMap[headA];ok {
			return headA
		}
		if _,ok := nodeMap[headB];ok {
			return headB
		}
		if headA != nil {
			nodeMap[headA]= struct{}{}
			headA = headA.Next
		}
		if headB != nil {
			nodeMap[headB]= struct{}{}
			headB = headB.Next
		}
	}
	return nil
}
