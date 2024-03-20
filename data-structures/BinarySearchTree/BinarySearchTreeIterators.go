package bst

import "github.com/jayantasamaddar/quick-reference-golang/data-structures/ds"

/**********************************************************************************/
// (3) Binary Search Tree Iterators
/**********************************************************************************/

/**********************************************************************************/
// (3c) Iterator for Binary Search Tree PreOrder Traversal
/**********************************************************************************/
type PreOrderIterator[T ds.Number] struct {
	stack []*TreeNode[T]
	root  *TreeNode[T]
}

// Check if there are more nodes to visit
func (iter *PreOrderIterator[T]) HasNext() bool {
	return len(iter.stack) > 0
}

// Get the next node value in pre-order traversal
func (iter *PreOrderIterator[T]) Next() T {
	// Pop the top node from the stack
	node := iter.stack[len(iter.stack)-1]
	iter.stack = iter.stack[:len(iter.stack)-1]

	// Push right child first so that it is visited after the left child
	if node.Right != nil {
		iter.stack = append(iter.stack, node.Right)
	}
	if node.Left != nil {
		iter.stack = append(iter.stack, node.Left)
	}

	return node.Value
}

// Reset the Iterator to its initialization
func (iter *PreOrderIterator[T]) Reset() {
	iter.stack = []*TreeNode[T]{iter.root}
}

/**********************************************************************************/
// (3d) Iterator for Binary Search Tree PostOrder Traversal
/**********************************************************************************/
type PostOrderIterator[T ds.Number] struct {
	stack []*postOrderStackNode[T]
	root  *TreeNode[T]
}

type postOrderStackNode[T ds.Number] struct {
	node   *TreeNode[T]
	visits int
}

// Check if there are more nodes to visit
func (iter *PostOrderIterator[T]) HasNext() bool {
	return len(iter.stack) > 0
}

// Get the next node value in pre-order traversal
func (iter *PostOrderIterator[T]) Next() T {
	var value T

	for len(iter.stack) > 0 {
		top := iter.stack[len(iter.stack)-1]

		if top.node == nil {
			iter.stack = iter.stack[:len(iter.stack)-1]
			continue
		}

		if top.visits == 2 {
			value = top.node.Value
			iter.stack = iter.stack[:len(iter.stack)-1]
			return value
		}

		if top.visits == 0 {
			iter.stack = append(iter.stack, &postOrderStackNode[T]{node: top.node.Right, visits: 0})
		}

		top.visits++
		if top.visits == 1 {
			iter.stack = append(iter.stack, &postOrderStackNode[T]{node: top.node.Left, visits: 0})
		}
	}

	return value
}

// Reset the Iterator to its initialization
func (iter *PostOrderIterator[T]) Reset() {
	stn := &postOrderStackNode[T]{node: iter.root, visits: 0}
	iter.stack = []*postOrderStackNode[T]{stn}
}
