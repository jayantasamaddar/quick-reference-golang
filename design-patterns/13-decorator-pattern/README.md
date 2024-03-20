# Table of Contents

- [Table of Contents](#table-of-contents)
- [Decorator Pattern](#decorator-pattern)
  - [Introduction](#introduction)
  - [Implementing Decorators in Go](#implementing-decorators-in-go)
- [Summary](#summary)

---

# Decorator Pattern

## Introduction

The Decorator Design Pattern is a structural pattern that allows behavior to be added to an individual object, either statically or dynamically, without affecting the behavior of other objects from the same class. It's achieved by creating a set of decorator classes that are used to wrap concrete components. This pattern is useful for extending functionalities in a flexible and reusable way.

---

## Implementing Decorators in Go

Let's implement a situation where a `Dragon` is both a `Bird` and a `Lizard`.

To do that we need to first create the `Bird` and the `Lizard` struct:

```go
package main

import (
	"fmt"
)

/*****************************************/
// Bird
/*****************************************/
type Bird struct {
	Age int
}
func (b *Bird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying!")
	}
}

/*****************************************/
// Lizard
/*****************************************/
type Lizard struct {
	Age int
}
func (l *Lizard) Climb() {
	if l.Age < 60 {
		fmt.Println("Climbing!")
	}
}
```

In canonical Go, we would simply aggregate these two inside a Dragon struct like:

```go
type Dragon struct {
    Bird
    Lizard
}
```

Unfortunately, this can cause problems if there are ambiguity in field names, i.e. multiple structs have the same field names. In this case, both the `Bird` and the `Lizard` structs have the `Age` field.

```go
func main() {
    d := Dragon{}
    d.Age = 30 // ambiguous selector d.AgeInMonths
}
```

To avoid this issue we could do the following:

```go
func main() {
	d := Dragon{}
	d.Bird.Age = 30
	d.Lizard.Age = 30
	d.Climb() // Climbing!
	d.Fly() // Flying!
}
```

However, this approach can introduce inconsistencies as we have to set the `Age` twice, once on `Bird` and once on `Lizard`. Also, this wouldn't scale too well. Also, this is a single field, setting it twice is ridiculous.

We can solve this by adding a common interface `AgedInterface` and using the `Dragon` struct as a decorator.

```go
package main

import (
	"fmt"
)

// Getter and Setter: Not idomatic go but for implementing the decorator there's very few options
type AgedInterface interface {
	Age() int
	SetAge(age int)
}

/*****************************************/
// Bird
/*****************************************/
type Bird struct {
	age int
}

func (b *Bird) Age() int       { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }
func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying!")
	}
}

/*****************************************/
// Lizard
/*****************************************/
type Lizard struct {
	age int
}

func (l *Lizard) Age() int       { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }
func (l *Lizard) Climb() {
	if l.age < 60 {
		fmt.Println("Climbing!")
	}
}

/*****************************************/
// Dragon
/*****************************************/
type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) Age() int       { return d.bird.age }
func (d *Dragon) SetAge(age int) { d.bird.age = age; d.lizard.age = age }
func (d *Dragon) Climb()         { d.lizard.Climb() }
func (d *Dragon) Fly()           { d.bird.Fly() }

func NewDragon(age int) *Dragon {
	d := &Dragon{Bird{}, Lizard{}}
	d.SetAge(age)
	return d
}

/*****************************************/
// Main Function
/*****************************************/
func main() {
	d := NewDragon(0)
    fmt.Println("Dragon's age is:", d.Age()) // Dragon's age is: 0
	d.SetAge(30)
	fmt.Println("Dragon's age is:", d.Age()) // Dragon's age is: 30
	d.Climb()                                // Climbing!
	d.Fly()                                  // Flying!
}
```

---

# Summary
