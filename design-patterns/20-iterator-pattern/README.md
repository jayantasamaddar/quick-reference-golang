# Table of Contents

- [Table of Contents](#table-of-contents)
- [Iterator Pattern](#iterator-pattern)
	- [Introduction](#introduction)
	- [Implementation in Go](#implementation-in-go)
		- [In-built Iteration for arrays, slices and maps](#in-built-iteration-for-arrays-slices-and-maps)
		- [Implementation using Channel and Goroutine as Generator](#implementation-using-channel-and-goroutine-as-generator)
		- [Implemnenting Iterator Pattern using Structs](#implemnenting-iterator-pattern-using-structs)
	- [Tree Traversal](#tree-traversal)
		- [Using Iterator Pattern](#using-iterator-pattern)
		- [Using Idiomatic Go (Channel + Goroutine)](#using-idiomatic-go-channel--goroutine)
- [Summary](#summary)

---

# Iterator Pattern

## Introduction

The Iterator Pattern is a behavioral design pattern that provides a way to access elements of a collection sequentially without exposing the underlying details of the collection. It defines an interface for iterating over a collection of objects without exposing the underlying implementation.

**Characteristics of an Iterator**:

- An iterator is an object that provides a way to traverse or iterate over elements of a collection sequentially.
- It typically defines methods like `HasNext()` to check if there are more elements, and `Next()` to retrieve the next element in the iteration.
- Iterators are usually used to access elements of a collection in a controlled manner, one at a time.

---

## Implementation in Go

### In-built Iteration for arrays, slices and maps

Go has an inbuilt iterator support with the `range` keyword to iterate over the following types:

1. **Arrays**: When you use range with an array, it iterates over each element in the array, allowing access to the index and value.
2. **Slices**: range iterates over each element in a slice, , allowing access to the index and value.
3. **Strings**: range iterates over each Unicode character in a string, allowing access to the index and value.
4. **Maps**: With maps, range iterates over key-value pairs.
5. **Channels**: When used with channels, `range` iterates over values sent to the channel until the channel is closed.

```go
package main

import "fmt"

func main() {
	slice := []int{15, 25, 35, 45, 55}
	mapped := map[string]int{
		"Rohit":  35,
		"Adraha": 55,
		"Ravi":   20,
	}
	str := "hello"
	ch := make(chan int)

	fmt.Println(std.Yellow + "Using range with slices (Works similarly for arrays):" + std.Reset)
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %v\n", index, value)
	}
	/*
		Index: 0, Value: 15
		Index: 1, Value: 25
		Index: 2, Value: 35
		Index: 3, Value: 45
		Index: 4, Value: 55
	*/

	fmt.Println(std.Yellow + "Using range with maps:" + std.Reset)
	for key, value := range mapped {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	/*
		Key: Rohit, Value: 35
		Key: Adraha, Value: 55
		Key: Ravi, Value: 20
	*/

	fmt.Println(std.Yellow + "Using range with strings:" + std.Reset)
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
	/*
		Index: 0, Character: h
		Index: 1, Character: e
		Index: 2, Character: l
		Index: 3, Character: l
		Index: 4, Character: o
	*/

	fmt.Println(std.Yellow + "Using range with channels:" + std.Reset)
	// Run Goroutine to push integers to channel
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	for value := range ch {
		fmt.Println(value)
	}
	/*
		Using range with channels:
		0
		1
		2
	*/
}
```

---

### Implementation using Channel and Goroutine as Generator

A generator is a function or method that can be paused and resumed, returning a sequence of values lazily.
It produces values on-the-fly, potentially indefinitely, **without needing to precompute all values upfront**.
Generators are often used to produce sequences of values in a more concise and readable manner, especially when dealing with potentially large or infinite sequences.

We can achieve a generator by using function that returns a channel and then iterating using the `range` keyword over the channel. The iteration continues until the channel is closed, which, is controlled by a goroutine. This is considered the most idiomatic way of iterating complex structures in Go.

Here is an example of how this would work:

```go
package main

import "fmt"

type Person struct {
    FirstName, MiddleName, LastName string
}

func (p *Person) NamesGenerator() <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        out <- p.FirstName
        if len(p.MiddleName) > 0 {
            out <- p.MiddleName
        }
        out <- p.LastName
    }()
    return out
}

func main() {
    p := Person{"Alexander", "Graham", "Bell"}
    for name := range p.NamesGenerator() {
        fmt.Println(name)
    }
}
```

---

### Implemnenting Iterator Pattern using Structs

In Go, for objects, there is no explicit interface for iterators like in some other languages, but we can create our own iterator pattern using a combination of a custom type and methods.

Let's create a simple example to demonstrate the Iterator Pattern in Go:

This is the classic way of doing an Iterator just like in C++. Typically when we talk about the iterator design pattern, we mainly talk about explicitly constructed iterators like this one. When we talk of Iterators, we talk about separate structures which are used to track the position of where we are in the object that's being iterated and have a pointer to that object so that we can go into it and get some information that we actually need to.

A Basic Example:

```go
package main

import "fmt"

// Iterable defines the interface for the collection that needs to be iterated.
type Iterable interface {
	// Iterator returns an iterator for the collection.
	Iterator() Iterator
}

// Iterator defines the interface for iterating over the collection.
type Iterator interface {
	// HasNext returns true if there are more elements to iterate over.
	HasNext() bool
	// Next returns the next element in the iteration.
	Next() interface{}
}

// Concrete implementation of an iterable collection (e.g., a list).
type MyCollection struct {
	data *[]int
}

// Concrete implementation of an iterator for MyCollection.
type MyCollectionIterator struct {
	collection *MyCollection
	index      int
}

// NewMyCollection creates a new instance of MyCollection with initial data.
func NewMyCollection(data *[]int) *MyCollection {
	return &MyCollection{data}
}

// Iterator returns an iterator for MyCollection.
func (c *MyCollection) Iterator() Iterator {
	return &MyCollectionIterator{collection: c, index: 0}
}

// HasNext returns true if there are more elements to iterate over.
func (i *MyCollectionIterator) HasNext() bool {
	return i.index < len(*i.collection.data)
}

// Next returns the next element in the iteration.
func (i *MyCollectionIterator) Next() Iterable {
	if i.HasNext() {
		value := (*i.collection.data)[i.index]
		i.index++
		return value
	}
	return nil
}

func main() {
	collection := NewMyCollection(&[]int{1, 2, 3, 4, 5})
	iterator := collection.Iterator()

	fmt.Println("Iterating over the collection:")
	for iterator.HasNext() {
		element := iterator.Next()
		fmt.Println(element)
	}
}
```

Let's look at a more complex example in a Tree Traversal case.

---

## Tree Traversal

### Using Iterator Pattern

Traversing a tree is a situation where you cannot get away without using an Iterator.

```go
func main() {
    // See the ds package for the BinarySearchTree data structure
    bst := ds.NewBinarySearchTree[int](50, 40, 70, 45, 35, 80, 60, 55, 65, 63, 62)

    // Tree Traversal (Returns all values by either doing Depth-First Search or Breadth-First-Search)
	fmt.Println("In Order (Ascending):", bst.InOrder(""))      // In Order (Ascending): [35 40 45 50 55 60 62 63 65 70 80]
	fmt.Println("In Order (Descending):", bst.InOrder("desc")) // In Order (Descending): [80 70 65 63 62 60 55 50 45 40 35]

	fmt.Println("In Order (Pre-Order):", bst.PreOrder())   // In Order (Pre-Order): [50 40 35 45 70 60 55 65 63 62 80]
	fmt.Println("In Order (Post-Order):", bst.PostOrder()) // In Order (Post-Order): [35 45 40 55 62 63 65 60 80 70 50]

    fmt.Println("Level Order:", bst.LevelOrder()) // Level Order: [50 40 70 35 45 60 80 55 65 63 62]
}
```

In the above default, `BinarySearchTree` implementation, the Depth-First-Search and Breadth-First Search results in the values in the Binary Search Tree returned as a slice via their respective traversal methods.

What if we want to lazy generate these values? Why would we do that? Well, we could then run some complex conditional logic, targetting specific values. This is when we need an Iterator or a Generator.

In the `BinarySearchTree.go`,

```go
type Iterable[T Type] interface {
	Iterator[T]
}
type Iterator[T Type] interface {
	HasNext() bool
	Next() T
}
type PreOrderIterator[T Number] struct {
	stack []*TreeNode[T]
    root *TreeNode[T]
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

// Factory Function
func NewPreOrderIterator[T Number](node *TreeNode[T]) *PreOrderIterator[T] {
	return &PreOrderIterator[T]{stack: []*TreeNode[T]{node}}
}

// Method for a BinarySearchTree to generate a PreOrderIterator
func (this *BinarySearchTree[T]) PreOrderIterator() Iterator[T] {
	return NewPreOrderIterator[T](this.Root)
}

func main() {
    // Using an Iterator
	iter := bst.PreOrderIterator()
	for iter.HasNext() {
        // Run conditional logic (if any)
		fmt.Println(iter.Next())
	}
    /*
        50
        40
        35
        45
        70
        60
        55
        65
        63
        62
        80
    */
}
```

Here, The `PreOrderIterator` implementation provided earlier can be seen as a generator.

In this implementation:

- The `PreOrderIterator` struct encapsulates the state required for generating values (nodes of the tree) in preorder.
- The `HasNext` method checks if the iteration is possible by checking whether there still are nodes to traverse to.
- The `Next` method lazily generates the next value by traversing the tree in preorder fashion, popping nodes from the stack as needed and pushing their children onto the stack.
- Thus, the `Next` method behaves like a generator that produces values (nodes of the tree in preorder) on the fly as it's called, making it a lazy generator implementation.

This pattern of generating values on demand is characteristic of generators, and it's similar to how generators work in other languages where you can lazily generate values as needed.

---

### Using Idiomatic Go (Channel + Goroutine)

While the Iterator pattern can be implemented in Go, it is not the idiomatic way to perform iterations in Go. We can use channel with goroutines to make this Go-like. For instance, the `PostOrderIterator` can get quite complicated to implement. ([Check the implementation](../../data-structures/BinarySearchTree/BinarySearchTreeIterators.go))

Channels and Goroutines solve this:

```go
func (this *BinarySearchTree[T]) traversePostOrder(node *TreeNode[T], ch chan<- T) {
	if node == nil {
		return
	}
	this.traversePostOrder(node.Left, ch)
	this.traversePostOrder(node.Right, ch)
	ch <- node.Value
}

func (this *BinarySearchTree[T]) PostOrderTraversal() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		this.traversePostOrder(this.Root, ch)
	}()
	return ch
}

func main() {
	tree := bst.NewBinarySearchTree[int](50, 40, 70, 45, 35, 80, 60, 55, 65, 63, 62)
	generatorPostOrder := tree.PostOrderTraversal()
	result := make([]int, 0)
	for value := range generatorPostOrder {
		result = append(result, value)
	}
	fmt.Println("Post-order Traversal, using Channel + Goroutine:", result)
}
```

---

# Summary

- An Iterator specifies how you can traverse an object
- Moves along the iterated collection (using a method called `Next`), indicating when the last element has been reached (usually a method called `HasNext`)
- Not idiomatic in Go (no standard Iterable interface), however can be built. Channels with Goroutines are a viable alternative.
