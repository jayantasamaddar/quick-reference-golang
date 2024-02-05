# Table of Contents

- [Table of Contents](#table-of-contents)
- [Bridge Pattern](#bridge-pattern)
  - [Introduction](#introduction)
  - [Bridge Pattern in Go](#bridge-pattern-in-go)
  - [Bridge with Factory](#bridge-with-factory)
- [Summary](#summary)

---

# Bridge Pattern

## Introduction

The Bridge Design Pattern is a structural design pattern that separates an abstraction from its implementation so that the two can vary independently. It is part of the Gang of Four (GoF) design patterns and is used to decouple abstraction and implementation by providing a bridge structure between them.

Here's a breakdown of the key components:

- **Abstraction**: This defines the interface or abstraction's high-level functionality. It is the client-facing interface that clients interact with.

- **Implementor**: This is the interface for the implementation classes. It provides the low-level functionality that the abstraction uses. There can be multiple implementations of this interface.

- **Refined Abstraction**: This is a subclass of the abstraction that extends or adjusts the abstraction's interface.

- **Concrete Implementor**: This is a subclass of the implementor that provides a specific implementation.

The main idea is to allow the abstraction and implementation to vary independently. This enables you to change the abstraction or implementation separately without affecting the other, promoting flexibility and easier maintenance.

In Go, a Bridge is a mechanism that decouples an abstraction (interface) from an implementation.

---

## Bridge Pattern in Go

Imagine a scenario, where we want to render shapes like a Circle using either Vector or Raster methods. Later we might introduce new shapes like Square or Triangle and other Polygons. Instead of having a `RasterCircle`, `VectorCircle` and so on, which would explode exponentially as the different abstractions and implementations increase, we want to use the Bridge Pattern to ensure the `Shape` (abstractions) with its `Draw` and `Resize` methods stay independent of the `Renderer` (implementor).

In this case, we will just have the `Circle` as the refined abstraction, but we should be able to add other shapes too.
Likewise, we will have the `VectorRenderer` and the `RasterRenderer` as the concrete implementors.

```go
package main

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
}

type RasterRenderer struct {
	Dpi int
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels of a circle of radius", radius)
}

// Shape interface
type Shape interface {
	Draw()
	Resize(factor float32)
}

type Circle struct {
	radius   float32
	renderer Renderer
	Shape
}

func NewCircle(radius float32) *Circle {
	return &Circle{radius: radius}
}

// SetRenderer sets the renderer for the Circle
func (c *Circle) SetRenderer(renderer Renderer) {
	c.renderer = renderer
}

// Draw the circle, irrespective of renderer
func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	// Create instances of renderers
	rr := RasterRenderer{}
	vr := VectorRenderer{}

	// Create a circle with RasterRenderer
	circle := NewCircle(5)
	circle.SetRenderer(&rr)
	circle.Draw() // Drawing pixels of a circle of radius 5
	circle.Resize(2)
	circle.Draw() // Drawing pixels of a circle of radius 10

	// Create a circle with RasterRenderer
	circleWithVector := NewCircle(7)
	circleWithVector.SetRenderer(&vr)
	circleWithVector.Draw() // Drawing a circle of radius 7
	circleWithVector.Resize(3)
	circleWithVector.Draw() // Drawing a circle of radius 21
}
```

In the abiove implementation, we can categorize the components based on the Bridge pattern:

1. **Abstraction:**

   - The `Shape` interface is the abstraction. It defines the high-level functionality for shapes, which includes methods like `Draw` and `Resize`.

2. **Implementor:**

   - The `Renderer` interface is the implementor. It declares the low-level functionality for rendering, specifically the `RenderCircle` method.

3. **Refined Abstraction:**

   - The `Circle` struct is the refined abstraction. It extends the `Shape` interface and includes methods to draw and resize a circle. It also holds a reference to the `Renderer` interface.

4. **Concrete Implementor:**
   - The `VectorRenderer` and `RasterRenderer` are concrete implementors. They implement the `Renderer` interface and provide specific implementations for rendering circles.

To summarize:

- **Abstraction:** `Shape` interface
- **Implementor:** `Renderer` interface
- **Refined Abstraction:** `Circle` struct
- **Concrete Implementor:** `VectorRenderer` and `RasterRenderer`

We have the `Renderer` interface serving as the abstraction, and `VectorRenderer` and `RasterRenderer` as the concrete implementors while the `Circle` struct is not to be aware of any specific renderer implementations during construction. The `Circle` struct (class) has a `SetRenderer` method to dynamically associate a renderer with a circle instance after construction, making it more in line with the Bridge pattern.

---

## Bridge with Factory

In our previous implementation, we still are creating shapes (here: `Circle`) separately with no way of implementing the `Shape` interface unless we add it to the `Circle` struct. We might want a way where we can create Shapes, any shapes, and have the abstraction methods, i.e. `Shape` interface's methods (`Draw` and `Resize` method) available to each of them without adding the `Shape` interface manually to the new refined abstraction structs.
In addition, the `SetRenderer` seems to be a method, we should include within the Shape interface considering it'll apply to every `Shape`. This is to adhere more closely to the principles of the Bridge pattern, where abstraction and implementation are independent and can vary independently.
Finally, we might want to hide the fields of the refined abstraction, so there's no direct modification accidentally.

- The `Shape` interface should now correctly abstracts all methods necessary for any refined abstraction like a `Circle`. In the future any other refined abstractions like `Square` can be coerced to follow this abstraction.

We can do so by using a `NewShape` factory function:

```go
package main

import "fmt"

// (1) Implementor: Renderer interface
type Renderer interface {
	RenderCircle(radius float32)
}

// (2) Concrete Implementors
type VectorRenderer struct{}
type RasterRenderer struct {
	Dpi int
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels of a circle of radius", radius)
}

// (3) Abstraction: Shape interface
type Shape interface {
	Draw()
	Resize(factor float32)
	SetRenderer(renderer Renderer)
}

// (4) Refined Abstraction: Circle
type Circle struct {
	radius   float32
	renderer Renderer
}

func NewCircle(radius float32) *Circle {
	return &Circle{radius: radius}
}

// SetRenderer sets the renderer for the Circle
func (c *Circle) SetRenderer(renderer Renderer) {
	c.renderer = renderer
}

// Draw method for Circle: Decoupled from renderer
func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

// (5) Shape Factory function
func NewShape(shape Shape) Shape {
	return shape
}

func main() {
	// Create instances of renderers
	rr := RasterRenderer{}
	vr := VectorRenderer{}

	// Create a circle with RasterRenderer
	circle := NewShape(NewCircle(5))
	circle.SetRenderer(&rr)
	circle.Draw() // Drawing pixels of a circle of radius 5
	circle.Resize(2)
	circle.Draw() // Drawing pixels of a circle of radius 10

	// Create a circle with RasterRenderer
	circleWithVector := NewShape(NewCircle(7))
	circleWithVector.SetRenderer(&vr)
	circleWithVector.Draw() // Drawing a circle of radius 7
	circleWithVector.Resize(3)
	circleWithVector.Draw() // Drawing a circle of radius 21
}
```

Here are some advantages of using the `NewShape` factory function:

- **Dynamic Association**: With the `NewShape` function, you can dynamically associate different shapes (implementing the `Shape` interface) with different renderers. This promotes flexibility by allowing you to change the renderer of a shape after it has been created.

- **Separation of Concerns**: The `NewShape` function separates the creation of the shape from its rendering details. This adheres more closely to the principles of the Bridge pattern, where abstraction and implementation are independent and can vary independently.

- **Ease of Extension**: If you want to add new shapes in the future, you can simply create new implementations of the `Shape` interface, and the existing rendering logic using the `Renderer` interface can be reused.

Overall, using a factory function to create and associate shapes with renderers aligns better with the principles of the Bridge pattern and provides a more modular and extensible design.

---

# Summary

- Decouple abstraction from implementation
- Both the abstraction and implementation can exist as separate hierarchies
- A stronger form of encapsulation
