package ds

import (
	"fmt"
)

type Queue[T Type] struct {
	Items []T
}

// Enqueue: Add to the start of the Queue.
func (this *Queue[T]) Enqueue(items ...T) int {
	// Reverse the order of the items received
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
	this.Items = append(items, this.Items...)
	return len(this.Items)
}

// Dequeue: Remove an element from the start of the Queue, and return it.
func (this *Queue[T]) Dequeue() T {
	item := this.Items[0]
	this.Items = this.Items[1:]
	return item
}

// Map: Modifies the elements of the stack with a mapping function that applies to each element. Returns a pointer to a new Slice.
func (this *Queue[T]) Map(callback func(index int, value T) T) *Queue[T] {
	var newQueue Queue[T]
	for index, value := range this.Items {
		newQueue.Items = append(newQueue.Items, callback(index, value))
	}
	return &newQueue
}

// Filter: Takes in a Filter function that evaluates whether true or false on each item of the stack. Returns a new Queue of filtered items.
func (this *Queue[T]) Filter(callback func(index int, value T) bool) *Queue[T] {
	var newQueue Queue[T]
	for index, value := range this.Items {
		if callback(index, value) {
			newQueue.Items = append(newQueue.Items, value)
		}
	}
	return &newQueue
}

// ForEach: Takes in a Callback function to run for each element in the Queue. Returns nil always.
func (this *Queue[T]) ForEach(callback func(index int, value T) error) *Queue[T] {
	for index, value := range this.Items {
		callback(index, value)
	}
	return nil
}

// Slice: Takes a start and end index and returns a pointer to a new slice
func (this *Queue[T]) Slice(start, end int) *Queue[T] {
	var sliced Queue[T]
	sliced.Items = append(sliced.Items, this.Items[start:end+1]...)
	return &sliced
}

// Splice: Mutates original Queue by removing elements starting at index. Takes three parameters total.
// Two parameters: `start` and `count` mutates the original Queue by removing items from a start index upto `count` number of items.
// The third is an optional variadic parameters that takes items of the same type as the Queue that it will be inserted at the index of removal.
// Returns the removed items as a new Queue
func (this *Queue[T]) Splice(start, count int, items ...T) *Queue[T] {
	var spliced Queue[T]
	if this.Items != nil {
		// Get the required Splice
		spliced = Queue[T]{
			Items: append(spliced.Items, this.Items[start:start+count]...),
		}
		// Reverse the items as they need to be inserted for the Queue
		for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
			items[i], items[j] = items[j], items[i]
		}
		this.Items = append(append(append([]T{}, this.Items[:start]...), items...), this.Items[start+count:]...)
	}

	return &spliced
}

// Get: Getter Function. Returns the value at Index
func (this *Queue[T]) Get(index int) T {
	return this.Items[index]
}

// Set: Setter Function. Takes in an index and a value.
func (this *Queue[T]) Set(index int, value T) {
	this.Items[index] = value
}

// Size: Return the length of the stack
func (this *Queue[T]) Size() int {
	return len(this.Items)
}

// IndexOf: Returns the index of the first matching item in the Queue. If not found returns -1
func (this *Queue[T]) IndexOf(item T) int {
	for index, value := range this.Items {
		if item == value {
			return index
		}
	}
	return -1
}

// LastIndexOf: Returns the index of the last matching item in the Queue. If not found returns -1
func (this *Queue[T]) LastIndexOf(item T) int {
	var lastIndex int
	for index, value := range this.Items {
		if item == value {
			lastIndex = index
		}
	}
	if this.Items[lastIndex] == item {
		return lastIndex
	}
	return -1
}

// Includes: Find whether the Queue contains an item. Returns true if found, false if not found
func (this *Queue[T]) Includes(item T) bool {
	for _, v := range this.Items {
		if item == v {
			return true
		}
	}
	return false
}

// Find: Takes in a callback function that acts as an evaluator for each element.
// Returns the item if the callback function evaluates to true, zero-value of the Type otherwise
func (this *Queue[T]) Find(callback func(index int, value T) bool) T {
	for index, value := range this.Items {
		if callback(index, value) {
			return value
		}
	}
	var zeroValue T
	return zeroValue
}

// FindIndex: Takes in a callback function that acts as an evaluator for each element.
// Returns the index (position) of the first element that evaluates to true. Returns -1 if no element passes evaluation.
func (this *Queue[T]) FindIndex(callback func(index int, value T) bool) int {
	for index, value := range this.Items {
		if callback(index, value) {
			return index
		}
	}
	return -1
}

// Some: Takes in a callback function that acts as an evaluator for each element. Returns a boolean.
// Returns `true` on encountering an element that the callback evaluates as true. Stops and exits method.
// Returns `false` if the callback returns false on all the elements.
func (this *Queue[T]) Some(callback func(index int, value T) bool) bool {
	for index, value := range this.Items {
		if callback(index, value) {
			return true
		}
	}
	return false
}

// Every: Takes in a callback function that acts as an evaluator for each element. Returns a boolean.
// Returns `true` if all the callback on all elements returns true.
// Returns `false` on encountering an element that the callback evaluates as false. Stops and exits method.
func (this *Queue[T]) Every(callback func(index int, value T) bool) bool {
	for index, value := range this.Items {
		if !callback(index, value) {
			return false
		}
	}
	return true
}

// Join: Takes in a Separator string, joins the elements of the stack into a single string and returns it
func (this *Queue[T]) Join(sep string) string {
	len := len(this.Items)
	joined := ""
	for i, val := range this.Items {
		joined += fmt.Sprintf("%v", val)
		if i < len-1 {
			joined += sep
		}
	}
	return joined
}

// String: Default formatting method.
func (this *Queue[T]) String() string {
	len := len(this.Items)
	out := ""
	for i := 0; i < len; i++ {
		out += fmt.Sprintf("%v", this.Items[i])
		if i < len-1 {
			out += ", "
		}
	}
	return out
}
