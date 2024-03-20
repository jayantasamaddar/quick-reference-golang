package bst

import (
	"encoding/json"
	"fmt"
	"strings"

	"log"

	"github.com/jayantasamaddar/quick-reference-golang/data-structures/ds"
)

/**********************************************************************************/
// (1) Tree Node
/**********************************************************************************/
type TreeNode[T ds.Number] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// Default formatting method
func (node *TreeNode[T]) String() string {
	return fmt.Sprintf("%v", node.Value)
}

/**********************************************************************************/
// (2) Binary Search Tree
/**********************************************************************************/
type BinarySearchTree[T ds.Number] struct {
	Root *TreeNode[T]
	size int
}

// Adds values to Binary Search Tree
func (this *BinarySearchTree[T]) Add(values ...T) {
	for _, value := range values {
		if this.Root == nil {
			this.Root = NewTreeNode(value)
		} else {
			this.searchAndInsert(this.Root, value)
		}
		this.size++
	}
}

// Removes values from the Binary Search Tree
func (this *BinarySearchTree[T]) Remove(values ...T) {
	for _, value := range values {
		this.Root = this.searchAndRemove(this.Root, value)
		this.size--
	}
}

// Checks if the tree is empty.
func (this *BinarySearchTree[T]) IsEmpty() bool {
	return this.Root == nil
}

// Removes all nodes from the tree.
func (this *BinarySearchTree[T]) Clear() {
	this.Root = nil
}

// Returns the number of nodes in the tree.
func (this *BinarySearchTree[T]) Size() int {
	return this.size
}

// Returns the node with the minimum value for the current Binary Search Tree
func (this *BinarySearchTree[T]) Min() *TreeNode[T] {
	return this.findMin(this.Root)
}

// Returns the node with the maximum value for the current Binary Search Tree
func (this *BinarySearchTree[T]) Max() *TreeNode[T] {
	return this.findMax(this.Root)
}

func (this *BinarySearchTree[T]) Contains(value T) bool {
	return this.contains(this.Root, value)
}

// Returns the MaxHeight: The distance between the root node and the bottom-most node
func (this *BinarySearchTree[T]) MaxHeight() int {
	return this.maxHeight(this.Root)
}

// Returns the MinHeight: The distance between the root node and the first node that doesn't have two children.
func (this *BinarySearchTree[T]) MinHeight() int {
	return this.minHeight(this.Root)
}

// A balanced tree is one where the difference between the `MinHeight` and the `MaxHeight` is at most 1.
func (this *BinarySearchTree[T]) IsBalanced() bool {
	return this.MaxHeight()-this.MinHeight() <= 1
}

/**************************************************************************************/
// Tree Traversal
/**************************************************************************************/

// Depth-first search (In-Order)
func (this *BinarySearchTree[T]) InOrder(order string) []T {
	if this.Root == nil {
		return nil
	}
	result := []T{}
	switch strings.ToLower(order) {
	case "d", "desc", "descending":
		this.descendingOrder(&result, this.Root)
	default:
		this.ascendingOrder(&result, this.Root)
	}
	return result
}

// Depth-first search (Pre-Order)
func (this *BinarySearchTree[T]) PreOrder() []T {
	return this.preOrder(this.Root)
}

// Depth-first search (Post-Order)
func (this *BinarySearchTree[T]) PostOrder() []T {
	return this.postOrder(this.Root)
}

// Breadth-first search (Level-Order)
func (this *BinarySearchTree[T]) LevelOrder() []T {
	return this.levelOrder(this.Root)
}

/*******************************************************************************************************/
// Iterators: Using the Iterator Pattern to traverse the Binary Search Tree.
// Returns: Iterator that uses the `HasNext` and `Next` to traverse the Binary Search Tree
/*******************************************************************************************************/
func (this *BinarySearchTree[T]) InOrderIterator() ds.Iterator[T] {
	return NewPreOrderIterator[T](this.Root)
}

func (this *BinarySearchTree[T]) PreOrderIterator() ds.Iterator[T] {
	return NewPreOrderIterator[T](this.Root)
}

func (this *BinarySearchTree[T]) PostOrderIterator() ds.Iterator[T] {
	return NewPreOrderIterator[T](this.Root)
}

func (this *BinarySearchTree[T]) LevelOrderIterator() ds.Iterator[T] {
	return NewPreOrderIterator[T](this.Root)
}

/*******************************************************************************************************/
// Generators: Using a Channel and Goroutine to traverse the Binary Search Tree.
// Returns a channel that can be used to iterate across values
/*******************************************************************************************************/
// InOrderTraversal: Returns a generator (receive-only channel) that can be used to iterate across values generated via in-order traversal
// Passing "d", "desc" or "descending" returns a generator, generating values in the descending order
// Anything else, including an empty string will return a generator, generating values in the ascending order
func (this *BinarySearchTree[T]) InOrderTraversal(order string) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		switch strings.ToLower(order) {
		case "d", "desc", "descending":
			this.traverseInOrderDescending(this.Root, ch)
		default:
			this.traverseInOrderAscending(this.Root, ch)
		}
	}()
	return ch
}

// PreOrderTraversal: Returns a generator (receive-only channel) that can be used to iterate across values generated via pre-order traversal
func (this *BinarySearchTree[T]) PreOrderTraversal() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		this.traversePreOrder(this.Root, ch)
	}()
	return ch
}

// PostOrderTraversal: Returns a generator (receive-only channel) that can be used to iterate across values generated via post-order traversal
func (this *BinarySearchTree[T]) PostOrderTraversal() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		this.traversePostOrder(this.Root, ch)
	}()
	return ch
}

// LevelOrderTraversal: Returns a generator (receive-only channel) that can be used to iterate across values generated via level-order traversal
// Level Order Traversal implements a breadth-first search mechanism where elements of the same order are returned (left element first).
func (this *BinarySearchTree[T]) LevelOrderTraversal() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		this.traverseLevelOrder(this.Root, ch)
	}()
	return ch
}

/**********************************************************************************/
// Serializations
/**********************************************************************************/
// Serialize Binary Search Tree in JSON
func (this *BinarySearchTree[T]) JSON() string {
	b, err := json.Marshal(this)
	if err != nil {
		log.Panicln("error marshalling BinarySearchTree", this)
	}
	return string(b)
}

// Serialize Binary Search Tree in JSON with indentation
func (this *BinarySearchTree[T]) JSONPretty() string {
	b, err := json.MarshalIndent(this, "", "  ")
	if err != nil {
		log.Panicln("error marshalling BinarySearchTree", this)
	}
	return string(b)
}

// String returns a string representation of the BinarySearchTree in a tree-like structure
func (this *BinarySearchTree[T]) String() string {
	var builder strings.Builder
	this.buildString(this.Root, &builder, "")
	return builder.String()
}

// Helper function to recursively build the string representation of the tree in a tree-like structure
func (this *BinarySearchTree[T]) buildString(node *TreeNode[T], builder *strings.Builder, prefix string) {
	if node == nil {
		return
	}

	// Print the current node value
	fmt.Fprintf(builder, "%s├── %v\n", prefix, node.Value)

	// Prepare the prefixes for the child nodes
	var childPrefix string
	if node.Left != nil {
		childPrefix = prefix + "│   "
	} else {
		childPrefix = prefix + "    "
	}

	// Recursively build the string for the left subtree
	this.buildString(node.Left, builder, childPrefix)

	// Prepare the prefixes for the right child nodes
	var rightChildPrefix string
	if node.Right != nil {
		rightChildPrefix = prefix + "│   "
	} else {
		rightChildPrefix = prefix + "    "
	}

	// Recursively build the string for the right subtree
	this.buildString(node.Right, builder, rightChildPrefix)
}

/**************************************************************************************/
// Helper Methods: Private (internal) methods that are used by the BST's main methods
/**************************************************************************************/
// Helper function that recursively inserts a value into the tree
func (this *BinarySearchTree[T]) searchAndInsert(node *TreeNode[T], value T) {
	if value == node.Value {
		return // Value already exists do nothing
	}
	if value < node.Value {
		if node.Left == nil {
			node.Left = NewTreeNode(value)
			return
		}
		this.searchAndInsert(node.Left, value)
	} else {
		if node.Right == nil {
			node.Right = NewTreeNode(value)
			return
		}
		this.searchAndInsert(node.Right, value)
	}
}

// Helper function that takes in a pointer to a TreeNode and value and removes it from the Binary Search Tree
func (this *BinarySearchTree[T]) searchAndRemove(node *TreeNode[T], value T) *TreeNode[T] {
	if node == nil {
		return nil
	}
	// If value matches current node value
	if value == node.Value {
		// (1) Scenario 1: Node is a leaf node
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// (2) Scenario 2: Node is a unary / single-child node
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		// (3) Scenario 3: Node has two children
		// Get left node's Max Node (This is the node whose value has to be set as the new node value)
		tempNode := node.Left
		for tempNode.Right != nil {
			tempNode = tempNode.Right
		}

		// Replace the current node Value with this tempNode value
		node.Value = tempNode.Value

		// Cleanup the replaced tempNode
		node.Left = this.searchAndRemove(node.Left, tempNode.Value)
	} else if value > node.Value {
		node.Right = this.searchAndRemove(node.Right, value)
	} else {
		node.Left = this.searchAndRemove(node.Left, value)
	}
	return node
}

// Takes in a TreeNode and returns the node with the minimum value
func (this *BinarySearchTree[T]) findMin(node *TreeNode[T]) *TreeNode[T] {
	if node == nil {
		return nil
	}
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Takes in a TreeNode and returns the node with the maximum value
func (this *BinarySearchTree[T]) findMax(node *TreeNode[T]) *TreeNode[T] {
	if node == nil {
		return nil
	}
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}

// Takes in a TreeNode and a value and returns if it exists in the Tree
func (this *BinarySearchTree[T]) contains(node *TreeNode[T], value T) bool {
	for node != nil {
		if value > node.Value {
			node = node.Right
		} else if value < node.Value {
			node = node.Left
		} else {
			return true
		}
	}
	return false
}

// Returns the MinHeight when given a TreeNode: The distance between the node and the first node that doesn't have two children.
func (this *BinarySearchTree[T]) minHeight(node *TreeNode[T]) int {
	if node == nil {
		return -1
	}

	left := this.minHeight(node.Left)
	right := this.minHeight(node.Right)

	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}

// Returns the MaxHeight when given a TreeNode: The distance between the node and the bottom-most node
func (this *BinarySearchTree[T]) maxHeight(node *TreeNode[T]) int {
	if node == nil {
		return -1
	}

	left := this.maxHeight(node.Left)
	right := this.maxHeight(node.Right)

	// Return the max of the two
	return max(left, right) + 1
}

// Takes in a destination slice buffer and a TreeNode pointer and does the in-order traversal in ascending order
func (this *BinarySearchTree[T]) ascendingOrder(buf *[]T, node *TreeNode[T]) {
	if node.Left != nil {
		this.ascendingOrder(buf, node.Left)
	}
	*buf = append(*buf, node.Value)
	if node.Right != nil {
		this.ascendingOrder(buf, node.Right)
	}
}

// Takes in a destination slice buffer and a TreeNode pointer and does the in-order traversal in descending order
func (this *BinarySearchTree[T]) descendingOrder(buf *[]T, node *TreeNode[T]) {
	if node.Right != nil {
		this.descendingOrder(buf, node.Right)
	}
	*buf = append(*buf, node.Value)
	if node.Left != nil {
		this.descendingOrder(buf, node.Left)
	}
}

// Takes in a node and returns a slice of values using pre-order traversal.
// Performs breadth-first search
func (this *BinarySearchTree[T]) levelOrder(node *TreeNode[T]) []T {
	if node == nil {
		return nil
	}
	result := []T{}
	queue := []*TreeNode[T]{node}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		result = append(result, currentNode.Value)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}
	return result
}

// Takes in a node and returns a slice of values using pre-order traversal.
// Visit the current node, traverse the left subtree, then the right subtree.
func (this *BinarySearchTree[T]) preOrder(node *TreeNode[T]) []T {
	if node == nil {
		return nil
	}
	// Initialize result slice with the current node's value
	result := []T{node.Value}
	left := this.preOrder(node.Left)
	right := this.preOrder(node.Right)

	// Concatenate the slices: result, left, right
	result = append(result, left...)
	result = append(result, right...)

	return result
}

// Takes in a node and returns a slice of values using post-order traversal.
// Traverse the left subtree, then the right subtree, and finally visit the current node.
func (this *BinarySearchTree[T]) postOrder(node *TreeNode[T]) []T {
	if node == nil {
		return nil
	}
	// Initialize result slice
	result := []T{}
	left := this.postOrder(node.Left)
	right := this.postOrder(node.Right)

	// Concatenate the slices: left, right, current node value
	result = append(result, left...)
	result = append(result, right...)
	result = append(result, node.Value)

	return result
}

func (this *BinarySearchTree[T]) traverseInOrderAscending(node *TreeNode[T], ch chan<- T) {
	if node == nil {
		return
	}
	this.traverseInOrderAscending(node.Left, ch)
	ch <- node.Value
	this.traverseInOrderAscending(node.Right, ch)
}

func (this *BinarySearchTree[T]) traverseInOrderDescending(node *TreeNode[T], ch chan<- T) {
	if node == nil {
		return
	}
	this.traverseInOrderDescending(node.Right, ch)
	ch <- node.Value
	this.traverseInOrderDescending(node.Left, ch)
}

func (this *BinarySearchTree[T]) traversePreOrder(node *TreeNode[T], ch chan<- T) {
	if node == nil {
		return
	}
	ch <- node.Value
	this.traversePreOrder(node.Left, ch)
	this.traversePreOrder(node.Right, ch)
}

func (this *BinarySearchTree[T]) traversePostOrder(node *TreeNode[T], ch chan<- T) {
	if node == nil {
		return
	}
	this.traversePostOrder(node.Left, ch)
	this.traversePostOrder(node.Right, ch)
	ch <- node.Value
}

func (this *BinarySearchTree[T]) traverseLevelOrder(node *TreeNode[T], ch chan<- T) {
	if node == nil {
		return
	}
	queue := []*TreeNode[T]{node}
	for len(queue) > 0 {
		// Do a dequeue operation
		currentNode := queue[0]
		queue = queue[1:]

		// Push value to channel
		ch <- currentNode.Value

		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}
}
