# Table of Contents

- [Table of Contents](#table-of-contents)
- [Prototype Pattern](#prototype-pattern)
  - [Introduction](#introduction)
  - [Deep Copying](#deep-copying)
    - [Reflection-based DeepCopy](#reflection-based-deepcopy)
    - [Serialization-Based DeepCopy](#serialization-based-deepcopy)
- [Summary](#summary)

---

# Prototype Pattern

## Introduction

Complicated objects are rarely created from scratch. For example an iPhone 15 maybe created out of an existing iPhone base struct. There has to be an existing design or prototype which can then be copied and customized. This requires "deep copy" support, which means, if a struct is copied, what happens to the pointers, because then you have separate pointers, pointing to the same object. Deep copy means creating a completely separate set of pointers with new references for the new struct.

Thus, a Prototype is a partially or fully initialized object that you copy (clone) and make use of.

---

## Deep Copying

In addition to structs, slices, arrays, and maps, you might also encounter other types that need to be deep-copied depending on your specific use case. Here are a few more types to consider:

1. **Channels**: Channels are not directly copyable. If your struct contains channels or if you have a map with channels as values, you might need to handle channel copying separately.

2. **Pointers**: If your struct has fields that are pointers, you need to decide whether you want to create new instances for the pointed values or keep references to the same objects. The deep copy should be adjusted accordingly.

3. **Interfaces**: If your struct has fields of interface type, you need to handle the underlying types correctly. Deep copying interfaces might involve creating new instances of the underlying types.

4. **Custom Types**: If you have custom types that require special handling during copying, you should consider those as well.

---

### Reflection-based DeepCopy

**Advantages**:

- **Flexibility**: The reflection-based approach allows you to handle a wide variety of types, including structs, slices, arrays, pointers, maps, etc., without needing specific serialization logic for each type.

**Considerations**:

- **Performance**: Reflection can be slower compared to serialization, especially for large data structures, due to the runtime type checking and the overhead associated with it.
- **Complexity**: Handling different types using reflection can make the code more complex and harder to understand.

Here's how it can be done:

In **`lib/ReflectionBasedDeepCopy.go`**:

```go
package lib

import "reflect"

type DeepCopyable interface {
	DeepCopy() DeepCopyable
}

// DeepCopy recursively copies the given struct, slices, arrays, and maps
func DeepCopy(src interface{}) interface{} {
	srcValue := reflect.ValueOf(src)

	switch srcValue.Kind() {
	case reflect.Slice, reflect.Array:
		// Handle slices and arrays
		dst := reflect.MakeSlice(srcValue.Type(), srcValue.Len(), srcValue.Len())
		for i := 0; i < srcValue.Len(); i++ {
			dst.Index(i).Set(reflect.ValueOf(DeepCopy(srcValue.Index(i).Interface())))
		}
		return dst.Interface()

	case reflect.Map:
		// Handle maps
		dst := reflect.MakeMap(srcValue.Type())
		for _, key := range srcValue.MapKeys() {
			dstKey := reflect.ValueOf(DeepCopy(key.Interface()))
			dstValue := reflect.ValueOf(DeepCopy(srcValue.MapIndex(key).Interface()))
			dst.SetMapIndex(dstKey, dstValue)
		}
		return dst.Interface()

	case reflect.Ptr:
		// Handle pointers
		if srcValue.IsNil() {
			return nil
		}
		return reflect.New(srcValue.Type().Elem()).Elem().Interface()

	case reflect.Struct:
		// Handle structs
		dst := reflect.New(srcValue.Type()).Elem()
		for i := 0; i < srcValue.NumField(); i++ {
			srcField := srcValue.Field(i)
			dstField := dst.Field(i)

			if srcField.Kind() == reflect.Struct || srcField.Kind() == reflect.Slice || srcField.Kind() == reflect.Array || srcField.Kind() == reflect.Map {
				dstField.Set(reflect.ValueOf(DeepCopy(srcField.Interface())))
			} else {
				dstField.Set(srcField)
			}
		}
		return dst.Interface()

	default:
		// For other types, return the original value
		return src
	}
}
```

In `main.go`:

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/08-prototype-pattern/lib"
)

// Person struct with nested Address struct
type Person struct {
	Name    string
	Age     int
	Address Address
	Hobbies []string
}

// Address struct
type Address struct {
	City  string
	State string
}

// Attach
func (p Person) DeepCopy() lib.DeepCopyable {
	return lib.DeepCopy(p).(lib.DeepCopyable)
}

/*********************************************************************************************************/
// Main Function
/*********************************************************************************************************/
func main() {
	/*********************************************************************************************************/
	// (1) Reflection-Based Deep Copying
	/*********************************************************************************************************/
	// Creating a sample Person instance
	person1 := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			City:  "New York",
			State: "NY",
		},
		Hobbies: []string{"Reading", "Traveling"},
	}

	// Creating a deep copy of the Person instance (using standalone function)
	// person2 := lib.DeepCopy(person1).(Person)

	// Creating a deep copy of the Person instance (using struct method that implements the standalone function)
	person2 := person1.DeepCopy().(Person)

	// Modifying the original instance
	person2.Name = "Jane Doe"
	person2.Age = 25
	person2.Address.City = "San Francisco"
	person2.Hobbies = []string{"Dancing", "Gardening"}

	// Displaying the results
	fmt.Printf("Original Type (%s): %v\n", reflect.TypeOf(person1), person1)     // Original Type (main.Person): {John Doe 30 {New York NY} [Reading Traveling]}
	fmt.Printf("Deep Copy of Type (%s): %v\n", reflect.TypeOf(person1), person2) // Deep Copy of Type (main.Person): {Jane Doe 25 {San Francisco NY} [Dancing Gardening]}
}
```

---

### Serialization-Based DeepCopy

---

# Summary

- To implement a Prototype, partially construct an object and store it somewhere.
- Deep copy the prototype.
- Customize the resulting instance.
- A prototype factory provides a convenient API for using Prototypes.

---
