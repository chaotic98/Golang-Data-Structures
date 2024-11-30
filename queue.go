package main

import "fmt"

// Queue represents a queue data structure
type Queue struct {
	elements []int
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the element from the front of the queue
// Returns a value and a boolean indicating success
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false // Queue is empty
	}
	front := q.elements[0]
	q.elements = q.elements[1:]
	return front, true
}

// Peek returns the front element without removing it
func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false // Queue is empty
	}
	return q.elements[0], true
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.elements)
}

// Print displays the elements of the queue
func (q *Queue) Print() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println("Queue elements:", q.elements)
}

func main() {
	queue := &Queue{}

	// Enqueue elements
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	fmt.Println("After enqueueing 10, 20, 30:")
	queue.Print()

	// Peek at the front element
	if front, ok := queue.Peek(); ok {
		fmt.Println("Peek:", front)
	} else {
		fmt.Println("Queue is empty")
	}

	// Dequeue elements
	for !queue.IsEmpty() {
		if value, ok := queue.Dequeue(); ok {
			fmt.Println("Dequeued:", value)
		}
	}

	// Check if queue is empty
	fmt.Println("Is queue empty?", queue.IsEmpty())
}
