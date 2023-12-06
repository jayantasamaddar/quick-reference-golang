Types should be open for expansion, closed for modification.

**Example**: In this scenario, you want to add a Product Filter by Size, Filter by Price

```go
package main

import (
    "fmt"
)

type Color int

const (
    red Color = iota
    green
    blue
)

type Size int

const (
    s Size = iota
    m
    l
    xl
    xxl
)

type Product struct {
    name string
    color Color
    size Size
}

type Filter struct {

}

type Specification interface {
    IsSatisfied(p *product) bool
}

type ColorSpecification struct {
    color Color
}

type SizeSpecification struct {
    size Size
}

type BetterFilter struct {}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
    return p.color == c.color
}

func (s sizeSpecification) IsSatisfied(s *Size) bool {
    return p.size == s.size
}

func (f *Filter) FilterByColor(products []Product, size Size, color Color) []*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if v.size == size && v.color == color {
            result = append(result, &products[i])
        }
    }
    return result
}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
    result := make([]*Product, 0)
    for i, v := range products {
        if spec.isSatsified(&v) {
            result = append(result, &products[i])
        }
    }
    return result
}

func main() {
    apple := Product{"Apple", green, small}
    tree := Product{"Tree", green, large}
    house := Product{"House", white, xl}

    products := []Products{apple, tree, house}
    fmt.Println("Green products (old):")
    f := Filter{}
    for _, v := range f.FilterByColor(products, green) {
        fmt.Println("%s is green", v.name)
    }

    fmt.Println("Green products (new):")
    greenSpecification := ColorSpecification{green}
    bf := BetterFilter{}

    for _, v := range bf.Filter(products, greenSpecification) {
        fmt.Println("%s is green", v.name)
    }
}
```

In this case the interface `Specification` is open for extension, but it's closed for modification (unlikely to modify).
Similarly, we are unlikely to modify `Filter` method on the `BetterFilter` struct.
