# Table of Contents

- [Table of Contents](#table-of-contents)
- [Liskov's Substitution Principle](#liskovs-substitution-principle)

---

# Liskov's Substitution Principle

The Liskov's Substitution Principle states that if you have some API that takes a base class and works correctly with the base class, it should also work correctly with the derived class.

But in Go, we do not have Parent class and children classes that inherit from the parent. This concept doesn't exist in Go.

However, we could do a variation of it using Interfaces. It doesn't work exactly as intended because Inheritance in Go doesn't work the same way as an object oriented programming language like Java or a prototypal language like JavaScript.

**Example**: We are trying to deal with quadrilaterals (four sided polygons like rectangles, squares, rhombuses). We want the `BaseQuadrilateral` to have some common functions and properties that we can use for either a `Square` or `Rectangle` without breaking them.

```go
package main

import (
	"fmt"
)

type Quadrilateral interface {
	SetSides(a, b, c, d int)
	GetSides() [4]int
	Perimeter() int
}

type BaseQuadrilateral struct {
	SideA int
	SideB int
	SideC int
	SideD int
	Quadrilateral
}

type Square struct {
	BaseQuadrilateral
}

type Rectangle struct {
	BaseQuadrilateral
}

/* ******************************************** */
// Base Quadrilateral
/* ******************************************** */
// Set Sides
func (quad *BaseQuadrilateral) SetSides(a, b, c, d int) {
	quad.SideA = a
	quad.SideB = b
	quad.SideC = c
	quad.SideD = d
}

// Get Sides
func (quad *BaseQuadrilateral) GetSides() [4]int {
	return [4]int{quad.SideA, quad.SideB, quad.SideC, quad.SideD}
}

// Perimeter
func (quad *BaseQuadrilateral) Perimeter() int {
	return quad.SideA + quad.SideB + quad.SideC + quad.SideD
}

/* ******************************************** */
// Square
/* ******************************************** */
// Validity Check
func (sq *Square) IsValid() bool {
	if sq.SideA == sq.SideB && sq.SideB == sq.SideC && sq.SideC == sq.SideD {
		return true
	}
	return false
}

// Area
func (sq *Square) Area() int {
	if sq.IsValid() {
		return sq.SideA * sq.SideB
	}
	return -1
}

/* ******************************************** */
// Rectangle
/* ******************************************** */
// Validity check
func (sq *Rectangle) IsValid() bool {
	if sq.SideA == sq.SideC && sq.SideB == sq.SideD {
		return true
	}
	return false
}

// Area
func (sq *Rectangle) Area() int {
	if sq.IsValid() {
		return sq.SideA * sq.SideB
	}
	return -1
}

/* ******************************************** */
// Main Function
/* ******************************************** */
func main() {
	/** BaseQuadrilateral **/
	quad := &BaseQuadrilateral{}
	sq := &Square{}
	rect := &Rectangle{}

	quad.SetSides(4, 5, 6, 7)
	sq.SetSides(4, 4, 4, 4)
	rect.SetSides(4, 6, 4, 6)

	// Get Sides
	fmt.Println("Sides (Quadrilateral):", quad.GetSides())
	fmt.Println("Sides (Square):", sq.GetSides())
	fmt.Println("Sides (Rectangle):", rect.GetSides())

	// Perimeter
	fmt.Println("Perimeter (Quadrilateral):", quad.Perimeter())
	fmt.Println("Perimeter (Square):", sq.Perimeter())
	fmt.Println("Perimeter (Rectangle):", rect.Perimeter())

	// Area
	fmt.Println("Area (Square):", sq.Area())
	fmt.Println("Area (Rectangle):", rect.Area())
}
```
