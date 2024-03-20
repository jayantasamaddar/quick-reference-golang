# Table of Contents

- [Table of Contents](#table-of-contents)
- [Binary Tree](#binary-tree)
- [Binary Search Tree](#binary-search-tree)
- [Terminology related to Trees](#terminology-related-to-trees)
- [Properties](#properties)
	- [Node Properties](#node-properties)
	- [Binary Tree Properties](#binary-tree-properties)
- [Methods](#methods)
	- [Add](#add)
	- [Remove](#remove)
	- [Minimum and Maximum Value](#minimum-and-maximum-value)
		- [Building the `min()` method](#building-the-min-method)
		- [Building the `max()` method](#building-the-max-method)
	- [`contains(value T)`: Finding a value within the Tree](#containsvalue-t-finding-a-value-within-the-tree)
		- [Recursive Approach](#recursive-approach)
		- [Iterative Approach](#iterative-approach)
	- [Height](#height)
		- [Building the `minHeight()` method:](#building-the-minheight-method)
		- [Building the `maxHeight()` method:](#building-the-maxheight-method)
	- [Balanced Tree](#balanced-tree)
		- [Building the `isBalanced()` method:](#building-the-isbalanced-method)
	- [Tree Traversal](#tree-traversal)
		- [Tree Traversal: Idiomatic approach in Go](#tree-traversal-idiomatic-approach-in-go)
		- [Tree Traversal: Depth-First Search](#tree-traversal-depth-first-search)
			- [Building the `InOrderTraversal(order string)` method to search the Binary Search Tree](#building-the-inordertraversalorder-string-method-to-search-the-binary-search-tree)
			- [Building the `preOrder()` method to search the Binary Search Tree](#building-the-preorder-method-to-search-the-binary-search-tree)
			- [Building the `postOrder()` method to search the Binary Search Tree](#building-the-postorder-method-to-search-the-binary-search-tree)
		- [Tree Traversal: Breadth-First Search](#tree-traversal-breadth-first-search)
			- [Building the `levelOrder()` method to search the Binary Tree](#building-the-levelorder-method-to-search-the-binary-tree)
		- [Tree Traversal: Use Cases](#tree-traversal-use-cases)
			- [In-Order Traversal](#in-order-traversal)
			- [Pre-Order Traversal](#pre-order-traversal)
			- [Post-Order Traversal](#post-order-traversal)
			- [Level-Order Traversal](#level-order-traversal)
- [List of Solved Binary Tree / Binary Search Tree Algorithms](#list-of-solved-binary-tree--binary-search-tree-algorithms)
	- [Leetcode](#leetcode)
- [References](#references)

---

# Binary Tree

A **Binary Tree** is a node-based tree data structure that has at most branches for every single node.
Other tree data structures can have any number of branches for a single node. However a Binary Tree can have at most two branches for every single node.

All Binary Trees have the following characteristics:

- Has exactly one root node.
- Each parent node has at most two children.
- There can only be one single unique path from the root to any node.

The **Binary Search Tree** is one type of a Binary Tree.
The **Heap** is an partially ordered Binary Tree, that can be either,

- **Min Heap**: The parent nodes <= child nodes.
- **Max Heap**: The parent nodes are >= child nodes.

**Trie** and **State-Space Tree** are examples of non-binary trees.

---

# Binary Search Tree

A **Binary Search Tree** is a specific node-based binary tree data structure that has the following characteristics:

- The left subtree of a node contains only nodes with keys < the node’s key.
- The right subtree of a node contains only nodes with keys > the node’s key.
- The left and right subtree each must also be a binary search tree.

**Note:** Usually a Binary Tree doesn't allow duplicates. However, there are variants of Binary Search Trees that may allow duplicates.

In that case, two approaches can be adapted:

1. The left subtree of nodes will only contain keys with value **_lesser than or equal to_** (<=) the node's key and the right subtree of nodes will only contain keys with value **_greater than or equal to_** (>=) the node's key. However, duplicates may be separated by any number of levels, so checking for a duplicates existence is not a very optimum operation.

2. In order to avoid the above issue to not represent duplicates structurally (as separate nodes), is to use a counter that counts the number of occurences of the key. This simplifies lookup, removal and insertion operations, at the expense of some extra bytes and counter operations.

![Binary Search Tree](https://media.geeksforgeeks.org/wp-content/uploads/BSTSearch.png)

Let's take a look at the terminologies related to trees in general.

---

# Terminology related to Trees

- **Nodes**: The data points in a tree are called nodes. For a binary tree, each node can have a left and a right node attached to it.
- **Root Node**: The topmost node is called the root node.
- **Parent Node**: A node with branches leading to other nodes (subtree) is called a parent node.
- **Child Node**: A node with a parent node is called a child node. Children are parents of their own subtree.
- **Leaf Node**: A node that do not have any children is called a Leaf Node.
- **Unary Node**: A node that has only one child (either Left or Right) is called an Unary Node.
- **Siblings**: Nodes sharing the same parent node.
- **Height**: The height in a tree represents the distance from the root node and any given node.
  - **minHeight**: The distance between the root node and _the first node that doesn't have two children_.
  - **maxHeight**: The distance between the root node and the _bottom-most node_.
- **Balanced Tree**: A balanced tree is one where the difference between the `minHeight` and the `maxHeight` is at most 1.

---

# Properties

## Node Properties

- **`Value`** − The value stored by the node.
- **`Left`** − The left child node of this node.
- **`Right`** − The right child node of this node.

We can encapsulate this in a struct **`TreeNode`** and use the struct to create Node objects:

```go
type TreeNode[T ds.Number] struct {
    Value T
    Left *TreeNode[T]
    Right *TreeNode[T]
}

// Default formatting method
func (node *TreeNode[T]) String() string {
	return fmt.Sprintf("%v", node.Value)
}

// Constructor
func NewTreeNode[T ds.Number](value T) *TreeNode[T] {
	return &TreeNode[T]{Value: value}
}
```

## Binary Tree Properties

- **`Root`** − The root node of the Binary Tree.

The main **`BinarySearchTree`** struct may look something like this:

```go
type BinarySearchTree[T ds.Number] struct {
    Root *TreeNode[T]
}

// Constructor
func NewBinarySearchTree[T ds.Number](initialValues ...T) *BinarySearchTree[T] {
	bst := &BinarySearchTree[T]{}
	bst.Add(initialValues...)
	return bst
}
```

---

# Methods

**Basic Operations**

- **`Add`**: Adds an element to the tree.
- **`Remove`**: Removes an element from the tree.
- **`Min`**: Returns the node with the minimum value.
- **`Max`**: Returns the node with the maximum value.
- **`MinHeight`**: Returns the minimum height of the tree
- **`MaxHeight`**: Returns the maximum height of the tree.
- **`Contains`**: Returns whether value exists in the tree.
- **`IsBalanced`**: Returns whether the tree is balanced.
- **`Size`**: Returns the number of nodes in the tree.
- **`IsEmpty`**: Checks if the tree is empty.
- **`Clear`**: Removes all nodes from the tree
- **`String`**: Returns a string representation of the tree.

**Advanced Operations**

- **`InOrder`**: Depth First Search that traverses the tree in-order (left-root-right) and returns values.
- **`PreOrder`**: Depth First Search that traverses the tree pre-order (root-left-right) and returns values.
- **`PostOrder`**: Depth First Search that traverses the tree post-order (left-right-root) and returns values.
- **`LevelOrder`**: Breadth First Search that traverses the tree level-order and returns values
- **`Clone`**: Creates a copy of the tree.
- **`Union`**: Combines two trees into a new tree.
- **`Intersection`**: Returns a new tree containing common elements of two trees.
- **`Difference`**: Returns a new tree containing elements that are present in one tree but not the other.
- **`FindSuccessor`**: Finds the successor of a given node.
- **`FindPredecessor`**: Finds the predecessor of a given node.
- **`RotateLeft`**: Performs a left rotation on the tree.
- **`RotateRight`**: Performs a right rotation on the tree.

---

## Add

```go
// Adds values to Binary Search Tree
func (this *BinarySearchTree[T]) Add(values ...T) {
	for _, value := range values {
		if this.Root == nil {
			this.Root = NewTreeNode(value)
		} else {
			this.searchAndInsert(this.Root, value)
		}
	}
}

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
```

---

## Remove

The process involves:

- Finding the node to be removed.
- Choosing a path to traverse, in order to find another value to replace it in the tree.

Let's take the following Binary Tree as an example:

```
        50
       /  \
      /    \
     /      \
   40        70
  /  \      /  \
35    45  60    80
```

If `50` is to be removed, the replacement value can be found by either going left or right.

1. When going left, we find the rightmost node and replace the to-be-removed node's value with this value. If no node to the right exists, then the value of the left node is taken as the replacement.

**To illustrate**,

```bash
        45
       /  \
      /    \
     /      \
   40        70
  /         /  \
35        60    80
```

Let's write the code for removal of a key using the left traversal path for a Binary Search Tree:

```go
// Removes values from the Binary Search Tree
func (this *BinarySearchTree[T]) Remove(values ...T) {
	for _, value := range values {
		this.Root = this.searchAndRemove(this.Root, value)
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
```

2. When going right, we find the leftmost node and replace the to-be-removed node's value with this value. If no node to the left exists, then the value of the right node is taken as the replacement.

**To illustrate**,

```
        60
       /  \
      /    \
     /      \
   40        70
  /  \         \
35    45        80
```

Let's write the code for removal of a key using the right traversal path for a Binary Search Tree:

```go
// Removes values from the Binary Search Tree
func (this *BinarySearchTree[T]) Remove(values ...T) {
	for _, value := range values {
		this.Root = this.searchAndRemove(this.Root, value)
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
		tempNode := node.Right
		for tempNode.Left != nil {
			tempNode = tempNode.Left
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
```

We can pick any of these two traversal paths for removal. Both result in valid Binary Search Trees.

---

## Minimum and Maximum Value

### Building the `min()` method

The minimum value of a Binary Search Tree is the value of the leftmost node.

To do this, we have to do a traversal towards the leftmost node until a leaf node is reached.

```go
// Returns the node with the minimum value for the current Binary Search Tree
func (this *BinarySearchTree[T]) Min() *TreeNode[T] {
	return this.findMin(this.Root)
}

// Takes in a TreeNode and returns the node with the minimum value
func (this *BinarySearchTree[T]) findMin(node *TreeNode[T]) *TreeNode[T] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
```

---

### Building the `max()` method

The maximum value of a Binary Search Tree is value of the rightmost node.

To do this, we have to do a traversal towards the rightmost node until a leaf node is reached.

```go
// Returns the node with the maximum value for the current Binary Search Tree
func (this *BinarySearchTree[T]) Max() *TreeNode[T] {
	return this.findMax(this.Root)
}

// Takes in a TreeNode and returns the node with the maximum value
func (this *BinarySearchTree[T]) findMax(node *TreeNode[T]) *TreeNode[T] {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}
```

---

## `contains(value T)`: Finding a value within the Tree

This method returns a boolean whether the value passed as argument, exists or not.

We can take two approaches to solving this:

1. Recursive Approach
2. Iterative Approach

---

### Recursive Approach

```go
func (this *BinarySearchTree[T]) Contains(value T) bool {
	return this.contains(this.Root, value)
}

// Takes in a TreeNode and a value and returns if it exists in the Tree
func (this *BinarySearchTree[T]) contains(node *TreeNode[T], value T) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return this.contains(node.Left, value)
	} else if value > node.Value {
		return this.contains(node.Right, value)

	} else {
		return true
	}
}
```

The above solution traverses the tree recursively based on the value being searched for, comparing it with the values in each node.

However, there are a couple of considerations to keep in mind:

1. **Efficiency**: The provided implementation has a time complexity of O(h), where h is the height of the tree. In the worst-case scenario, when the tree is unbalanced and resembles a linked list, the time complexity can become O(n), where n is the number of nodes in the tree. If you have a balanced tree, the time complexity is O(log n).

2. **Iterative Implementation**: While the recursive approach is straightforward and easy to understand, it might not be the most efficient for very deep trees due to potential stack overflow issues. An alternative approach is to implement an iterative version using a loop and a stack or queue data structure.

---

### Iterative Approach

```go
func (this *BinarySearchTree[T]) Contains(value T) bool {
	return this.contains(this.Root, value)
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
```

---

## Height

Height in a tree represents the distance from the root node to any given node.

- **minHeight**: The distance between the root node and _the first node that doesn't have two children_.
- **maxHeight**: The distance between the root node and the _bottom-most node_.

---

### Building the `minHeight()` method:

**Tree**:

```bash
		    50
	       /  \
	      /    \
	     /      \
	   40        70
	  /  \      /  \
	35    45  60    80
                    \
                    100
```

**Result**: 2

```go
// Returns the MinHeight when given a TreeNode: The distance between the node and the first node that doesn't have two children.
func (this *BinarySearchTree[T]) minHeight(node *TreeNode[T]) int {
	if node == nil {
		return -1
	}

	left := this.minHeight(node.Left)
	right := this.minHeight(node.Right)
   fmt.Println("For Value:", node.Value, "|", "Left:", left, ",", "Right:", right)

	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}

// Returns the MinHeight of the BST: The distance between the root node and the first node that doesn't have two children.
func (this *BinarySearchTree[T]) MinHeight() int {
	return this.minHeight(this.Root)
}
```

**Explanation**:

For Value: 35 | Left: -1 , Right: -1
For Value: 45 | Left: -1 , Right: -1
For Value: 40 | Left: 0 , Right: 0
For Value: 60 | Left: -1 , Right: -1
For Value: 100 | Left: -1 , Right: -1
For Value: 80 | Left: -1 , Right: 0
For Value: 70 | Left: 0 , Right: 0
For Value: 50 | Left: 1 , Right: 1

---

### Building the `maxHeight()` method:

**Tree**:

```bash
		    50
	       /  \
	      /    \
	     /      \
	   40        70
	  /  \      /  \
	35    45  60    80
                     \
                     100
```

**Result**: 3

```go
// Returns the MaxHeight when given a TreeNode: The distance between the node and the bottom-most node
func (this *BinarySearchTree[T]) maxHeight(node *TreeNode[T]) int {
	if node == nil {
		return -1
	}

	left := this.maxHeight(node.Left)
	right := this.maxHeight(node.Right)
   fmt.Println("For Value:", node.Value, "|", "Left:", left, ",", "Right:", right)

	// Return the max of the two
	return max(left, right) + 1
}

// Returns the MaxHeight: The distance between the root node and the bottom-most node
func (this *BinarySearchTree[T]) MaxHeight() int {
	return this.maxHeight(this.Root)
}
```

**Explanation**:

For Value: 35 | Left: -1 , Right: -1
For Value: 45 | Left: -1 , Right: -1
For Value: 40 | Left: 0 , Right: 0
For Value: 60 | Left: -1 , Right: -1
For Value: 100 | Left: -1 , Right: -1
For Value: 80 | Left: -1 , Right: 0
For Value: 70 | Left: 0 , Right: 1
For Value: 50 | Left: 1 , Right: 2

---

## Balanced Tree

A balanced tree is one where the difference between the `minHeight` and the `maxHeight` is at most 1.

### Building the `isBalanced()` method:

This method finds out if a Binary Search Tree is balanced.

```go
// A balanced tree is one where the difference between the `MinHeight` and the `MaxHeight` is at most 1.
func (this *BinarySearchTree[T]) IsBalanced() bool {
	return this.MaxHeight()-this.MinHeight() <= 1
}
```

---

## Tree Traversal

Tree Traversal methods can be used to explore tree data structures and basically find all the values in the tree.

There are two approaches to tree traversal:

1. **Depth-First Search**: In depth-first search, a given subtree is explored as deeply as possible before the search continues on another subtree.

There are three ways Depth-First search can be done:-

- **In-Order Traversal**: In-Order search returns the values of the Binary Search Tree, either in ascending or descending order.
- **Pre-Order Traversal**:
- **Post-Order Traversal**:

2. **Breadth-First Search**: A breadth-first search explores all the nodes in the given level before continuing to the next level.

   - **Level-Order Traversal**:

---

### Tree Traversal: Idiomatic approach in Go

We can implement tree traversal in Go using various mechanisms:

1. Simply return **ALL** the values as a slice, either via the various Depth-First Search methods (In-Order, Pre-Order, Post-Order) or Breadth-First Search (Level-Order). This mechanism doesn't allow any control over the returned values, or manipulating a particular value, as we simply get a whole slice at the end. This method always traverses the entire tree.
2. Using the Iterator design pattern. The traditional iterator pattern has a `HasNext()` method to check if there are more values and uses it to iterate using the `Next()` method. This may work, however Go has a better, more idiomatic way of implementing the Iterator design pattern using channels that do not require these methods.
3. Using Channels in Go, we can return the channel as a generator for a particular method and then `range` over the channel to get the values one-by-one. This allows us to exit the traversal on a condition, manipulate each value (mapping function) and store it into another data structure. This allows for maximum flexibility and is the idiomatic way of Tree Traversal in Go.

---

### Tree Traversal: Depth-First Search

#### Building the `InOrderTraversal(order string)` method to search the Binary Search Tree

In-Order search returns the values of the Binary Search Tree, either in ascending or descending order.

The **`InOrderTraversal(order string)`** takes a string for the order we want to return. This is either "d", "desc", "descending" for descending. Anything else including an empty string, returns in ascending order.

```
               50
             /    \
            /      \
           /        \
          /          \
         /            \
        /              \
       /                \
      40                70
    /    \            /    \
   35    45         60      80
                  /    \
                 /      \
                55      65
                       /
                      63
                     /
                    62
```

> **Result:**
>
> Ascending Order: **[ 35, 40, 45, 50, 55, 60, 62, 63, 65, 70, 80 ]**
>
> Descending Order: **[ 80, 70, 65, 63, 62, 60, 55, 50, 45, 40, 35 ]**

**Procecdure:**

- **For Ascending order**

  - It does a depth-first search, left-first, recursively. What that means is for every node it finds out its left node and returns the value until there is no left node left.
  - Thus, for each node, this is above operation is done for both the left node and the right node with the left node, completing it's operations first (left-first).
  - Since it is a recursive function - the first value is the leftmost value and it goes up from there.

- **For Descending Order**
  - It does a depth-first search, right-first, recursively. What that means is for every node it finds out its right node and returns the value until there is no right node left.
  - Thus, for each node, this is done for both the right node and the left node with the right node, completing it's operations first (right-first).
  - Since it is a recursive function - the first value is the rightmost value and it goes up from there.

**Syntax:**

```go
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

func main() {
  // In-order Traversal (Ascending): Idiomatic Go (channel + goroutine) Implementation
	generatorAscending := tree.InOrderTraversal("asc")
	result := make([]int, 0)
	for value := range generatorAscending {
		result = append(result, value)
	}
	fmt.Println("In-order Traversal (Ascending), using Channel + Goroutine:", result)

	// In-order Traversal (Descending): Idiomatic Go (channel + goroutine) Implementation
	generatorDescending := tree.InOrderTraversal("desc")
	result = make([]int, 0)
	for value := range generatorDescending {
		result = append(result, value)
	}
	fmt.Println("In-order Traversal (Descending), using Channel + Goroutine:", result)
}
```

---

#### Building the `preOrder()` method to search the Binary Search Tree

This method returns the array of values of the Binary Search in the pre-order.

```
               50
             /    \
            /      \
           /        \
          /          \
         /            \
        /              \
       /                \
      40                70
    /    \            /    \
   35    45         60      80
                  /    \
                 /      \
                55      65
                       /
                      63
                     /
                    62
```

> **Result: [ 50, 40, 35, 45, 70, 60, 55, 65, 63, 62, 80 ]**

**Procedure:**

- Visit the root.
- Traverse the left subtree, i.e., call Preorder(left-subtree)
- Traverse the right subtree, i.e., call Preorder(right-subtree)

**Syntax:**

```go
// PreOrderTraversal: Returns a generator (receive-only channel) that can be used to iterate across values generated via pre-order traversal
func (this *BinarySearchTree[T]) PreOrderTraversal() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		this.traversePreOrder(this.Root, ch)
	}()
	return ch
}

func (this *BinarySearchTree[T]) traversePreOrder(node *TreeNode[T], ch chan<- T) {
	if node == nil {
		return
	}
	ch <- node.Value
	this.traversePreOrder(node.Left, ch)
	this.traversePreOrder(node.Right, ch)
}

func main() {
  // Pre-order Traversal: Idiomatic Go (channel + goroutine) Implementation
	generatorPreOrder := tree.PreOrderTraversal()
	result = make([]int, 0)
	for value := range generatorPreOrder {
		result = append(result, value)
	}
	fmt.Println("Pre-order Traversal, using Channel + Goroutine:", result)
}
```

#### Building the `postOrder()` method to search the Binary Search Tree

This method returns the array of values of the Binary Search in the post-order.

```
               50
             /    \
            /      \
           /        \
          /          \
         /            \
        /              \
       /                \
      40                70
    /    \            /    \
   35    45         60      80
                  /    \
                 /      \
                55      65
                       /
                      63
                     /
                    62
```

> **Result: [ 35, 45, 40, 55, 62, 63, 65, 60, 80, 70, 50 ]**

**Procedure:**

- Traverse the left subtree, i.e., call Postorder(left-subtree)
- Traverse the right subtree, i.e., call Postorder(right-subtree)
- Visit the root.

**Syntax:**

```go
// PostOrderTraversal: Returns a generator (receive-only channel) that can be used to iterate across values generated via post-order traversal
func (this *BinarySearchTree[T]) PostOrderTraversal() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		this.traversePostOrder(this.Root, ch)
	}()
	return ch
}

func (this *BinarySearchTree[T]) traversePostOrder(node *TreeNode[T], ch chan<- T) {
	if node == nil {
		return
	}
	this.traversePostOrder(node.Left, ch)
	this.traversePostOrder(node.Right, ch)
	ch <- node.Value
}

func main() {
  // Post-order Traversal: Idiomatic Go (channel + goroutine) Implementation
	generatorPostOrder := tree.PostOrderTraversal()
	result = make([]int, 0)
	for value := range generatorPostOrder {
		result = append(result, value)
	}
	fmt.Println("Post-order Traversal, using Channel + Goroutine:", result)
}
```

---

### Tree Traversal: Breadth-First Search

A breadth-first search explores all the nodes in the given level before continuing to the next level. The next level is usually traversed left to right.

#### Building the `levelOrder()` method to search the Binary Tree

```
               50
             /    \
            /      \
           /        \
          /          \
         /            \
        /              \
       /                \
      40                70
    /    \            /    \
   35    45         60      80
                  /    \
                 /      \
                55      65
                       /
                      63
                     /
                    62
```

> **Result: [ 50, 40, 70, 35, 45, 60, 80, 55, 65, 63, 62 ]**

**Procedure:**

- We declare a Queue array to load the current node and a results array to store values.
- We start by adding the root node to the Queue.
- Begin a loop on the Queue where we dequeue the first node in the queue and add it's value to the results array.
- Inspect the node's left node - if it is not null, enqueue this left node.
- Inspect the node's right node - if it is not null, enqueue this right node.
- This process continues until the queue is empty (Q.length === 0).

**Syntax:**

```go
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

func main() {
  	// Level-order Traversal: Idiomatic Go (channel + goroutine) Implementation
	generatorLevelOrder := tree.LevelOrderTraversal()
	result = make([]int, 0)
	for value := range generatorLevelOrder {
		result = append(result, value)
	}
	fmt.Println("Level-order Traversal, using Channel + Goroutine:", result)
}
```

---

### Tree Traversal: Use Cases

#### In-Order Traversal

In the case of binary search trees (BST), In-Order traversal gives nodes in non-decreasing order. To get nodes of BST in non-increasing order, a variation of Inorder traversal where Inorder traversal is reversed can be used.

#### Pre-Order Traversal

Preorder traversal is used to create a copy of the tree. Preorder traversal is also used to get prefix expression on an expression tree. Please see http://en.wikipedia.org/wiki/Polish_notation to know why prefix expressions are useful.

#### Post-Order Traversal

Postorder traversal is used to delete the tree. Please see the question for the deletion of a tree for details. Postorder traversal is also useful to get the postfix expression of an expression tree. Please see http://en.wikipedia.org/wiki/Reverse_Polish_notation for the usage of postfix expression.

#### Level-Order Traversal

Breadth-first search can be used to solve many problems in graph theory, for example:

- Copying garbage collection, Cheney's algorithm
- Finding the shortest path between two nodes u and v, with path length measured by number of edges ([an advantage over depth-first search](https://web.archive.org/web/20141029100806/http://opendatastructures.org/versions/edition-0.1e/ods-java/12_3_Graph_Traversal.html#SECTION001531000000000000000))
- (Reverse) Cuthill–McKee mesh numbering
- Ford–Fulkerson method for computing the maximum flow in a flow network
- Serialization/Deserialization of a binary tree vs serialization in sorted order, allows the tree to be re-constructed in an efficient manner.
- Construction of the failure function of the Aho-Corasick pattern matcher.
- Testing bipartiteness of a graph.
- Implementing parallel algorithms for computing a graph's transitive closure.

---

# List of Solved Binary Tree / Binary Search Tree Algorithms

## Leetcode

- [704 - Binary Search](../../leetcode-challenges/javascript/704%20-%20Binary%20Search/index.js)
- [374 - Guess Number Higher or Lower](../../leetcode-challenges/javascript/374%20-%20Guess%20Number%20Higher%20or%20Lower/index.js)
- [700 - Search in a Binary Search Tree](../../leetcode-challenges/javascript/700%20-%20Search%20in%20a%20BST/index.js)
- [104 - Maximum Depth of Binary Tree](../../leetcode-challenges/javascript/104%20-%20Maximum%20Depth%20of%20Binary%20Tree/index.js)
- [111 - Minimum Depth of Binary Tree](../../leetcode-challenges/javascript/111%20-%20Minimum%20Depth%20of%20a%20Binary%20Tree/index.js)
- [701 - Insert into a Binary Search Tree](../../leetcode-challenges/javascript/701%20-%20Insert%20into%20a%20Binary%20Search%20Tree/index.js)
- [222 - Count Complete Tree Nodes](../../leetcode-challenges/javascript/222%20-%20Count%20Complete%20Tree%20Nodes/index.js)
- [450 - Delete Node in a BST](../../leetcode-challenges/javascript/450%20-%20Delete%20Node%20in%20a%20BST/index.js)

---

# References

- [Binary Search Tree - Data Structure Construction](BinarySearchTree.js)
- **@datastructures-js/binary-search-tree | [Github](https://github.com/datastructures-js/binary-search-tree)**
- [Binary Tree Algorithms for Technical Interviews by freecodecamp.org](https://www.youtube.com/watch?v=fAAZixBzIAI)
- [Heap - Data Structure Construction]
