package ds

import (
	"fmt"
)

type Stack[T Type] struct {
	Items []T
}

// Push: Push items to the end of the stack
func (this *Stack[T]) Push(items ...T) int {
	this.Items = append(this.Items, items...)
	return len(this.Items)
}

// Pop: Remove an element from the end of the stack, and return it.
func (this *Stack[T]) Pop() T {
	l := len(this.Items) - 1
	item := this.Items[l]
	this.Items = this.Items[:l]
	return item
}

// Map: Modifies the elements of the stack with a mapping function that applies to each element. Returns a pointer to a new Slice.
func (this *Stack[T]) Map(callback func(index int, value T) T) *Stack[T] {
	var newStack Stack[T]
	for index, value := range this.Items {
		newStack.Items = append(newStack.Items, callback(index, value))
	}
	return &newStack
}

// Filter: Takes in a Filter function that evaluates whether true or false on each item of the stack. Returns a new Stack of filtered items.
func (this *Stack[T]) Filter(callback func(index int, value T) bool) *Stack[T] {
	var newStack Stack[T]
	for index, value := range this.Items {
		if callback(index, value) {
			newStack.Items = append(newStack.Items, value)
		}
	}
	return &newStack
}

// ForEach: Takes in a Callback function to run for each element in the Stack. Returns nil always.
func (this *Stack[T]) ForEach(callback func(index int, value T) error) *Stack[T] {
	for index, value := range this.Items {
		callback(index, value)
	}
	return nil
}

// Slice: Takes a start and end index and returns a pointer to a new slice
func (this *Stack[T]) Slice(start, end int) *Stack[T] {
	var sliced Stack[T]
	sliced.Items = append(sliced.Items, this.Items[start:end+1]...)
	return &sliced
}

// Splice: Mutates original Stack by removing elements starting at index. Takes three parameters total.
// Two parameters: `start` and `count` mutates the original Stack by removing items from a start index upto `count` number of items.
// The third is an optional variadic parameters that takes items of the same type as the Stack that it will be inserted at the index of removal.
// Returns the removed items as a new Stack
func (this *Stack[T]) Splice(start, count int, items ...T) *Stack[T] {
	var spliced Stack[T]
	if this.Items != nil {
		// Get the required Splice
		spliced = Stack[T]{
			Items: append(spliced.Items, this.Items[start:start+count]...),
		}
		this.Items = append(append(append([]T{}, this.Items[:start]...), items...), this.Items[start+count:]...)
	}

	return &spliced
}

// Get: Getter Function. Returns the value at Index
func (this *Stack[T]) Get(index int) T {
	return this.Items[index]
}

// Set: Setter Function. Takes in an index and a value.
func (this *Stack[T]) Set(index int, value T) {
	this.Items[index] = value
}

// Size: Return the length of the stack
func (this *Stack[T]) Size() int {
	return len(this.Items)
}

// IndexOf: Returns the index of the first matching item in the Stack. If not found returns -1
func (this *Stack[T]) IndexOf(item T) int {
	for index, value := range this.Items {
		if item == value {
			return index
		}
	}
	return -1
}

// LastIndexOf: Returns the index of the last matching item in the Stack. If not found returns -1
func (this *Stack[T]) LastIndexOf(item T) int {
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

// Includes: Find whether the Stack contains an item. Returns true if found, false if not found
func (this *Stack[T]) Includes(item T) bool {
	for _, v := range this.Items {
		if item == v {
			return true
		}
	}
	return false
}

// Find: Takes in a callback function that acts as an evaluator for each element.
// Returns the item if the callback function evaluates to true, zero-value of the Type otherwise
func (this *Stack[T]) Find(callback func(index int, value T) bool) T {
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
func (this *Stack[T]) FindIndex(callback func(index int, value T) bool) int {
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
func (this *Stack[T]) Some(callback func(index int, value T) bool) bool {
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
func (this *Stack[T]) Every(callback func(index int, value T) bool) bool {
	for index, value := range this.Items {
		if !callback(index, value) {
			return false
		}
	}
	return true
}

// Join: Takes in a Separator string, joins the elements of the stack into a single string and returns it
func (this *Stack[T]) Join(sep string) string {
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

// String: Default formatting method
func (this *Stack[T]) String() string {
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
