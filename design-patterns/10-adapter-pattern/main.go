package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
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

// The adapter
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
// Adapter Caching
/************************************************************************************/
// Point cache: 16 byte array that will store a MD5 hash
var pointCache = map[[16]byte][]Point{}

// Hash function
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

/************************************************************************************/
// Main Function
/************************************************************************************/
func main() {
	rc := NewRectangle(6, 4)

	/************************************************************************************/
	// (1) Demonstration of Adapter Pattern (Not Cached)
	/************************************************************************************/
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	fmt.Println(std.Purple + "\tDemo: Adapter Pattern (Not cached)" + std.Reset)
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
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

	/************************************************************************************/
	// (2) Demonstration of Adapter Pattern (Cached)
	/************************************************************************************/
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	fmt.Println(std.Purple + "\tDemo: Adapter Pattern (Cached)" + std.Reset)
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
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
