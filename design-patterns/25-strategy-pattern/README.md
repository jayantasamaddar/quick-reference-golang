# Table of Contents

- [Table of Contents](#table-of-contents)
- [Strategy Pattern](#strategy-pattern)
  - [Introduction](#introduction)
  - [Use Cases](#use-cases)
  - [Implementation in Go](#implementation-in-go)
    - [Basic Implementation](#basic-implementation)
    - [Changing Text Formats](#changing-text-formats)
- [Summary](#summary)

---

# Strategy Pattern

## Introduction

The Strategy design pattern is a behavioral design pattern that allows you to define a family of algorithms, encapsulate each one of them, and make them interchangeable. The Strategy pattern lets the algorithm vary independently from the clients that use it.

- Many algorithms can be decomposed into higher and lower level parts.
- Example: Making Tea can be decomposed into:
  - The process of making a hot beverage
    - boil water
    - pour it into a cup
  - Tea specific things:
    - Put teabag in water
- The higher level algorithm can then be reused for making coffee or hot chocolate
  - Supported by beverage-specific strategies

The idea is to separate an algorithm into its skeleton and concrete implementation steps, which can be varied at runtime.

---

## Use Cases

The Strategy pattern is commonly used in scenarios where you need to implement different algorithms or behaviors that can be interchangeable at runtime. Some practical use cases of the Strategy pattern include:

1. Sorting algorithms: You can use the Strategy pattern to encapsulate different sorting algorithms (e.g., quicksort, mergesort, bubblesort) and switch between them dynamically depending on factors like the size of the data or its characteristics.

2. Payment processing: In a payment processing system, you may have different strategies for handling payments (e.g., credit card, PayPal, bank transfer). Each payment method can be encapsulated as a strategy, allowing you to easily switch between them without modifying the core payment processing logic.

3. Compression algorithms: You can use the Strategy pattern to implement different compression algorithms (e.g., gzip, deflate, lz4) and switch between them based on factors like performance requirements or the type of data being compressed.

4. User authentication: Different authentication methods (e.g., username/password, OAuth, fingerprint) can be implemented as separate strategies, allowing you to switch between them based on user preferences or security requirements.

5. Routing in web frameworks: Web frameworks often use the Strategy pattern for routing requests to different controllers or handlers based on URL patterns or request attributes.

Popular libraries that use the Strategy pattern include:

- **Java Collections Framework**: The `sort()` method of collections in Java allows you to specify a comparator strategy to customize the sorting behavior.
- **.NET Framework**: The `System.IO.Compression` namespace in .NET provides classes like `DeflateStream` and `GZipStream`, which implement different compression algorithms using the Strategy pattern.
- **Ruby on Rails**: Rails controllers often use the Strategy pattern for routing requests to actions based on the requested URL or HTTP method.
- **Django**: Django's middleware system allows you to define middleware classes that implement different behaviors for processing HTTP requests and responses, similar to the Strategy pattern.

---

## Implementation in Go

### Basic Implementation

Here's an implementation of the Strategy pattern in Go:

```go
package main

import "fmt"

// Strategy defines the interface for the different algorithms
type Strategy interface {
	Execute()
}

// ConcreteStrategyA implements a specific algorithm
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute() {
	fmt.Println("Executing Strategy A")
}

// ConcreteStrategyB implements a specific algorithm
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute() {
	fmt.Println("Executing Strategy B")
}

// Context uses the Strategy interface to call the algorithm
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

func main() {
	context := Context{}

	// Using Strategy A
	context.SetStrategy(&ConcreteStrategyA{})
	context.ExecuteStrategy()

	// Changing to Strategy B
	context.SetStrategy(&ConcreteStrategyB{})
	context.ExecuteStrategy()
}
```

**In this implementation**:

- `Strategy` is an interface that defines the algorithm's contract.
- `ConcreteStrategyA` and `ConcreteStrategyB` are concrete implementations of the `Strategy` interface, each encapsulating a specific algorithm.
- `Context` is the class that uses a strategy. It holds a reference to a `Strategy` object and can change the strategy dynamically.
- In the `main()` function, we demonstrate how the context can switch between different strategies (`ConcreteStrategyA` and `ConcreteStrategyB`) and execute them interchangeably.

---

### Changing Text Formats

Implementation of the strategy pattern where we print a list of text items in different formats:

```go
package main

import (
	"fmt"
	"strings"
)

const (
	Markdown int = iota
	Html
)

type ListStrategy interface {
	Prefix(builder *strings.Builder)
	Suffix(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

// Markdown strategy
type MarkdownList struct{}

func (m *MarkdownList) Prefix(builder *strings.Builder) {}

func (m *MarkdownList) Suffix(builder *strings.Builder) {}

func (m *MarkdownList) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("- " + item + "\n")
}

// HTML Strategy
type HTMLList struct{}

func (h *HTMLList) Prefix(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h *HTMLList) Suffix(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (h *HTMLList) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("  <li>" + item + "</li>\n")
}

// TextProcessor
type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

// Constructor
func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{strings.Builder{}, listStrategy}
}

func (t *TextProcessor) SetStrategy(format int) {
	switch format {
	case Markdown:
		t.listStrategy = &MarkdownList{}
	case Html:
		t.listStrategy = &HTMLList{}
	}
}

// Add an item to the list
func (t *TextProcessor) AppendList(items ...string) {
	s := t.listStrategy
	s.Prefix(&t.builder)
	for _, item := range items {
		s.AddListItem(&t.builder, item)
	}
	s.Suffix(&t.builder)
}

// Reset the string builder
func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

// Default formatting option
func (t *TextProcessor) String() string {
	return t.builder.String()
}

func main() {
	tp := NewTextProcessor(&MarkdownList{})
	tp.AppendList("Apple", "Bat", "Cat", "Dog", "Elephant")
	fmt.Println(tp)
	/*
		- Apple
		- Bat
		- Cat
		- Dog
		- Elephant
	*/

	tp.Reset()
	tp.SetStrategy(Html)
	tp.AppendList("Apple", "Bat", "Cat", "Dog", "Elephant")
	fmt.Println(tp)
	/*
		<ul>
		  <li>Apple</li>
		  <li>Bat</li>
		  <li>Cat</li>
		  <li>Dog</li>
		  <li>Elephant</li>
		</ul>
	*/
}
```

---

# Summary

- Define an algorithm at a high level
- Define the common interface you want the strategies to follow
- Support the injection of the strategy into the high-level algorithm

---
