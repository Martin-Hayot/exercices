package main

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	next *Node
	data any
}

func newLinkedList() *LinkedList {
	return &LinkedList{}
}

func newNode(data any, next ...*Node) *Node {
	if len(next) > 0 {
		return &Node{
			data: data,
			next: next[0],
		}
	}
	return &Node{
		data: data,
	}
}

func (l *LinkedList) InsertFirst(data any) {
	l.head = newNode(data, l.head)
}

func (l *LinkedList) Size() int {
	counter := 0
	node := l.head

	for node != nil {
		counter++
		node = node.next
	}

	return counter
}

func (l *LinkedList) GetFirst() *Node {
	return l.head
}

func (l *LinkedList) GetLast() *Node {
	node := l.head
	for node != nil {
		if node.next == nil {
			return node
		}
		node = node.next
	}

	return node
}

func (l *LinkedList) Clear() {
	l.head = nil
}

func (l *LinkedList) RemoveFirst() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}

func (l *LinkedList) RemoveLast() {
	if l.head == nil {
		return
	}

	if l.head.next == nil {
		l.head = nil
		return
	}

	var prev *Node
	node := l.head
	for node.next != nil {
		prev = node
		node = node.next
	}
	prev.next = nil
}

func Midpoint(l *LinkedList) *Node {
	if l.head == nil {
		return nil
	}

	slow := l.head
	fast := l.head

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func Circular(l *LinkedList) bool {
	if l.head == nil {
		return false
	}

	slow := l.head
	fast := l.head

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}

func FromLast(l *LinkedList, n int) *Node {
	if l.head == nil {
		return nil
	}

	slow := l.head
	fast := l.head
	for i := 0; i < n; i++ {
		fast = fast.next
	}

	if fast == nil {
		return nil
	}

	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}

	return slow
}

func main() {
	list := newLinkedList()
	list.InsertFirst("a")
	list.InsertFirst("b")
	list.InsertFirst("c")
	list.InsertFirst("d")
	list.InsertFirst("e")
	list.InsertFirst("f")

	fmt.Println(FromLast(list, 6))
}
