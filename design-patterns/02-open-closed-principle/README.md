# Table of Contents

- [Table of Contents](#table-of-contents)
- [Open-Closed Principle](#open-closed-principle)
- [Examples](#examples)
  - [Example 1: Product Filter](#example-1-product-filter)
    - [Case 1: Product Filter (that does not follow the Open-Closed Principle)](#case-1-product-filter-that-does-not-follow-the-open-closed-principle)
    - [Case 2: Product Filter (that follows the Open-Closed Principle)](#case-2-product-filter-that-follows-the-open-closed-principle)
- [Summary](#summary)

---

# Open-Closed Principle

Types should be open for expansion, closed for modification.

# Examples

## Example 1: Product Filter

**Description**: In this scenario, you want to add a Product Filter by Size, Filter by Price

### Case 1: Product Filter (that does not follow the Open-Closed Principle)

**Implementation**: Not following the Open-Closed Principle means the need to keep modifying code.

```go
package lib

type Color int
type Size int

// Declare Constants for Color and Size
const (
	Red Color = iota
	Green
	Blue
	White
)
const (
	Sm Size = iota
	M
	Lg
	Xl
	Xxl
)

// Declare Product struct
type Product struct {
	Name  string
	Color Color
	Size  Size
}

// Declare Filter struct
type Filter struct{}

// Method on Filter that filters by color (Does not follow Open-Closed Principle as we have to keep modifying based on changes)
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.Color == color {
			result = append(result, &products[i])
		}
	}
	return result
}
```

---

### Case 2: Product Filter (that follows the Open-Closed Principle)

In **`lib/types.go`**:

```go
package lib

type Color int
type Size int

// Declare Constants for Color and Size
const (
	Red Color = iota
	Green
	Blue
	White
)
const (
	Sm Size = iota
	M
	Lg
	Xl
	Xxl
)

// Declare Product struct
type Product struct {
	Name  string
	Color Color
	Size  Size
}

// Specification structs
type ColorSpecification struct{ Color Color }
type SizeSpecification struct{ Size Size }

// Base Specification interface that can be extended
type Specification interface {
	IsSatisfied(p *Product) bool
}

// Filter struct that obeys Open-Closed principle
type BetterFilter struct{}

// Method on ColorSpecification that checks if ColorSpecification is satisfied
func (s ColorSpecification) IsSatisfied(p *Product) bool {
	return p.Color == s.Color
}

// Method on SizeSpecification that checks if SizeSpecification is satisfied
func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size == s.Size
}

// Method on BetterFilter that takes in a list of products and a number of specifications. Does AND operation.
// Follows Open-Closed principle as this doesn't have to be modified.
func (f *BetterFilter) AndFilter(products []Product, specs ...Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		isSatisfied := true
		for _, spec := range specs {
			isSatisfied = isSatisfied && spec.IsSatisfied(&v)
			if !isSatisfied {
				break
			}
		}
		if isSatisfied {
			result = append(result, &products[i])
		}
	}
	return result
}

// Method on BetterFilter that takes in a list of products and a number of specifications. Does OR operation.
// Follows Open-Closed principle as this doesn't have to be modified.
func (f *BetterFilter) OrFilter(products []Product, specs ...Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		isSatisfied := false
		for _, spec := range specs {
			isSatisfied = isSatisfied || spec.IsSatisfied(&v)
			if !isSatisfied {
				break
			}
		}
		if isSatisfied {
			result = append(result, &products[i])
		}
	}
	return result
}
```

In **`main.go`**:

```go
package main

import (
	"fmt"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/02-open-closed-principle/lib"
)

func main() {
	apple := lib.Product{Name: "Apple", Color: lib.Green, Size: lib.Sm}
	tree := lib.Product{Name: "Tree", Color: lib.Green, Size: lib.Lg}
	house := lib.Product{Name: "House", Color: lib.White, Size: lib.Xl}

	products := []lib.Product{apple, tree, house}
	fmt.Println("Green products (old):")
	f := lib.Filter{}
	for _, v := range f.FilterByColor(products, lib.Green) {
		fmt.Printf("%s is green\n", v.Name)
	}

	fmt.Println("Green products (new):")
	greenSpecification := lib.ColorSpecification{Color: lib.Green}
	bf := lib.BetterFilter{}

	for _, v := range bf.AndFilter(products, greenSpecification) {
		fmt.Printf("%s is green\n", v.Name)
	}

	/** Extended to have an AND specification */
	fmt.Println("Green AND Large products:")
	lgSpecification := lib.SizeSpecification{Size: lib.Lg}

	for _, v := range bf.AndFilter(products, greenSpecification, lgSpecification) {
		fmt.Printf("%s is large and green\n", v.Name)
	}

	/** Extended to have an OR specification */
	fmt.Println("Green OR Large products:")

	for _, v := range bf.OrFilter(products, greenSpecification, lgSpecification) {
		fmt.Printf("%s is either large or green\n", v.Name)
	}
}
```

---

# Summary

In this case the interface `Specification` is open for extension, but it's closed for modification (unlikely to modify).
Similarly, we are unlikely to modify `Filter` method on the `BetterFilter` struct.
