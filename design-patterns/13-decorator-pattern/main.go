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
