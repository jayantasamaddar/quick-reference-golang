package main

import (
	"fmt"

	bst "github.com/jayantasamaddar/quick-reference-golang/data-structures/BinarySearchTree"
	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
)

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
func (i *MyCollectionIterator) Next() interface{} {
	if i.HasNext() {
		value := (*i.collection.data)[i.index]
		i.index++
		return value
	}
	return nil
}

func main() {
	/***********************************************************************************************/
	// (1) Default Iteration in Go using `range`: Works on arrays, slices, strings, maps, channels
	/***********************************************************************************************/
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	fmt.Println(std.Purple + "\t(1) Default Iteration in Go using `range`: Works on arrays, slices, strings, maps, channel" + std.Reset)
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)

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

	/***********************************************************************************************/
	// (2) Idiomatic iteration in Go using Goroutine and channels to create a Generator
	/***********************************************************************************************/

	/***********************************************************************************************/
	// (3) Iterator Pattern implementation in Go
	/***********************************************************************************************/

	collection := NewMyCollection(&[]int{1, 2, 3, 4, 5})
	iterator := collection.Iterator()

	fmt.Println("Using range with the collection:")
	for iterator.HasNext() {
		element := iterator.Next()
		fmt.Println(element)
	}

	/********************************************************************************/
	// (4) Tree Traversal
	/********************************************************************************/
	// tree := ds.NewBinarySearchTree[int](50, 40, 45, 35, 70, 60, 80, 100)
	tree := bst.NewBinarySearchTree[int](50, 40, 70, 45, 35, 80, 60, 55, 65, 63, 62)

	fmt.Println(tree)
	/*
		├── 50
		│   ├── 40
		│   │   ├── 35
		│   │   ├── 45
		│   ├── 70
		│   │   ├── 60
		│   │   │   ├── 55
		│   │   │   ├── 65
		│   │   │   │   ├── 63
		│   │   │   │   │   ├── 62
		│   │   ├── 80
	*/
	/*
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
	*/

	/********************************************************************************/
	// (4a) Tree Traversal: Using Traditional Approach
	/********************************************************************************/

	// Breadth-first Search (BFS): Level Order
	fmt.Println("Level Order:", tree.LevelOrder()) // Level Order: [50 40 70 35 45 60 80 55 65 63 62]

	// Depth-first Search (DFS): In-Order
	fmt.Println("In Order (Ascending):", tree.InOrder(""))      // In Order (Ascending): [35 40 45 50 55 60 62 63 65 70 80]
	fmt.Println("In Order (Descending):", tree.InOrder("desc")) // In Order (Descending): [80 70 65 63 62 60 55 50 45 40 35]

	// Depth-first Search (DFS): Pre-Order and Post-Order
	fmt.Println("In Order (Pre-Order):", tree.PreOrder())   // In Order (Pre-Order): [ 50 40 35 45 70 60 55 65 63 62 80 ]
	fmt.Println("In Order (Post-Order):", tree.PostOrder()) // In Order (Post-Order): [ 35 45 40 55 62 63 65 60 80 70 50 ]

	/********************************************************************************/
	// (4b) Tree Traversal: Using Iterator Pattern
	/********************************************************************************/
	// Pre-Order Traversal: Using an Iterator
	iter := tree.PreOrderIterator()
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}

	iter.Reset()
	fmt.Println("*************")

	for iter.HasNext() {
		fmt.Println(iter.Next())
	}

	/********************************************************************************/
	// (4c) Tree Traversal: Using Idiomatic Go
	/********************************************************************************/
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

	// Pre-order Traversal: Idiomatic Go (channel + goroutine) Implementation
	generatorPreOrder := tree.PreOrderTraversal()
	result = make([]int, 0)
	for value := range generatorPreOrder {
		result = append(result, value)
	}
	fmt.Println("Pre-order Traversal, using Channel + Goroutine:", result)

	// Post-order Traversal: Idiomatic Go (channel + goroutine) Implementation
	generatorPostOrder := tree.PostOrderTraversal()
	result = make([]int, 0)
	for value := range generatorPostOrder {
		result = append(result, value)
	}
	fmt.Println("Post-order Traversal, using Channel + Goroutine:", result)
}
