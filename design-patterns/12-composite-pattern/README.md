# Table of Contents

- [Table of Contents](#table-of-contents)
- [Composite Pattern](#composite-pattern)
  - [Introduction](#introduction)
  - [Implementation in Go](#implementation-in-go)
    - [Composing Graphics](#composing-graphics)
    - [Neural Networks](#neural-networks)
- [Summary](#summary)

---

# Composite Pattern

## Introduction

The Composite Pattern is a structural design pattern that lets you compose objects into tree-like structures to represent part-whole hierarchies. This pattern allows clients to treat individual objects and compositions of objects uniformly.

In simpler terms, the Composite Pattern is used to treat both individual objects and compositions of objects in a uniform manner. This is achieved by creating a common interface for both types of objects. Clients can then use this interface to interact with individual objects or complex compositions of objects without having to distinguish between them.

**Key components of the Composite Pattern**:

1. **Component:** This is the common interface or abstract class that declares the methods for both leaf objects (individual objects) and composite objects (objects that contain other objects).

2. **Leaf:** This represents the individual objects that do not have any child elements. They implement the Component interface.

3. **Composite:** This represents the complex objects that can have child elements. It also implements the Component interface but may contain a collection of child components.

The main advantages of the Composite Pattern include:

- **Uniformity:** Clients can treat individual objects and compositions of objects uniformly, simplifying the client code.
- **Flexibility:** It makes it easy to add new types of components without changing the existing code.
- **Hierarchy:** The pattern allows you to represent part-whole hierarchies, making it suitable for tree-like structures.

The Composite Pattern is commonly used in graphical user interface frameworks, document structures, and in any scenario where you need to represent part-whole hierarchies of objects.

---

## Implementation in Go

### Composing Graphics

Imagine you have a `GraphicObject`. It can either be a `Square` or `Rectangle` or `Circle`. It has the `Name` and `Color` property, but a GraphicObject can also be comprised of other shapes, i.e. other `GraphicObjects` as `Children`. Thus, this is a recursive relationship, i.e. A `GraphicObject` can have a list of `GraphicObject` as `Children`, each of which can have other `GraphicObject` as `Children` and this can go down to theoretically infinite levels.

This is the basic idea of the Composite Pattern, that you can embed an object in another and still have uniformity.

```go
package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

// print is an utility method, that is called by the String method to print out the GraphicObject by tracking its depth
func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("+ ", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

func main() {
    /*******************************************************************************************************************/
	// (1) Demo: Simulate a drawing that is a composite made up of multiple shapes. The depth is denoted by the "+ "
    /*******************************************************************************************************************/
	drawing := GraphicObject{"My Drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewSquare("Red"), *NewCircle("Yellow"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"), *NewSquare("Red"))

	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())
	/*
		My Drawing
		+ Red Square
		+ Yellow Circle
		+ Group 1
		+ + Blue Circle
		+ + Red Square
	*/
}
```

---

### Neural Networks

A Neural Network is comprised of `Neurons` which are interconnected. Multiple `Neurons` can be grouped in a `NeuronLayer`. We want to ensure that using a single `Connect` function, whether it is a `Neuron` or a `NeuronLayer`, either of these can connect to the other.

Additionally, the operation for connecting `NeuronLayers` together, actually means, every single Neuron of the first `NeuronLayer` has to be connected to every single `Neuron` of the second `NeuronLayer`

How can we ensure we can do this with one single function, considering we have to iterate every single neuron inside a Neuron and a NeuronLayer?

```go
package main

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

// Factory Function to create a Neuron Layer comprised of `count` neurons
func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

// NeuronInterface that Iterates over Neuron and NeuronLayer objects and returns a list of Neuron pointers
type NeuronInterface interface {
	Iter() []*Neuron
}

// Implemention on the NeuronLayer Collection object
func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)

	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}

	return result
}

// Implementation on the Neuron scalar object
func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

func main() {
	/*******************************************************************************************************************/
	// Demo: Composite Design Pattern in Neural Network
	/*******************************************************************************************************************/
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	nLayer1, nLayer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	// We want to have one method to connect neurons to Neurons, Neurons to NeuronLayers, NeuronLayers to Neurons and NeuronLayers to NeuronLayers
	Connect(neuron1, neuron2)
	Connect(neuron1, nLayer1)
	Connect(nLayer2, neuron1)

	// Connect every single Neuron of the first NeuronLayer and connect it to every single Neuron of the second NeuronLayer
	Connect(nLayer1, nLayer2)

}
```

---

# Summary
