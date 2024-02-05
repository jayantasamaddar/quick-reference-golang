# Table of Contents

- [Table of Contents](#table-of-contents)
- [Builder Pattern](#builder-pattern)
  - [Introduction](#introduction)
  - [Buider Patterns in Go](#buider-patterns-in-go)
    - [`strings.Builder`](#stringsbuilder)
    - [Implementing a custom `HTMLBuilder` using `strings.Builder`](#implementing-a-custom-htmlbuilder-using-stringsbuilder)
    - [Using Fluent Interfaces with the Builder Pattern](#using-fluent-interfaces-with-the-builder-pattern)
  - [Builder Facets](#builder-facets)
  - [Builder Parameter](#builder-parameter)
  - [Functional Builder](#functional-builder)
- [Summary](#summary)

---

# Builder Pattern

## Introduction

The motivation for the Builder Pattern is that some objects (`structs` in Go) are simple and can be created in a single constructor call or a Factory function or by initializing the fields.

In some situations, this works fine. In other situations, objects require some ceremony to create.

For e.g. Having a factory function with 10 arguments is not very productive.

We can work around it and make the construction process, a kind of a multi-stage process, i.e. we construct an object piece-wise than try to do everything in a single factory call.

The Builder Pattern solves a situation when object construction is complicated, by providing an API for constructing an object step-by-step succinctly.

---

## Buider Patterns in Go

### `strings.Builder`

The strings builder is an example of the Builder pattern implementation in the core language of Go. It needs to be imported from the `strings` module.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "Hello World!"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")

	fmt.Println(sb.String())

	words := []string{"Jack", "Paul", "Sam"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())
	sb.Reset()
}
```

This may seem a bit complicated to build a HTML element. Ideally we would want to simply add this to some sort of a Tree Data structure representing HTML nodes and then construct the HTML out of it.

---

### Implementing a custom `HTMLBuilder` using `strings.Builder`

```go
package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

/*********************************************************************************************************/
// Implementation of a single HTMLElement
/*********************************************************************************************************/
type HTMLElement struct {
	name, text string
	elements   []HTMLElement
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))

	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
	}
	sb.WriteString(e.text)
	sb.WriteString("\n")

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))

	return sb.String()
}

/*********************************************************************************************************/
// Implementation of HTMLBuilder
/*********************************************************************************************************/
type HTMLBuilder struct {
	// Name of the root element
	rootName string
	root     HTMLElement
}

func NewHTMLBuilder(rootName string) *HTMLBuilder {
	return &HTMLBuilder{rootName, HTMLElement{rootName, "", []HTMLElement{}}}
}

func (b *HTMLBuilder) String() string {
	return b.root.String()
}

func (b *HTMLBuilder) AddChild(childName, childText string) {
	e := HTMLElement{childName, childText, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
}

func main() {
	// Demonstration of HTMLBuilder
	b := NewHTMLBuilder("ul")
	fruits := []string{"Apple", "Banana", "Mango"}
	for _, v := range fruits {
		b.AddChild("li", v)
	}
	fmt.Println(b.String())
}
```

---

### Using Fluent Interfaces with the Builder Pattern

A Fluent Interface is an interface that allows you to chain calls together. Chaining calls in Go, isn't super convenient as we leave the dot hanging in the end, instead of the beginning:

```go
// Clear HTMLElements for the HTMLBuilder root
func (b *HTMLBuilder) Clear() {
	b.root = HTMLElement{b.rootName, "", []HTMLElement{}}
}

// Return the Builder itself for chaining
func (b *HTMLBuilder) AddChildFluent(childName, childText string) *HTMLBuilder {
	e := HTMLElement{childName, childText, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	// (3) Demonstration of HTMLBuilder with Fluent Interface (AddChildFluent)
	b.Clear()
	b.AddChildFluent("li", "One").AddChildFluent("li", "Two").AddChildFluent("li", "Three")
	fmt.Println(b.String())
}
```

---

## Builder Facets

There are situations where you need more than one builder, where you need to somehow separate the process of building up the different aspects of a particular type.

Let's try to build a sort of DSL to build a Person struct that has two sets of information that can have their own information:

- `Person`: Built using a `PersonBuilder`
  - Address Information: Built using a `PersonAddressBuilder`
  - Work Information: Built using a `PersonWorkBuilder`

```go
package main

/*********************************************************************************************************/
// (3) Builder Facets: PersonBuilder with PersonAddressBuilder and PersonWorkBuilder
/*********************************************************************************************************/
type Person struct {
	Name string

	// Address Information
	Street, City, Zip string

	// Work Information
	Company, Designation string
	AnnualIncome         int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder(name string) *PersonBuilder {
	return &PersonBuilder{&Person{Name: name}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}
type PersonWorkBuilder struct {
	PersonBuilder
}

// Utility Methods to access the PersonAddressBuilder and PersonWorkBuilder
func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}
func (pb *PersonBuilder) Works() *PersonWorkBuilder {
	return &PersonWorkBuilder{*pb}
}

// PersonAddressBuilder specific methods
func (pb *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	pb.person.Street = street
	return pb
}
func (pb *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pb.person.City = city
	return pb
}
func (pb *PersonAddressBuilder) Zip(zip string) *PersonAddressBuilder {
	pb.person.Zip = zip
	return pb
}

// PersonWorkBuilder specific methods
func (pb *PersonWorkBuilder) At(company string) *PersonWorkBuilder {
	pb.person.Company = company
	return pb
}
func (pb *PersonWorkBuilder) Designation(designation string) *PersonWorkBuilder {
	pb.person.Designation = designation
	return pb
}
func (pb *PersonWorkBuilder) AnnualIncome(income int) *PersonWorkBuilder {
	pb.person.AnnualIncome = income
	return pb
}

// Build the actual Person
func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

/*********************************************************************************************************/
// Main Function
/*********************************************************************************************************/
func main() {
	// Demonstration of Builder Facets
	pb := NewPersonBuilder("Sherlock Holmes")
	pb.Lives().At("221B Baker Street").In("London").Zip("NW16XE").Works().At("Science of Deduction").Designation("Private Investigator").AnnualIncome(200000)

	fmt.Println(pb.Build())
}
```

At any point of the building process we can switch to another builder.

---

## Builder Parameter

One question we might have is, "How do I get the users of my API to use my builders, as opposed to stop messing with the objects directly?".
In other words, we want to force the client to use the Builder as opposed to providing some sort of incomplete object for initialization.

The way to get this to work is to use hide objects that you don't want the users to touch.

In **`lib/PersonBuilder.go`**:

```go
package lib

import "strings"

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

// Builder methods
func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain '@'")
	}
	eb.email.from = from
	return eb
}
func (eb *EmailBuilder) To(to string) *EmailBuilder {
	eb.email.to = to
	return eb
}
func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.subject = subject
	return eb
}
func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.body = body
	return eb
}

// Send the email
func sendMailImpl(email *email) {
	// Implementation
}

// Builder Parameter: Takes a builder and does something with it, typically calls something on the builder
type build func(*EmailBuilder)

// Publicly exposed function that people are meant to use
func SendEmail(action build) {
	b := EmailBuilder{}
	action(&b)
	sendMailImpl(&b.email)
}
```

In **`main.go`**:

```go
package main

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").To("bar@baz.com").Subject("Meeting!").Body("Hello do you want to meet?")
	})
}
```

---

## Functional Builder

One way of extending a builder is by using a functional programming approach.
The benefit of this setup is that it is very easy to extend the builder with additional build actions without messing about with making new builders which aggregate the current builder.

Let's try to apply this to a `Country` struct.

In **`lib/CountryBuilder.go`**:

```go
package lib

type Country struct {
	name, currency string
}

type countryMod func(*Country)
type CountryBuilder struct {
	actions []countryMod
}

func (cb *CountryBuilder) Name(name string) *CountryBuilder {
	cb.actions = append(cb.actions, func(c *Country) {
		c.name = name
	})
	return cb
}

func (cb *CountryBuilder) Currency(currency string) *CountryBuilder {
	cb.actions = append(cb.actions, func(c *Country) {
		c.currency = currency
	})
	return cb
}

func (cb *CountryBuilder) Build() *Country {
	c := &Country{}
	for _, action := range cb.actions {
		action(c)
	}
	return c
}
```

In **`main.go`**:

```go
package main

import "fmt"

func main() {
	// Demonstration of Functional Builder
	cb := lib.CountryBuilder{}
	fmt.Println(cb.Name("India").Currency("INR").Build()) // &{India INR}
}
```

Using this method, the builder, instead of doing modifications in-place, keeps a list of actions which are executed when the `Build` method is called to execute the slice of actions.

---

# Summary

- A Builder is a separate component used for building an object.
- To make a builder fluent, return the receiver - allows chaining.
- Different facets of an object can be built with different builders working in tandem via a common struct they are aggregating to.
