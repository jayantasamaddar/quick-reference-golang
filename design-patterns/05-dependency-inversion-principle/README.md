# Dependency Inversion Principle

The Dependency Inversion Principle states that High-level modules should not depend on Low-level modules and both of them should depend on abstractions (in Go, that would mean interfaces).

The idea behind the principle is that, any breaking changes to the Low-level module would break the high-level module depending on it.

Let's say you are doing some genealogy research and you want to model relationships between different people.

```go
package main

import (
	"fmt"
)

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Storage for relationships (Could be database): Low level module
type Relationships struct {
	relations []Info
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// Perform research on the data: High-level module
type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate(name string) {
	for _, child := range r.browser.FindAllChildrenOf(name) {
		fmt.Printf("%v has a child called %s.\n", name, child.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Steve"}
	child2 := Person{"Chris"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate("John")
}
```

The above implementation has been correctly done with **Dependency Inversion Principle**. We are now protected against changes now. For example, if we were to change the storage mechanism of relations from a slice to something more sophisticated, we would only be modifying the methods of `Relationships`. We would not be modifying for example, the methods of `Research`, because it doesn't depend on the low-level detail.
