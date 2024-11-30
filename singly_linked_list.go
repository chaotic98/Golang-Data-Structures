package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	head *Node
}

// InsertAtEnd adds a new node with the given value at the end of the list
func (ll *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// InsertAtBeginning adds a new node at the beginning of the list
func (ll *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{value: value, next: ll.head}
	ll.head = newNode
}

// Delete removes the first node with the given value
func (ll *LinkedList) Delete(value int) {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}

	// If the head node needs to be deleted
	if ll.head.value == value {
		ll.head = ll.head.next
		return
	}

	current := ll.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next == nil {
		fmt.Printf("Value %d not found in the list\n", value)
		return
	}

	current.next = current.next.next
}

// Traverse prints all the elements in the linked list
func (ll *LinkedList) Traverse() {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{}

	// Insert elements
	ll.InsertAtEnd(10)
	ll.InsertAtEnd(20)
	ll.InsertAtEnd(30)
	ll.InsertAtBeginning(5)

	// Print the linked list
	fmt.Println("Linked list after insertion:")
	ll.Traverse()

	// Delete an element
	ll.Delete(20)
	fmt.Println("Linked list after deleting 20:")
	ll.Traverse()

	// Delete a non-existing element
	ll.Delete(100)

	// Delete the head element
	ll.Delete(5)
	fmt.Println("Linked list after deleting the head (5):")
	ll.Traverse()
}
