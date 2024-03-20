package ll

import "github.com/jayantasamaddar/quick-reference-golang/data-structures/ds"

func NewListNode[T ds.Type](item T) *ListNode[T] {
	return &ListNode[T]{
		Value: item,
	}
}

func NewLinkedList[T ds.Type](items ...T) *LinkedList[T] {
	llist := &LinkedList[T]{}
	llist.Append(items...)
	return llist
}
