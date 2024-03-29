package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 0}
	next := head
	for i := 1; i <= 10; i++ {
		next.Next = &ListNode{Val: i}
		next = next.Next
	}
	next = head.Next
	println(head.Val)
	for next != nil {
		println(next.Val)
		next = next.Next
	}

	head = reverseKGroupFromBack(head, 3)

	next = head.Next
	println(head.Val)
	for next != nil {
		println(next.Val)
		next = next.Next
	}
}

//链表从尾部分组反转
func reverseKGroupFromBack(head *ListNode, k int) *ListNode {
	head = reverseListNode2(head)
	head = reverseKGroup(head, k)
	head = reverseListNode2(head)
	return head
}

//链表从头部分组反转
func reverseKGroup(head *ListNode, k int) *ListNode {
	temp := head
	for i := 1; i < k; i++ {
		if temp == nil {
			break
		}
		temp = temp.Next
	}
	if temp == nil {
		return head
	}

	next := temp.Next

	temp.Next = nil

	newListNode := reverseListNode2(head)
	newGroup := reverseKGroup(next, k)
	head.Next = newGroup

	return newListNode
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
