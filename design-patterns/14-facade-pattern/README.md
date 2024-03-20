# Table of Contents

- [Table of Contents](#table-of-contents)
- [Facade Design Pattern](#facade-design-pattern)
	- [Introduction](#introduction)
	- [Implementation in Go](#implementation-in-go)
- [Summary](#summary)

---

# Facade Design Pattern

## Introduction

The Facade design pattern is a structural design pattern that provides a simplified and unified interface to a set of interfaces in a subsystem. It involves creating a higher-level interface that makes it easier to use a complex system by providing a simplified interface to a set of interfaces or functionalities.

The main goal of the Facade pattern is to hide the complexities of a subsystem and present a simplified and unified interface to the client or user. This can improve the usability and maintainability of a system by reducing the dependencies between the client code and the subsystem components.

**Key components of the Facade pattern**:

1. **Facade**: This is the main entry point for the client and provides a simplified interface. The facade delegates client requests to the appropriate subsystem objects.

2. **Subsystem Classes**: These are the individual components or classes that make up the subsystem. The facade delegates the client's requests to these classes, but the client is shielded from the complexities of interacting with these classes directly.

---

## Implementation in Go

Here's an example in Go that uses the Facade pattern with multiple `Buffers` and `Viewports`, initializing a single viewport and buffer initially:

In this example:

- The `Console` facade is extended to handle multiple buffers and viewports created from `Buffer` and `Viewport` subsystems.
- The `CreateBuffer` method allows creating new buffers, and the `CreateViewport` method allows creating new viewports associated with the latest buffer.
- The `SwitchViewport` method allows switching between active viewports.
- The client code demonstrates creating multiple buffers and viewports and writing to the selected viewport. The `Render` method displays the content of the - currently active viewport.

```go
package main

import "fmt"

// Subsystem: Buffer
type Buffer struct {
	content string
}

func (b *Buffer) Write(text string) {
	b.content += text
}

// Subsystem: Viewport
type Viewport struct {
	buffer *Buffer
	width  int
	height int
}

func NewViewport(buffer *Buffer, width, height int) *Viewport {
	return &Viewport{
		buffer: buffer,
		width:  width,
		height: height,
	}
}

func (v *Viewport) Render() string {
	return fmt.Sprintf("Viewport (%dx%d):\n%s", v.width, v.height, v.buffer.content)
}

// Facade: Console
type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	active    int
}

func NewConsole() *Console {
	// Initialize with a single buffer and viewport
	buffer := &Buffer{}
	viewport := NewViewport(buffer, 80, 40)

	return &Console{
		buffers:   []*Buffer{buffer},
		viewports: []*Viewport{viewport},
		active:    0,
	}
}

func (c *Console) CreateBuffer() {
	buffer := &Buffer{}
	c.buffers = append(c.buffers, buffer)
}

func (c *Console) CreateViewport(width, height int) {
	buffer := c.buffers[len(c.buffers)-1]
	viewport := NewViewport(buffer, width, height)
	c.viewports = append(c.viewports, viewport)
	c.active = len(c.viewports) - 1
}

func (c *Console) SwitchViewport(index int) {
	if index >= 0 && index < len(c.viewports) {
		c.active = index
	}
}

func (c *Console) Write(text string) {
	buffer := c.buffers[c.active]
	buffer.Write(text)
}

func (c *Console) Render() string {
	viewport := c.viewports[c.active]
	return viewport.Render()
}

// Client code
func main() {
	console := NewConsole()

	// Writing to the initial viewport
	console.Write("Hello from the initial viewport!\n")

	// Creating a new buffer and viewport
	console.CreateBuffer()
	console.CreateViewport(120, 60)

	// Writing to the new viewport
	console.SwitchViewport(1)
	console.Write("Hello from the new viewport!\n")

	// Rendering the current viewport
	result := console.Render()
	fmt.Println(result)
	/*
		Viewport (120x60):
		Hello from the new viewport!
	*/
}
```

---

# Summary

- Provide a simplified API over a set of components
- May wish to (optionally) expose internals through the facade (for power users)
- May allow users to "escalate" to use more complex APIs if they need to

---
