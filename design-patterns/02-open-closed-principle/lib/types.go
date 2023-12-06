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

// Base Specification interface that can be extended
type Specification interface {
	IsSatisfied(p *Product) bool
}

// Specification structs
type ColorSpecification struct{ Color Color }
type SizeSpecification struct{ Size Size }

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

// Method on BetterFilter that takes in a single specification. Follows Open-Closed principle as this doesn't have to be modified.
func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

// Method on BetterFilter that takes in a number of specifications. Does AND operation.
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

// Method on BetterFilter that takes in a number of specifications. Does OR operation.
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
