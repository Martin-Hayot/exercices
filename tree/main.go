package main

import "fmt"

type Node struct {
	data     any
	children []*Node
}

type Tree struct {
	root *Node
}

func newTree(root *Node) *Tree {
	return &Tree{
		root: root,
	}
}

func newNode(data any, children ...*Node) *Node {
	return &Node{
		data:     data,
		children: children,
	}
}

func (t *Tree) print() {
	fmt.Println(t.root.data)
	for _, node := range t.root.children {
		node.print()
	}
}

func (n *Node) print() {
	fmt.Println(n.data)
	for _, node := range n.children {
		node.print()
	}
}

func main() {
	tree := newTree(newNode(1,
		newNode(2,
			newNode(4),
			newNode(5),
		),
		newNode(3,
			newNode(6),
			newNode(7),
		),
	))

	tree.print()
}
