package main

import "fmt"

// Stack structure using a slice
type Stack struct {
	elements []int
}

// Push adds an element to the top of the stack
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
// Returns a value and a boolean indicating success
func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false // Stack is empty
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top, true
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false // Stack is empty
	}
	return s.elements[len(s.elements)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.elements)
}

func main() {
	stack := &Stack{}

	// Push elements onto the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println("Stack after pushing 10, 20, 30:", stack.elements)

	// Peek at the top element
	if top, ok := stack.Peek(); ok {
		fmt.Println("Peek:", top)
	} else {
		fmt.Println("Stack is empty")
	}

	// Pop elements from the stack
	for !stack.IsEmpty() {
		if value, ok := stack.Pop(); ok {
			fmt.Println("Popped:", value)
		}
	}

	// Check if stack is empty
	fmt.Println("Is stack empty?", stack.IsEmpty())
}
