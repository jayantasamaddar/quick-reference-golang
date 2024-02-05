package lib

import "fmt"

type person struct {
	name string
	age  int
}

type Person interface {
	SayHello()
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old.\n", p.name, p.age)
}

// In this implementation, since we return the Person interface, it doesn't matter if the struct fields themselves are global or not
func NewPerson(name string, age int) Person {
	return &person{name, age}
}
