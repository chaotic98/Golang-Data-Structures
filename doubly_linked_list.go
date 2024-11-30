package main

import "fmt"

// Node2 represents a single node in the doubly linked list
type Node2 struct {
	value int
	next  *Node2
	prev  *Node2
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
	head *Node2
	tail *Node2
}

// InsertAtEnd adds a new node with the given value at the end of the list
func (dll *DoublyLinkedList) InsertAtEnd(value int) {
	newNode := &Node2{value: value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}
	dll.tail.next = newNode
	newNode.prev = dll.tail
	dll.tail = newNode
}

// InsertAtBeginning adds a new node at the beginning of the list
func (dll *DoublyLinkedList) InsertAtBeginning(value int) {
	newNode := &Node2{value: value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}
	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode
}

// Delete removes the first node with the given value
func (dll *DoublyLinkedList) Delete(value int) {
	if dll.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := dll.head
	for current != nil && current.value != value {
		current = current.next
	}

	if current == nil {
		fmt.Printf("Value %d not found in the list\n", value)
		return
	}

	// If the node to be deleted is the head
	if current == dll.head {
		dll.head = current.next
		if dll.head != nil {
			dll.head.prev = nil
		} else {
			dll.tail = nil
		}
		return
	}

	// If the node to be deleted is the tail
	if current == dll.tail {
		dll.tail = current.prev
		dll.tail.next = nil
		return
	}

	// If the node to be deleted is in the middle
	current.prev.next = current.next
	current.next.prev = current.prev
}

// TraverseForward prints all the elements from head to tail
func (dll *DoublyLinkedList) TraverseForward() {
	if dll.head == nil {
		fmt.Println("List is empty")
		return
	}
	current := dll.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

// TraverseBackward prints all the elements from tail to head
func (dll *DoublyLinkedList) TraverseBackward() {
	if dll.tail == nil {
		fmt.Println("List is empty")
		return
	}
	current := dll.tail
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.prev
	}
	fmt.Println("nil")
}

func main() {
	dll := &DoublyLinkedList{}

	// Insert elements
	dll.InsertAtEnd(10)
	dll.InsertAtEnd(20)
	dll.InsertAtEnd(30)
	dll.InsertAtBeginning(5)

	// Print the list forward
	fmt.Println("List traversed forward:")
	dll.TraverseForward()

	// Print the list backward
	fmt.Println("List traversed backward:")
	dll.TraverseBackward()

	// Delete an element
	dll.Delete(20)
	fmt.Println("\nList after deleting 20:")
	dll.TraverseForward()

	// Delete the head
	dll.Delete(5)
	fmt.Println("\nList after deleting the head (5):")
	dll.TraverseForward()

	// Delete the tail
	dll.Delete(30)
	fmt.Println("\nList after deleting the tail (30):")
	dll.TraverseForward()
}
