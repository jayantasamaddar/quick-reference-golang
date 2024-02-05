package lib

const (
	Male = iota
	Female
)

// Prototype interface
type Cloneable interface {
	CloneInto(gender int, married bool) Cloneable
}

// Concrete prototype
type Candidate struct {
	Name, Title string
	Married     bool
	Age         int
}

func NewCandidate(name string, gender, age int, married bool) *Candidate {
	c := &Candidate{Name: name, Age: age, Married: married}
	switch gender {
	case Male:
		c.Title = "Mr."
	case Female:
		if c.Married {
			c.Title = "Mrs."
		} else {
			c.Title = "Ms."
		}
	default:
		panic("gender should either be Male or Female")
	}
	return c
}

func (c *Candidate) CloneInto(gender int, married bool) Cloneable {
	ca := &Candidate{Name: c.Name, Married: married, Age: c.Age}
	switch gender {
	case Male:
		ca.Title = "Mr."
	case Female:
		if ca.Married {
			ca.Title = "Mrs."
		} else {
			ca.Title = "Ms."
		}
	default:
		panic("gender should either be Male or Female")
	}
	return ca
}
