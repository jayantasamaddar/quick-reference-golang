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
