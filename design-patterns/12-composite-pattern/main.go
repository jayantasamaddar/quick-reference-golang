package main

import (
	"fmt"
	"strings"
)

/*******************************************************************************************************************/
// (1) Composing Graphics
/*******************************************************************************************************************/

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

/*******************************************************************************************************************/
// (2) Neural Networks
/*******************************************************************************************************************/
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
	/*******************************************************************************************************************/
	// (2) Demo: Composite Design Pattern in Neural Network
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
