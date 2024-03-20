# Table of Contents

- [Table of Contents](#table-of-contents)
- [Visitor Pattern](#visitor-pattern)
  - [Introduction](#introduction)
  - [Use Cases](#use-cases)
  - [Implementation in Go](#implementation-in-go)
    - [Basic Implementation](#basic-implementation)
    - [Intrusive Visitor](#intrusive-visitor)
- [Summary](#summary)

---

# Visitor Pattern

## Introduction

The Visitor pattern is a behavioural design pattern where a component (visitor) is allowed to traverse the entire hierarchy of types. Implemented by propagating a single `Accept()` method throughout the entire hierarchy. This approach is often used for traversal as an alternative to the Iterator pattern.

<!-- The Visitor design pattern is a behavioral design pattern that allows adding new behaviors to existing classes without modifying their structure. It separates the algorithm from the object structure on which it operates. -->

---

## Use Cases

The Visitor pattern is particularly useful in scenarios where:

1. **Adding new operations to classes without modifying them**: When you have a stable class hierarchy but need to add new operations to these classes frequently, the Visitor pattern allows you to encapsulate these operations in separate Visitor classes, thus avoiding modification of existing classes.

2. **Processing complex object structures**: If you have a complex object structure composed of different types of objects and need to perform various operations on them, the Visitor pattern helps to separate the algorithm from the object structure, making it easier to manage and extend.

3. **Type-specific operations**: When you have operations that are specific to certain types within a class hierarchy, the Visitor pattern allows you to define these operations in separate Visitor implementations for each type.

Popular libraries and frameworks that use the Visitor pattern include:

1. **Java's Abstract Syntax Tree (AST) processing libraries**: Many AST processing libraries in Java, such as JavaParser or ANTLR, utilize the Visitor pattern to traverse and manipulate the AST nodes.

2. **Apache Commons Collections**: The Apache Commons Collections library provides visitors for various collection types, allowing users to perform operations on elements of collections without exposing the internal structure of the collections.

3. **Gson library for Java (Google's JSON library)**: Gson provides a `JsonElementVisitor` interface that allows users to define custom visitors to traverse and process JSON elements.

4. **Hibernate ORM**: Hibernate, a popular ORM framework for Java, uses the Visitor pattern in its Criteria API for constructing queries. Visitors are used to represent query criteria, which are then translated into SQL queries by the framework.

These are just a few examples, but the Visitor pattern is a common pattern in software development and can be found in many libraries and frameworks across different languages and domains.

---

## Implementation in Go

### Basic Implementation

The Visitor design pattern is a behavioral design pattern that allows adding new behaviors to existing classes without modifying their structure. It separates the algorithm from the object structure on which it operates.

Here's a simple implementation of the Visitor pattern in Go:

```go
package main

import "fmt"

// Visitable interface defines the accept method that takes a Visitor as an argument.
type Visitable interface {
	Accept(Visitor)
}

// Visitor interface declares a visit method for each concrete type.
type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

// ConcreteVisitor1 implements the Visitor interface and provides concrete implementation for Visit methods.
type ConcreteVisitor1 struct{}

func (c *ConcreteVisitor1) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor1 visited ConcreteElementA")
}

func (c *ConcreteVisitor1) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor1 visited ConcreteElementB")
}

// ConcreteVisitor2 implements the Visitor interface and provides concrete implementation for Visit methods.
type ConcreteVisitor2 struct{}

func (c *ConcreteVisitor2) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor2 visited ConcreteElementA")
}

func (c *ConcreteVisitor2) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor2 visited ConcreteElementB")
}

// ConcreteElementA and ConcreteElementB implement the Visitable interface.
type ConcreteElementA struct{}

func (c *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(c)
}

type ConcreteElementB struct{}

func (c *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(c)
}

// ObjectStructure maintains a collection of Visitable elements.
type ObjectStructure struct {
	elements []Visitable
}

func (o *ObjectStructure) Attach(element Visitable) {
	o.elements = append(o.elements, element)
}

func (o *ObjectStructure) Detach(element Visitable) {
	for i, e := range o.elements {
		if e == element {
			o.elements = append(o.elements[:i], o.elements[i+1:]...)
			break
		}
	}
}

func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range o.elements {
		element.Accept(visitor)
	}
}

func main() {
	objectStructure := &ObjectStructure{}
	objectStructure.Attach(&ConcreteElementA{})
	objectStructure.Attach(&ConcreteElementB{})

	visitor1 := &ConcreteVisitor1{}
	objectStructure.Accept(visitor1)

	visitor2 := &ConcreteVisitor2{}
	objectStructure.Accept(visitor2)
}
```

In this example, we have:

- `Visitor` interface: Declares visit methods for each concrete element.
- `ConcreteVisitor1` and `ConcreteVisitor2`: Implement the `Visitor` interface, providing concrete implementations for visit methods.
- `Visitable` interface: Declares the `Accept` method that takes a `Visitor` as an argument.
- `ConcreteElementA` and `ConcreteElementB`: Implement the `Visitable` interface.
- `ObjectStructure`: Maintains a collection of `Visitable` elements and provides methods to attach, detach, and accept visitors.
- In the `main` function, we demonstrate how the visitor pattern works by attaching elements to the object structure and accepting different visitors.

---

### Intrusive Visitor

---

# Summary
