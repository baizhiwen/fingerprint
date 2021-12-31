package main

type Node struct {
	Value int
	Next  *Node
}

func main() {
	head := &Node{Value: 0}
	next := head
	for i := 1; i <= 10; i++ {
		next.Next = &Node{Value: i}
		next = next.Next
	}
	next = head.Next
	println(head.Value)
	for next != nil {
		println(next.Value)
		next = next.Next
	}

	head = reverseKGroupFromBack(head, 3)

	next = head.Next
	println(head.Value)
	for next != nil {
		println(next.Value)
		next = next.Next
	}
}

//链表从尾部分组反转
func reverseKGroupFromBack(head *Node, k int) *Node {
	head = reverseNode2(head)
	head = reverseKGroup(head, k)
	head = reverseNode2(head)
	return head
}

//链表从头部分组反转
func reverseKGroup(head *Node, k int) *Node {
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

	newNode := reverseNode2(head)
	newGroup := reverseKGroup(next, k)
	head.Next = newGroup

	return newNode
}

//链表反转：递归
func reverseNode2(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newNode := reverseNode2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newNode
}

//链表反转：头插法
func reverseNode(head *Node) *Node {
	var p, next *Node
	for head != nil {
		next = head.Next
		head.Next = p
		p = head
		head = next
	}

	return p
}
