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
