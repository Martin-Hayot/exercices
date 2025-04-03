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

type Queue struct {
	first  *Stack
	second *Stack
}

func (q *Queue) Add(item any) {
	q.first.Push(item)
}

func (q *Queue) Remove() any {
	for range len(q.first.Array) {
		q.second.Push(q.first.Pop())
	}
	item := q.second.Pop()
	for range len(q.second.Array) {
		q.first.Push(q.second.Pop())
	}
	return item
}

func (q *Queue) Peek() any {
	for range len(q.first.Array) {
		q.second.Push(q.first.Pop())
	}
	item := q.second.Peek()
	for range len(q.second.Array) {
		q.first.Push(q.second.Pop())
	}
	return item
}

func NewQueue() *Queue {
	return &Queue{
		first:  NewStack(),
		second: NewStack(),
	}
}

func main() {
	q := NewQueue()

	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Remove()
	q.Add(5)
	q.Remove()
	fmt.Println(q.Peek())
}
