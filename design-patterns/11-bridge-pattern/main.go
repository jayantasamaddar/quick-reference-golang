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
