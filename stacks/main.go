package main

import "fmt"

type Stack struct {
	Array []any
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item any) {
	s.Array = append(s.Array, item)
}

func (s *Stack) Pop() any {
	if len(s.Array) == 0 {
		return nil
	}
	item := s.Array[len(s.Array)-1]
	s.Array = s.Array[:len(s.Array)-1]
	return item
}

func (s *Stack) Peek() any {
	if len(s.Array) == 0 {
		return nil
	}
	return s.Array[len(s.Array)-1]
}

func main() {
	s := NewStack()
	s.Push("hey")
	s.Push(5)
	s.Push(120)
	s.Pop()
	s.Pop()

	fmt.Println(s.Peek())
}
