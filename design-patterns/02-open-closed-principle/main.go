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
