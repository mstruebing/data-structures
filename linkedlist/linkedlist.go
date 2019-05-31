package linkedlist

import (
	"fmt"
)

// Node struct representing an Item of the LinkedList
type Node struct {
	value string
	next  *Node
}

// Returns an Empty LinkedList
func CreateEmpty() *Node {
	return &Node{
		next: nil,
	}
}

// Tests if a LinkedList is Empty
func (n *Node) IsEmpty() bool {
	return n.next == nil && n.value == ""
}

// Returns the length of a LinkedList
func (n *Node) Length() int {
	count := 0

	for n.next != nil {
		count++
		n = n.next
	}

	// if the value is set the element counts
	// even if it has no successor
	if n.value != "" {
		count++
	}

	return count
}

// Inserts an element at the end of the list
func (n *Node) Insert(value string) {
	for n.next != nil {
		n = n.next
	}

	newNode := &Node{
		value: value,
		next:  nil,
	}

	if n.value == "" {
		// it's not possible to set n = newNode - don't know why
		n.value = newNode.value
	} else {
		n.next = newNode
	}
}

// Inserts an element at a specific index inside the list
// returns an error if the index is out of bounds
// TODO: Simplify code
func (n *Node) InsertAt(index int, value string) error {
	if n.Length() < index && index != 0 {
		return fmt.Errorf("Index out of bound")
	}

	if index == 0 {
		if n.value == "" && n.next == nil {
			n.value = value
			n.next = nil
		} else if n.next == nil && n.value != "" {
			newNode := &Node{
				value: n.value,
				next:  nil,
			}

			n.next = newNode
			n.value = value
		} else {
			newNode := &Node{
				value: n.value,
				next:  n.next,
			}
			n.next = newNode
			n.value = value
		}

		return nil
	}

	currentIndex := 0

	for currentIndex < index-1 {
		currentIndex++
		n = n.next
	}

	if n.next != nil && n.next.next != nil {
		newNode := &Node{
			value: value,
			next:  n.next,
		}

		n.next = newNode
	} else if n.next != nil {
		newNode := &Node{
			value: value,
			next:  n.next,
		}

		n.next = newNode
	} else {
		newNode := &Node{
			value: value,
			next:  nil,
		}

		n.next = newNode
	}

	return nil
}

// Removes an element from the end of the list
func (n *Node) Remove() {
	for n.next != nil && n.next.next != nil {
		n = n.next
	}

	n.next = nil
}

// Removes an element from a specific index
// returns an error if the index is out of bounds
// TODO: Simplify code
func (n *Node) RemoveAt(index int) error {
	if n.IsEmpty() || n.Length()-1 < index {
		return fmt.Errorf("Index out of bound")
	}

	currentIndex := 0

	for currentIndex < index-1 {
		currentIndex++
		n = n.next
	}

	if index == 0 && n.next != nil {
		// first element
		n.value = n.next.value
		n.next = n.next.next
	} else if index != 0 && n.next != nil && n.next.next != nil {
		// some index in between
		n.next.value = n.next.next.value
		n.next.next = n.next.next.next
	} else {
		// last element
		n.next.value = ""
		n.next = nil
	}

	return nil
}

// Removes all elements from the list
func (n *Node) Truncate() {
	n.next = nil
	n.value = ""
}

// Check if a specific value is inside the list
func (n *Node) Contains(value string) bool {
	for n.next != nil {
		if n.value == value {
			break
		}

		n = n.next
	}

	return n.value == value
}

// Gets the item of a specific index
// returns an error if the index is out of bound
func (n *Node) Get(index int) (string, error) {
	currentIndex := 0

	for n.next != nil {
		if currentIndex == index {
			return n.value, nil
		}

		n = n.next
		currentIndex++
	}

	// We need to check the last element
	if currentIndex == index {
		return n.value, nil
	}

	return "", fmt.Errorf("Index %d doesn't exist", index)
}

// Gets the first element of the list
func (n *Node) GetFirst() string {
	return n.value
}

// Gets the last element of a list
func (n *Node) GetLast() string {
	for n.next != nil {
		n = n.next
	}

	return n.value
}

// Returns the index of a value inside the list
// returns -1 if the value is not present
func (n *Node) IndexOf(value string) int {
	currentIndex := 0

	for n.next != nil {
		if n.value == value {
			return currentIndex
		}

		currentIndex++
		n = n.next
	}

	// We need to check the last element
	if n.value == value {
		return currentIndex
	}

	return -1
}

// Returns the last index of a value inside the list
// returns -1 if the value is not present
func (n *Node) LastIndexOf(value string) int {
	currentIndex := 0
	foundIndex := -1

	for n.next != nil {
		if n.value == value {
			foundIndex = currentIndex
		}

		n = n.next
		currentIndex++
	}

	if n.value == value {
		foundIndex = currentIndex
	}

	return foundIndex
}

// Converts a LinkedList to an Array
func (n *Node) ToArray() []string {
	length := n.Length()
	arr := make([]string, length)

	for i := 0; i < length; i++ {
		arr[i], _ = n.Get(i)
	}

	return arr
}

// Creates a LinkedList of an Array
func FromArray(array []string) *Node {
	list := CreateEmpty()

	for _, value := range array {
		list.Insert(value)
	}

	return list
}

// Prints the LinkedList
func (n *Node) Print() {
	if n.IsEmpty() {
		return
	}

	currentIndex := 1

	for n.next != nil && n.value != "" {
		fmt.Printf("%d: %s\n", currentIndex, n.value)
		n = n.next
		currentIndex++
	}

	fmt.Printf("%d: %s\n", currentIndex, n.value)
}
