package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(Val int) *ListNode {
	return &ListNode{Val: Val, Next: nil}
}

/* 在链表的节点 n0 之后插入节点 P */
func insertNode(n *ListNode, p *ListNode) {
	p.Next = n.Next
	n.Next = p
}

/* 删除链表的节点 n 之后的首个节点 */
func deleteNode(n *ListNode) {
	if n.Next == nil {
		return
	}
	n.Next = n.Next.Next
}

/* 访问链表中索引为 index 的节点 */
func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head.Next == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

/* 在链表中查找值为 target 的首个节点,输出节点在链表中的索引 */
func findNode(head *ListNode, target int) (index int) {
	p := head
	index = 0
	for p != nil {
		if p.Val == target {
			return
		}
		index++
		p = p.Next
	}
	return -1
}

func printLink(head *ListNode) {
	p := head
	str := ""
	for p != nil {
		if p.Next == nil {
			str += fmt.Sprintf("%d", p.Val)

		} else {
			str += fmt.Sprintf("%d -> ", p.Val)
		}
		p = p.Next
	}
	fmt.Println(str)
}
func main() {
	/* 初始化链表 1 -> 3 -> 2 -> 5 -> 4 */
	// 初始化各个节点
	n0 := NewListNode(1)
	n1 := NewListNode(3)
	n2 := NewListNode(2)
	n3 := NewListNode(5)
	n4 := NewListNode(4)
	// 构建引用指向
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	head := n0
	printLink(head)
	insertNode(n0, NewListNode(10))
	deleteNode(n2)
	printLink(head)
	fmt.Println(findNode(head, 4))
	fmt.Println(access(head, 3).Val)

	//fmt.Println(findNode(head, 1))
}
