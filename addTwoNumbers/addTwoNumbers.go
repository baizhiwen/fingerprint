package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nextVal := 0
	var head, tail *ListNode
	for l1 != nil || l2 != nil {
		val1 := 0
		val2 := 0
		//nextNode := &ListNode{}
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := nextVal + val1 + val2
		nextVal = sum / 10
		val := sum % 10
		if head == nil {
			head = &ListNode{}
			head.Val = val
			tail = head
		} else {
			tail.Next = &ListNode{Val: val}
			tail = tail.Next
		}
	}
	if nextVal > 0 {
		tail.Next = &ListNode{Val: nextVal}
	}
	return head
}
