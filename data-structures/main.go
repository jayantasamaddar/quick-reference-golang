package main

import (
	"fmt"

	"github.com/jayantasamaddar/quick-reference-golang/data-structures/ds"
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
	stack := &ds.Stack[int]{Items: []int{100, 200}}

	// Print items of the stack
	fmt.Println(PrintC(Green, fmt.Sprintf("INITIALIZE STACK (type: %T): ", stack.Items)), PrintC(Yellow, stack.String()))

	// Push: Add items onto the end of the stack
	stack.Push(300, 400)
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER PUSH OPERATION (type: %T): ", stack.Items)), PrintC(Yellow, stack.String()))

	// Pop: Remove the last item from the stack
	item := stack.Pop()
	fmt.Println(PrintC(Green, fmt.Sprintf("POP OPERATION (ITEM) (type: %T): ", item)), PrintC(Yellow, fmt.Sprintf("%d", item)))
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER POP OPERATION (type: %T): ", stack.Items)), PrintC(Yellow, stack.String()))

	// Map: Modify elements of the stack and return pointer to a new stack
	fmt.Println()
	newStack := stack.Map(func(i, v int) int {
		return v * (i + 1)
	})
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER MAP `v * (i + 1)` OPERATION (type: %T): ", newStack.Items)), PrintC(Yellow, newStack.String()))
	fmt.Println(PrintC(Green, fmt.Sprintf("SHOW ORIGINAL UNAFFECTED AFTER MAP OPERATION (type: %T): ", stack.Items)), PrintC(Yellow, stack.String()))

	// Filter: Filter elements of the stack onto a new stack
	fmt.Println()
	newStack = stack.Filter(func(_, val int) bool {
		return val > 100
	})
	fmt.Println(PrintC(Green, fmt.Sprintf("ITEMS AFTER FILTER `v > 100` OPERATION (type: %T): ", newStack.Items)), PrintC(Yellow, newStack.String()))
	fmt.Println(PrintC(Green, fmt.Sprintf("SHOW ORIGINAL UNAFFECTED AFTER FILTER OPERATION (type: %T): ", stack.Items)), PrintC(Yellow, stack.String()))

	// ForEach: Run a callback against each element of the stack
	fmt.Println()
	fmt.Println(PrintC(Green, "RUNNING FOR EACH ON:"), PrintC(Yellow, stack.String()))
	stack.ForEach(func(i, v int) error {
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
	fmt.Println(PrintC(Green, "RUNNING JOIN:"), PrintC(Yellow, stack.Join(" | --- | ")))

	// Slice: Returns part of the original Stack as a new Stack
	fmt.Println()
	toSplice := &ds.Stack[int]{}
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
	queue := &ds.Queue[int]{Items: []int{100, 200}}

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
	ll := &ds.LinkedList[int]{}

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

	bst := &ds.BinarySearchTree[int]{}

	// Add elements to the Binary Tree
	bst.Add(50, 30, 15, 80, 110, 75, 25)
}
