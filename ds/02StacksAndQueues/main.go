package main

import "fmt"

// stacks represents stack that hold a slice

type Stack struct {
	items []int
}

// push will add the value at the top
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove the value at the top
// and RETURNS the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

type Queue struct {
	items []int
}

// Enqueue will add the value at the top
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue will remove the value at the bottom
// and RETURNS the removed value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	// myStack := Stack{}
	// fmt.Println(myStack)
	// myStack.Push(100)
	// myStack.Push(43)
	// myStack.Push(34234)
	// myStack.Push(23434)
	// fmt.Println(myStack)
	// myStack.Pop()
	// fmt.Println(myStack)

	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(23)
	myQueue.Enqueue(243)
	myQueue.Enqueue(256)
	myQueue.Enqueue(267)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
