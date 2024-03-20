package main

import (
	"fmt"

	bst "github.com/jayantasamaddar/quick-reference-golang/data-structures/BinarySearchTree"
	ll "github.com/jayantasamaddar/quick-reference-golang/data-structures/LinkedList"
	queue "github.com/jayantasamaddar/quick-reference-golang/data-structures/Queue"
	stack "github.com/jayantasamaddar/quick-reference-golang/data-structures/Stack"
)

// Terminal Colours
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func PrintC(color string, s string) string {
	return color + s + Reset
}

func main() {
	/********************************************************************************************************************/
	// STACK
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: Stack" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	// Initialize Stack with default items
	s := &stack.Stack[int]{Items: []int{100, 200}}

	// Print items of the stack
	fmt.Println(PrintC(Green, fmt.Sprintf("INITIALIZE STACK (type: %T): ", s.Items)), PrintC(Yellow, s.String()))

	// Push: Add items onto the end of the stack
	s.Push(300, 400)
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER PUSH OPERATION (type: %T): ", s.Items)), PrintC(Yellow, s.String()))

	// Pop: Remove the last item from the stack
	item := s.Pop()
	fmt.Println(PrintC(Green, fmt.Sprintf("POP OPERATION (ITEM) (type: %T): ", item)), PrintC(Yellow, fmt.Sprintf("%d", item)))
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER POP OPERATION (type: %T): ", s.Items)), PrintC(Yellow, s.String()))

	// Map: Modify elements of the stack and return pointer to a new stack
	fmt.Println()
	newStack := s.Map(func(i, v int) int {
		return v * (i + 1)
	})
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER MAP `v * (i + 1)` OPERATION (type: %T): ", newStack.Items)), PrintC(Yellow, s.String()))
	fmt.Println(PrintC(Green, fmt.Sprintf("SHOW ORIGINAL UNAFFECTED AFTER MAP OPERATION (type: %T): ", s.Items)), PrintC(Yellow, s.String()))

	// Filter: Filter elements of the stack onto a new stack
	fmt.Println()
	newStack = s.Filter(func(_, val int) bool {
		return val > 100
	})
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER FILTER `v > 100` OPERATION (type: %T): ", newStack.Items)), PrintC(Yellow, newStack.String()))
	fmt.Println(PrintC(Green, fmt.Sprintf("SHOW ORIGINAL UNAFFECTED AFTER FILTER OPERATION (type: %T): ", s.Items)), PrintC(Yellow, s.String()))

	// ForEach: Run a callback against each element of the stack
	fmt.Println()
	fmt.Println(PrintC(Green, "RUNNING FOR EACH ON:"), PrintC(Yellow, s.String()))
	s.ForEach(func(i, v int) error {
		fmt.Println(PrintC(Blue, fmt.Sprintf("The value at index %d, is: %s%v%s", i, Yellow, v, Yellow)))
		return nil
	})

	// Get: Getter function to get the value at index
	fmt.Println()
	fmt.Printf(PrintC(Green, "GETTER: VALUE AT INDEX `0` for %v (type: %T): "), newStack, newStack.Get(0))
	fmt.Printf(PrintC(Yellow, "%v\n"), newStack.Get(0)) // 200

	// Set: Setter function to get the value at index
	fmt.Println()
	fmt.Printf(PrintC(Green, "SETTER: SET `201` AT INDEX `0` for %v: "), PrintC(Yellow, newStack.String()))
	newStack.Set(0, 201)
	fmt.Printf(PrintC(Green, "SHOW MODIFIED ITEMS (type: %T): "), newStack.Items)
	fmt.Println(PrintC(Yellow, fmt.Sprintf("%v", newStack)))

	// Join: Joins the elements of the stack into a single string using a separator
	fmt.Println()
	fmt.Println(PrintC(Green, "RUNNING JOIN:"), PrintC(Yellow, s.Join(" | --- | ")))

	// Slice: Returns part of the original Stack as a new Stack
	fmt.Println()
	toSplice := &stack.Stack[int]{}
	toSplice.Push(0, 1, 2, 3, 4, 5)
	fmt.Println(PrintC(Green, fmt.Sprintf("SLICED: Index 3 and 4 for %v: ", toSplice.Items)), PrintC(Yellow, toSplice.Slice(3, 4).String()))
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER SLICE OPERATION (type: %T): %v", toSplice.Items, PrintC(Yellow, toSplice.String()))))

	// Splice: Mutates original stack by removing items
	fmt.Println()
	// Splice: (a) Removal
	fmt.Println(PrintC(Green, fmt.Sprintf("SPLICE: Remove Index 3 and 4 for %v: ", toSplice.Items)), PrintC(Yellow, toSplice.Splice(3, 2).String()))
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER SPLICE OPERATION (type: %T): %v", toSplice.Items, PrintC(Yellow, toSplice.String()))))
	// Splice: (b) Removal and insertion
	fmt.Println(PrintC(Green, fmt.Sprintf("SPLICE: Insert 3, 4 at index 3 %v: ", toSplice.Items)), PrintC(Yellow, toSplice.Splice(3, 0, 3, 4).String()))
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER SPLICE OPERATION (type: %T): %v", toSplice.Items, PrintC(Yellow, toSplice.String()))))

	/********************************************************************************************************************/
	// QUEUE
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: Queue" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	// Initialize Stack with default items
	queue := &queue.Queue[int]{Items: []int{100, 200}}

	// Print items of the queue
	fmt.Println(PrintC(Green, fmt.Sprintf("INITIALIZE QUEUE (type: %T): %v", queue.Items, queue.Items)))

	// Enqueue: Add items onto the start of the queue
	queue.Enqueue(300, 400)
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER ENQUEUE OPERATION (type: %T): %v", queue.Items, queue.Items)))

	// Dequeue: Remove the first item from the queue, stores it in a variable and returns its pointer
	item = queue.Dequeue()

	fmt.Println(PrintC(Green, fmt.Sprintf("DEQUEUE OPERATION (ITEM) (type: %T): %v", item, item)))
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER DEQUEUE OPERATION (type: %T): %v", queue.Items, queue.Items)))

	/********************************************************************************************************************/
	// LINKED LIST
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: Linked List" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	// Initialize LinkedList
	ll := &ll.LinkedList[int]{}

	// Append
	ll.Append(1, 2, 3, 4, 5, 6)
	fmt.Println(PrintC(Green, fmt.Sprintf("AFTER APPEND LINKEDLIST (type: %T):", ll)), PrintC(Yellow, ll.String()))

	// Prepend
	ll.Prepend(0, -1, -2, -3, -4, -5)
	fmt.Println(PrintC(Green, fmt.Sprintf("AFTER PREPEND LINKEDLIST (type: %T):", ll)), PrintC(Yellow, ll.String()))

	// RemoveFirstOccurence
	fmt.Println()
	fmt.Println(PrintC(Green, "LINKEDLIST SIZE:"), PrintC(Yellow, fmt.Sprintf("%d", ll.Size)))
	fmt.Println(PrintC(Green, "REMOVE FIRST OCCURENCE (-5):"), PrintC(Blue, fmt.Sprintf("%t", ll.RemoveFirstOccurence(-5))))
	fmt.Println(PrintC(Green, "CURRENT LINKEDLIST:"), PrintC(Yellow, ll.String()))
	fmt.Println(PrintC(Green, "LINKEDLIST SIZE:"), PrintC(Yellow, fmt.Sprintf("%d", ll.Size)))
	fmt.Println(PrintC(Green, "REMOVE FIRST OCCURENCE (0):"), PrintC(Blue, fmt.Sprintf("%t", ll.RemoveFirstOccurence(0))))
	fmt.Println(PrintC(Green, "CURRENT LINKEDLIST:"), PrintC(Yellow, ll.String()))
	fmt.Println(PrintC(Green, "LINKEDLIST SIZE:"), PrintC(Yellow, fmt.Sprintf("%d", ll.Size)))
	fmt.Println(PrintC(Green, "ERROR REMOVE FIRST OCCURENCE (100):"), PrintC(Red, fmt.Sprintf("%t", ll.RemoveFirstOccurence(100))))

	// RemoveLastOccurence
	fmt.Println()
	ll.Append(4)
	fmt.Println(PrintC(Green, "LINKEDLIST SIZE:"), PrintC(Yellow, fmt.Sprintf("%d", ll.Size)))
	fmt.Println(PrintC(Green, "REMOVE LAST OCCURENCE (-4):"), PrintC(Blue, fmt.Sprintf("%t", ll.RemoveLastOccurence(-4))))
	fmt.Println(PrintC(Green, "CURRENT LINKEDLIST:"), PrintC(Yellow, ll.String()))
	fmt.Println(PrintC(Green, "LINKEDLIST SIZE:"), PrintC(Yellow, fmt.Sprintf("%d", ll.Size)))
	fmt.Println(PrintC(Green, "REMOVE LAST OCCURENCE (4):"), PrintC(Blue, fmt.Sprintf("%t", ll.RemoveLastOccurence(4))))
	fmt.Println(PrintC(Green, "CURRENT LINKEDLIST:"), PrintC(Yellow, ll.String()))
	fmt.Println(PrintC(Green, "LINKEDLIST SIZE:"), PrintC(Yellow, fmt.Sprintf("%d", ll.Size)))
	fmt.Println(PrintC(Green, "ERROR REMOVE LAST OCCURENCE (100):"), PrintC(Red, fmt.Sprintf("%t", ll.RemoveLastOccurence(100))))

	// RemoveFirst
	fmt.Println()
	first, _ := ll.RemoveFirst()
	fmt.Println(PrintC(Green, "REMOVE FIRST (expecting -3):"), PrintC(Yellow, fmt.Sprintf("%d", first)))
	fmt.Println(PrintC(Green, "AFTER REMOVE FIRST:"), PrintC(Yellow, ll.String()))
	fmt.Println(PrintC(Green, "LINKEDLIST SIZE:"), PrintC(Yellow, fmt.Sprintf("%d", ll.Size)))

	// RemoveLast
	fmt.Println()
	last, _ := ll.RemoveLast()
	fmt.Println(PrintC(Green, "REMOVE LAST (expecting 6):"), PrintC(Yellow, fmt.Sprintf("%d", last)))
	fmt.Println(PrintC(Green, "AFTER REMOVE LAST:"), PrintC(Yellow, ll.String()))
	fmt.Println(PrintC(Green, "LINKEDLIST SIZE:"), PrintC(Yellow, fmt.Sprintf("%d", ll.Size)))

	/********************************************************************************************************************/
	// BINARY SEARCH TREE
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: Binary Search Tree" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	/*
		Tree to be created:

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
	tree := bst.NewBinarySearchTree[int]()

	// Add elements to the Binary Tree
	tree.Add(50, 40, 70, 45, 35, 80, 60, 55, 65, 63, 62)

	fmt.Println(tree.JSON())
	// {"Root":{"Value":50,"Left":{"Value":40,"Left":{"Value":35,"Left":null,"Right":null},"Right":{"Value":45,"Left":null,"Right":null}},"Right":{"Value":70,"Left":{"Value":60,"Left":{"Value":55,"Left":null,"Right":null},"Right":{"Value":65,"Left":{"Value":63,"Left":{"Value":62,"Left":null,"Right":null},"Right":null},"Right":null}},"Right":{"Value":80,"Left":null,"Right":null}}}}
	fmt.Println(tree.JSONPretty())
	/*
		{
		  "Root": {
		    "Value": 50,
		    "Left": {
		      "Value": 40,
		      "Left": {
		        "Value": 35,
		        "Left": null,
		        "Right": null
		      },
		      "Right": {
		        "Value": 45,
		        "Left": null,
		        "Right": null
		      }
		    },
		    "Right": {
		      "Value": 70,
		      "Left": {
		        "Value": 60,
		        "Left": {
		          "Value": 55,
		          "Left": null,
		          "Right": null
		        },
		        "Right": {
		          "Value": 65,
		          "Left": {
		            "Value": 63,
		            "Left": {
		              "Value": 62,
		              "Left": null,
		              "Right": null
		            },
		            "Right": null
		          },
		          "Right": null
		        }
		      },
		      "Right": {
		        "Value": 80,
		        "Left": null,
		        "Right": null
		      }
		    }
		  }
		}
	*/

	// Remove
	tree.Remove(70)
	fmt.Println(tree)
	/*
		├── 50
		│   ├── 40
		│   │   ├── 35
		│   │   ├── 45
		│   ├── 65
		│   │   ├── 60
		│   │   │   ├── 55
		│   │   │   ├── 63
		│   │   │   │   ├── 62
		│   │   ├── 80
	*/

	// Min
	fmt.Println("Min:", tree.Min()) // Min: 35

	// Max
	fmt.Println("Max:", tree.Max()) // Max: 80

	// Size
	fmt.Println("Size:", tree.Size()) // Size: 10

	// MinHeight
	fmt.Println("MinHeight:", tree.MinHeight()) // MinHeight: 2

	// MaxHeight
	fmt.Println("MaxHeight:", tree.MaxHeight()) // MaxHeight: 4

	// IsBalanced
	fmt.Println("Is Balanced:", tree.IsBalanced()) // Is Balanced: false
	tree.Remove(63, 62)                            // Remove nodes to balance tree
	fmt.Println("Is Balanced:", tree.IsBalanced()) // Is Balanced: true
	tree.Add(63, 62)                               // Add them back

	// Contains
	fmt.Println("Contains:", tree.Contains(63))   // Contains: true
	fmt.Println("Contains:", tree.Contains(1000)) // Contains: false

	// Breadth-first Search (BFS): Level Order
	fmt.Println("Level Order:", tree.LevelOrder()) // Level Order: [50 40 70 35 45 60 80 55 65 63 62]

	// Depth-first Search (DFS): In-Order
	fmt.Println("In Order (Ascending):", tree.InOrder(""))      // In Order (Ascending): [35 40 45 50 55 60 62 63 65 70 80]
	fmt.Println("In Order (Descending):", tree.InOrder("desc")) // In Order (Descending): [80 70 65 63 62 60 55 50 45 40 35]

	// Depth-first Search (DFS): Pre-Order and Post-Order
	fmt.Println("In Order (Pre-Order):", tree.PreOrder())   // In Order (Pre-Order): [ 50 40 35 45 70 60 55 65 63 62 80 ]
	fmt.Println("In Order (Post-Order):", tree.PostOrder()) // In Order (Post-Order): [ 35 45 40 55 62 63 65 60 80 70 50 ]

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

	// Level-order Traversal: Idiomatic Go (channel + goroutine) Implementation
	generatorLevelOrder := tree.LevelOrderTraversal()
	result = make([]int, 0)
	for value := range generatorLevelOrder {
		result = append(result, value)
	}
	fmt.Println("Level-order Traversal, using Channel + Goroutine:", result)

	// IsEmpty
	fmt.Println("IsEmpty:", tree.IsEmpty()) // IsEmpty: false
	tree.Clear()                            // Clears all nodes
	fmt.Println("IsEmpty:", tree.IsEmpty()) // IsEmpty: true
}
