package bst

import "github.com/jayantasamaddar/quick-reference-golang/data-structures/ds"

/**********************************************************************************/
// Factory Functions (Constructors)
/**********************************************************************************/
func NewBinarySearchTree[T ds.Number](initialValues ...T) *BinarySearchTree[T] {
	bst := &BinarySearchTree[T]{}
	bst.Add(initialValues...)
	return bst
}

func NewTreeNode[T ds.Number](value T) *TreeNode[T] {
	return &TreeNode[T]{Value: value}
}

// Initialize and return a new PreOrderIterator
func NewPreOrderIterator[T ds.Number](node *TreeNode[T]) *PreOrderIterator[T] {
	return &PreOrderIterator[T]{stack: []*TreeNode[T]{node}, root: node}
}

// Initialize and return a new PostOrderIterator
func NewPostOrderIterator[T ds.Number](node *TreeNode[T]) *PostOrderIterator[T] {
	stn := &postOrderStackNode[T]{node: node, visits: 0}
	return &PostOrderIterator[T]{stack: []*postOrderStackNode[T]{stn}, root: node}
}
