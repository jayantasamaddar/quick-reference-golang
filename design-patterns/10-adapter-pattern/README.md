# Table of Contents

- [Table of Contents](#table-of-contents)
- [Adapter Pattern](#adapter-pattern)
  - [Introduction](#introduction)
  - [Implementing Adapter in Go](#implementing-adapter-in-go)
  - [Adapter Caching](#adapter-caching)
- [Summary](#summary)

---

# Adapter Pattern

## Introduction

The Adapter Pattern is a design pattern that allows the interface of an existing class to be used as another interface using an "adapter" class. It is a structural pattern that enables incompatible interfaces to work together. This pattern involves a single class called the adapter, which is responsible for joining functionalities of independent or incompatible interfaces.

The Adapter Pattern is particularly useful when you have existing code or components that you want to reuse, but their interfaces are not compatible with the rest of your system. Instead of modifying the existing code, you can create an adapter class that acts as a bridge between the existing code and the desired interface.

A real life example to envision this is the Adapter Plug that can be used to connect a device to the socket where the device interface is incompatible. For e.g. A laptop from India that has a 240V interface will not work with the 120V socket in the US unless we use an adapter.

In Go, an Adapter is a construct which adapts an existing interface X to conform to the required interface Y.

---

## Implementing Adapter in Go

Let's imagine a scenario where you are working with some sort of API for rendering of graphical objects and that API is completely vector based (all images are defined as a bunch of lines). Let's say we are given this API by an external developer and we have a function for making new graphical objects.

```go
type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

// Function for making a rectangle
func NewRectangle(width, height int) *VectorImage {
	// Zero-based. 5 means, position 0 to position 4
	width -= 1
	height -= 1

	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}}
}
```

Let's say, we are given the `NewRectangle` interface. However, we cannot work with the interface as we do not way of rendering the rectangle as an output.

Now, we may have an interface that is meant for a Raster Image using Points

```go
// Interface we have
type Point struct {
	X, Y int
}

// An Image defined by points. We have this as an interface to have different implementations.
type RasterImage interface {
	GetPoints() []Point
}

// Utility function for defining the points as an image (as string)
func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate
	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.X][point.Y] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}
```

So, the conundrum is, the only way to create a Rectangle is using the API provided, which is a `VectorImage`, and the only way to print it is providing a `RasterImage`. So how do we fix this?

We will need an Adapter, that takes in the `VectorImage` and somehow adapts it into something that has a bunch of points in it, so that they can be fed into the `RasterImage` and eventually into the `DrawPoints` function.

Our final implementation will look something like this:

```go
package main

import "fmt"

func main() {
	rc = NewRectangle(6, 4)
	// An Adapter that converts Vector to Raster
	adapted := VectorToRaster(rc)

	// We can now use the adapted RasterImage and feed it into the DrawPoints function
	fmt.Println(DrawPoints(a))
}
```

Thus, what remains now is building the solution:

```go
package main

import (
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height}}}
}

// Interface we have

type Point struct {
	X, Y int
}

// An Image defined by points. We have this as an interface to have different implementations.
type RasterImage interface {
	GetPoints() []Point
}

// Utility function for defining the points as an image (as string)
func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate
	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

/************************************************************************************/
// Solution
/************************************************************************************/
func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// The adapter that will represent the RasterImage
type vectorToRasterAdapter struct {
	points []Point
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

// Convert a line into points
func (a *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	fmt.Println("generated", len(a.points), "points")
}

// Convert a VectorImage into a RasterImage via an Adapter
func VectorToRaster(v *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range v.Lines {
		adapter.addLine(line)
	}

	return adapter // as RasterImage
}

/************************************************************************************/
// Main Function
/************************************************************************************/
func main() {
	rc := NewRectangle(6, 4)
	// An Adapter that converts Vector to Raster
	a := VectorToRaster(rc)

	// We can now use the adapted RasterImage and feed it into the DrawPoints function
	fmt.Print(DrawPoints(a))
	/*
    generated 6 points
    generated 10 points
    generated 14 points
    generated 20 points
	******
	*    *
	*    *
	******

	 */
}
```

---

## Adapter Caching

One of the things to aware of, when working with adapters that we have to be careful about, is the creation of too many temporary objects. In the above example, to draw the lines, we had to convert every single line into bunch of points. This becomes a bit of a problem, if we try to do this more than once.

To illustrate, let's just add another adapter:

```go
func main() {
	rc := NewRectangle(6, 4)
	// An Adapter that converts Vector to Raster
	a := VectorToRaster(rc)
    // Second adapter
	_ = VectorToRaster(rc)

	// We can now use the adapted RasterImage and feed it into the DrawPoints function
	fmt.Print(DrawPoints(a))
	/*
			generated 6 points
		    generated 10 points
		    generated 14 points
		    generated 20 points
            generated 6 points
		    generated 10 points
		    generated 14 points
		    generated 20 points
			******
			*    *
			*    *
			******

	*/
}
```

We can see that the same adapter has been initialized twice, and also ran the `addLine` computations twice. As we can see, these operations are not really necessary. Although one may think, you might modify one adapter and the other adapter stays the same, so why don't we keep the extra data? Well, we can avoid running this extra computation and storing it, if you don't need it by caching the results and using the old results if there are no changes.

To implement caching, we can build a very simple cache:

```go
/************************************************************************************/
// Adapter Caching
/************************************************************************************/
// Point cache: 16 byte array that will store a MD5 hash
var pointCache = map[[16]byte][]Point{}

// Hash function to calculate the hash of each Line
func hash(obj interface{}) [16]byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic("failed to marshal line")
	}
	return md5.Sum(bytes)
}

// Convert a line into points
func (a *vectorToRasterAdapter) addLineCached(line Line) {
	// Calculate hash
	h := hash(line)
    // Don't add points if they have already been generated
	if points, ok := pointCache[h]; ok {
		// Instead of calculating points, simply add them to this new adapter
		a.points = append(a.points, points...)
		return
	}

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

    // Store points to the hash map in the cache
	pointCache[h] = a.points
	fmt.Println("generated", len(a.points), "points")
}

// Convert a VectorImage into a RasterImage via an Adapter
func VectorToRasterCached(v *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range v.Lines {
		adapter.addLineCached(line)
	}

	return adapter // as RasterImage
}

func main() {
    rc := NewRectangle(6, 4)
    /************************************************************************************/
	// Demonstration of Adapter Pattern (Cached)
	/************************************************************************************/
	// An Adapter that converts Vector to Raster
	b := VectorToRasterCached(rc)
	// Second adapter
	_ = VectorToRasterCached(rc)

	// We can now use the adapted RasterImage and feed it into the DrawPoints function
	fmt.Print(DrawPoints(b))
	/*
			generated 6 points
		    generated 10 points
		    generated 14 points
		    generated 20 points
			******
			*    *
			*    *
			******

	*/
}
```

As you can see, any subsequent initialization of an adapter with the same `VectorImage` retrieves from the cache instead of processing the line to points conversion, saving on processing time and avoiding the same calculations.

---

# Summary

- Determine the API you have (source) and the API you need (destination)
- Create a construct which aggregates (has a pointer to) the adaptee (destination)
- Intermediate representation can pile up: use caching and other optimizations
