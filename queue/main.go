package main

import "fmt"

type Queue struct {
	Array []any
	front int
}

func (q *Queue) Add(item any) {
	q.Array = append(q.Array, item)
}

func (q *Queue) Remove() {
	if q.front < len(q.Array) {
		q.front++
	}
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

func main() {
	q1 := NewQueue()
	q2 := NewQueue()

	q1.Add("hey")
	q1.Add(10)
	q1.Remove()
	q1.Remove()
	q2.Remove()
	q2.Add("Helloo")
	q2.Add(5)
	q2.Add("foo")
	fmt.Println(q1.Peek())
	fmt.Println(q2.Peek())
}
