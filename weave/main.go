package main

import "fmt"

type Queue struct {
	Array []any
	front int
}

func (q *Queue) Add(item any) {
	q.Array = append(q.Array, item)
}

func (q *Queue) Remove() any {
	if q.front < len(q.Array) {
		q.front++
		return q.Array[q.front-1]
	}
	return nil
}

func (q *Queue) Peek() any {
	if q.front >= len(q.Array) {
		return nil
	}
	return q.Array[q.front]
}

func NewQueue() *Queue {
	return &Queue{}
}

func Weave(first, second Queue) *Queue {
	q := NewQueue()

	for first.Peek() != nil || second.Peek() != nil {
		if first.Peek() != nil {
			q.Add(first.Remove())
		}
		if second.Peek() != nil {
			q.Add(second.Remove())
		}
	}
	return q
}

func main() {
	q1 := NewQueue()
	q2 := NewQueue()

	q1.Add("hey")
	q1.Add(10)
	q2.Add("Helloo")
	q2.Add(5)
	q2.Add("foo")

	q := Weave(*q1, *q2)
	fmt.Println(q.Array...)
}
