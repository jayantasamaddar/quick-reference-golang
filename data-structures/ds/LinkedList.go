package ds

import (
	"fmt"
)

type ListNode[T Type] struct {
	Value T
	Next  *ListNode[T]
}

type LinkedList[T Type] struct {
	Head *ListNode[T]
	Size int
}

// Append: Add items to the end of the LinkedList. Does not change the Head of the LinkedList. Returns the new length of the List.
func (this *LinkedList[T]) Append(items ...T) int {
	for _, item := range items {
		node := &ListNode[T]{Value: item}
		currentNode := this.Head

		if currentNode == nil {
			this.Head = node
		} else {
			for currentNode.Next != nil {
				currentNode = currentNode.Next
			}
			currentNode.Next = node
		}
		this.Size++
	}
	return this.Size
}

// Prepend: Add Items to the start of the LinkedList. The Head value withis change to the last item provided. Returns the new length of the List.
func (this *LinkedList[T]) Prepend(items ...T) int {
	for _, item := range items {
		head := this.Head
		this.Head = &ListNode[T]{Value: item}
		this.Head.Next = head
		this.Size++
	}
	return this.Size
}

// RemoveFirstOccurence: Remove the first occurence of the node with the same value as the item.
// Returns `true` if successful, `false` if not successful
func (this *LinkedList[T]) RemoveFirstOccurence(item T) bool {
	currentNode := this.Head
	if currentNode != nil {
		// If initial Head Value matches item, handle that
		if currentNode.Value == item {
			// Remove if first linked list item's value matches with
			this.Head = this.Head.Next
			this.Size--
			return true
		}
		for i := 0; i < this.Size; i++ {
			next := currentNode.Next
			if next != nil && next.Value == item {
				// Remove
				currentNode.Next = next.Next
				this.Size--
				return true
			}
			currentNode = currentNode.Next
		}
	}
	return false
}

// RemoveLastOccurence: Remove the last occurence of the node with the same value as the item
// Returns `true` if successful, `false` if unsuccessful
func (this *LinkedList[T]) RemoveLastOccurence(item T) bool {
	currentNode := this.Head
	var i int
	if currentNode != nil {
		for j := 0; j < this.Size; j++ {
			if currentNode.Value == item {
				i = j
			}
			currentNode = currentNode.Next
		}

		// Reset currentNode to loop again to reach last index
		currentNode = this.Head
		if i > 0 {
			// Item found in an index other than the first
			for j := 0; j < i-1; j++ {
				currentNode = currentNode.Next
			}
			currentNode.Next = currentNode.Next.Next
		} else if i == 0 && currentNode.Value == item {
			// Item found in the first index itself
			this.Head = currentNode.Next
		} else {
			// Case when i == 0 as initiated, however no item was found, so it still is 0
			return false
		}
		this.Size--
		return true
	}
	return false
}

// RemoveFirst: Remove the first node in the LinkedList. Returns the value of the Node.
func (this *LinkedList[T]) RemoveFirst() (T, error) {
	head := this.Head
	if head != nil {
		this.Head = this.Head.Next
		this.Size--
		return head.Value, nil
	}
	var zeroVal T
	return zeroVal, fmt.Errorf("linked list is initiated but has no head value")
}

// RemoveLast: Removes the last node in the LinkedList. Returns the value of the Node.
func (this *LinkedList[T]) RemoveLast() (T, error) {
	currentNode := this.Head
	if currentNode == nil {
		var zeroVal T
		return zeroVal, fmt.Errorf("linked list is initiated but has no head value")
	} else if currentNode.Next == nil {
		this.Head = nil
		this.Size--
		return currentNode.Value, nil
	} else {
		for currentNode.Next.Next != nil {
			currentNode = currentNode.Next
		}
		tail := currentNode.Next.Value
		currentNode.Next = nil
		this.Size--
		return tail, nil
	}
}

// String: Default formatting method for the LinkedList.
func (this *LinkedList[T]) String() string {
	head := this.Head
	out := ""
	for i := this.Size; i > 0; i-- {
		out += fmt.Sprintf("%v", head.Value)
		head = head.Next
		if i > 1 {
			out += " -> "
		}
	}
	return out
}
